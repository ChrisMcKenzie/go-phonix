package soundex

import "testing"

func TestMatch(t *testing.T) {
	tests := []struct {
		Words       [2]string
		Expectation bool
	}{
		{
			[2]string{
				"robert",
				"rupert",
			},
			true,
		},
		{
			[2]string{
				"robert",
				"rubin",
			},
			false,
		},
		{
			[2]string{
				"ashcraft",
				"ashcroft",
			},
			true,
		},
	}

	for _, test := range tests {
		result := Match(test.Words[0], test.Words[1])
		if result != test.Expectation {
			t.Errorf("Test Failed Expected %s and %s to return %t", test.Words[0], test.Words[1], test.Expectation)
			t.Fail()
		}
	}
}

func TestEncode(t *testing.T) {
	tests := []struct {
		Word   string
		Result string
	}{
		{
			"rubin",
			"r150",
		},
		{
			"robert",
			"r163",
		},
		{
			"ashcroft",
			"a261",
		},
		{
			"ashcraft",
			"a261",
		},
	}

	for _, test := range tests {
		if result := Encode(test.Word); result != test.Result {
			t.Errorf("Expected %s to be encoded to %s got %s", test.Word, test.Result, result)
		}
	}
}
