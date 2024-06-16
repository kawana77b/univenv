package translator

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/kawana77b/univenv/internal/common/container"
	"github.com/kawana77b/univenv/internal/common/task"
	"github.com/kawana77b/univenv/internal/config/item"
	"github.com/kawana77b/univenv/internal/config/item/itype"
	"github.com/kawana77b/univenv/internal/templates"
)

type ItemTranslator struct {
	tpl *templates.ScriptTemplate
}

func NewItemTranslator(tpl *templates.ScriptTemplate) *ItemTranslator {
	return &ItemTranslator{
		tpl: tpl,
	}
}

func (t *ItemTranslator) Translate(item item.Item) ([]byte, error) {
	result := container.NewBytesList()
	buf, err := t.createBuffer(item)
	if err != nil {
		return result.Bytes(), err
	}
	result.Append(buf.Bytes()...)
	return result.Bytes(), nil
}

func (t *ItemTranslator) TranslateAll(items []item.Item) ([]byte, error) {
	result := container.NewBytesList()
	for idx, item := range items {
		if b, err := t.Translate(item); err != nil {
			return nil, err
		} else {
			result.Append(b...)
		}
		if idx < len(items)-1 {
			result.Append('\n')
		}
	}
	return result.Bytes(), nil
}

func (t *ItemTranslator) createBuffer(i item.Item) (*bytes.Buffer, error) {
	if i.Type.Variant().IsCondition() {
		return t.createConditionBuffer(i)
	}
	return t.createItemBuffer(i)
}

func (t *ItemTranslator) createItemBuffer(i item.Item) (*bytes.Buffer, error) {
	res, err := task.NewTask(func() (*bytes.Buffer, error) {
		// 1. gen item segment
		buf, err := t.parseItem(i)
		if err != nil {
			return &bytes.Buffer{}, err
		}
		return buf, nil
	}).Then(func(buf *bytes.Buffer) (*bytes.Buffer, error) {
		// 2. gen directory or command wrapper
		buf, err := t.withTest(i, buf)
		if err != nil {
			return buf, err
		}
		return buf, nil
	}).Then(func(buf *bytes.Buffer) (*bytes.Buffer, error) {
		// 3. gen general properties
		buf, err := t.withGeneralProps(i, buf)
		if err != nil {
			return buf, err
		}
		return buf, nil
	}).Result()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (t *ItemTranslator) createConditionBuffer(i item.Item) (*bytes.Buffer, error) {
	res, err := task.NewTask(func() (*bytes.Buffer, error) {
		// 1. gen item segment
		buf, err := t.parseCondition(i)
		if err != nil {
			return &bytes.Buffer{}, err
		}
		return buf, nil
	}).Then(func(buf *bytes.Buffer) (*bytes.Buffer, error) {
		// 2. gen general properties
		buf, err := t.withGeneralProps(i, buf)
		if err != nil {
			return buf, err
		}
		return buf, nil
	}).Result()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (t *ItemTranslator) parseItem(item item.Item) (*bytes.Buffer, error) {
	switch item.Type {
	case itype.ENV:
		return t.tpl.Env(item.Name, item.Value)
	case itype.ALIAS:
		return t.tpl.Alias(item.Name, item.Value)
	case itype.COMMENT:
		return t.tpl.Comment(item.Value)
	case itype.PATH:
		return t.tpl.PATH(item.Value)
	case itype.SOURCE:
		return t.tpl.Source(item.Value)
	case itype.RAW:
		return t.tpl.Raw(item.Value)
	default:
		return &bytes.Buffer{}, fmt.Errorf("unknown item type: %s", item.Type)
	}
}

func (t *ItemTranslator) parseItems(items []item.Item) ([]string, error) {
	results := []string{}
	for _, i := range items {
		if buf, err := t.createItemBuffer(i); err != nil {
			return results, fmt.Errorf("failed to parse item: %s", i.Type)
		} else {
			results = append(results, buf.String())
		}
	}
	return results, nil
}

func (t *ItemTranslator) parseCondition(item item.Item) (*bytes.Buffer, error) {
	if !item.Type.Variant().IsCondition() {
		return &bytes.Buffer{}, fmt.Errorf("invalid item type: %s", item.Type)
	}

	bodies, err := t.parseItems(item.Items)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	switch item.Type {
	case itype.IF_DIRECTORY:
		return t.tpl.If_Directory(item.Value, bodies...)
	case itype.IF_COMMAND:
		return t.tpl.If_Command(item.Value, bodies...)
	}
	return &bytes.Buffer{}, fmt.Errorf("unknown item type: %s", item.Type)
}

func (t *ItemTranslator) withGeneralProps(item item.Item, b *bytes.Buffer) (*bytes.Buffer, error) {
	res, err := task.NewTask(func() (*bytes.Buffer, error) {
		// 1. gen comment above the item
		buf, err := t.withTitle(item, b)
		if err != nil {
			return buf, err
		}
		return buf, nil
	}).Then(func(buf *bytes.Buffer) (*bytes.Buffer, error) {
		// 2. set line break below the item
		buf, err := t.withLF(item, buf)
		if err != nil {
			return buf, err
		}
		return buf, nil
	}).Result()
	if err != nil {
		return res, err
	}

	return res, nil
}

func (t *ItemTranslator) withTitle(item item.Item, b *bytes.Buffer) (*bytes.Buffer, error) {
	if item.Title == "" {
		return b, nil
	}

	var result bytes.Buffer
	comment, err := t.tpl.Comment(item.Title)
	if err != nil {
		return b, err
	}

	// 1. Add one line of white space at the top.
	result.WriteString("\n")
	// 2. Add a comment.
	result.Write(comment.Bytes())
	result.WriteString("\n")
	// 3. Set the body text.
	result.Write(b.Bytes())
	return &result, nil
}

func (t *ItemTranslator) withTest(item item.Item, b *bytes.Buffer) (*bytes.Buffer, error) {
	if item.Command != "" {
		return t.withTestCommand(item, b)
	} else if item.Directory != "" {
		return t.withTestDirectory(item, b)
	}
	return b, nil
}

func (t *ItemTranslator) withTestCommand(item item.Item, b *bytes.Buffer) (*bytes.Buffer, error) {
	if item.Command == "" {
		return b, nil
	}
	return t.tpl.Command(item.Command, b.String())
}

func (t *ItemTranslator) withTestDirectory(item item.Item, b *bytes.Buffer) (*bytes.Buffer, error) {
	if item.Directory == "" {
		return b, nil
	}
	return t.tpl.Directory(item.Directory, b.String())
}

func (t *ItemTranslator) withLF(item item.Item, b *bytes.Buffer) (*bytes.Buffer, error) {
	if item.LF < 1 {
		return b, nil
	}
	b.WriteString(strings.Repeat("\n", item.LF))
	return b, nil
}
