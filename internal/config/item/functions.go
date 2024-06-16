package item

import (
	"github.com/kawana77b/univenv/internal/sysutil/ostype"
	"github.com/kawana77b/univenv/internal/sysutil/shell"
)

func GetEnabledItems(items []Item, o ostype.OS, sh shell.Shell) []Item {
	var result []Item
	for _, i := range items {
		if !isEnable(i, o, sh) {
			continue
		}
		value, values := i, []Item{}
		value.Items = append(values, GetEnabledItems(i.Items, o, sh)...)
		result = append(result, value)
	}
	return result
}

func isEnable(i Item, o ostype.OS, sh shell.Shell) bool {
	return !i.Disabled && osMatch(i, o) && shellMatch(i, sh)
}

func shellMatch(i Item, sh shell.Shell) bool {
	if len(i.Shell) == 0 {
		return true
	}
	return sh.Match(i.Shell...)
}

func osMatch(i Item, o ostype.OS) bool {
	if len(i.OS) == 0 {
		return true
	}
	return o.Match(i.OS...)
}
