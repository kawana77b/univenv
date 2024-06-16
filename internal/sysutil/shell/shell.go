package shell

import (
	"slices"

	"github.com/kawana77b/univenv/internal/common"
)

type Shell string

const (
	BASH = Shell("bash")
	FISH = Shell("fish")
	PWSH = Shell("pwsh")
)

func (i Shell) String() string {
	return string(i)
}

func (s Shell) Match(sh ...Shell) bool {
	return slices.Contains(sh, s)
}

func (s Shell) Validate() error {
	if !Contains(s.String()) {
		return common.NewUnknownError("shell")
	}
	return nil
}

func All() []Shell {
	return []Shell{BASH, FISH, PWSH}
}

func Contains(v string) bool {
	return slices.Contains(All(), Shell(v))
}

func FromString(s string) (Shell, bool) {
	if Contains(s) {
		return Shell(s), true
	}
	return "", false
}
