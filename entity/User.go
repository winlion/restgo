package entity

import (
	"time"
)

type User struct {
	ID int64 `xorm:"pk autoincr 'id'" json:"id"`
	CreateAt time.Time `xorm:"created" json:"create_at"  time_format:"2006-01-02 15:04:05"`
	Stat int `json:"stat"`
	UserName string `xorm:"varchar(40)" json:"user_name"`
	Passwd string `xorm:"varchar(40)" json:"-"`
	NickName string `xorm:"varchar(40)" json:"nick_name"`
	Avatar string `xorm:"varchar(180)" json:"avatar"`
}

