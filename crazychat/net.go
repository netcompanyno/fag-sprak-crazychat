package crazychat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/mesan/fag-sprak-crazychat/crazychat/rx"
)

func ListenForMessages(port string, ch chan rx.Message) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "Unknown method", http.StatusMethodNotAllowed)
			fmt.Printf(ColRed+"Unknown method %s\n"+ColReset, r.Method)
			return
		}

		if !strings.HasPrefix(r.Header.Get("Content-Type"), "application/json") {
			http.Error(w, "Unsupported media type", http.StatusUnsupportedMediaType)
			fmt.Printf(ColRed+"Unsupported media type %s\n"+ColReset, r.Header.Get("Content-Type"))
			return
		}

		msg := rx.Message{}
		err := json.NewDecoder(r.Body).Decode(&msg)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			fmt.Printf(ColRed + "Invalid JSON\n" + ColReset)
			return
		}

		fmt.Fprintf(w, "%#v", msg)

		ch <- msg
	})
	go http.ListenAndServe(":"+port, nil)
}

func SendMessage(msg rx.Message, address string) {
	buf := bytes.Buffer{}
	json.NewEncoder(&buf).Encode(msg)
	_, err := http.Post(address, "application/json", &buf)
	if err != nil {
		fmt.Println(ColRed + err.Error() + ColReset)
	}
}
