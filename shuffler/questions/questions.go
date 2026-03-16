// questions/questions.go
package questions

type Question struct {
	Country string
	Capital string
}

func All() []Question {
	return []Question{
		{"Франции", "Париж"},
		{"Японии", "Токио"},
		{"Бразилии", "Бразилиа"},
		{"Египта", "Каир"},
		{"Канады", "Оттава"},
		{"Австралии", "Канберра"},
		{"Германии", "Берлин"},
		{"Индии", "Нью-Дели"},
		{"Южной Кореи", "Сеул"},
		{"Аргентины", "Буэнос-Айрес"},
	}
}
