package tsContain

import (
	"tsEngine/tsRand"
	"errors"
)

type SliceMgr struct {
	Datas []interface{}
	index int
}

func NewSliceMgr(l int64, c int64)(res *SliceMgr) {
	res = new(SliceMgr)
	res.Datas = make([]interface{}, l, c)
	return
}


func (this *SliceMgr)Count()(int) {
	return len(this.Datas)
}

func (this *SliceMgr)First()(int, interface{}) {
	if len(this.Datas)==0 {
		return -1,nil
	}
	this.index = 0
	return this.index, this.Datas[this.index]
}

func (this *SliceMgr)Next()(int, interface{}) {
	if this.index+1>=len(this.Datas) {
		return -1, nil
	}
	this.index++
	return this.index, this.Datas[this.index]
}

func (this *SliceMgr)Add(pos int, v interface{})(err error) {
	count := len(this.Datas)
	
	if pos==-1 || pos==count-1 {
		this.Datas = append(this.Datas, v)
		return
	}
	
	if pos==0 {
		t := []interface{}{v}
		this.Datas = append(t, this.Datas...)
		return
	}
	
	if pos>=count {
		err = errors.New("out range")
		return 
	}
	this.Datas = append(this.Datas[:pos], append([]interface{}{v}, this.Datas[pos:]...)...)
	
	return
}

func (this *SliceMgr)Del(pos int)(interface{}) {
	count := len(this.Datas)
	if pos>=count {
		return nil
	}
	v:=this.Datas[pos]
	if count==1 {
		this.Datas = make([]interface{}, 0, 0)
		return v
	}
	if pos==0 {
		this.Datas = this.Datas[pos+1:]
		return v
	}
	if pos==count-1 {
		this.Datas = this.Datas[:pos]
		return v
	}
	this.Datas = append(this.Datas[:pos], this.Datas[pos+1:]...)
	return v
}

func (this *SliceMgr)Set(pos int, v interface{})(err error) {
	if pos>=len(this.Datas) {
		err = errors.New("out range")
		return 
	}
	this.Datas[pos] = v
	return
}

func (this *SliceMgr)Get(pos int)(interface{}) {
	if pos>=len(this.Datas) {
		return nil
	}	
	return this.Datas[pos]
}

func (this *SliceMgr)Clear()() {
	count := len(this.Datas)
	for i:=0; i<count; i++ {
		this.Datas[i] = nil
	}
}

func (this *SliceMgr)Destroy()() {
	count := len(this.Datas)
	for i:=0; i<count; i++ {
		this.Datas[i] = nil
	}
	
	this.Datas = make([]interface{}, 0, 0)
}

func (this *SliceMgr)Change(pos1 int, pos2 int)(err error) {
	count := len(this.Datas)
	if pos1>=count || pos2>=count || pos1<0 || pos2<0 {
		err = errors.New("pos out range")
		return 
	}
	
	if pos1==pos2 {
		return
	}
	
	t:=this.Datas[pos1]
	this.Datas[pos1] = this.Datas[pos2]
	this.Datas[pos2] = t
	
	return
}

func (this *SliceMgr)NoNilInMiddle()() {
	count := len(this.Datas)
	nil_pos := -1
	for i:=0; i<count; i++ {
		if this.Datas[i]==nil {
			if nil_pos==-1 {
				nil_pos = i
			}
		} else {
			if nil_pos!=-1 {
				this.Datas[nil_pos] = this.Datas[i]
				this.Datas[i] = nil
				nil_pos++
			}
		}
	}
}

func (this *SliceMgr)NilPos(pos int)(interface{}) {
	count := len(this.Datas)
	if pos>=count {
		return nil
	}
	v:=this.Datas[pos]
	this.Datas[pos]=nil
	return v
}

func (this *SliceMgr)Rand()() {
	count := len(this.Datas)
	if count<=1 {
		return
	}
	for i:=0; i<count; i++ {
		pos := tsRand.RandInt(0, count-1)
		this.Change(i, pos)
	}
}

/*
 * @brief 数组元素后移，队尾元素将丢失
 *
 * @param pos 位置，该位置的值将等于nil
 */
func (this *SliceMgr)MoveBack(pos int, num int)() {
	count := len(this.Datas)
	if pos>=count {
		return 
	}
	if count==1 {
		return
	}
	if pos==count-1 {
		this.Datas[pos] = nil
		return
	}
	for i:=count-1-1; i>=pos; i-- {
		this.Datas[i+1] = this.Datas[i]
		this.Datas[i] = nil
	}
}

/*
 * @brief 元素添加到nil的位置
 *
 * @param v 元素
 */
func (this *SliceMgr)AddToNilPos(v interface{})(err error) {
	count := len(this.Datas)
	for i:=0; i<count; i++ {
		if this.Datas[i]==nil {
			this.Datas[i]=v
			return
		}
	}
	return errors.New("no nil pos")
}

func (this *SliceMgr)LastNoNilIndex()(int) {
	count := len(this.Datas)
	for i:=count-1; i>=0; i-- {
		if this.Datas[i]!=nil {
			return i
		}
	}
	return -1
}

func (this *SliceMgr)LastNoNil()(interface{}) {
	count := len(this.Datas)
	for i:=count-1; i>=0; i-- {
		if this.Datas[i]!=nil {
			return this.Datas[i]
		}
	}
	return nil
}

func (this *SliceMgr)NoNilCount()(int) {
	count := len(this.Datas)
	num := 0
	for i:=0; i<count; i++ {
		if this.Datas[i]!=nil {
			num++
		}
	}
	return num
}

func (this *SliceMgr)NoNilArray()([]interface{}) {
	count := this.NoNilCount()
	Datas := make([]interface{}, count, count)
	j := 0
	for i:=0; i<len(this.Datas); i++ {
		if this.Datas[i]!=nil {
			Datas[j] = this.Datas[i]
			j++
		}
	}
	return Datas
}