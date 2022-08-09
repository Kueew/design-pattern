package facctory

import (
	"reflect"
	"testing"
)

func TestNewRuleConfigParser(t *testing.T) {

	type args struct {
		t string
	}
	tests := []struct {
		name   string
		args   args
		target IRuleConfigParser
	}{
		{
			name:   "json",
			args:   args{t: "json"},
			target: &jsonRuleConfigParser{},
		},
		{
			name:   "xml",
			args:   args{t: "xml"},
			target: &xmlRuleConfigParser{},
		},
		{
			name:   "yaml",
			args:   args{t: "yaml"},
			target: &yamlRuleConfigParser{},
		},
		{
			name:   "property",
			args:   args{t: "property"},
			target: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if parser := NewRuleConfigParser(tt.args.t); !reflect.DeepEqual(&parser, &tt.target) {
				t.Errorf("NewIRuleConfigParser() = %v ,target %v", parser, tt.target)
			}
		})
	}

}
