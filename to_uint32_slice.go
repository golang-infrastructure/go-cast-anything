package cast_anything

func ToUint32SliceE(value any, options ...Options) ([]uint32, error) {
	return toSliceE(value, func(value any, options ...Options) (uint32, error) {
		return ToUint32E(value, options...)
	}, options...)
}

func ToUint32Slice(value any, options ...Options) []uint32 {
	e, _ := ToUint32SliceE(value, options...)
	return e
}

func ToUint32SliceOrDefault(value any, defaultValue []uint32, options ...Options) []uint32 {
	e, err := ToUint32SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
