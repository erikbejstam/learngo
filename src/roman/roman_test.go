package roman

import (
	"reflect"
	"testing"
)

func TestToRoman(t *testing.T) {
	t.Run("ones", func(t *testing.T) {
		arabic := Arabic{Ones: 7}
		want := "VII"
		got := arabic.toRoman()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("tens", func(t *testing.T) {
		arabic := Arabic{Ones: 2, Tens: 4}
		want := "XLII"
		got := arabic.toRoman()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("hundreds", func(t *testing.T) {
		arabic := Arabic{Ones: 9, Tens: 7, Hundreds: 8}
		want := "DCCCLXXIX"
		got := arabic.toRoman()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("thousands", func(t *testing.T) {
		arabic := Arabic{Ones: 3, Tens: 3, Hundreds: 3, Thousands: 3}
		want := "MMMCCCXXXIII"
		got := arabic.toRoman()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func TestParseInt(t *testing.T) {
	t.Run("ones", func(t *testing.T) {
		data := 9
		want := Arabic{Ones: 9}
		got := parseInt(data)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("tens", func(t *testing.T) {
		data := 78
		want := Arabic{Ones: 8, Tens: 7}
		got := parseInt(data)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
	t.Run("hundres", func(t *testing.T) {
		data := 490
		want := Arabic{Ones: 0, Tens: 9, Hundreds: 4}
		got := parseInt(data)

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}
