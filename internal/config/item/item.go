package item

import (
	"github.com/go-playground/validator/v10"
	"github.com/kawana77b/univenv/internal/common"
	"github.com/kawana77b/univenv/internal/config/item/itype"
	"github.com/kawana77b/univenv/internal/config/item/variant"
	"github.com/kawana77b/univenv/internal/sysutil/ostype"
	"github.com/kawana77b/univenv/internal/sysutil/shell"
)

var itemValidateFuncs = map[string]validator.Func{
	"item_type": func(fl validator.FieldLevel) bool {
		v := fl.Field().String()
		return itype.Contains(v)
	},
	"os": func(fl validator.FieldLevel) bool {
		v := fl.Field().String()
		return ostype.Contains(v)
	},
	"shell": func(fl validator.FieldLevel) bool {
		v := fl.Field().String()
		return shell.Contains(v)
	},
}

type Item struct {
	//-------------------------------
	// Basic
	//-------------------------------

	Type  itype.ItemType `yaml:"type" validate:"required,item_type"`
	Name  string         `yaml:"name,omitempty"`
	Value string         `yaml:"value" validate:"required"`
	Items []Item         `yaml:"items,omitempty"`

	//-------------------------------
	// Condition
	//-------------------------------

	// Add a branch that checks for the existence of the corresponding command
	Command string `yaml:"command,omitempty"`
	// Add a branch that checks for the existence of the corresponding directory
	Directory string `yaml:"directory,omitempty"`

	// Output scripts only if applicable to the OS on which they were executed
	OS []ostype.OS `yaml:"os,omitempty" validate:"dive,os"`
	// Output scripts only if applicable to the shell on which they were executed
	Shell []shell.Shell `yaml:"shell,omitempty" validate:"dive,shell"`

	//-------------------------------
	// User
	//-------------------------------

	// Inject a title comment into the script
	Title string `yaml:"title,omitempty"`
	// Ignore this item
	Disabled bool `yaml:"disabled,omitempty"`
	// Add a blank line break below this item
	LF int `yaml:"lf,omitempty" validate:"lt=20"`
}

var itemValidator *validator.Validate = func() *validator.Validate {
	vali := validator.New(validator.WithRequiredStructEnabled())
	for k, v := range itemValidateFuncs {
		vali.RegisterValidation(k, v)
	}
	return vali
}()

func (i Item) Validate() error {
	return i.validate()
}

func (i Item) validate() error {
	if err := itemValidator.Struct(i); err != nil {
		return err
	}
	for _, item := range i.Items {
		// NOTE: Assumes that items cannot be nested consecutively
		if item.Type.Variant() != variant.ITEM {
			return common.NewInvalidError("items")
		}
		if err := item.Validate(); err != nil {
			return err
		}
	}
	return nil
}
