package tsDb

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"tsEngine/tsPagination"
)

type DbBase struct {
	db            orm.Ormer
	isTransaction bool
	isRollback    bool
}

func NewDbBase() *DbBase {
	res := new(DbBase)
	res.db = orm.NewOrm()

	return res
}

func (this *DbBase) SetRollback(is_rollback bool) {
	this.isRollback = is_rollback
}

/*
//*******************例子1*************************
	db := tsDb.NewDbBase()
	db.Transaction()
	defer db.TransactionEnd()

	room :=new(models.Room)
	room.Name = "测试01"
	err = db.DbInsert(room)
	db.SetRollback(true)

//*******************例子2*************************
	db := tsDb.NewDbBase()
	db.Transaction()

	room :=new(models.Room)
	room.Name = "测试01"
	db.DbInsert(room)
	db.SetRollback(true)
	db.TransactionEnd()
*/
func (this *DbBase) Transaction() (err error) {
	err = this.db.Begin()
	if err == nil {
		this.isTransaction = true
	}
	return
}

func (this *DbBase) TransactionEnd() {
	if this.isTransaction {
		if this.isRollback {
			this.db.Rollback()
		} else {
			this.db.Commit()
		}
	}
	this.isTransaction = false
	this.isRollback = false
}

//获取单条记录
func (this *DbBase) DbGet(obj interface{}, fields ...string) (err error) {
	err = this.db.Read(obj, fields...)
	return
}

//获取单条记录
func (this *DbBase) DbRead(obj interface{}, fields ...string) (err error) {
	err = this.db.Read(obj, fields...)
	return
}

func (this *DbBase) DbListObj(obj interface{}, array interface{}, fields ...interface{}) (err error) {
	qt := this.db.QueryTable(obj)
	length := len(fields)
	count := length / 2

	if length%2 != 0 {
		err = errors.New("error fields count")
		return
	}

	for i := 0; i < count; i++ {
		qt = qt.Filter(fields[i*2+0].(string), fields[i*2+1])
	}

	qt = qt.Limit(-1)

	_, err = qt.All(array)

	return
}

func (this *DbBase) DbListObjOrder(obj interface{}, array interface{}, order []string, fields ...interface{}) (err error) {
	qt := this.db.QueryTable(obj)
	length := len(fields)
	count := length / 2

	if length%2 != 0 {
		err = errors.New("error fields count")
		return
	}

	for i := 0; i < count; i++ {
		qt = qt.Filter(fields[i*2+0].(string), fields[i*2+1])
	}

	qt = qt.OrderBy(order...)

	qt = qt.Limit(-1)

	_, err = qt.All(array)

	return
}

//获取记录数量
func (this *DbBase) DbCount(obj interface{}, filters ...interface{}) (count int64, err error) {

	length := len(filters)
	filters_count := length / 2

	if length%2 != 0 {
		err = errors.New("error fields count")
		return
	}

	qt := this.db.QueryTable(obj)

	// 过滤条件
	for i := 0; i < filters_count; i++ {
		qt = qt.Filter(filters[i*2+0].(string), filters[i*2+1])
	}

	count, err = qt.Count()

	return
}

//获取列表记录
func (this *DbBase) DbList(obj interface{}, filters ...interface{}) (data []orm.Params, err error) {

	qt := this.db.QueryTable(obj)
	length := len(filters)
	count := length / 2

	if length%2 != 0 {
		err = errors.New("error filters count")
		return
	}

	for i := 0; i < count; i++ {
		qt = qt.Filter(filters[i*2+0].(string), filters[i*2+1])
	}

	qt = qt.Limit(-1)

	_, err = qt.Values(&data)

	return data, err
}

//获取列表记录
func (this *DbBase) DbListOrder(obj interface{}, order []string, filter ...interface{}) (data []orm.Params, err error) {

	qt := this.db.QueryTable(obj)

	length := len(filter)

	count := length / 2

	if length%2 != 0 {
		err = errors.New("error filter count")
		return
	}

	for i := 0; i < count; i++ {
		qt = qt.Filter(filter[i*2+0].(string), filter[i*2+1])
	}

	qt = qt.OrderBy(order...)

	qt = qt.Limit(-1)

	_, err = qt.Values(&data)

	return data, err
}

//获取列表记录
func (this *DbBase) DbListFields(obj interface{}, fields []string, order []string, filter ...interface{}) (data []orm.Params, err error) {

	qt := this.db.QueryTable(obj)

	length := len(filter)

	count := length / 2

	if length%2 != 0 {
		err = errors.New("error filter count")
		return
	}

	for i := 0; i < count; i++ {
		qt = qt.Filter(filter[i*2+0].(string), filter[i*2+1])
	}

	qt = qt.OrderBy(order...)

	qt = qt.Limit(-1)

	_, err = qt.Values(&data, fields...)

	return data, err
}

//获取列表记录
func (this *DbBase) DbListOrderLimt(obj interface{}, limt int, order []string, filter ...interface{}) (data []orm.Params, err error) {

	qt := this.db.QueryTable(obj)
	length := len(filter)
	count := length / 2

	if length%2 != 0 {
		err = errors.New("error filter count")
		return
	}

	for i := 0; i < count; i++ {
		qt = qt.Filter(filter[i*2+0].(string), filter[i*2+1])
	}

	qt = qt.OrderBy(order...)

	qt = qt.Limit(limt)

	_, err = qt.Values(&data)

	return data, err
}

/*
 * @bried 分页读取内容
 *
 * @param obj 对象
 * @param page 第几页 >=1
 * @param page_size 页面大小
 * @param fields 读取的数据列
 * @param order 排序设置
 * @param filters 过滤条件
 */
func (this *DbBase) DbListPage(obj interface{},
	page int64, page_size int64,
	fields []string, order []string, filters ...interface{}) (data []orm.Params, pagination *tsPagination.Pagination, err error) {

	length := len(filters)
	filters_count := length / 2

	if length%2 != 0 {
		err = errors.New("error filters count")
		return
	}

	qt := this.db.QueryTable(obj)

	// 过滤条件
	for i := 0; i < filters_count; i++ {
		qt = qt.Filter(filters[i*2+0].(string), filters[i*2+1])
	}

	qt = qt.OrderBy(order...)

	count, err := qt.Count()
	if err != nil {
		return data, nil, err
	}
	pagination = tsPagination.NewPagination(page, page_size, count)

	qt = qt.Limit(page_size, pagination.GetOffset())

	qt.Values(&data, fields...)

	return data, pagination, err
}

//通过 in 的方式获取数据
func (this *DbBase) DbInIds(obj interface{}, field string, ids interface{}, fields ...string) (data []orm.Params, err error) {

	qt := this.db.QueryTable(obj)

	qt = qt.Filter(field+"__in", ids)

	qt = qt.Limit(-1)

	_, err = qt.Values(&data, fields...)

	return data, err
}

//单条数据插入
func (this *DbBase) DbInsert(obj interface{}) (id int64, err error) {
	id, err = this.db.Insert(obj)
	return
}

//多条数据插入
func (this *DbBase) DbInsertMulti(array interface{}, lenght int) error {

	if _, err := this.db.InsertMulti(lenght, array); err != nil {
		return err
	}
	return nil
}

//数据更新
func (this *DbBase) DbUpdate(obj interface{}, fields ...string) (err error) {
	_, err = this.db.Update(obj, fields...)
	return
}

//数据删除
func (this *DbBase) DbDel(obj interface{}, filters ...interface{}) (err error) {
	length := len(filters)

	if length <= 0 {
		_, err = this.db.Delete(obj)
	} else {
		count := length / 2

		if length%2 != 0 {
			err = errors.New("error filters count")
			return
		}

		qt := this.db.QueryTable(obj)
		for i := 0; i < count; i++ {
			qt = qt.Filter(filters[i*2+0].(string), filters[i*2+1])
		}
		_, err = qt.Delete()
	}

	return
}

//数据逻辑删除
func (this *DbBase) DbDelLogic(obj interface{}) (err error) {
	err = this.DbUpdate(obj, "IsDel")
	return
}
