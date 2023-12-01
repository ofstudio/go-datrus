package datrus

const lenAny = 0

func validateNumber(number string, expectedLen int) error {
	runes := []rune(number)
	if len(runes) == 0 {
		return ErrLen
	}
	if expectedLen != lenAny && len(runes) != expectedLen {
		return ErrLen
	}
	for _, char := range runes {
		if char < '0' || char > '9' {
			return ErrChar
		}
	}
	return nil
}
