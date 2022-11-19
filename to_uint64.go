package cast_anything

func ToUint64E(value any, options ...Options) (uint64, error) {
	return ToUnsignedE[uint64](value, options...)
}

func ToUint64(value any, options ...Options) uint64 {
	v, _ := ToUint64E(value, options...)
	return v
}

func ToUint64OrDefault(value any, defaultValue uint64, options ...Options) uint64 {
	v, err := ToUint64E(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}
