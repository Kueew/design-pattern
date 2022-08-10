/*
 * 在简单工厂中，可能存在很多 if 分之逻辑，如果我们非得将 if 分之逻辑去掉，比较经典的方法就是利用多态。
 * 实际上，这也是工厂方法模式的典型代码实现。
 * 工厂方法实际就是为工厂类再创建一个简单工厂，也就是工厂的工厂，用来创建工厂类对象。
 */

package factory

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

type IRuleConfigParserFactory interface {
	CreateParse() IRuleConfigParser
}

type jsonRuleConfigParserFactory struct {
}

func (f *jsonRuleConfigParserFactory) CreateParse() IRuleConfigParser {
	return &jsonRuleConfigParser{}
}

type xmlRuleConfigParserFactory struct {
}

func (f *xmlRuleConfigParserFactory) CreateParse() IRuleConfigParser {
	return &xmlRuleConfigParser{}
}

//创建规则配置解析器工厂
func NewRuleConfigParserFactory(configFormat string) IRuleConfigParserFactory {
	switch configFormat {
	case "json":
		return &jsonRuleConfigParserFactory{}
	case "xml":
		return &xmlRuleConfigParserFactory{}
	default:
		return nil
	}
}
