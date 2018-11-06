package witai

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"summer/nonthings/football"
)

func makeGetCall(message string) []byte {
	client := &http.Client{}
	url := fmt.Sprintf("https://api.wit.ai/message")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("q", message)
	req.URL.RawQuery = q.Encode()
	req.Header.Add("Authorization", os.Getenv("WIT_AI_ACCESS_TOKEN"))
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		panic(err)
	}

	defer resp.Body.Close()
	respBody, _ := ioutil.ReadAll(resp.Body)
	return respBody
}

// ExtractMessage extracts the meaning using wit ai
func ExtractMessage(message string) string {
	byt := makeGetCall(message)
	responseMessage := Response{}
	if err := json.Unmarshal(byt, &responseMessage); err != nil {
		panic(err)
	}
	if len(responseMessage.Entities.Question) > 0 && responseMessage.Entities.Question[0].Value == "football matches" {
		return football.SendMatches(responseMessage.Entities.Datetime[0].Value.Format("2006-01-02"))
	}
	return "No action for this message yet"
}
