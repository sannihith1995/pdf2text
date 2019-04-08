package main

import (
	"os"

	fitz "github.com/gen2brain/go-fitz"
)

func main() {
	doc, err := fitz.New("Thai.pdf")
	if err != nil {
		panic(err)
	}

	defer doc.Close()

	// Extract pages as text
	for n := 0; n < doc.NumPage(); n++ {
		text, err := doc.Text(n)
		if err != nil {
			panic(err)
		}

		file, err := os.Create("out.txt")
		defer file.Close()
		file.WriteString(text)
	}

	// Extract pages as html
	for n := 0; n < doc.NumPage(); n++ {
		html, err := doc.HTML(n, true)
		if err != nil {
			panic(err)
		}

		file, err := os.Create("out.html")
		defer file.Close()
		file.WriteString(html)
	}
}
