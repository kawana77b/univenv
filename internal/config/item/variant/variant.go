package variant

type Variant string

const (
	ITEM      = Variant("item")
	CONDITION = Variant("condition")
)

func (v Variant) String() string {
	return string(v)
}

func (v Variant) IsItem() bool {
	return v == ITEM
}

func (v Variant) IsCondition() bool {
	return v == CONDITION
}
