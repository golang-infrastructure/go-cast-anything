package cast_anything

func ToUint16E(value any, options ...Options) (uint16, error) {
	return ToUnsignedE[uint16](value, options...)
}

func ToUint16(value any, options ...Options) uint16 {
	v, _ := ToUint16E(value, options...)
	return v
}

func ToUint16OrDefault(value any, defaultValue uint16, options ...Options) uint16 {
	v, err := ToUint16E(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}
