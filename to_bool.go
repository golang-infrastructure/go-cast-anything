package cast_anything

import (
	"encoding/json"
	"fmt"
	"github.com/golang-infrastructure/go-maths"
	reflect_utils "github.com/golang-infrastructure/go-reflect-utils"
	"reflect"
	"strconv"
)

// DefaultBoolOptions 默认类型转换默认处理规则
var DefaultBoolOptions = Options{
	Float:       ToDefaultZeroValue,
	EmptyString: ToDefaultZeroValue,
	Nil:         NilToZeroValue,
}

// ------------------------------------------------ ---------------------------------------------------------------------

func ToBoolE(value any, options ...Options) (bool, error) {
	options = append(options, DefaultBoolOptions)
	value = indirect(value)

	// 先尝试按照直接类型转换
	switch _type := value.(type) {
	case nil:
		switch options[0].Nil {
		case NilToZeroValue:
			// 把nil转换为布尔类型的零值，即false
			return false, nil
		case NilToError:
			// 不能转换，直接拒绝转换
			return false, fmt.Errorf("unable to cast type %T nil to bool", value)
		}
	case bool:
		return _type, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		// 如果是integer类型的话，非0转换为true，0转换为false
		return _type != 0, nil
	case float32, float64:
		// 如果是浮点类型的话，非零转换为true，零值转换为false
		if maths.IsZero(ToFloat64(_type)) {
			return false, nil
		} else {
			return true, nil
		}
	case uintptr:
		// 认为是在判断指针类型不为空，因此只要不是0就可以了
		return _type != 0, nil
	case string:
		// 空字符串的时候特殊处理
		if _type == "" {
			switch options[0].EmptyString {
			case ToDefaultZeroValue:
				// 空字符串转为false
				return false, nil
			case ToError:
				// 空字符串无法处理，直接拒绝
				return false, fmt.Errorf("unable to cast %#v of type %T to bool", value, value)
			}
		} else {
			// 非空字符串尝试去解析一下，但是并不一定能解析得出来，这里暂时不提供将非空字符串直接转为true的选项了
			return strconv.ParseBool(value.(string))
		}
	case json.Number:
		v, err := ToInt64E(_type)
		if err != nil {
			return false, fmt.Errorf("unable to cast %#v of type %T to bool", value, value)
		}
		return v != 0, nil
		// TODO 复数怎么转换
		//case complex128, complex64:
	}

	// 然后尝试按照反射类型转换
	reflectValue := reflect.ValueOf(value)
	switch reflectValue.Kind() {
	case reflect.Invalid:
		return false, nil
	// TODO 复数怎么处理啊
	//case reflect.Complex64:
	//case reflect.Complex128:
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Slice:
		// 如果是这几个有长度的容器的话，则看下其是否为空，如果没有元素认为是false,有元素则认为是true
		return reflectValue.Len() != 0, nil
	case reflect.Func:
		// 函数类型始终转为true
		return true, nil
	case reflect.Interface, reflect.Struct:
		// 如果为空的话转为false，否则转为true
		return reflect_utils.IsNotNil(value), nil
		//case reflect.Pointer, reflect.UnsafePointer:
	}

	return false, fmt.Errorf("unable to cast %#v of type %T to bool", value, value)
}

func ToBool(value any, options ...Options) bool {
	v, _ := ToBoolE(value, options...)
	return v
}

func ToBoolOrDefault(value any, defaultValue bool, options ...Options) bool {
	v, err := ToBoolE(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}
