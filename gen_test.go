package colornames

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

var testdiv = `<div class="a">a<div class="b">b1</div><div class="b">b2</div><div class="c">c</div></div>`

func TestFindChildren(t *testing.T) {
	n, err := html.Parse(bytes.NewBufferString(testdiv))
	if err != nil {
		t.Fatal(err)
	}

	ret := findChildren(n, map[string]string{matchAtom: atom.Div.String(), atom.Class.String(): "b"})
	t.Log(ret)
	if len(ret) != 2 {
		t.Error("Did not find the right items")
	}
}
