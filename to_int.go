package cast_anything

import (
	"encoding/json"
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
	"time"
)

// ------------------------------------------------ ---------------------------------------------------------------------

// 转换为有符号整数
func toSignedE[T constraints.Signed](value any, options ...Options) (T, error) {
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
			return T(v), nil
		}
	case json.Number:
		e, err := ToIntE(string(_type))
		if err != nil {
			return zero, err
		}
		return T(e), nil
	case bool:
		if _type {
			return 1, nil
		}
		return 0, nil
	case time.Weekday:
		return T(_type), nil
	case time.Month:
		return T(_type), nil
	case nil:
		switch options[0].Nil {
		case NilToZeroValue:
			return zero, nil
		case NilToError:
			return zero, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
		}
	}
	return 0, fmt.Errorf("unable to cast %#v of type %T to %T", value, value, zero)
}

// ------------------------------------------------ ---------------------------------------------------------------------

func ToIntE(value any, options ...Options) (int, error) {
	return toSignedE[int](value, options...)
}

func ToInt(value any) int {
	v, _ := ToIntE(value)
	return v
}

func ToIntOrDefault(value any, defaultValue int) int {
	v, err := ToIntE(value)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}

// ------------------------------------------------ ---------------------------------------------------------------------

func trimZeroDecimal(s string) string {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}

// ------------------------------------------------- --------------------------------------------------------------------
