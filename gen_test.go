package colornames

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestMatchNode(t *testing.T) {

	tests := []struct {
		html  string
		match map[string]string
	}{
		{`<div class="abcd />"`, map[string]string{matchNodeType: "div", "class": "abcd"}},
	}

	for _, test := range tests {
		n, err := html.Parse(bytes.NewBufferString(test.html))
		if err != nil {
			t.Fatal(err)
		}
		if !matchNode(n, test.match) {
			t.Errorf("failed %s", test.html)
		}
	}

}
