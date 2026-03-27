package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// var gameName string = strings.Join(os.Args[1:], "%20")

func main() {
	// userInput := readInput("Which game do you want to know about: \n")
	// fmt.Println(userInput)

	// gameName := os.Args[1] + "%20" + os.Args[2]

	http.HandleFunc("/metadata", ScraperHandler) // this means that any request to /metadata will triger the func, like writing something in the query param's val will send the req to our server which will trigger handler.
	// http.HandleFunc("/humidity", tempScraperHandler)
	// http.HandleFunc("/petrol", tempScraperHandler)

	fmt.Println("starting http server...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}


func ScraperHandler(w http.ResponseWriter, r *http.Request) {
	game := r.FormValue("name")
	game = url.QueryEscape(game)

	// fmt.Println(game)

	title, releaseDate, imageUrl := ExtractGameDetails(game)

	fmt.Println(title, releaseDate, imageUrl)

	w.Write([]byte("Title: " + title + "\n" + "releaseDate: " + releaseDate + "\n" + "Image URL: " + imageUrl + "\n"))
}

func ExtractGameDetails(gameName string )  (string, string, string) {
	fmt.Printf("looking up for %s game review...\n", gameName)

	res, err := http.Get("https://www.metacritic.com/search/" + gameName + "/")
	if err != nil {
		fmt.Printf("[http.Get()] error: %s\n", err.Error())
	}

	resBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[io.ReadAll()] error: %s\n", err.Error())
	}
	res.Body.Close()

	rawHtml := string(resBytes)
	if rawHtml != "" {
		fmt.Println("http GET request returned response...")
	}

	html, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		panic(err)
	}

	title := html.Find(`p.c-search-item__title`).First().Text()
	releaseDate := html.Find(`li.c-search-product-meta__release-date`).First().Text()
	imageUrl := html.Find(`div.c-search-item__image img`).AttrOr("src", "")


	// text := html.Find(`div[class="c-search-product-meta c-search-product-meta--search-page"] ul[class="c-search-product-meta__list"] span`).First().Text()
	fmt.Printf("title: %s, release-date: %s, image-url: %s", title, releaseDate, imageUrl)

	// fmt.Printf("Game Name: %s and Game Data: %s\n", gameName, gameData)

	return title, releaseDate, imageUrl

}

// func ExtractGameDetails() {
// 	title := html.Find(`p.c-search-item__title`).First().Text()
// 	releaseDate := html.Find(`li.c-search-product-meta__release-date`).First().Text()
// 	imageUrl := html.Find(`div.c-search-item__image img`).AttrOr("src", "")
// 	// text := html.Find(`div[class="c-search-product-meta c-search-product-meta--search-page"] ul[class="c-search-product-meta__list"] span`).First().Text()

// 	fmt.Printf("title: %s, release-date: %s, image-url: %s", title, releaseDate, imageUrl)
// }

// func readInput(prompt string) string {
// 	fmt.Printf("%s", prompt)

// 	var userInput string
// 	fmt.Scan(&userInput)

// 	return userInput
// }




//just to commit smth yooooooooooooooooooo



