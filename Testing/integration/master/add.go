package master

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Sum(i, j int) int {
	return i + j
}

type MathClient struct {
	Token string
	Host  string
}

type AddResult struct {
	Result int `json:"result"`
}

func (c *MathClient) APISum(i, j int) (int, error) {
	query := fmt.Sprintf("http://%s/add?a=%v&b=%v&authtoken=%v", c.Host, i, j, c.Token)

	response, err := http.Get(query)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	a := AddResult{}
	err = json.Unmarshal(data, &a)
	if err != nil {
		return 0, err
	}

	return a.Result, nil
}
