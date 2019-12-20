package service

import (
	"newbug/model"
	"reflect"
)

type ValidResult struct {
	Valid bool
	ErrorMessage string
}

type ValidSlice []ValidResult

func Store(user *model.User) (int64,error) {
	createResult := model.Connect().Create(&user)
	return createResult.RowsAffected,createResult.Error
}

func FindByEmail(user *model.User) {
	model.Connect().Where("email = ?", user.Email).First(user)
}

/**
 * 参数不为空验证
 */
func Valid(user interface{}) ValidSlice {
	var validSlice ValidSlice

	v := reflect.ValueOf(user)
	t := reflect.TypeOf(user)

	for i:=0;i<t.NumField();i++{
		value := v.Field(i)
		tag := t.Field(i).Tag
		notnull := tag.Get("notnull")
		if len(notnull) > 0 {
			if IsBlank(value) {
				validSlice = append(validSlice,ValidResult{false,notnull})
			}
		}
	}

	return  validSlice
}

/**
 * 判断Value是否为空
 */
func IsBlank(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return value.Len() == 0
	case reflect.Bool:
		return !value.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return value.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return value.IsNil()
	}

	return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}