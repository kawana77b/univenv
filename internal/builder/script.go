package builder

import (
	"fmt"
	"io/fs"
	"strings"

	"github.com/kawana77b/univenv/internal/common"
	"github.com/kawana77b/univenv/internal/config"
	"github.com/kawana77b/univenv/internal/sysutil/ostype"
	"github.com/kawana77b/univenv/internal/sysutil/shell"
	"github.com/kawana77b/univenv/internal/templates"
	"github.com/kawana77b/univenv/internal/translator"
)

type ScriptBuilderOptions struct {
	file        string
	isNoComment bool
	target      string
}

type ScriptBuilder struct {
	os    ostype.OS
	shell shell.Shell
	fs    fs.FS

	options ScriptBuilderOptions
}

func (b *ScriptBuilder) SetShell(os string) {
	if v, ok := shell.FromString(os); ok {
		b.shell = v
	}
}

func (b *ScriptBuilder) DetectOS() {
	if os, ok := ostype.GetOS(); ok {
		b.os = os
	}
}

func (b *ScriptBuilder) SetFS(fs fs.FS) {
	b.fs = fs
}

func (b *ScriptBuilder) SetTarget(target string) {
	b.options.target = target
}

func (b *ScriptBuilder) SetFile(target string) {
	b.options.file = target
}

func (b *ScriptBuilder) SetNoComment(isNoComment bool) {
	b.options.isNoComment = isNoComment
}

func (b *ScriptBuilder) validate() error {
	return common.Validate(b.os, b.shell)
}

var insertComment = map[string]string{
	"start": fmt.Sprintf("# %s Created By univenv %s", strings.Repeat("-", 20), strings.Repeat("-", 20)),
	"end":   fmt.Sprintf("# %s End Of Created By univenv %s", strings.Repeat("-", 20), strings.Repeat("-", 20)),
}

func (b *ScriptBuilder) Build() (string, error) {
	if err := b.validate(); err != nil {
		return "", err
	}
	opener := getConfigOpner(b.options.file, b.options.target)
	conf, err := opener.Open()
	if err != nil {
		return "", err
	}

	tpl, err := templates.LoadScriptTemplate(b.fs, b.shell, b.os)
	if err != nil {
		return "", err
	}

	items := conf.GetEnableItems(b.shell, b.os)
	t := translator.NewItemTranslator(tpl)
	buf, err := t.TranslateAll(items)
	if err != nil {
		return "", err
	}
	output := string(buf)
	if !b.options.isNoComment {
		output = strings.Join([]string{
			insertComment["start"],
			output,
			insertComment["end"],
		}, "\n")
	}
	return output, nil
}

func getConfigOpner(file, target string) *config.ConfigOpener {
	opener := &config.ConfigOpener{}
	opener.SetFile(file)
	opener.SetTarget(target)
	return opener
}
