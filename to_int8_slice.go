package cast_anything

func ToInt8SliceE(value any, options ...Options) ([]int8, error) {
	return toSliceE(value, func(value any, options ...Options) (int8, error) {
		return ToInt8E(value, options...)
	}, options...)
}

func ToInt8Slice(value any, options ...Options) []int8 {
	e, _ := ToInt8SliceE(value, options...)
	return e
}

func ToInt8SliceOrDefault(value []any, defaultValue []int8, options ...Options) []int8 {
	e, err := ToInt8SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
