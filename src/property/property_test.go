package property

import (
	"testing"
	"testing/quick"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		Description string
		Arabic      uint16
		Want        string
	}{
		{"1 converts to I", 1, "I"},
		{"2 converts to II", 2, "II"},
		{"3 converts to III", 3, "III"},
		{"4 converts to IV", 4, "IV"},
		{"5 converts to V", 5, "V"},
		{"7 converts to VII", 7, "VII"},
		{"8 converts to VIII", 8, "VIII"},
		{"9 converts to IX", 9, "IX"},
		{"1984 converts to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"3119 converts to MMMCXIX", 3119, "MMMCXIX"},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			want := test.Want

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}

		})
	}
}

func TestConvertToArabic(t *testing.T) {
	cases := []struct {
		Description string
		Roman       string
		Want        uint16
	}{
		{"I converts to 1", "I", 1},
		{"II converts to 2", "II", 2},
		{"III converts to 3", "III", 3},
		{"IV converts to 4", "IV", 4},
		{"V converts to 5", "V", 5},
		{"VI converts to 6", "VI", 6},
		{"VII converts to 7", "VII", 7},
		{"VII converts to 8", "VIII", 8},
		{"IX converts to 9", "IX", 9},
		{"MMMCXIX converts to 3119", "MMMCXIX", 3119},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := ConvertToArabic(test.Roman)
			want := test.Want

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}

		})
	}
}

func TestRecursive(t *testing.T) {
	cases := []struct {
		Description string
		Roman       string
		Want        uint16
	}{
		{"I converts to 1", "I", 1},
		{"II converts to 2", "II", 2},
		{"III converts to 3", "III", 3},
		{"IV converts to 4", "IV", 4},
		{"V converts to 5", "V", 5},
		{"VI converts to 6", "VI", 6},
		{"VII converts to 7", "VII", 7},
		{"VII converts to 8", "VIII", 8},
		{"IX converts to 9", "IX", 9},
		{"MMMCXIX converts to 3119", "MMMCXIX", 3119},
		{"XCI converts to 91", "XCI", 91},
	}

	for _, test := range cases {
		t.Run("", func(t *testing.T) {
			got := Recursive(test.Roman)
			want := test.Want

			if got != want {
				t.Errorf("got %v want %v", got, want)
			}

		})
	}
}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursive("III")
	}
}

func BenchmarkConvertToArabic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ConvertToArabic("III")
	}
}

// Property tests

func TestPropertiesOfConversion(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			t.Log("testing", arabic)
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed checks", err)
	}
}
