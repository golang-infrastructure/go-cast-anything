package cast_anything

func ToBoolSliceE(value any, options ...Options) ([]bool, error) {
	return toSliceE(value, func(value any, options ...Options) (bool, error) {
		return ToBoolE(value, options...)
	}, options...)
}

func ToBoolSlice(value any, options ...Options) []bool {
	v, _ := ToBoolSliceE(value, options...)
	return v
}

func ToBoolSliceOrDefault(value any, defaultValue []bool, options ...Options) []bool {
	slice, err := ToBoolSliceE(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return slice
	}
}
