package itype_test

import (
	"testing"

	"github.com/kawana77b/univenv/internal/config/item/itype"
)

func Test_Variant(t *testing.T) {
	for _, v := range itype.All() {
		if v.Variant() == "" {
			t.Errorf("item type %s has no variant", v)
		}
	}
}
