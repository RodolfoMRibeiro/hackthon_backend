package pluggy

import (
	"encoding/json"
	"fmt"
	"hackthon/adapter/env"
	"hackthon/pluggy/dto"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateItem() dto.CreateItemUserDto {

	url := "https://api.pluggy.ai/items"

	payload := strings.NewReader("{\"parameters\":{\"user\":\"user-ok\",\"password\":\"password-ok\"},\"connectorId\":8}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	createItemUserDto := dto.CreateItemUserDto{}

	json.Unmarshal(body, &createItemUserDto)
	fmt.Println(createItemUserDto)
	return createItemUserDto
}

func DeleteItem(uuid string) {

	url := "https://api.pluggy.ai/items/" + uuid

	req, _ := http.NewRequest("DELETE", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
