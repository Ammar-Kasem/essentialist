package internal

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestOpenDeck(t *testing.T) {
	d, err := NewDeckFromFile("samples/testdata/test-1.md")
	if err != nil {
		t.Fatal(err)
	}
	cards := d.SelectBefore(time.Now())
	if len(cards) != 5 {
		t.Errorf("Missing cards: %d", len(cards))
	}
}

func TestMissingAnswer(t *testing.T) {
	_, err := NewDeckFromFile("samples/testdata/test-2.md")
	if err == nil {
		t.Error("missing error")
	}
}

func TestSaveAndLoadProgressJSON(t *testing.T) {
	dir := t.TempDir()
	filename := filepath.Join(dir, "deck.md")
	err := os.WriteFile(filename, []byte(`
## question 1
answer 1
## question 2
answer 2
`), 0600)
	if err != nil {
		t.Fatalf("write error: %s", err)
	}
	d, err := NewDeckFromFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	d.Cards[0].Review(ScoreGood)
	err = d.SaveDeckMeta()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(dir, "deck.json")); err != nil {
		t.Fatal(err)
	}
	d, err = NewDeckFromFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	if d.Cards[0].Meta.FSRSReps != 1 {
		t.Errorf("saved review count = %d, want 1", d.Cards[0].Meta.FSRSReps)
	}
	if d.Cards[1].Meta.FSRSReps != 0 {
		t.Errorf("unreviewed card count = %d, want 0", d.Cards[1].Meta.FSRSReps)
	}
}
