package models

import "time"

type Post struct {
	PostID      int64     `db:"post_id" json:"post_id"`           // 帖子 ID
	Title       string    `db:"title" json:"title"`               // 帖子名称
	Content     string    `db:"content" json:"content"`           // 帖子内容
	UserId      int64     `db:"user_id" json:"user_id"`           // 用户ID
	CommunityId int64     `db:"community_id" json:"community_id"` // 社区ID
	Status      uint8     `db:"status" json:"status"`             // 帖子状态(默认1)
	CreateTime  time.Time `db:"create_time" json:"create_time"`   // 创建时间
	UpdateTime  time.Time `db:"update_time" json:"update_time"`   // 更新时间
}
