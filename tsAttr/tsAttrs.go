package tsAttr

import (
)

type Attrs struct {
	Field map[string]interface{}
}

func NewAttrs()(res *Attrs) {
	res = &Attrs{
		Field: make(map[string]interface{}),
		}
	return
}

func (this *Attrs)SetObj(name string, v interface{})(err error) {
	this.Field[name] = v
	return	
}

func (this *Attrs)SetInt32(name string, v int32)(err error) {
	this.Field[name] = v
	return
}

func (this *Attrs)SetInt64(name string, v int64)(err error) {
	this.Field[name] = v
	return
}

func (this *Attrs)SetUint32(name string, v uint32)(err error) {
	this.Field[name] = v
	return
}

func (this *Attrs)SetUint64(name string, v uint64)(err error) {
	this.Field[name] = v
	return
}

func (this *Attrs)SetFloat64(name string, v float64)(err error) {
	this.Field[name] = v
	return
}

func (this *Attrs)SetFloat32(name string, v float32)(err error) {
	this.Field[name] = v
	return
}

func (this *Attrs)SetString(name string, v string)(err error) {
	this.Field[name] = v
	return
}

func (this *Attrs)Clear() {
	this.Field = make(map[string]interface{})
}