package common

type KisMode string

const (
	// V为校验特征的KisFunction
	// 主要进行数据的过滤，验证，字段梳理，幂等等前置数据梳理
	V KisMode = "Verify"

	// S为存储特征的KisFunction
	// S会通过NsConnector将数据进行存储，数据的临时声明周期为NsWindow
	S KisMode = "Save"

	// L为加载特征的KisFunction
	// L会通过NsConnector将数据进行加载，通过该Function可以从逻辑上与对应的S Function进行并流
	L KisMode = "Load"

	// C为计算特征的KisFunction
	// C会通过KisFlow中的数据计算，生成新的字段，将数据流传递给下游S进行存储，或者自己也可以直接通过KisConnector进行存储
	C KisMode = "Calculate"

	// E为扩展特征的KisFunction
	// 作为流式计算的自定义特征Function，如Notify调度器触发任务的消息发送，删除一些数据，重置状态等
	E KisMode = "Expand"
)

type KisOnOff int // 0: off, 1: on

const (
	FlowEnable  KisOnOff = 1
	FlowDisable KisOnOff = 0
)

type KisConnType string // connector的类型

const (
	REDIS KisConnType = "redis"
	MYSQL KisConnType = "mysql"
	KAFKA KisConnType = "kafka"
	TIDB  KisConnType = "tidb"
	ES    KisConnType = "es"
)
