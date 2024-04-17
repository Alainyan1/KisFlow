package config

import (
	"errors"
	"fmt"

	common "github.com/Alainyan1/KisFlow/common"
)

type KisConnConfig struct {
	// 配置类型
	KisType string `yaml:"kistype"`
	// 连接器名称，唯一描述标识
	CName string `yaml:"cname"`
	// 基础存储媒介地址
	AddrString string `yaml:"addr"`
	// 存储媒介类型
	Type common.KisConnType `yaml:"type"`
	// 一次存储的标识， 如redis为key的名称，mysql为table的名称
	Key string `yaml:"key"`
	// 配置信息中的自定义参数
	Params map[string]string `yaml:"params"`
	// 存储读取所绑定的NsFunction的ID
	Load []string `yaml:"load"`
	Save []string `yaml:"save"`
}

// NewConnConfig 创建一个新的Connectorcel配置对象，用于描述一个KisConnector信息
func NewConnConfig(cName string, addr string, t common.KisConnType, key string, param FParam) *KisConnConfig {
	config := new(KisConnConfig)
	config.CName = cName
	config.AddrString = addr
	config.Type = t
	config.Key = key
	config.Params = param

	return config
}

// WithFunc 将Connector与Function绑定
func (c *KisConnConfig) WithFunc(fConfig *KisFuncConfig) error {

	switch common.KisMode(fConfig.FMode) {
	case common.S:
		c.Save = append(c.Save, fConfig.FName)
	case common.L:
		c.Load = append(c.Load, fConfig.FName)
	default:
		return errors.New(fmt.Sprintf("Invalid KisMode %s", fConfig.FMode))
	}

	return nil
}
