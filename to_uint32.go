package cast_anything

func ToUint32E(value any, options ...Options) (uint32, error) {
	return ToUnsignedE[uint32](value, options...)
}

func ToUint32(value any, options ...Options) uint32 {
	v, _ := ToUint32E(value, options...)
	return v
}

func ToUint32OrDefault(value any, defaultValue uint32, options ...Options) uint32 {
	v, err := ToUint32E(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}
