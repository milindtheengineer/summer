package messaging

import (
	"encoding/json"
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

// IncomingMessage is the struct that comes from Telegram
type IncomingMessage struct {
	UpdateID int `json:"update_id"`
	Message  struct {
		MessageID int `json:"message_id"`
		From      struct {
			ID           int    `json:"id"`
			IsBot        bool   `json:"is_bot"`
			FirstName    string `json:"first_name"`
			LastName     string `json:"last_name"`
			LanguageCode string `json:"language_code"`
		} `json:"from"`
		Chat struct {
			ID        int    `json:"id"`
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Type      string `json:"type"`
		} `json:"chat"`
		Date int    `json:"date"`
		Text string `json:"text"`
	} `json:"message"`
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
func (t *Telegram) ProcessMessage(message []byte) (string, int) {
	incomingMessage := IncomingMessage{}
	if err := json.Unmarshal(message, &incomingMessage); err != nil {
		panic(err)
	}
	return incomingMessage.Message.Text, incomingMessage.Message.Chat.ID
}
