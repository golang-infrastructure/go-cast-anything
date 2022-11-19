package cast_anything

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
)

// DefaultBoolOptions 默认类型转换默认处理规则
var DefaultBoolOptions = Options{
	Float:       ToDefaultZeroValue,
	EmptyString: ToDefaultZeroValue,
	Nil:         NilRuleToZeroValue,
}

// ------------------------------------------------ ---------------------------------------------------------------------

func ToBoolE(value any, options ...Options) (bool, error) {
	options = append(options, DefaultBoolOptions)
	value = indirect(value)

	switch _type := value.(type) {
	case nil:
		switch options[0].Nil {
		case NilRuleToZeroValue:
			return false, nil
		case NilRuleToError:
			return false, fmt.Errorf("unable to cast type %T nil to bool", value)
		}
	case bool:
		return _type, nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		// 如果是integer类型的话，非0转换为true，0转换为false
		return _type != 0, nil
	case float32, float64:
		// 如果是浮点类型的话，非零转换为true，零值转换为false
		if math.Abs(ToFloat64(_type)) < 0.00001 {
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
				return false, nil
			case ToError:
				return false, fmt.Errorf("unable to cast %#v of type %T to bool", value, value)
			}
		} else {
			// 非空字符串尝试去解析一下
			return strconv.ParseBool(value.(string))
		}
	case json.Number:
		v, err := ToInt64E(_type)
		if err == nil {
			return v != 0, nil
		}
		return false, fmt.Errorf("unable to cast %#v of type %T to bool", value, value)
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
