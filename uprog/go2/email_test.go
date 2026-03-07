package email

import "testing"

func TestIsValidEmail(t *testing.T) {

	// таблица тестов
	tests := []struct {
		email string
		valid bool
	}{
		{"bbbaav@ya.ru", true},
		{"bbbaa@bk.ru", true},
		{"test@gmail.com", true},
		{"user@mail.com", true},

		{"wrong-email", false},
		{"test@", false},
		{"@mail.com", false},
		{"", false},
	}

	// перебор всех тестов
	for _, tt := range tests {

		t.Run(tt.email, func(t *testing.T) {

			result := IsValidEmail(tt.email)

			if result != tt.valid {
				t.Errorf(
					"IsValidEmail(%s) = %v; want %v",
					tt.email,
					result,
					tt.valid,
				)
			}

		})

	}
}