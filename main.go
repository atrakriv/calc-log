package main

import (
	"log"
	"net/http"
	"os"

	//"strconv"
	"bytes"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel
var cnt = 0
var msgs []string

// Configure the upgrader
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
type Message struct {
	//Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func getPort() string {
	p := os.Getenv("PORT")
	if p != "" {
		return ":" + p
	}
	// Start the server on localhost port 8080 and log any errors
	log.Println("http server started on :8080")
	return ":8080"
}

func main() {
	// Create a simple file server
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	go handleMessages()

	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	defer ws.Close()

	// Register our new client
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel

		if len(msgs) == 10 {
			msgs = msgs[1:10]
		}
		msgs = append(msgs, msg.Message)

		var s bytes.Buffer
		for i := 0; i < len(msgs); i++ {
			s.WriteString("</div>")
			s.WriteString("</div>")
			s.WriteString(msgs[i])
			s.WriteString("<br/>")
		}

		msg.Message = s.String()
		//msg += " : "
		//msg.Username = strconv.Itoa(cnt)
		//msg += strconv.Itoa(cnt)
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		msg := <-broadcast
		// Send it out to every client that is currently connected
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
