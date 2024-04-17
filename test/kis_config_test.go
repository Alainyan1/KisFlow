package test

import (
	"testing"

	common "github.com/Alainyan1/KisFlow/common"
	config "github.com/Alainyan1/KisFlow/config"
	log "github.com/Alainyan1/KisFlow/log"
)

func TestNewFuncConfig(t *testing.T) {
	source := config.KisSource{
		Name: "test1",
		Must: []string{"order_id", "user_id"},
	}
	option := config.KisFuncOption{
		CName:         "Connector1",
		RetryTimes:    3,
		RetryDuration: 300,
		Params: config.FParam{
			"Param1": "value1",
			"Param2": "value2",
		},
	}

	myFunc1 := config.NewFuncConfig("myFunc1", common.S, &source, &option)

	log.Logger().InfoF("myFunc1: %+v\n", myFunc1)
}

func TestNewFlowConfig(t *testing.T) {
	flowFuncParams1 := config.KisFlowFunctionParam{
		FuncName: "myFlow1",
		Params: config.FParam{
			"flowParam1": "value1",
			"flowParam2": "value2",
		},
	}

	flowFuncParams2 := config.KisFlowFunctionParam{
		FuncName: "myFlow2",
		Params: config.FParam{
			"default": "value1",
		},
	}

	myFlow1 := config.NewFlowConfig("flow1", common.FlowEnable)
	myFlow1.AppendFunctionParam(flowFuncParams1)
	myFlow1.AppendFunctionParam(flowFuncParams2)

	log.Logger().InfoF("myFlow1: %+v\n", myFlow1)
}

func TestNewConnConfig(t *testing.T) {

	source := config.KisSource{
		Name: "test1",
		Must: []string{"order_id", "user_id"},
	}
	option := config.KisFuncOption{
		CName:         "Connector1",
		RetryTimes:    3,
		RetryDuration: 300,
		Params: config.FParam{
			"Param1": "value1",
			"Param2": "value2",
		},
	}

	myFunc1 := config.NewFuncConfig("myFunc1", common.S, &source, &option)

	connParams := config.FParam{
		"connParam1": "value1",
		"connParam2": "value2",
	}

	myConn1 := config.NewConnConfig("Connector1", "0.0.0.0:9987, 0.0.0.0:9997", common.REDIS, "key", connParams)

	if err := myConn1.WithFunc(myFunc1); err != nil {
		log.Logger().ErrorF("WithFunc error: %s\n", err.Error())
	}

	log.Logger().InfoF("myConn1: %+v\n", myConn1)
}
