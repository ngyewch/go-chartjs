package chartjs

import (
	"encoding/json"
	"testing"
)

func TestFontWeight(t *testing.T) {
	t.Logf("FontWeight: bold -> %s\n", dumpJson(FontWeightBold))
	t.Logf("FontWeight: lighter -> %s\n", dumpJson(FontWeightLighter))
	t.Logf("FontWeight: 900 -> %s\n", dumpJson(FontWeight(900)))
}

func TestFontSpec(t *testing.T) {
	fontSpec := FontSpec{
		Family: "sans-serif",
	}
	t.Logf("FontSpec: %s\n", dumpJson(fontSpec))
	fontSpec = FontSpec{
		Family: "serif",
		Style:  FontStyleItalic,
		Weight: FontWeightBolder,
	}
	t.Logf("FontSpec: %s\n", dumpJson(fontSpec))
}

func dumpJson(v any) string {
	jsonBytes, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(jsonBytes)
}
