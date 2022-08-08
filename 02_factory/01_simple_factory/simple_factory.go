package facctory

//需求：根据配置文件的后缀（json、xml、yaml、properties）
//选择不同的解析器（JsonRuleConfigParser、XmlRuleConfigParser……）
//将存储在文件中的配置解析成内存对象 RuleConfig

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

type yamlRuleConfigParser struct {
}

func (parser *yamlRuleConfigParser) Parser(data []byte) RuleConfig {
	panic("NotImplementedException")
}

//创建规则配置解析器
func NewRuleConfigParser(configFormat string) IRuleConfigParser {
	switch configFormat {
	case "json":
		return &jsonRuleConfigParser{}
	case "xml":
		return &xmlRuleConfigParser{}
	case "yaml":
		return &yamlRuleConfigParser{}
	default:
		return nil
	}
}
