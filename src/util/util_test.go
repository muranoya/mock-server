package util_test

import (
	"testing"

	"github.com/muranoya/mock-server/src/util"
)

func TestContaines(t *testing.T) {
	if util.Containes([]string{}, "") {
		t.Fatal()
	}

	if util.Containes([]string{"apple", "lemon", "grape"}, "orange") {
		t.Fatal()
	}

	if !util.Containes([]string{"apple", "lemon", "grape"}, "apple") {
		t.Fatal()
	}
}
