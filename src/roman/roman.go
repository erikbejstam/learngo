package roman

const (
	i    = "I"
	v    = "V"
	x    = "X"
	l    = "L"
	c    = "C"
	d    = "D"
	m    = "M"
	vHat = "vHat"
	xHat = "xHat"
)

type Arabic struct {
	Thousands int
	Hundreds  int
	Tens      int
	Ones      int
}

func (a *Arabic) toRoman() string {
	var roman string

	roman += digitToRoman(a.Thousands, vHat, m, xHat)
	roman += digitToRoman(a.Hundreds, d, c, m)
	roman += digitToRoman(a.Tens, l, x, c)
	roman += digitToRoman(a.Ones, v, i, x)

	return roman
}

func digitToRoman(number int, mid, modifier, higher string) string {
	var roman string

	switch {
	case number < 4:
		for number > 0 {
			number--
			roman = modifier + roman
		}
	case number == 4:
		roman = modifier + mid
	case number == 5:
		roman = mid
	case number > 5 && number < 9:
		roman = mid
		number -= 5
		for number > 0 {
			number--
			roman = roman + modifier
		}
	case number == 9:
		roman = modifier + higher
	}

	return roman
}

func parseInt(number int) Arabic {
	return Arabic{
		Ones:      number % 10,
		Tens:      (number / 10) % 10,
		Hundreds:  (number / 100) % 10,
		Thousands: (number / 1000) % 10,
	}
}
