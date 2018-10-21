package brain

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"summer/messaging"

	"github.com/gorilla/mux"
)

func justHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from Milind :)")
}

func messageProcessor(w http.ResponseWriter, r *http.Request) {
	t := messaging.Telegram{QueryURL: "https://api.telegram.org/bot", Token: os.Getenv("BOT_TOKEN")}
	t.SendMessage("Hello, we are running", os.Getenv("CHAT_ID"))
}

//ServeRequests is kind of the main routing function
func ServeRequests() {
	r := mux.NewRouter()
	r.HandleFunc("/", justHello)
	r.HandleFunc(fmt.Sprintf("/%s", os.Getenv("BOT_TOKEN")), messageProcessor).Methods("POST")
	log.Fatal(http.ListenAndServe(":8080", r))
}
