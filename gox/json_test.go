package gox

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var req int
	jsons, _ := json.Marshal(req)
	var servers int
	reqs, _ := http.NewRequest("POST", "", bytes.NewReader(jsons))
	resp, _ := http.DefaultClient.Do(reqs)
	defer resp.Body.Close()
	// -----------------------------------------------------------
	body, _ := ioutil.ReadAll(resp.Body)
	_ = json.Unmarshal(body, &servers)
}

func TestDecoder(t *testing.T) {
	var req int
	jsons, _ := json.Marshal(req)
	var servers int
	reqs, _ := http.NewRequest("POST", "", bytes.NewReader(jsons))
	resp, _ := http.DefaultClient.Do(reqs)
	defer resp.Body.Close()
	// -----------------------------------------------------------
	_ = json.NewDecoder(resp.Body).Decode(&servers)

}
