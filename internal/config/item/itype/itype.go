package itype

import (
	"fmt"
	"slices"

	"github.com/kawana77b/univenv/internal/common"
	vari "github.com/kawana77b/univenv/internal/config/item/variant"
)

type ItemType string

const (
	COMMENT = ItemType("comment")
	ENV     = ItemType("env")
	PATH    = ItemType("path")
	ALIAS   = ItemType("alias")
	SOURCE  = ItemType("source")
	RAW     = ItemType("raw")

	IF_DIRECTORY = ItemType("if-directory")
	IF_COMMAND   = ItemType("if-command")
)

var toItemType = map[vari.Variant][]ItemType{
	vari.ITEM:      {COMMENT, ENV, PATH, ALIAS, SOURCE, RAW},
	vari.CONDITION: {IF_DIRECTORY, IF_COMMAND},
}

var toVariant = map[ItemType]vari.Variant{
	COMMENT: vari.ITEM,
	ENV:     vari.ITEM,
	PATH:    vari.ITEM,
	ALIAS:   vari.ITEM,
	SOURCE:  vari.ITEM,
	RAW:     vari.ITEM,

	IF_DIRECTORY: vari.CONDITION,
	IF_COMMAND:   vari.CONDITION,
}

var all = slices.Concat(toItemType[vari.ITEM], toItemType[vari.CONDITION])

func (i ItemType) String() string {
	return string(i)
}

func (i ItemType) Variant() vari.Variant {
	v, ok := toVariant[i]
	if !ok {
		panic(fmt.Sprintf("unknown item vairiant!!, ItemType: %s", i.String()))
	}
	return v
}

func (i ItemType) Validate() error {
	if !Contains(i.String()) {
		return common.NewUnknownError("item type")
	}
	return nil
}

func All() []ItemType {
	return all
}

func Contains(v string) bool {
	return slices.Contains(All(), ItemType(v))
}

func FromString(s string) (ItemType, bool) {
	if Contains(s) {
		return ItemType(s), true
	}
	return "", false
}
