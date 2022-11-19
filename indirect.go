package cast_anything

import "reflect"

// 消除指针，寻找真实的值
func indirect(value any) any {
	if IsNil(value) {
		return nil
	}
	if t := reflect.TypeOf(value); t.Kind() != reflect.Ptr {
		// Avoid creating value reflect.Value if it's not value pointer.
		return value
	}
	reflectValue := reflect.ValueOf(value)
	for reflectValue.Kind() == reflect.Ptr && !reflectValue.IsNil() {
		reflectValue = reflectValue.Elem()
	}
	return reflectValue.Interface()
}

// IsNil 判断参数是否为nil，当只有type但是value为nil的时候会认为是nil
func IsNil(x any) bool {

	if x == nil {
		return true
	}

	v := reflect.ValueOf(x)
	if !v.IsValid() {
		return true
	}

	switch v.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Slice:
		return v.IsNil()
	case reflect.Ptr:
		elem := v.Elem()
		if elem.IsValid() {
			return IsNil(elem.Interface())
		} else {
			return true
		}
	default:
		return false
	}
}

// IsNotNil 判断参数是否不为nil
func IsNotNil(v any) bool {
	return !IsNil(v)
}
