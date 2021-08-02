package snowflake

import (
	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

func Init() (err error) {
	node, err = sf.NewNode(1)
	if err != nil {
		return
	}
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
