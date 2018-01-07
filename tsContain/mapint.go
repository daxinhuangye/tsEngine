package tsContain

import (
	"errors"
)


type MapIntMgr struct {
	ObjMgr map[int64]interface{}
	count int
}

func NewMapIntMgr()(*MapIntMgr) {
	res := new(MapIntMgr)
	res.ObjMgr = make(map[int64]interface{})
	
	return res
}

func (this *MapIntMgr)Add(key int64, obj interface{})(error) {
	_, ok := this.ObjMgr[key]
	if ok {
		return errors.New("have obj")
	}
	this.ObjMgr[key] = obj
	this.count++
	return nil
}

func (this *MapIntMgr)Get(key int64)(interface{}) {
	obj, ok := this.ObjMgr[key]
	if !ok {
		return nil
	}
	if obj==nil {
		return nil
	}
	return obj
}

func (this *MapIntMgr)Have(key int64)(bool) {
	_, ok := this.ObjMgr[key]
	return ok
}

func (this *MapIntMgr)Del(key int64)() {
	if this.Have(key) {
		delete(this.ObjMgr, key)
		this.count--
	}
}

func (this *MapIntMgr)Count()(count int) {
	return this.count
}

func (this *MapIntMgr)ToArray()(res []interface{}) {
	res = make([]interface{}, 0)
	for _,v:=range this.ObjMgr {
		res = append(res, v)
	}
	return
}

func (this *MapIntMgr)Clear()() {
	for _,v:=range this.ObjMgr {
		if v!=nil {
			v = nil
		}
	}
	this.ObjMgr = make(map[int64]interface{})
	this.count = 0
}