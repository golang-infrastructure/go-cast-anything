package cast_anything

func ToUint64SliceE(value any, options ...Options) ([]uint64, error) {
	return toSliceE(value, func(value any, options ...Options) (uint64, error) {
		return ToUint64E(value, options...)
	}, options...)
}

func ToUint64Slice(value any, options ...Options) []uint64 {
	e, _ := ToUint64SliceE(value, options...)
	return e
}

func ToUint64SliceOrDefault(value []any, defaultValue []uint64, options ...Options) []uint64 {
	e, err := ToUint64SliceE(value, options...)
	if err != nil {
		return defaultValue
	}
	return e
}
