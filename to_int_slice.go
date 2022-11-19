package cast_anything

func ToIntSliceE(value any, options ...Options) ([]int, error) {
	return toSliceE(value, func(value any, options ...Options) (int, error) {
		return ToIntE(value, options...)
	}, options...)
}

func ToIntSlice(value any, options ...Options) []int {
	e, _ := ToIntSliceE(value, options...)
	return e
}

func ToIntSliceOrDefault(value []any, defaultValue []int, options ...Options) []int {
	e, err := ToIntSliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
