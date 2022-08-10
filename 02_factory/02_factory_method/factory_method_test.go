package factory

import (
	"reflect"
	"testing"
)

func TestNewRuleConfigParserFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name   string
		args   args
		target IRuleConfigParserFactory
	}{
		{
			name:   "json",
			args:   args{t: "json"},
			target: &jsonRuleConfigParserFactory{},
		},
		{
			name:   "xml",
			args:   args{t: "xml"},
			target: &xmlRuleConfigParserFactory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if parser := NewRuleConfigParserFactory(tt.args.t); !reflect.DeepEqual(&parser, &tt.target) {
				t.Errorf("NewRuleConfigParserFactory() = %v ,target %v", parser, tt.target)
			}
		})
	}

}
