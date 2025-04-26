package v1

import "testing"

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      int
		Want        string
	}{
		{"1 to roman", 1, "I"},
		{"2 to roman", 2, "II"},
		{"3 to roman", 3, "III"},
		{"4 to roman not adding `I`", 4, "IV"},
		{"5 to roman", 5, "V"},
		{"6 to roman", 6, "VI"},
		{"7 to roman", 7, "VII"},
		{"8 to roman", 8, "VIII"},
		{"9 to roman", 9, "IX"},
		{"10 gets converted to X", 10, "X"},
		{"14 gets converted to XIV", 14, "XIV"},
		{"18 gets converted to XVIII", 18, "XVIII"},
		{"20 gets converted to XX", 20, "XX"},
		{"39 gets converted to XXXIX", 39, "XXXIX"},
		{"40 gets converted to XL", 40, "XL"},
		{"47 gets converted to XLVII", 47, "XLVII"},
		{"49 gets converted to XLIX", 49, "XLIX"},
		{"50 gets converted to L", 50, "L"},
	}

	for _, test := range cases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			if got != test.Want {
				t.Errorf("got %q want %q ", got, test.Want)
			}
		})
	}
}
