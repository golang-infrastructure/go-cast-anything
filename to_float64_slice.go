package cast_anything

func ToFloat64SliceE(value any, options ...Options) ([]float64, error) {
	return toSliceE(value, func(value any, options ...Options) (float64, error) {
		return ToFloat64E(value, options...)
	}, options...)
}

func ToFloat64Slice(value any, options ...Options) []float64 {
	e, _ := ToFloat64SliceE(value, options...)
	return e
}

func ToFloat64SliceOrDefault(value []any, defaultValue []float64, options ...Options) []float64 {
	e, err := ToFloat64SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
