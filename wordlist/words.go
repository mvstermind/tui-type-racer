package wordlist

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error: %v", err)
		return "", err
	}

	bdy := strip(body)
	return string(bdy), nil
}

func strip(b []byte) string {
	body := string(b)
	body = strings.Replace(body, `"`, "", len(body))
	body = strings.Replace(body, `,`, " ", len(body))
	body = strings.TrimLeft(body, "[")
	body = strings.TrimRight(body, "]")
	return body

}
