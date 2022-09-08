package snowflake

import (
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

var node *sf.Node

// Init 初始化
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	st, err = time.Parse("2006-01-02", startTime)
	if err != nil {
		zap.L().Error("Init snowflake failed")
		fmt.Println(err)
		return
	}
	// set time to start time
	sf.Epoch = st.UnixNano() / 1000000
	// new  Node
	node, err = sf.NewNode(machineID)
	return
}

// GenID get snowID
func GenID() int64 {
	return node.Generate().Int64()
}
