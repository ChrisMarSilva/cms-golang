package services

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

type Validation struct {
	BadWordsTotal int `json:"bad_words_total"`
	BadWordsList  []struct {
		Original string `json:"original"`
		Word     string `json:"word"`
	} `json:"bad_words_list"`
	Message string `json:"message"`
}

func (v *Validation) BadWords() string {
	ws := []string{}
	for _, w := range v.BadWordsList {
		ws = append(ws, w.Word)
	}

	return strings.Join(ws, ",")
}

func Validate(words string) (Validation, error) {
	var result Validation
	url := "https://api.apilayer.com/bad_words?censor_character=censor_character=*"
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(words))
	if err != nil {
		//return result, err
	}
	req.Header.Set("apikey", os.Getenv("BAD_WORDS_API_KEY"))

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		//return result, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		//return result, err
	}

	if len(result.Message) != 0 {
		//return result, fmt.Errorf("%s", result.Message)
	}

	return result, nil
}
