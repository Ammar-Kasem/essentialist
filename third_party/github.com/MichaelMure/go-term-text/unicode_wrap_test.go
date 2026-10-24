package text

import "testing"

func TestWrapUnicodeWords(t *testing.T) {
	for _, tc := range []struct {
		name  string
		text  string
		width int
		want  string
	}{
		{"accented word", "En fråga kommer här", 6, "En\nfråga\nkommer\nhär"},
		{"smart quotes", "En ”fråga” kommer här", 8, "En\n”fråga”\nkommer\nhär"},
		{"combining mark", "En fra\u030aga kommer här", 6, "En\nfra\u030aga\nkommer\nhär"},
		{"styled word", "En \x1b[1mfråga\x1b[0m kommer här", 6, "En\n\x1b[1mfråga\x1b[0m\nkommer\nhär"},
		{"long word", "frågetecken", 5, "fråge\ntecke\nn"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got, _ := Wrap(tc.text, tc.width); got != tc.want {
				t.Errorf("wrapped text = %q, want %q", got, tc.want)
			}
		})
	}
}
