package utils

import "strings"

func CheckPassword(password string) string {
	if len(password) < 9 {
		return "Password must be at least 9 characters long"
	}
	if strings.Contains(password, " ") {
		return "Password must not contain spaces"
	}
	// check if have numbers
	if !strings.ContainsAny(password, "0123456789") {
		return "Password must contain at least one number"
	}
	// check if have special characters
	if !strings.ContainsAny(password, "!@#$%^&*()_+") {
		return "Password must contain at least one special character"
	}
	return ""
}
