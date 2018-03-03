package service

import (
	"restgo/entity"
	"restgo/restgo"
	"restgo/model"
)

type UserService struct {}
//根据userId 获取用户编号
func (service *UserService)FindOne(userId int64)(entity.User){
	var user entity.User
	orm := restgo.OrmEngin("ds1")
	orm.Id(userId).Get(&user)
	return  user
}

func (service *UserService)Query(arg model.UserArg)([]entity.User){
	var users []entity.User = make([]entity.User , 0)
	orm := restgo.OrmEngin("ds1")
	t := orm.Where("id>0")
	if (0<len(arg.Kword)){
		t = t.Where("name like ?","%"+arg.Kword+"%")
	}

	if (!arg.Datefrom.IsZero()){
		t = t.Where("create_at >= ?",arg.Datefrom)
	}
	if (!arg.Dateto.IsZero()){
		t = t.Where("create_at <= ?",arg.Dateto)
	}
	t.Limit(arg.GetPageFrom()).Find(&users)
	return  users
}