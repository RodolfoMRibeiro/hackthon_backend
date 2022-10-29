package pluggy

import (
	"fmt"
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

	url := "https://api.pluggy.ai/connect_token"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add("X-API-KEY", "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjoiMzAwNDRjNmU1MGMyODg1ZjJhOTZkNGNhYzlhMjQzNTE6N2ViZmQ4NDUwMjlhMzg0ZTM2YzkyYjYwNWRiZTY5MmYxNTYyMWQ4OTkzZDg0MDc2YWQ0NjlkNTUwZjgyODg5NTFjYTcwMWU1OTM3ZTAwMTc5Y2ExODE0NjU2N2ExNWUzMjE5N2E0MTIwNjkyZmFlODAyNjI1YTgyYWUwODM0ZTUwZDNlZjU1MzY5MTk5ZDViZTMwY2IwNTFjOTliY2EyMyIsImlhdCI6MTY2NzA1OTEwNiwiZXhwIjoxNjY3MDY2MzA2fQ.dC7jRZ1i4QPEN-VOGSgtWgY0bc9AzWiUQKjLTlz6Z-ugIPkgGvvWTY5upl8KiPZzL9018dtfKK7V24E5tL2wS3Y1tkugySRn6xpgAz11fTreXvOfq9fCWTfAvnvpBvEJaC6In3F1tdnxeMpxo4hyCJhNwFhAIL0QOuOYdu3dR4ElYMnCNl8aB-kZJd8p0PeTdkk0i7-CLVRxfQbur0Ks7kGgGcLnxPd3EVx2JtuuDwhdubu3WiyQ31-S3SBIhtt4_a88odli7HIYMJHJjNHPSgY5autYLk5VvHaiOreZZUr0gTCpJa4xvmxloom6TydGxxO4vh1m0_57maO0Vv6-Aw")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
