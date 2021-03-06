// Code generated by "mapstructure-to-hcl2 -type Config"; DO NOT EDIT.
package exoscaleimport

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName         *string           `mapstructure:"packer_build_name" cty:"packer_build_name"`
	PackerBuilderType       *string           `mapstructure:"packer_builder_type" cty:"packer_builder_type"`
	PackerDebug             *bool             `mapstructure:"packer_debug" cty:"packer_debug"`
	PackerForce             *bool             `mapstructure:"packer_force" cty:"packer_force"`
	PackerOnError           *string           `mapstructure:"packer_on_error" cty:"packer_on_error"`
	PackerUserVars          map[string]string `mapstructure:"packer_user_variables" cty:"packer_user_variables"`
	PackerSensitiveVars     []string          `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables"`
	SkipClean               *bool             `mapstructure:"skip_clean" cty:"skip_clean"`
	SOSEndpoint             *string           `mapstructure:"sos_endpoint" cty:"sos_endpoint"`
	APIEndpoint             *string           `mapstructure:"api_endpoint" cty:"api_endpoint"`
	APIKey                  *string           `mapstructure:"api_key" cty:"api_key"`
	APISecret               *string           `mapstructure:"api_secret" cty:"api_secret"`
	ImageBucket             *string           `mapstructure:"image_bucket" cty:"image_bucket"`
	TemplateZone            *string           `mapstructure:"template_zone" cty:"template_zone"`
	TemplateName            *string           `mapstructure:"template_name" cty:"template_name"`
	TemplateDescription     *string           `mapstructure:"template_description" cty:"template_description"`
	TemplateUsername        *string           `mapstructure:"template_username" cty:"template_username"`
	TemplateDisablePassword *bool             `mapstructure:"template_disable_password" cty:"template_disable_password"`
	TemplateDisableSSHKey   *bool             `mapstructure:"template_disable_sshkey" cty:"template_disable_sshkey"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.BlockAttrsSpec{TypeName: "packer_user_variables", ElementType: cty.String, Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"skip_clean":                 &hcldec.AttrSpec{Name: "skip_clean", Type: cty.Bool, Required: false},
		"sos_endpoint":               &hcldec.AttrSpec{Name: "sos_endpoint", Type: cty.String, Required: false},
		"api_endpoint":               &hcldec.AttrSpec{Name: "api_endpoint", Type: cty.String, Required: false},
		"api_key":                    &hcldec.AttrSpec{Name: "api_key", Type: cty.String, Required: false},
		"api_secret":                 &hcldec.AttrSpec{Name: "api_secret", Type: cty.String, Required: false},
		"image_bucket":               &hcldec.AttrSpec{Name: "image_bucket", Type: cty.String, Required: false},
		"template_zone":              &hcldec.AttrSpec{Name: "template_zone", Type: cty.String, Required: false},
		"template_name":              &hcldec.AttrSpec{Name: "template_name", Type: cty.String, Required: false},
		"template_description":       &hcldec.AttrSpec{Name: "template_description", Type: cty.String, Required: false},
		"template_username":          &hcldec.AttrSpec{Name: "template_username", Type: cty.String, Required: false},
		"template_disable_password":  &hcldec.AttrSpec{Name: "template_disable_password", Type: cty.Bool, Required: false},
		"template_disable_sshkey":    &hcldec.AttrSpec{Name: "template_disable_sshkey", Type: cty.Bool, Required: false},
	}
	return s
}
