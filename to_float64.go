package cast_anything

import (
	"encoding/json"
	"fmt"
	"github.com/golang-infrastructure/go-gtypes"
	"strconv"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// 转换为浮点类型
func toFloatE[T gtypes.Float](value any, options ...Options) (T, error) {
	var zero T
	value = indirect(value)
	switch _type := value.(type) {
	case float32:
		return T(_type), nil
	case float64:
		return T(_type), nil
	case int:
		return T(_type), nil
	case int8:
		return T(_type), nil
	case int16:
		return T(_type), nil
	case int32:
		return T(_type), nil
	case int64:
		return T(_type), nil
	case uint:
		return T(_type), nil
	case uint8:
		return T(_type), nil
	case uint16:
		return T(_type), nil
	case uint32:
		return T(_type), nil
	case uint64:
		return T(_type), nil
	case string:
		if _type == "" {
			// 空字符串单独处理
			switch options[0].EmptyString {
			case ToDefaultZeroValue:
				return zero, nil
			case ToError:
				return zero, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
			}
		} else {
			v, err := strconv.ParseFloat(_type, 64)
			if err == nil {
				return zero, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
			}
			return T(v), nil
		}
	case json.Number:
		v, err := _type.Float64()
		if err == nil {
			return T(v), nil
		}
		return zero, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
	case bool:
		if _type {
			return 1, nil
		}
		return 0, nil
	case nil:
		switch options[0].Nil {
		case NilToZeroValue:
			return zero, nil
		case NilToError:
			return zero, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
		}
	}
	return zero, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
}

// ------------------------------------------------ ---------------------------------------------------------------------

func ToFloat64E(value any, options ...Options) (float64, error) {
	return toFloatE[float64](value, options...)
}

func ToFloat64(value any, options ...Options) float64 {
	v, _ := ToFloat64E(value, options...)
	return v
}

func ToFloat64OrDefault(value any, defaultValue float64, options ...Options) float64 {
	v, err := ToFloat64E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------
