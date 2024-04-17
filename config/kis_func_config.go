package config

import (
	"github.com/Alainyan1/KisFLow/common"
	"github.com/Alainyan1/KisFLow/log"
)

// FParam在当前Flow中Fucntion定制固定配置参数类型
type FPrarm map[string]string

// KisSource表示当前Function的业务源
type KisSource struct {
	Name string   `yaml:"name"` // 本层function但数据源描述
	Must []string `yaml:"must"` // source中必须包含的字段
}

// KisFuncOption可选配置
type KisFuncOption struct {
	CName         string `yaml:"cname"`          // 连接器Connector的名称
	RetryTimes    int    `yaml:"retry_times"`    // 可选，Function调度重试（不包括正常调度）最大次数
	RetryDuration int    `yaml:"retry_duration"` // 可选，Function每次重试最大间隔时间（ms）
	Params        FPrarm `yaml:"params"`         // 可选，在当前Flow中Function定制固定配置参数
}

// KisFuncConfig 一个KisFunction的策略配置
type KisFuncConfig struct {
	KisType string        `yaml:"kis_type"`
	Fname   string        `yaml:"fname"`
	FMode   string        `yaml:"fmode"`
	Source  KisSource     `yaml:"source"`
	Option  KisFuncOption `yaml:"option"`
}

// 构造方法, 创建一个Function的策略配置对象，用于描述一个KisFunction信息
func NewFuncConfig(funcName string, mode common.KisMode, source *KisSource, option *KisFuncOption) *KisFuncConfig {
	config := new(KisFuncConfig)
	config.Fname = funcName

	if source == nil {
		log.Logger().ErrorF("funcName NewConfig Error, source is nil, funcName is %s\n", funcName)
		return nil
	}

	config.Source = *source

	config.FMode = string(mode)

	// Function S/L必须需要传入KisConnector参数，因为S/L需要通过Connector进行建立流式关系
	// S/L说明这个Function有查询或者存储数据的行为，需要通过Connector进行数据的流式传递, 那么CName是必须的
	if mode == common.S || mode == common.L {
		if option == nil {
			log.Logger().ErrorF("Function S/L need option -> Cid\n")
			return nil
		} else if option.CName == "" {
			log.Logger().ErrorF("Function S/L need option -> Cid\n")
			return nil
		}
	}

	if option != nil {
		config.Option = *option
	}

	return config
}
