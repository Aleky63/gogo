package email

import "regexp"

// регулярное выражение для проверки email
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

// функция возвращает true если email корректный
func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}