package config

import "github.com/Alainyan1/KisFlow/common"

// KisFlowFunctionParam 一个Flow配置中的ID以及携带固定配置参数
type KisFlowFunctionParam struct {
	FuncName string `yaml:"fname"`  // 必填
	Params   FParam `yaml:"params"` //选填，在当前flow中function定制固定配置参数
}

// KisFlowConfig 用户贯穿整条流式计算上下文环境的对象
type KisFlowConfig struct {
	KisType  string                 `yaml:"kistype"`
	Status   int                    `yaml:"status"`
	FLowName string                 `yaml:"flowname"`
	Flows    []KisFlowFunctionParam `yaml:"flows"`
}

// NewFlowConfig 创建一个新的Flow策略配置对象，用于描述一个KisFlow信息
func NewFlowConfig(flowName string, enable common.KisOnOff) *KisFlowConfig {
	config := new(KisFlowConfig)
	config.FLowName = flowName
	config.Flows = make([]KisFlowFunctionParam, 0)
	config.Status = int(enable)
	return config
}

// AppendFunctionParam 向Flow配置中添加一个Function配置到当前Flow中
func (fConfig *KisFlowConfig) AppendFunctionParam(params KisFlowFunctionParam) {
	fConfig.Flows = append(fConfig.Flows, params)
}
