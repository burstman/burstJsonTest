package jsontest

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//GetTest use http.get to import for json getstate burst api
func GetTest() (map[string]interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}
	var result map[string]interface{}

	return result, json.NewDecoder(resp.Body).Decode(&result)
}
