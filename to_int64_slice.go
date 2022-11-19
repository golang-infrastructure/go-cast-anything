package cast_anything

func ToInt64SliceE(value any, options ...Options) ([]int64, error) {
	return toSliceE(value, func(value any, options ...Options) (int64, error) {
		return ToInt64E(value, options...)
	}, options...)
}

func ToInt64Slice(value any, options ...Options) []int64 {
	e, _ := ToInt64SliceE(value, options...)
	return e
}

func ToInt64SliceOrDefault(value []any, defaultValue []int64, options ...Options) []int64 {
	e, err := ToInt64SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
