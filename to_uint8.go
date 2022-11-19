package cast_anything

func ToUint8E(value any, options ...Options) (uint8, error) {
	return ToUnsignedE[uint8](value, options...)
}

func ToUint8(value any, options ...Options) uint8 {
	v, _ := ToUint8E(value, options...)
	return v
}

func ToUint8OrDefault(value any, defaultValue uint8, options ...Options) uint8 {
	v, err := ToUint8E(value, options...)
	if err != nil {
		return defaultValue
	}
	return v
}
