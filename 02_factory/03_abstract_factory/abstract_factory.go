package factory

/*
 * 需求：在简单工厂例子中，规则配置有多种配置方式，在此基础上新增配置有分类（规则配置、系统配置），每种分类有多种配置方式。
 * 针对这种需求我们可以采用抽象工厂，让一个工厂负责创建多个不同类型的对象（IRuleConfigParser、ISystemConfigParser 等）
 */

//规则配置对象
type RuleConfig struct {
}

//规则配置解析接口
type IRuleConfigParser interface {
	Parser(data []byte) RuleConfig
}

type jsonRuleConfigParser struct {
}

func (parser *jsonRuleConfigParser) Parser(data []byte) RuleConfig {
	panic("NotImplementedException")
}

type xmlRuleConfigParser struct {
}

func (parser *xmlRuleConfigParser) Parser(data []byte) RuleConfig {
	panic("NotImplementedException")
}

type SystemConfig struct {
}

type ISystemConfigParser interface {
	Parser(data []byte) SystemConfig
}

type jsonSystemConfigParser struct {
}

func (parser *jsonSystemConfigParser) Parser(data []byte) SystemConfig {
	panic("NotImplementedException")
}

type xmlSystemConfigParser struct {
}

func (parser *xmlSystemConfigParser) Parser(data []byte) SystemConfig {
	panic("NotImplementedException")
}

type IConfigParserFactory interface {
	CreateRuleParser() IRuleConfigParser
	CreateSystemParser() ISystemConfigParser
}

type jsonConfigParserFactory struct {
}

func (f *jsonConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return &jsonRuleConfigParser{}
}

func (f *jsonConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return &jsonSystemConfigParser{}
}

type xmlConfigParserFactory struct {
}

func (f *xmlConfigParserFactory) CreateRuleParser() IRuleConfigParser {
	return &xmlRuleConfigParser{}
}

func (f *xmlConfigParserFactory) CreateSystemParser() ISystemConfigParser {
	return &xmlSystemConfigParser{}
}

func NewConfigParserFactory(configFormat string) IConfigParserFactory {
	switch configFormat {
	case "json":
		return &jsonConfigParserFactory{}
	case "xml":
		return &xmlConfigParserFactory{}
	default:
		return nil
	}
}
