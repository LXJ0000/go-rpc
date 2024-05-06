package domain

import "gorm.io/gorm"

type Relation struct {
	gorm.Model
	RelationID int64 `gorm:"primaryKey;autoIncrement"`
	Followee   int64 `gorm:"not null;uniqueIndex:idx_followee_follower"`
	Follower   int64 `gorm:"not null;uniqueIndex:idx_followee_follower;index:idx_follower"`
	// 典型场景：某个人关注列表follower 某个人的粉丝列表followee 我都要
	Status bool // true 关注 false 取消关注
}

type RelationStat struct {
	gorm.Model
	UserID    int64 `gorm:"unique"`
	Follower  int
	Following int
}
