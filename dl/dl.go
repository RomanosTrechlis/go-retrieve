package dl

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func Download(url string, tokenEnvVar string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	if tokenEnvVar != "" {
		req.Header.Set("Authorization", fmt.Sprintf("token %s", os.Getenv(tokenEnvVar)))
	}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}
