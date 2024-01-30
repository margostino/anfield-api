package network

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/margostino/anfield-api/common"
)

var client = &http.Client{}

func Get(url string, responseData interface{}, wg *sync.WaitGroup, headers map[string]string) {
	req, _ := http.NewRequest("GET", url, nil)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	response, err := client.Do(req)
	common.Check(err)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	common.Check(err)

	err = json.Unmarshal(body, responseData)
	if err != nil {
		fmt.Println(string(body))
	}
	common.Check(err)
	if wg != nil {
		wg.Done()
	}
}
