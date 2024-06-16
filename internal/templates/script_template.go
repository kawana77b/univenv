package templates

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/kawana77b/univenv/internal/sysutil/ostype"
	"github.com/kawana77b/univenv/internal/sysutil/shell"
)

type ScriptTemplate struct {
	tpl   *template.Template
	shell shell.Shell
	os    ostype.OS
}

func NewScriptTemplate(tpl *template.Template, sh shell.Shell, o ostype.OS) *ScriptTemplate {
	return &ScriptTemplate{
		tpl:   tpl,
		shell: sh,
		os:    o,
	}
}

func (t *ScriptTemplate) cleanPath(path string) string {
	if t.shell == shell.PWSH {
		// Since PowerShell does not have the behavior of replacing ~ with an absolute path when used, $HOME is used instead.
		if strings.HasPrefix(path, "~") {
			path = strings.Replace(path, "~", "$HOME", 1)
		}
		if t.os == ostype.WINDOWS {
			// expected: foo/baz/bar -> foo\baz\bar
			return filepath.Clean(path)
		}
	}
	return filepath.ToSlash(path)
}

func (t *ScriptTemplate) fixComma(value string) string {
	if t.shell == shell.PWSH {
		// Powershell templates sometimes enclose the text in "" so that it can be removed in some situations.
		// cf. bash -> export FOO='bar', powershell -> $env:FOO = "bar"
		if strings.HasPrefix(value, "'") && strings.HasSuffix(value, "'") {
			return strings.Trim(value, "'")
		}
	}
	return value
}

func (t *ScriptTemplate) String() string {
	return fmt.Sprintf("ScriptTemplate{tpl: %s, shell: %s, os: %s}", t.tpl.Name(), t.shell, t.os)
}

func (t *ScriptTemplate) execute(name string, data any) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	err := t.tpl.ExecuteTemplate(&buf, fmt.Sprintf("%s:%s", t.shell, name), data)
	return &buf, err
}

func (t *ScriptTemplate) Raw(value string) (*bytes.Buffer, error) {
	var bytes bytes.Buffer
	bytes.WriteString(value)
	return &bytes, nil
}

func (t *ScriptTemplate) Comment(value string) (*bytes.Buffer, error) {
	return t.execute("comment", struct {
		Value string
	}{
		Value: value,
	})
}

func (t *ScriptTemplate) Env(name, value string) (*bytes.Buffer, error) {
	return t.execute("env", struct {
		Name  string
		Value string
	}{
		Name:  name,
		Value: t.fixComma(t.cleanPath(value)),
	})
}

var pathSeparators = map[shell.Shell]map[ostype.OS]string{
	shell.BASH: {
		ostype.WINDOWS: ":",
		ostype.LINUX:   ":",
		ostype.DARWIN:  ":",
	},
	shell.FISH: {
		ostype.WINDOWS: " ",
		ostype.LINUX:   " ",
		ostype.DARWIN:  " ",
	},
	// https://learn.microsoft.com/ja-jp/powershell/module/microsoft.powershell.core/about/about_environment_variables?view=powershell-7.4#set-environment-variables-in-your-profile
	shell.PWSH: {
		ostype.WINDOWS: ";",
		ostype.LINUX:   ":",
		ostype.DARWIN:  ":",
	},
}

func (t *ScriptTemplate) PATH(value string) (*bytes.Buffer, error) {
	var separator string = ":"
	if sep, ok := pathSeparators[t.shell][t.os]; ok {
		separator = sep
	}
	return t.execute("path", struct {
		Value     string
		Separator string
	}{
		Value:     t.cleanPath(value),
		Separator: separator,
	})
}

func (t *ScriptTemplate) Alias(name, value string) (*bytes.Buffer, error) {
	return t.execute("alias", struct {
		Name  string
		Value string
	}{
		Name:  name,
		Value: value,
	})
}

func (t *ScriptTemplate) Source(value string) (*bytes.Buffer, error) {
	return t.execute("source", struct {
		Value string
	}{
		Value: t.cleanPath(value),
	})
}

func (t *ScriptTemplate) Command(command, value string) (*bytes.Buffer, error) {
	b, err := t.type_(command)
	if err != nil {
		return b, err
	}
	return t.and(b.String(), value)
}

func (t *ScriptTemplate) Directory(path, value string) (*bytes.Buffer, error) {
	b, err := t.test_directory(path)
	if err != nil {
		return b, err
	}
	return t.and(b.String(), value)
}

func (t *ScriptTemplate) If_Command(command string, items ...string) (*bytes.Buffer, error) {
	b, err := t.type_(command)
	if err != nil {
		return b, err
	}
	return t.If(b.String(), items...)
}

func (t *ScriptTemplate) If_Directory(path string, items ...string) (*bytes.Buffer, error) {
	b, err := t.test_directory(path)
	if err != nil {
		return b, err
	}
	return t.If(b.String(), items...)
}

func (t *ScriptTemplate) type_(command string) (*bytes.Buffer, error) {
	return t.execute("type", struct {
		Command string
	}{
		Command: command,
	})
}

type test_type string

const (
	test_type_directory = test_type("directory")
	test_type_file      = test_type("file")
)

func (t test_type) String() string {
	return string(t)
}

// test -d
func (t *ScriptTemplate) test_directory(path string) (*bytes.Buffer, error) {
	return t.test(t.cleanPath(path), test_type_directory)
}

var testOperators = map[shell.Shell]map[test_type]string{
	shell.BASH: {
		test_type_file:      "-f",
		test_type_directory: "-d",
	},
	shell.FISH: {
		test_type_file:      "-f",
		test_type_directory: "-d",
	},
	shell.PWSH: {},
}

// cf. test -f, ftype: file, directory
func (t *ScriptTemplate) test(value string, ttype test_type) (*bytes.Buffer, error) {
	var operator string
	if op, ok := testOperators[t.shell][ttype]; ok {
		operator = op
	}
	return t.execute("test", struct {
		Operator string
		Value    string
	}{
		Operator: operator,
		Value:    value,
	})
}

// if block
func (t *ScriptTemplate) If(cond string, value ...string) (*bytes.Buffer, error) {
	return t.if_(cond, value...)
}

// if block
func (t *ScriptTemplate) if_(cond string, value ...string) (*bytes.Buffer, error) {
	return t.execute("if", struct {
		Condition string
		Values    []string
	}{
		Condition: cond,
		Values:    value,
	})
}

// &&
func (t *ScriptTemplate) and(left, right string) (*bytes.Buffer, error) {
	return t.execute("and", struct {
		Left  string
		Right string
	}{
		Left:  left,
		Right: right,
	})
}
