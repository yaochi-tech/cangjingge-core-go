package model

import "fmt"

// Functions 用于存储所有的方法函数，注册于此
var functions = make(map[string]Function)

func RegisterFunction(name string, function Function) error {
	if _, ok := functions[name]; ok {
		return fmt.Errorf("function %s is already registered", name)
	}
	functions[name] = function
	return nil
}

type FunctionParam struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Required    bool   `json:"required,omitempty"`
	Type        string `json:"type,omitempty"`
}

// Function 方法接口定义
type Function interface {
	//InputParams 输入参数定义，仅对params部分定义即可
	InputParams() []FunctionParam
	Input(map[string]interface{}) error
	Output() map[string]interface{}
	Do() error
}

const (
	//EngineParamName 灵泉引擎实例参数名
	EngineParamName = "lingquan.engine"
	//SchemaName 模型模式名称参数名
	SchemaName = "schema_name"
	//Params 参数参数名
	Params = "params"
)
