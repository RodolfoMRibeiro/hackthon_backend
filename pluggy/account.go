package pluggy

import (
	"encoding/json"
	"fmt"
	"hackthon/adapter/env"
	"hackthon/pluggy/dto"
	"io/ioutil"
	"net/http"
)

func ListAccounts(item_id string) dto.Accounts {
	url := "https://api.pluggy.ai/accounts?itemId=" + item_id

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	acc := dto.Accounts{}

	json.Unmarshal(body, &acc)

	return acc
}

func RetriveAccount(account_id string) dto.RetriveBankAccount {
	url := "https://api.pluggy.ai/accounts/" + account_id

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	retrive := dto.RetriveBankAccount{}

	json.Unmarshal(body, &retrive)

	fmt.Println(res)
	fmt.Println(string(body))
	return retrive
}
