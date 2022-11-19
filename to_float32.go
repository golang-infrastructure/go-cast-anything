package cast_anything

func ToFloat32E(value any, options ...Options) (float32, error) {
	return toFloatE[float32](value, options...)
}

func ToFloat32(value any, options ...Options) float32 {
	v, _ := ToFloat32E(value, options...)
	return v
}

func ToFloat32OrDefault(value any, defaultValue float32, options ...Options) float32 {
	v, err := ToFloat32E(value, options...)
	if err != nil {
		return defaultValue
	} else {
		return v
	}
}
