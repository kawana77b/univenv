package builder_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/kawana77b/univenv/internal/builder"
)

func Test_ScriptBuilder(t *testing.T) {
	fs := os.DirFS("../../")
	for _, sh := range []string{"bash", "fish", "pwsh"} {
		b := &builder.ScriptBuilder{}
		b.DetectOS()

		b.SetFS(fs)
		b.SetShell(sh)
		b.SetNoComment(false)
		b.SetTarget("")
		b.SetFile("../../tests/config.yml")

		text, err := b.Build()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if !strings.HasPrefix(text, "# --") {
			t.Errorf("If NoComment is false, then it should have a comment.")
		}
	}
}

func Test_ScriptBuilderNoComment(t *testing.T) {
	fs := os.DirFS("../../")
	for _, sh := range []string{"bash", "fish", "pwsh"} {
		b := &builder.ScriptBuilder{}
		b.DetectOS()

		b.SetFS(fs)
		b.SetShell(sh)
		// NoComment is true
		b.SetNoComment(true)
		b.SetTarget("")
		b.SetFile("../../tests/config.yml")

		text, err := b.Build()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
		if strings.HasPrefix(text, "# --") {
			t.Errorf("Prefix Comment Exists")
		}
	}
}

func Test_ScriptBuilderDemo(t *testing.T) {
	fs := os.DirFS("../../")
	for _, sh := range []string{"bash", "fish", "pwsh"} {
		// set UNIVENV_CONFIG_DIR
		wd, _ := os.Getwd()
		d := filepath.Join(wd, "../", "../", "tests")
		os.Setenv("UNIVENV_CONFIG_DIR", d)

		b := &builder.ScriptBuilder{}
		b.DetectOS()

		b.SetFS(fs)
		b.SetShell(sh)
		b.SetNoComment(false)
		b.SetTarget("demo") // tests/config.demo.yml

		_, err := b.Build()
		if err != nil {
			t.Errorf("Error: %v", err)
		}
	}
}
