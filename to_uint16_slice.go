package cast_anything

func ToUint16SliceE(value any, options ...Options) ([]uint16, error) {
	return toSliceE(value, func(value any, options ...Options) (uint16, error) {
		return ToUint16E(value, options...)
	}, options...)
}

func ToUint16Slice(value any, options ...Options) []uint16 {
	e, _ := ToUint16SliceE(value, options...)
	return e
}

func ToUint16SliceOrDefault(value any, defaultValue []uint16, options ...Options) []uint16 {
	e, err := ToUint16SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
