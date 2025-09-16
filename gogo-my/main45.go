package main

import (
	"fmt"
	"maps"
	"slices"
)



func printRecommendations(movies map[string]map[string]float64) {

	genresMoviesGood := make(map[string]map[string]float64)

	for genre, films := range movies {
		goodMovies :=  make(map[string]float64)

		for filmName, rating := range films {
			if rating >= 7.0 {
				goodMovies[filmName] = rating
			}
		}

		if len(goodMovies) > 0 {
			genresMoviesGood[genre] = goodMovies
		}
	}
	genres := slices.Collect(maps.Keys(genresMoviesGood))
	slices.Sort(genres)

for _, genre := range genres {
	films := genresMoviesGood[genre]
		movieNames := slices.Collect(maps.Keys(films))

slices.SortFunc(movieNames, func(a, b string) int {
		
			ratingA := films[a]
			ratingB := films[b]

          if ratingA != ratingB {
				if ratingA > ratingB {
					return -1
				}
				return 1
			}

			if a < b {
				return -1
			} else if a > b {
				return 1
			}
			return 0
		})
		for _, movieName := range movieNames {
			rating := films[movieName]
			fmt.Printf("  %s (%.2f)\n", movieName, rating)
		}
		fmt.Println() 
	}

}

func main45() {
movies := map[string]map[string]float64{
		"Экшен": {
			"Матрица":           8.7,
			"Терминатор 2":      8.5,
			"Джон Уик":          7.4,
			"Быстрый и яростный": 6.5,
		},
		"Драма": {
			"Побег из Шоушенка": 9.3,
			"Крестный отец":     9.2,
			"Форрест Гамп":      8.8,
			"Зеленая миля":      8.6,
		},
		"Комедия": {
			"Один дома":         7.7,
			"Маска":            6.9,
			"Такси":            7.0,
		},
		"Ужасы": {
			"Звонок":           6.2,
			"Пила":             6.1,
		},
		"Фантастика": {
			"Интерстеллар":     8.6,
			"Начало":           8.85,
			"Бегущий по лезвию": 8.1,
		},
	}
	
	fmt.Println("Рекомендации фильмов (рейтинг 7.0+):")
	fmt.Println("====================================")
	printRecommendations(movies)
}
// https://stepik.org/lesson/1517684/step/1?unit=1537967package main

