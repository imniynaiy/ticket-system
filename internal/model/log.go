package model

import "time"

type Log struct {
	LogID      uint      `json:"log_id" gorm:"column:log_id;primaryKey;autoIncrement" comment:"操作ログ番号"`
	UserID     uint      `json:"user_id" gorm:"column:user_id" comment:"ユーザID"`
	FunctionID uint      `json:"function_id" gorm:"column:function_id" comment:"機能番号"`
	Timestamp  time.Time `json:"timestamp" gorm:"column:timestamp" comment:"操作日時"`
	Result     int       `json:"result" gorm:"column:result" comment:"操作結果　０＝成功　１＝失敗"`
}

func (Log) TableName() string {
	return "log"
}
