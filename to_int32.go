package cast_anything

func ToInt32E(value any, options ...Options) (int32, error) {
	return toSignedE[int32](value, options...)
}

func ToInt32(value any, options ...Options) int32 {
	v, _ := ToInt32E(value, options...)
	return v
}

func ToInt32OrDefault(value any, defaultValue int32, options ...Options) int32 {
	v, err := ToInt32E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}
