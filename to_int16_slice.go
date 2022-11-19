package cast_anything

func ToInt16SliceE(value any, options ...Options) ([]int16, error) {
	return toSliceE(value, func(value any, options ...Options) (int16, error) {
		return ToInt16E(value, options...)
	}, options...)
}

func ToInt16Slice(value any, options ...Options) []int16 {
	e, _ := ToInt16SliceE(value, options...)
	return e
}

func ToInt16SliceOrDefault(value []any, defaultValue []int16, options ...Options) []int16 {
	e, err := ToInt16SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}

