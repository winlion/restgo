package models

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id       int       `orm:"column(id);auto"`
	Account  string    `orm:"column(account);size(30);null"`
	Passwd   string    `orm:"column(passwd);size(40);null"`
	Mobile   string    `orm:"column(mobile);size(15);null"`
	Ticket   string    `orm:"column(ticket);size(60);null"`
	Createat time.Time `orm:"column(createat);type(datetime)"`
	Nickname string    `orm:"column(nickname);size(50);null"`
	Email    string    `orm:"column(email);size(40);null"`
	Avatar   string    `orm:"column(avatar);size(200);null"`
	Salt     string    `orm:"column(salt);size(10);null"`
	Sex      string    `orm:"column(sex);size(4);null"`
	Visible  uint      `orm:"column(visible);null"`
	Roleid   int       `orm:"column(roleid);null"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

func (u *User) Register(account string, passwd string) (userId int64, err error) {
	u.Account = account
	u.Salt = strconv.FormatInt(rand.Int63n(899999)+100000, 10)
	u.Passwd = passwd + u.Salt
	h := md5.New()
	h.Write([]byte(u.Passwd))
	u.Passwd = hex.EncodeToString(h.Sum(nil))
	o := orm.NewOrm()
	user := o.Read(&User{Account: account})
	if user != nil {
		err = errors.New("该手机号已经存在")

	} else {
		userId, err = o.Insert(u)
	}
	return

}

func (u *User) RegisterWithAvatar(account string, passwd string, avatar string) (userId int64, err error) {
	u.Account = account
	u.Salt = strconv.FormatInt(rand.Int63n(899999)+100000, 10)
	u.Passwd = passwd + u.Salt
	h := md5.New()
	h.Write([]byte(u.Passwd))
	u.Passwd = hex.EncodeToString(h.Sum(nil))
	u.Avatar = avatar

	var users []User
	o := orm.NewOrm()
	sql := fmt.Sprintf("SELECT id FROM user WHERE account = %s", account)
	num, err := o.Raw(sql).QueryRows(&users)

	if num > 0 {
		err = errors.New("该手机号已经存在")

	} else if err != nil {
		//
	} else {
		userId, err = o.Insert(u)
	}
	return

}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []User
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
