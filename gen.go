package colornames

import (
	//"fmt"
	"golang.org/x/net/html"
	"io"

	"golang.org/x/net/html/atom"
)

//function matchToken checks if a html.Token is of the required atom type
//and the specified attributes
func matchToken(t html.Token, a atom.Atom, attrs []html.Attribute) bool {
	if t.DataAtom != a {
		return false
	}

OuterLoop:
	for _, attr := range attrs { //for each attribute we want to match
		for _, a := range t.Attr { //range through the token Attr to find it
			if a.Key == attr.Key && a.Val == attr.Val {
				continue OuterLoop //and when found, move to the next attribute
			}
		}
		return false //we did not detect a match since loop terminated, so return false
	}

	return true //loop terminated without returning, so we have a full match

}

// extractSVG11 extracts named colors from the SVG 1.1 spec document
//
// In the document, data for the named colors are contained in tabel
// cells in span elements. Spans with the class "prop-value" define the
// color name eg. aliceblue and Spans with the class "color-keyword-value"
// define the color as rgb(240, 248, 255). These spans are in adjacent
// cells. We use html.Tokenizer to run through the document to find these
// spans with matchToken and extract the data
func extractSVG11(r io.Reader) ([]string, []string, error) {
	name_span := []html.Attribute{html.Attribute{"", "class", "prop-value"}}
	value_span := []html.Attribute{html.Attribute{"", "class", "color-keyword-value"}}

	colornames := []string{}
	colorvals := []string{}
	nextfill := 0

	tknr := html.NewTokenizer(r)
	for {
		ttyp := tknr.Next()
		tkn := tknr.Token()

		switch true {
		case ttyp == html.StartTagToken && matchToken(tkn, atom.Span, name_span):
			nextfill = 1

		case ttyp == html.StartTagToken && matchToken(tkn, atom.Span, value_span):
			nextfill = 2

		case ttyp == html.TextToken:
			switch nextfill {
			case 1:
				colornames = append(colornames, tkn.Data)
			case 2:
				colorvals = append(colorvals, tkn.Data)
			}
			nextfill = 0

		case ttyp == html.ErrorToken:
			if tknr.Err() != io.EOF {
				return nil, nil, tknr.Err()
			}
			return colornames, colorvals, nil
		default:
			nextfill = 0
		}
	}
}
