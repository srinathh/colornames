package colornames

import (
	"bytes"
	"testing"

	"golang.org/x/net/html/atom"

	"golang.org/x/net/html"
)

func TestMatch(t *testing.T) {

	testcases := []struct {
		text string
		atm  atom.Atom
		attr []html.Attribute
	}{
		{`<li class="abcd">data</li>`, atom.Li, []html.Attribute{html.Attribute{"", "class", "abcd"}}},
		{`<span class="color-keyword-value">rgb( 0, 139, 139)</span>`, atom.Span, []html.Attribute{html.Attribute{"", "class", "color-keyword-value"}}},
	}

	for _, testcase := range testcases {
		tknr := html.NewTokenizer(bytes.NewBufferString(testcase.text))
		if tknr.Next() == html.ErrorToken {
			t.Fatalf("Got an error token in testcase %s", testcase.text)
		}
		tkn := tknr.Token()

		if !matchToken(tkn, testcase.atm, testcase.attr) {
			t.Errorf("Did not detect %s %s in %s", testcase.atm.String(), testcase.attr, testcase.text)
			t.Error(tkn)
		}
	}

}

func TestExtractSVG11(t *testing.T) {
	t.Log(extractSVG11(bytes.NewBufferString(testdata)))
}
