package main

import (
	"database/sql"
	_ "github.com/FogCreek/mini"
	_ "github.com/lib/pq"
	"log"
	_ "os"
	_ "os/user"
	_ "strconv"
	_ "sync"

	"./downladUrl"
	"./model"
	"./view"
	_ "image/jpeg"
)

func main() {

	links := downladUrl.Generator()
	log.Printf("\nСкачено %v ссылок \n", len(links))

	urlColor := makeUrlColor(links)

	var connStr = "user=postgres password=123 dbname=mydb sslmode=disable"
	var dbUrl, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
	defer dbUrl.Close()

	model.CreatTable(dbUrl)

	model.InsertUrl(dbUrl, urlColor)

	result := model.SelectUrl(dbUrl)

	view.Format(result)
}

func makeUrlColor(links []string) []view.RecordUrl {
	var resaltUrlColor []view.RecordUrl
	var resalt view.RecordUrl
	for _, url := range links {
		resalt.UrlImg = url
		resalt.Color = view.FindFromUrl(url)
		resaltUrlColor = append(resaltUrlColor, resalt)
	}
	return resaltUrlColor
}
