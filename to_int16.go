package cast_anything

func ToInt16E(value any, options ...Options) (int16, error) {
	return toSignedE[int16](value, options...)
}

func ToInt16(value any, options ...Options) int16 {
	v, _ := ToInt16E(value, options...)
	return v
}

func ToInt16OrDefault(value any, defaultValue int16, options ...Options) int16 {
	v, err := ToInt16E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}
