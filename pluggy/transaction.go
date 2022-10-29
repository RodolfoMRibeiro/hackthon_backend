package pluggy

import (
	"encoding/json"
	"fmt"
	"hackthon/adapter/env"
	"hackthon/pluggy/dto"
	"io/ioutil"
	"net/http"
)

func GetTransactionList(account_id, inicialDate, endDate string) dto.TransactionData {
	var url string
	if inicialDate == "" || endDate == "" {
		url = "https://api.pluggy.ai/transactions?accountId=" + account_id
	} else {
		url = "https://api.pluggy.ai/transactions?accountId=" + account_id + "&from=" + inicialDate + "&to=" + endDate
	}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	transactionData := dto.TransactionData{}
	json.Unmarshal(body, &transactionData)

	fmt.Println(res)
	fmt.Println(string(body))
	return transactionData

}

func RetriveTransacion(bank_account_id string) dto.AccountOfTransaction {

	url := "https://api.pluggy.ai/transactions/" + bank_account_id

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("X-API-KEY", env.PuggyApi.API_KEY)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	transactionRetrival := dto.AccountOfTransaction{}

	json.Unmarshal(body, &transactionRetrival)

	fmt.Println(res)
	fmt.Println(string(body))
	return transactionRetrival

}
