package messaging

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Telegram is the struct with important fields that are required for messaging
type Telegram struct {
	QueryURL string
	Token    string
}

// SendMessage will send the message to the user
func (t *Telegram) SendMessage(message string, userID string) {
	client := &http.Client{}
	url := fmt.Sprintf("%s%s/sendMessage", t.QueryURL, t.Token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("chat_id", userID)
	q.Add("text", message)
	req.URL.RawQuery = q.Encode()

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Errored when sending request to the server")
		panic(err)
	}

	defer resp.Body.Close()
	//respBody, _ := ioutil.ReadAll(resp.Body)
	//return respBody
	//Need to write the responseCode and ResponseBodyLogic
}

// ProcessMessage will process the json and return the message and chatID
func (t *Telegram) ProcessMessage(message string) {
	// Yet to write
}
