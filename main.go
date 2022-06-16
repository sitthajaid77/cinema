package main

import (
	"fmt"

	"github.com/sitthajaid77/cinema/advance"
	"github.com/sitthajaid77/cinema/movie"
)

func init() {
	fmt.Println("Init from main")
}

func main() {
	fmt.Println("This is cinema")
	movie.ReviewMovie("Avenger", 9.99)
	mName := movie.FindMovie("a")
	advance.TestAdvance()
	fmt.Println(mName)
}
