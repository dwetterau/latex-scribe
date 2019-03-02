package recognize

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	latexEndpoint = "https://api.mathpix.com/v3/latex"
)

type Recognizer interface {
	ToLatex(imageURL string) (string, error)
}

func New() Recognizer {
	id := os.Getenv("MATHPIX_APP_ID")
	key := os.Getenv("MATHPIX_APP_KEY")
	if len(id) == 0 || len(key) == 0 {
		panic("Must set env variables for mathpix")
	}
	return recognizerImpl{
		c: clientImpl{id, key},
	}
}

type client interface {
	post(url string, textBody string) ([]byte, error)
}

type clientImpl struct {
	appId  string
	appKey string
}

func (c clientImpl) post(url string, textBody string) ([]byte, error) {
	body := bytes.NewBufferString(textBody)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header["app_id"] = []string{c.appId}
	req.Header["app_key"] = []string{c.appKey}
	req.Header["Content-Type"] = []string{"application/json"}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(fmt.Sprintf("status code: %d", resp.StatusCode))
	}
	responseBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseBytes, nil
}

type recognizerImpl struct {
	c client
}

type latexResponse struct {
	DetectionList   []string     `json:"detection_list"`
	DetectionMap    detectionMap `json:"detection_map"`
	Error           string       `json:"error"`
	Latex           string       `json:"latex"`
	LatexConfidence float64      `json:"latex_confidence"`
	Position        position     `json:"position"`
}

type detectionMap struct {
	ContainsChart    float64 `json:"contains_chart"`
	ContainsDiagram  float64 `json:"contains_diagram"`
	ContainsGeometry float64 `json:"contains_geometry"`
	ContainsGraph    float64 `json:"contains_graph"`
	ContainsTable    float64 `json:"contains_table"`
	IsInverted       float64 `json:"is_inverted"`
	IsNotMath        float64 `json:"is_not_math"`
	IsPrinted        float64 `json:"is_printed"`
}

type position struct {
	Height   int `json:"height"`
	Width    int `json:"width"`
	TopLeftX int `json:"top_left_x"`
	TopLeftY int `json:"top_left_y"`
}

func (r recognizerImpl) ToLatex(imageURL string) (string, error) {
	bodyString := fmt.Sprintf("{"+
		"\"src\": \"%s\","+
		"\"ocr\": [\"math\", \"text\"]"+
		"}", imageURL)
	rawJSON, err := r.c.post(latexEndpoint, bodyString)
	if err != nil {
		return "", err
	}
	var resp latexResponse
	err = json.Unmarshal(rawJSON, &resp)
	if err != nil {
		return "", err
	}

	// TODO: Only return it if the confidence is high enough? Or also include the confidence or something

	return resp.Latex, nil
}
