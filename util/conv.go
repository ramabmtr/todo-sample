package util

func GetBoolPointer(b bool) *bool {
	return &b
}

func GetBool(b *bool) bool {
	if b == nil {
		return false
	}

	return *b
}

func GetStringPointer(s string) *string {
	return &s
}

func GetString(s *string) string {
	if s == nil {
		return ""
	}

	return *s
}
