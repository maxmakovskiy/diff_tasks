package not_very_secure

import (
	"testing"
)

/*
	Expect(alphanumeric(".*?")).To(Equal(false))
	Expect(alphanumeric("a")).To(Equal(true))
	Expect(alphanumeric("Mazinkaiser")).To(Equal(true))
	Expect(alphanumeric("hello world_")).To(Equal(false))
	Expect(alphanumeric("PassW0rd")).To(Equal(true))
	Expect(alphanumeric("     ")).To(Equal(false))
	Expect(alphanumeric("")).To(Equal(false))
	Expect(alphanumeric("\n\t\n")).To(Equal(false))
	Expect(alphanumeric("ciao\n$$_")).To(Equal(false))
	Expect(alphanumeric("__ * __")).To(Equal(false))
	Expect(alphanumeric("&)))(((")).To(Equal(false))
	Expect(alphanumeric("43534h56jmTHHF3k")).To(Equal(true))
*/

func TestAlphanumeric(t *testing.T) {
	tables := []struct {
		input  string
		result bool
	}{
		{".?*", false},
		{"a", true},
		{"Mazinkaiser", true},
		{"hello world_", false},
		{"PassW0rd", true},
		{"     ", false},
		{"", false},
		{"\n\t\n", false},
		{"ciao\n$$_", false},
		{"__ * __", false},
		{"&)))(((", false},
		{"43534h56jmTHHF3k", true},
	}

	for _, table := range tables {
		alphanumericResult := alphanumeric(table.input)
		if alphanumericResult != table.result {
			t.Errorf("INCORRECT alphanumeric(\"%s\") -> GOT: %t; WANT: %t", table.input, alphanumericResult, table.result)
		}
	}
}
