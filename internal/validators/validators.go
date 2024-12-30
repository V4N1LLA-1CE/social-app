package validators

import (
	"regexp"
	"strings"

	"github.com/V4N1LLA-1CE/netio"
)

var emailRx = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func ValidatePassword(v *netio.Validator, password string) *netio.Validator {
	v.Check(len(password) >= 8, "password_minlen", "password must contain at least 8 characters")
	v.Check(len(password) <= 255, "password_maxlen", "password cannot exceed 255 characters")

	v.Check(strings.ContainsAny(password, "abcdefghijklmnopqrstuvwxyz"), "password_lowercase", "password must contain at least one lowercase letter")
	v.Check(strings.ContainsAny(password, "ABCDEFGHIJKLMNOPQRSTUVWXYZ"), "password_uppercase", "password must contain at least one uppercase letter")
	v.Check(strings.ContainsAny(password, "0123456789"), "password_digit", "password must contain at least one number")
	v.Check(strings.ContainsAny(password, "@$!%*?&"), "password_special", "password must contain at least one special character (@$!%*?&)")

	return v
}

func ValidateUsername(v *netio.Validator, username string) *netio.Validator {
	v.Check(len(username) >= 3, "username_minlen", "username must contain at least 3 characters")
	v.Check(len(username) <= 30, "username_maxlen", "username cannot exceed 30 characters")

	return v
}

func ValidateEmail(v *netio.Validator, email string) *netio.Validator {
	// Length checks
	v.Check(len(email) > 0, "email_minlen", "email cannot be empty")
	v.Check(len(email) <= 255, "email_maxlen", "email cannot exceed 255 characters")

	// Format validation
	v.Check(netio.Matches(email, emailRx), "email_format", "invalid email format")

	return v
}
