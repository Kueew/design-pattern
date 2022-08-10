package factory

import (
	"reflect"
	"testing"
)

func TestNewConfigFactory(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name         string
		args         args
		target       IConfigParserFactory
		ruleParser   IRuleConfigParser
		systemParser ISystemConfigParser
	}{
		{
			name:         "json",
			args:         args{t: "json"},
			target:       &jsonConfigParserFactory{},
			ruleParser:   &jsonRuleConfigParser{},
			systemParser: &jsonSystemConfigParser{},
		},
		{
			name:         "xml",
			args:         args{t: "xml"},
			target:       &xmlConfigParserFactory{},
			ruleParser:   &xmlRuleConfigParser{},
			systemParser: &xmlSystemConfigParser{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if factory := NewConfigParserFactory(tt.args.t); reflect.DeepEqual(&factory, &tt.target) {
				if ruleParser := factory.CreateRuleParser(); !reflect.DeepEqual(&ruleParser, &tt.ruleParser) {
					t.Errorf("CreateRuleParser() = %v ,target %v", ruleParser, tt.ruleParser)
				}
				if systemParser := factory.CreateSystemParser(); !reflect.DeepEqual(&systemParser, &tt.systemParser) {
					t.Errorf("CreateSystemParser() = %v ,target %v", systemParser, tt.systemParser)
				}
			} else {
				t.Errorf("NewConfigParserFactory() = %v ,target %v", factory, tt.target)
			}
		})
	}
}
