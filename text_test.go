package pdf

import "testing"

// helper to create a Font with a specified space glyph width (in font units).
func newTestFont(spaceWidth float64) Font {
	r := &Reader{}
	widths := array{spaceWidth}
	d := dict{
		name("FirstChar"): int64(32),
		name("LastChar"):  int64(32),
		name("Widths"):    widths,
	}
	return Font{V: Value{r, objptr{}, d}}
}

func TestCanBeMerged_FontSpaceWidth(t *testing.T) {
	font := newTestFont(500) // space glyph width 500 units
	span1 := Span{Font: "F1", FontSize: 12, X: 0, Y: 0, W: 20}
	expected := 500.0 / 1000 * span1.FontSize
	span2 := Span{Font: "F1", FontSize: 12, X: span1.X + span1.W + expected, Y: 0, W: 10}
	if !canBeMerged(span1, span2, font, 0, 0, DefaultMergeOptions) {
		t.Fatalf("expected spans to merge")
	}

	// Gap much larger than allowed should prevent merging.
	span2.X = span1.X + span1.W + expected*DefaultMergeOptions.SpaceMultiplier + 1
	if canBeMerged(span1, span2, font, 0, 0, DefaultMergeOptions) {
		t.Fatalf("expected spans not to merge due to large gap")
	}
}

func TestCanBeMerged_WithSpacingOperators(t *testing.T) {
	font := newTestFont(250) // space glyph width 250 units
	span1 := Span{Font: "F1", FontSize: 12, X: 0, Y: 0, W: 15}
	charSpacing := 1.0
	wordSpacing := 2.0
	expected := 250.0/1000*span1.FontSize + charSpacing + wordSpacing
	span2 := Span{Font: "F1", FontSize: 12, X: span1.X + span1.W + expected - 0.1, Y: 0, W: 5}

	if !canBeMerged(span1, span2, font, charSpacing, wordSpacing, DefaultMergeOptions) {
		t.Fatalf("expected spans to merge with spacing adjustments")
	}

	// Without spacing operators the gap is too large and should not merge.
	if canBeMerged(span1, span2, font, 0, 0, DefaultMergeOptions) {
		t.Fatalf("expected spans not to merge without Tc/Tw")
	}
}
