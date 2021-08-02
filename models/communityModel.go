package models

type Community struct {
	CommunityID   int64  `db:"community_id" json:"community_id"`  // 社区 ID
	CommunityName string `db:"community_name" json:"community_name"`  // 社区名称
	Introduction string `db:"introduction" json:"introduction"`  // 社区介绍
}
