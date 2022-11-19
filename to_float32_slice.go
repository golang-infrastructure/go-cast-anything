package cast_anything

func ToFloat32SliceE(value any, options ...Options) ([]float32, error) {
	return toSliceE(value, func(value any, options ...Options) (float32, error) {
		return ToFloat32E(value, options...)
	}, options...)
}

func ToFloat32Slice(value any, options ...Options) []float32 {
	e, _ := ToFloat32SliceE(value, options...)
	return e
}

func ToFloat32SliceOrDefault(value []any, defaultValue []float32, options ...Options) []float32 {
	e, err := ToFloat32SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
