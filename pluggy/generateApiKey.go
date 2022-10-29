package pluggy

import (
	"encoding/json"
	"fmt"
	"hackthon/adapter/env"
	"io/ioutil"
	"net/http"
	"strings"
)

func GenerateApiKey() {

	url := "https://api.pluggy.ai/auth"

	payload := strings.NewReader("{\"clientId\":\"a46d25ed-25b4-4188-b8b1-4cecb935d84d\",\"clientSecret\":\"2199a9da-ccfd-4282-958c-acc0558b87ab\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func GenerateAccessToken() {
	tokener := AccessToken{}
	url := "https://api.pluggy.ai/connect_token"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(body, &tokener)

	env.PuggyApi.ACCESS_TOKEN = tokener.AccessToken

	fmt.Println(res)
	fmt.Println(string(body))

}
