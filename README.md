# cangjingge-core-go
瑶池-藏经阁-逻辑核心库go语言版本

# 藏经阁逻辑核心库
本库提供基础的逻辑处理能力，通过定义的逻辑json数据，执行对应的逻辑流程及逻辑方法

## 方法
方法是系统提供的最小逻辑单元，每一个方法都对应有注册的代码，可以是系统内置的，也可以是采用动态脚本编写注册到本系统中的

可用的方法会在后续的文档中详细说明

## 逻辑
逻辑是由多个方法组成的逻辑流程，逻辑中应包含方法或子逻辑

## 中间件
中间件是对逻辑的补充，可以在逻辑执行前后执行，主要针对触发点如http请求、数据库语句执行等
