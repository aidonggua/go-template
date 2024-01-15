package domain

import "time"

type Base struct {
	Id          int64     `json:"id"`
	GmtCreate   time.Time `json:"gmt_create"`   // 创建时间
	GmtModified time.Time `json:"gmt_modified"` // 更新时间
	Deleted     int32     `json:"deleted"`      // 0:未删除 1:已删除
}
