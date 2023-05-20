package main

import (
	"fmt"
	"net/http"

	"github.com/StalkR/imdb"
)

func main() {
	client := http.DefaultClient
	title, _ := imdb.NewTitle(client, "tt10896634")
	fmt.Printf("%v^%v^%v^%v\n", title.URL, title.Name, title.Year, title.Rating)
}
