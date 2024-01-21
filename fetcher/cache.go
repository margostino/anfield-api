package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/margostino/anfield-api/common"
)

var datasetCache = make(map[string]DatasetIndex)
var indexCache = loadIndex()

func loadIndex() map[string]string {
	var urls = make(map[string]string)
	metadataUrl := os.Getenv("METADATA_URL")
	datasetUrls := fmt.Sprintf("%s/datasets.yml", metadataUrl)
	resp, err := http.Get(datasetUrls)
	defer resp.Body.Close()
	common.Check(err)
	bodyBytes, err := io.ReadAll(resp.Body)
	common.UnmarshalYamlBytes(bodyBytes, &urls)
	return urls
}
