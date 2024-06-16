package ostype

import (
	"runtime"
	"slices"

	"github.com/kawana77b/univenv/internal/common"
)

type OS string

const (
	LINUX   = OS("linux")
	WINDOWS = OS("windows")
	DARWIN  = OS("darwin")
)

func (i OS) String() string {
	return string(i)
}

func (o OS) Match(os ...OS) bool {
	return slices.Contains(os, o)
}

func (o OS) Validate() error {
	if !Contains(o.String()) {
		return common.NewUnknownError("os")
	}
	return nil
}

func All() []OS {
	return []OS{LINUX, WINDOWS, DARWIN}
}

func Contains(v string) bool {
	return slices.Contains(All(), OS(v))
}

func FromString(s string) (OS, bool) {
	if Contains(s) {
		return OS(s), true
	}
	return "", false
}

func GetOS() (OS, bool) {
	o := runtime.GOOS
	if Contains(o) {
		return OS(o), true
	}
	return "", false
}
