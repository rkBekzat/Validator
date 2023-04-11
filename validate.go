package Validator

func CheckTitle(txt string) bool {
	if txt == "" || len(txt) >= 100 {
		return false
	}
	return true
}

func CheckText(txt string) bool {
	if txt == "" || len(txt) >= 500 {
		return false
	}
	return true
}
