package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Attach struct {
	Id        int       `orm:"column(id);auto"`
	AlbumId   int       `orm:"column(album_id)"`
	UserId    int       `orm:"column(user_id)"`
	Url       string    `orm:"column(url);size(255)"`
	FileSize  int       `orm:"column(file_size)"`
	Deleted   int       `orm:"column(deleted)"`
	ProjectId int       `orm:"column(projectId);null"`
	NodeId    int       `orm:"column(nodeId);null"`
	Memo      string    `orm:"column(memo);size(30);null"`
	FileId    string    `orm:"column(fileId);size(60);null"`
	CreateAt  time.Time `orm:"column(createAt);type(datetime);null"`
}

func (t *Attach) TableName() string {
	return "attach"
}

func init() {
	orm.RegisterModel(new(Attach))
}

// AddAttach insert a new Attach into database and returns
// last inserted Id on success.
func AddAttach(m *Attach) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetAttachById retrieves Attach by Id. Returns error if
// Id doesn't exist
func GetAttachById(id int) (v *Attach, err error) {
	o := orm.NewOrm()
	v = &Attach{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAttach retrieves all Attach matches certain condition. Returns empty list if
// no records exist
func GetAllAttach(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Attach))
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

	var l []Attach
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

// UpdateAttach updates Attach by Id and returns error if
// the record to be updated doesn't exist
func UpdateAttachById(m *Attach) (err error) {
	o := orm.NewOrm()
	v := Attach{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAttach deletes Attach by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAttach(id int) (err error) {
	o := orm.NewOrm()
	v := Attach{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Attach{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
