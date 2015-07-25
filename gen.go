package colornames

import "golang.org/x/net/html"

const (
	matchAtom     = "MatchAtom"
	matchNodeType = "MatchNodeType"
)

var ndtype = map[string]html.NodeType{
	"ErrorNode":    html.ErrorNode,
	"TextNode":     html.TextNode,
	"DocumentNode": html.DocumentNode,
	"ElementNode":  html.ElementNode,
	"CommentNode":  html.CommentNode,
	"DoctypeNode":  html.DoctypeNode,
}

func matchNode(t *html.Node, match map[string]string) bool {
	matchctr := 0
	for k, v := range match {
		switch k {
		case matchAtom:
			if t.DataAtom.String() == v {
				matchctr++
			}
		case matchNodeType:
			if t.Type == ndtype[v] {
				matchctr++
			}
		default:
			for _, attr := range t.Attr {
				if attr.Key == k && attr.Val == v {
					matchctr++
					break
				}
			}
		}
	}
	if matchctr != len(match) {
		return false
	}
	return true
}

/*
func matchNode(t *html.Node, a atom.Atom, attrs []html.Attribute) bool {
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
*/

// findChildren searches all the children of the provided node
// and returns a slice of pointers to the children matching
// the specified node atom and classname. if classname is empty
// the function matches only by atom
func findChildren(n *html.Node, match map[string]string) []*html.Node {
	ret := []*html.Node{}
	if matchNode(n, match) {
		ret = append(ret, n)
	}

	//traverse the tree under the children and add to queue
	for cur := n.FirstChild; cur != n.LastChild; cur = cur.NextSibling {
		ret = append(ret, findChildren(cur, match)...)
	}
	return ret
}
