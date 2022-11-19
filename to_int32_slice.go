package cast_anything

func ToInt32SliceE(value any, options ...Options) ([]int32, error) {
	return toSliceE(value, func(value any, options ...Options) (int32, error) {
		return ToInt32E(value, options...)
	}, options...)
}

func ToInt32Slice(value any, options ...Options) []int32 {
	e, _ := ToInt32SliceE(value, options...)
	return e
}

func ToInt32SliceOrDefault(value []any, defaultValue []int32, options ...Options) []int32 {
	e, err := ToInt32SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
