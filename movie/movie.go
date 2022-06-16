package movie

import "fmt"

func init() {
	fmt.Println("Init from movie")
}

func ReviewMovie(name string, rating float64) {
	fmt.Printf("This movie name = %v and rating = %.2f\n", name, rating)
}

func FindMovie(movieID string) string {
	switch movieID {
	case "a":
		return "Movie name Avenger"
	case "b":
		return "Movie name Batman"
	}
	return "not found"
}
