package tsReflect

import (
	"reflect"
	"errors"
)

func NewObj(t reflect.Type)(interface{}) {
	return reflect.New(t.Elem()).Interface()
}

func GetType(i interface{})(reflect.Type) {
	return reflect.TypeOf(i)
}

/*
 * @brief 给结构体的指针域，创建指针对象
 * struct {v *int}
 *
 * @param i 结构体
 * @param field 域名称
 */
func NewFieldPoint(i interface{}, field string)(err error) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() {
		err = errors.New(field + "invalid")
		return
	}
	value_field.Set(reflect.New(value_field.Type().Elem()))	
	return
}

func SetFieldInt(i interface{}, field string, v int64)(err error) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() {
		err = errors.New(field + "invalid")
		return
	}
	if value_field.CanAddr() {
		if value_field.IsNil() {
			value_field.Set(reflect.New(value_field.Type().Elem()))
		}
	}
	value_field.Elem().SetInt(v)

	return
}

func SetFieldFloat(i interface{}, field string, v float64)(err error) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() {
		err = errors.New(field + "invalid")
		return
	}
	if value_field.CanAddr() {
		if value_field.IsNil() {
			value_field.Set(reflect.New(value_field.Type().Elem()))
		}
	}
	value_field.Elem().SetFloat(v)

	return
}

func SetFieldBool(i interface{}, field string, v bool)(err error) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() {
		err = errors.New(field + "invalid")
		return
	}
	if value_field.CanAddr() {
		if value_field.IsNil() {
			value_field.Set(reflect.New(value_field.Type().Elem()))
		}
	}
	value_field.Elem().SetBool(v)

	return
}

func SetFieldString(i interface{}, field string, v string)(err error) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() {
		err = errors.New(field + "invalid")
		return
	}
	if value_field.CanAddr() {
		if value_field.IsNil() {
			value_field.Set(reflect.New(value_field.Type().Elem()))
		}
	}
	value_field.Elem().SetString(v)

	return
}

func NewFieldSlice(i interface{}, field string, count int64)(err error) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() {
		err = errors.New(field + "invalid")
		return
	}
	//t_data := reflect.TypeOf(data)
	//t_slice := reflect.MakeSlice(reflect.SliceOf(t_data), 0, 1)
	
	//value_field.Elem().Type()
	indirectStr := reflect.Indirect(value_field)
	t_slice := reflect.MakeSlice(indirectStr.Type(), int(count), int(count+1))

	value_field.Set(t_slice)
	return
}

func ChangSlice(i interface{}, field string, pos1 int, pos2 int)(err error) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() { 
		err = errors.New(field + "invalid")
		return
	}
	
	p1 := value_field.Index(pos1)
	p2 := value_field.Index(pos2)
	
	i_p1 := p1.Interface()
	i_p2 := p2.Interface()
	
	p1.Set(reflect.ValueOf(i_p2))
	p2.Set(reflect.ValueOf(i_p1))
	return
}

func GetInterfaceSlice(i interface{}, field string, pos int)(interface{}){
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() { 
		return nil
	}
	
	return value_field.Index(pos).Interface()
}

func SliceLen(i interface{}, field string)(int64) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() {
		return 0
	}
	
	return int64(value_field.Len())
}

func AppendFieldSlice(i interface{}, field string, data interface{})(err error) {
	value_i := reflect.ValueOf(i).Elem()
	value_field := value_i.FieldByName(field)
	if !value_field.IsValid() {
		err = errors.New(field + "invalid")
		return
	}
	
	value_field.Len()
	
	value_field1 := reflect.Append(value_field, reflect.ValueOf(data))
	value_field.Set(value_field1)
	return
}

func AppendFieldSliceInt(i interface{}, field string, data int64)(err error) {
	err = AppendFieldSlice(i, field, data) 
	return
}