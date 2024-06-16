package templates

import (
	"fmt"
	"io/fs"
	"text/template"

	"github.com/kawana77b/univenv/internal/sysutil/ostype"
	"github.com/kawana77b/univenv/internal/sysutil/shell"
)

const (
	script_go_tmpl = "template/script.go.tmpl"
)

func LoadScriptTemplate(fsys fs.FS, sh shell.Shell, o ostype.OS) (*ScriptTemplate, error) {
	tpl, err := template.ParseFS(fsys, script_go_tmpl)
	if err != nil {
		return nil, fmt.Errorf("failed to parse template: %s", script_go_tmpl)
	}
	return NewScriptTemplate(tpl, sh, o), nil
}
