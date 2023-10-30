package translate

import (
	"fmt"

	tr "github.com/snakesel/libretranslate"
)

func Translate(msg interface{}, targetLang string) (string, error) {
	translate := tr.New(tr.Config{
		Url: "http://127.0.0.1:5000",
	})

	var msgString string

	// Check the type of the message and convert it to a string if needed
	switch t := msg.(type) {
	case string:
		msgString = t
	case []byte:
		msgString = string(t)
	default:
		return "", fmt.Errorf("Unsupported message data type: %T", msg)
	}

	trtext, err := translate.Translate(msgString, "auto", targetLang)
	if err != nil {
		return "", err
	}

	return trtext, nil
}

func TranslateMsg(msg string, targetLang string) string {
	
	translate := tr.New(tr.Config{
		Url: "http://127.0.0.1:5000",
		Key: "",
	})


	trtext, err := translate.Translate(msg, "auto", targetLang)
	if err != nil {
		return ""
	}

	return trtext
}