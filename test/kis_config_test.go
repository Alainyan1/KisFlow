package test

import (
	"log"
	"testing"

	"github.com/Alainyan1/KisFlow/common"
	"github.com/Alainyan1/KisFlow/config"
	"github.com/Alainyan1/KisFlow/log"
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
		Params: config.FPrarm{
			"Param1": "value1",
			"Param2": "value2",
		},
	}

	myFunc1 := config.NewFuncConfig("myFunc1", common.S, &source, &option)

	log.Logger().InfoF("myFunc1: %+v\n", myFunc1)
}
