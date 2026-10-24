package main

import (
	"image"
	"reflect"
	"strings"
	"testing"

	ui "github.com/gizak/termui/v3"
)

func TestMarkdownAreaLineBreaks(t *testing.T) {
	const label = "[Val av preposition · MIC-W09]"
	const question = "Fyll i prepositionen: ”Jag litar … dig.”"
	tests := []struct {
		name  string
		text  string
		width int
		want  []string
	}{
		{"html break", label + "<br>" + question, 90, []string{label, question}},
		{"self closing break", label + "<br/>" + question, 90, []string{label, question}},
		{"spaced break", label + "<br />" + question, 90, []string{label, question}},
		{"uppercase break", label + "<BR>" + question, 90, []string{label, question}},
		{"wrapped question", label + "<br>" + question, 30, []string{"[Val av preposition ·", "MIC-W09]", "Fyll i prepositionen: ”Jag", "litar … dig.”"}},
		{"plain question", label + " " + question, 90, []string{label + " " + question}},
		{"inline code", "Use `<br>` literally", 90, []string{"Use <br> literally"}},
		{"escaped tag", `Use \<br> literally`, 90, []string{"Use <br> literally"}},
		{"formatted question", "**Label**<br>*Question*", 90, []string{"Label", "Question"}},
		{"Swedish word boundary", "En fråga kommer här", 8, []string{"En", "fråga", "kommer", "här"}},
		{"quoted word boundary", "En ”fråga” kommer här", 10, []string{"En", "”fråga”", "kommer", "här"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			area := NewMarkdownArea()
			area.Border = false
			area.Text = tt.text
			area.SetRect(0, 0, tt.width, 20)
			buf := ui.NewBuffer(area.Rectangle)
			area.Draw(buf)

			var rows []string
			for y := area.Inner.Min.Y; y < area.Inner.Max.Y; y++ {
				var row strings.Builder
				for x := area.Inner.Min.X; x < area.Inner.Max.X; x++ {
					row.WriteRune(buf.GetCell(image.Pt(x, y)).Rune)
				}
				if text := strings.TrimSpace(row.String()); text != "" {
					rows = append(rows, text)
				}
			}
			if !reflect.DeepEqual(rows, tt.want) {
				t.Errorf("rendered rows = %q, want %q", rows, tt.want)
			}
		})
	}
}
