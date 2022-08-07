package singleton

import goutils "github.com/typa01/go-utils"

//单例模式：
//饿汉模式
//线程安全
//不支持延迟加载，不适合构造占有资源较多或者初始化耗时较长的实力提前初始化
//Id 生成器
type IIdGenerator interface {
	Create() string
}
type guidGenerator struct {
}

var instance *guidGenerator

func init() {
	instance = &guidGenerator{}
}

//获取实例
func GetInstance() *guidGenerator {
	return instance
}

//创建Id
func (instance *guidGenerator) Create() string {
	return goutils.GUID()
}
