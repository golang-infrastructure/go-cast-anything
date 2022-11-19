package cast_anything

func ToUint8SliceE(value any, options ...Options) ([]uint8, error) {
	return toSliceE(value, func(value any, options ...Options) (uint8, error) {
		return ToUint8E(value, options...)
	}, options...)
}

func ToUint8Slice(value any, options ...Options) []uint8 {
	e, _ := ToUint8SliceE(value, options...)
	return e
}

func ToUint8SliceOrDefault(value any, defaultValue []uint8, options ...Options) []uint8 {
	e, err := ToUint8SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
