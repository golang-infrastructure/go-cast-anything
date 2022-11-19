package cast_anything

func ToInt8E(value any, options ...Options) (int8, error) {
	return toSignedE[int8](value, options...)
}

func ToInt8(value any, options ...Options) int8 {
	v, _ := ToInt8E(value, options...)
	return v
}

func ToInt8OrDefault(value any, defaultValue int8, options ...Options) int8 {
	v, err := ToInt8E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}
