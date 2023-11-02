package translate

import (
	tr "github.com/snakesel/libretranslate"
)

func Translate(msg interface{}, targetLang string) (string, error) {
	translate := tr.New(tr.Config{
		Url: "http://127.0.0.1:5000",
	})

	
	msgstr := msg.(string)

	trtext, err := translate.Translate(msgstr, "auto", targetLang)
	if err != nil {
		return "", err
	}

	return trtext, nil
}
