package address_test

import (
	. "tests/address"
	"math/rand"
	"testing"
)

type TestScence struct {
	insertedAddr   string
	expectedResult string
}

func TestGetAddressType(t *testing.T) {
	t.Parallel()

	testScences := []TestScence{
		{"rua dos coqueiros", "rua"},
		{"avenida dos coqueiros", "avenida"},
		{"estrada dos coqueiros", "estrada"},
		{"rodovia dos coqueiros", "rodovia"},
		{"coqueiros", "tipo invalido"},
	}

	for _, scence := range testScences {
		if GetAddressType(scence.insertedAddr) != scence.expectedResult {
			t.Errorf(
				"\nresult: %s\nexpected: %s",
				ValidTypes[rand.Intn(4)],
				scence.expectedResult,
			)
		}
	}
}
