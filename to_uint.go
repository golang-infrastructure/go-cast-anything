package cast_anything

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
)

// ---------------------------------------------------------------------------------------------------------------------

var DefaultUintOptions = Options{
	Float:       ToForceCast,
	EmptyString: ToDefaultZeroValue,
	Nil:         NilToZeroValue,
}

// ---------------------------------------------------------------------------------------------------------------------

// ToUnsignedE 把目标转换为无符号类型
func ToUnsignedE[T constraints.Unsigned](value any, options ...Options) (T, error) {
	var zero T
	options = append(options, DefaultUintOptions)
	value = indirect(value)
	switch _type := value.(type) {
	case nil:
		switch options[0].Nil {
		case NilToZeroValue:
			return zero, nil
		case NilToError:
			return zero, fmt.Errorf("unable to cast nil of type %T to %T", value, zero)
		}
	case string:
		// 空字符串单独处理
		if _type == "" {
			switch options[0].EmptyString {
			case ToDefaultZeroValue:
				return zero, nil
			case ToError:
				return zero, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
			}
		} else {
			v, err := strconv.ParseInt(trimZeroDecimal(_type), 0, 0)
			if err != nil {
				return zero, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
			}
			return signedCheck[int64, T](v)
		}
	case json.Number:
		return ToUnsignedE[T](string(_type))
	case int:
		return signedCheck[int, T](_type)
	case int8:
		return signedCheck[int8, T](_type)
	case int16:
		return signedCheck[int16, T](_type)
	case int32:
		return signedCheck[int32, T](_type)
	case int64:
		return signedCheck[int64, T](_type)
	case uint:
		return T(_type), nil
	case uint64:
		return T(_type), nil
	case uint32:
		return T(_type), nil
	case uint16:
		return T(_type), nil
	case uint8:
		return T(_type), nil
	case float32:
		return signedCheck[float32, T](_type)
	case float64:
		return signedCheck[float64, T](_type)
	case bool:
		if _type {
			return 1, nil
		} else {
			return 0, nil
		}
	}
	return zero, fmt.Errorf("unable to cast %#v of type %T to uint", value, value)
}

// 有符号类型向无符号类型转换的时候都要经过这个检查，看看数据范围是否合法，主要是进行一个符号的数据范围检查
func signedCheck[Source constraints.Signed | constraints.Float, Target constraints.Unsigned](source Source) (Target, error) {
	if source < 0 {
		return 0, fmt.Errorf("unable to cast negative value")
	} else {
		return Target(source), nil
	}
}

// ---------------------------------------------------------------------------------------------------------------------

func ToUintE(value any, options ...Options) (uint, error) {
	return ToUnsignedE[uint](value, options...)
}

func ToUint(value any, options ...Options) uint {
	v, _ := ToUintE(value, options...)
	return v
}

func ToUintOrDefault(value any, defaultValue uint, options ...Options) uint {
	v, err := ToUintE(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}

// ---------------------------------------------------------------------------------------------------------------------
