package snowflake

import (
	"go.uber.org/zap"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init() (err error) {
	node, err = sf.NewNode(1)
	if err != nil {
		zap.L().Error("init snowflake config error", zap.Error(err))
		return
	}
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
