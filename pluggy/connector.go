package pluggy

import (
	"fmt"
	"hackthon/adapter/env"
	"io/ioutil"
	"net/http"
)

func ListConnectors(companyType string) {
	var url string

	if companyType != "" {
		url = "https://api.pluggy.ai/connectors?countries=BR&types=" + companyType + "&sandbox=true"
	} else {
		url = "https://api.pluggy.ai/connectors?countries=BR&sandbox=true"
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.ACCESS_TOKEN)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}

func RetriveConnectors(connector_id string) {

	url := "https://api.pluggy.ai/connectors/" + connector_id

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.ACCESS_TOKEN)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
