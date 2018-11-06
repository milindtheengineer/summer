package brain

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"summer/brain/witai"
	"summer/messaging"

	"github.com/gorilla/mux"
)

func justHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Milind :)")
}

func messageProcessor(w http.ResponseWriter, r *http.Request) {
	t := messaging.Telegram{QueryURL: "https://api.telegram.org/bot", Token: os.Getenv("BOT_TOKEN")}
	messageBytes, _ := ioutil.ReadAll(r.Body)
	text, messageChatID := t.ProcessMessage(messageBytes)
	chatID, _ := strconv.Atoi(os.Getenv("CHAT_ID"))
	if messageChatID != chatID {
		t.SendMessage("This bot doesn't belong to you!", strconv.Itoa(messageChatID))
		return
	}
	responseMessage := witai.ExtractMessage(text)
	t.SendMessage(responseMessage, os.Getenv("CHAT_ID"))
}

//ServeRequests is kind of the main routing function
func ServeRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", justHello)
	r.HandleFunc(fmt.Sprintf("/%s", os.Getenv("BOT_TOKEN")), messageProcessor).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
