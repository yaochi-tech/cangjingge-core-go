package repository

import (
	"errors"
	"github.com/yaochi-tech/cangjingge-core-go/model"
	"github.com/yaochi-tech/lingquan-core-go/db"
)

// 与数据库操作相关的方法

// DatabaseInsert 插入数据
// @param Engine *db.Engine 灵泉引擎初始化后得到的对象
// @param SchemaName string 模型模式名称，该名称必须为在灵泉引擎中注册过的模型模式名称
// 使用方法：<br>
// 1. 初始化一个DatabaseInsert对象<br>
// 2. 调用Input/InputParams方法设置参数<br>
// 3. 调用Do方法执行<br>
// 4. 调用Output方法获取结果: effected - 影响的行数
type DatabaseInsert struct {
	params map[string]interface{}
	result map[string]interface{}
}

func (f *DatabaseInsert) InputParams() []model.FunctionParam {
	return []model.FunctionParam{}

}
func (f *DatabaseInsert) Input(params map[string]interface{}) error {
	if params == nil {
		return errors.New("insert params is nil")
	}
	if params[model.EngineParamName] == nil {
		return errors.New("engine is nil")
	}
	if params[model.SchemaName] == nil {
		return errors.New("schema name is nil")
	}
	if params[model.Params] == nil {
		return errors.New("params is nil")
	}
	f.params = params
	return nil
}

func (f *DatabaseInsert) Output() map[string]interface{} {
	return f.result
}
func (f *DatabaseInsert) Do() error {
	if f.params[model.EngineParamName] == nil {
		return errors.New("lingquan engine is nil")
	}
	if f.params[model.SchemaName] == nil {
		return errors.New("schema name is nil")
	}
	if f.params[model.Params] == nil {
		return errors.New("database insert params is nil")
	}
	f.result = make(map[string]interface{})

	engine, ok := f.params[model.EngineParamName].(*db.Engine)
	if !ok {
		return errors.New("engine is not Lingquan Engine instance")
	}
	schemaName, ok := f.params[model.SchemaName].(string)
	if !ok {
		return errors.New("schema name is not string")
	}

	effected, err := engine.Insert(schemaName, f.params)
	if err != nil {
		return err
	}
	f.result["effected"] = effected
	return nil
}

func init() {
	_ = model.RegisterFunction("system.database.insert", &DatabaseInsert{})
}
