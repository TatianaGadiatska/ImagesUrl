package view

import (
	"fmt"
	"github.com/cenkalti/dominantcolor"
	"golang.org/x/image/draw"
	"html/template"
	"image"
	"log"
	"net/http"
)

type RecordUrl struct {
	Id     int
	UrlImg string
	Color  string
}

type ViewData struct {
	UrlImgData []RecordUrl
}

func Format(rs []RecordUrl) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		data := ViewData{
			UrlImgData: rs,
		}

		tmpl, _ := template.ParseFiles("index.html")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8181", nil)
}

func FindFromUrl(url string) string {

	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Print(err)
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		log.Print(err)
	}
	// создаём пустое изображение для
	// записи необходимого размера
	dst := image.NewRGBA(image.Rect(0, 0, 200, 200))
	// изменение размера
	draw.CatmullRom.Scale(dst, dst.Bounds(), img, img.Bounds(), draw.Over, nil)

	return dominantcolor.Hex(dominantcolor.Find(dst))
}
