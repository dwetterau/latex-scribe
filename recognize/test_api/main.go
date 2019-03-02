package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/dwetterau/latex-scribe/recognize"
)

func main() {
	// Load in the image
	b, err := ioutil.ReadFile("../text_and_latex_test.png")
	if err != nil {
		panic(err)
	}
	b64Image := base64.StdEncoding.EncodeToString(b)
	url := fmt.Sprintf("data:image/png;base64,%s", b64Image)

	r := recognize.New()
	result, err := r.ToLatex(url)
	if err != nil {
		panic(err)
	}
	fmt.Println("got response", result)
}
