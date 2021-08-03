package models

import (
	"time"
)

type Community struct {
	CommunityID   int64  `db:"community_id" json:"community_id"`     // 社区 ID
	CommunityName string `db:"community_name" json:"community_name"` // 社区名称
	Introduction  string `db:"introduction" json:"introduction"`     // 社区介绍
	// 这里是 time.Time 类型，数据库保存的为时间戳类型，如果需要转换，在连接数据库的时候加上 parseTime 参数
	CreateTime time.Time `db:"create_time" json:"create_time"` // 创建时间
	UpdateTime time.Time `db:"update_time" json:"update_time"` // 更新时间
}
