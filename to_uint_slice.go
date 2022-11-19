package cast_anything

import (
	"fmt"
	"reflect"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// 将值转换为切片
func toSliceE[T any](value any, toFuncE func(value any, options ...Options) (T, error), options ...Options) ([]T, error) {
	var zero T
	if IsNil(value) {
		return []T{}, fmt.Errorf("unable to cast nil of type %T to []%T", value, zero)
	}

	// 如果直接就是这种类型的话，可以直接返回了
	switch v := value.(type) {
	case []T:
		return v, nil
	}

	// 否则的话则将列表结构中的元素逐个转换
	kind := reflect.TypeOf(value).Kind()
	switch kind {
	case reflect.Slice, reflect.Array:
		reflectValue := reflect.ValueOf(value)
		resultSlice := make([]T, reflectValue.Len())
		for index := 0; index < reflectValue.Len(); index++ {
			newValue, err := toFuncE(reflectValue.Index(index).Interface(), options...)
			if err != nil {
				return []T{}, fmt.Errorf("unable to cast %#v of type %T to []%T", value, value, zero)
			}
			resultSlice[index] = newValue
		}
		return resultSlice, nil
	default:
		return []T{}, fmt.Errorf("unable to cast %#v of type %T to []bool", value, value)
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func ToUintSliceE(value any, options ...Options) ([]uint, error) {
	return toSliceE(value, func(value any, options ...Options) (uint, error) {
		return ToUintE(value, options...)
	}, options...)
}

func ToUintSlice(value any, options ...Options) []uint {
	e, _ := ToUintSliceE(value, options...)
	return e
}

func ToUintSliceOrDefault(value any, defaultValue []uint, options ...Options) []uint {
	e, err := ToUintSliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}

// ------------------------------------------------ ---------------------------------------------------------------------
