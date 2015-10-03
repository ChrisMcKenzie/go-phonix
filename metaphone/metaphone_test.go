package metaphone

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		test        string
		expectation string
	}{
		{
			"cheese",
			"xhese",
		},
		{
			"cicatrice",
			"sikatrise",
		},
		{
			"manchester",
			"manxhester",
		},
		{
			"numb",
			"num",
		},
		{
			"gnu",
			"nu",
		},
		{
			"pnuemonia",
			"nuemonia",
		},
		{
			"pride",
			"prite",
		},
		{
			"school",
			"skhol",
		},
	}

	for _, test := range tests {
		if result := Encode(test.test); result != test.expectation {
			t.Errorf("Expected %s to encode to %s Got: %s", test.test, test.expectation, result)
		}
	}
}
