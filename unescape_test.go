package unescape

import "testing"

func TestUnescape(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{`a\nb`, "a\nb"},
		{`a\tb`, "a\tb"},
		{`\"q\"`, "\"q\""},
		{`\\`, "\\"},
		{`\x41`, "A"},
		{`\x4a`, "J"},
		{`\101`, "A"}, // 八进制
		{`no\escape`, `no\escape`},
		{`\q`, `\q`},
	}
	for _, c := range cases {
		if got := Unescape(c.in); got != c.want {
			t.Errorf("Unescape(%q)=%q want %q", c.in, got, c.want)
		}
	}
}
