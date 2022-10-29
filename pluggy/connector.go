package pluggy

import (
	"encoding/json"
	"fmt"
	"hackthon/adapter/env"
	"hackthon/pluggy/dto"
	"io/ioutil"
	"net/http"
)

func ListConnectors(companyType string) dto.CreateItemUserDto {
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

	connectorsList := dto.CreateItemUserDto{}

	json.Unmarshal(body, &connectorsList)

	fmt.Println(res)
	fmt.Println(string(body))
	return connectorsList
}

func RetriveConnectors(connector_id string) dto.UserItem {

	url := "https://api.pluggy.ai/connectors/" + connector_id

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	connector := dto.UserItem{}

	json.Unmarshal(body, &connector)

	fmt.Println(res)
	fmt.Println(string(body))
	return connector
}
