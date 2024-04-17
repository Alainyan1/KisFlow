package config

import (
	"errors"
	"fmt"

	common "github.com/Alainyan1/KisFlow/common"
)

type KisConnConfig struct {
	KisType    string             `yaml:"kistype"`
	CName      string             `yaml:"cname"`
	AddrString string             `yaml:"addr"`
	Type       common.KisConnType `yaml:"type"`
	Key        string             `yaml:"key"`
	Params     map[string]string  `yaml:"params"`
	Load       []string           `yaml:"load"`
	Save       []string           `yaml:"save"`
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
