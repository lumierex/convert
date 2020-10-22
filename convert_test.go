package convert_test

import (
	"github.com/sfpxm/convert"
	"testing"
)

func TestStrTo_Int(t *testing.T) {
	s := "2"

	i, err := convert.StrTo(s).Int()

	if err != nil {
		t.Log(err)
	}

	if i != 2 {
		t.Errorf("str to int failed")
	}
}