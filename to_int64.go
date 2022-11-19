package cast_anything

func ToInt64E(value any, options ...Options) (int64, error) {
	return toSignedE[int64](value, options...)
}

func ToInt64(value any, options ...Options) int64 {
	v, _ := ToInt64E(value, options...)
	return v
}

func ToInt64OrDefault(value any, defaultValue int64, options ...Options) int64 {
	v, err := ToInt64E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}
