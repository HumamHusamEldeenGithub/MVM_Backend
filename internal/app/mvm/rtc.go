package mvm

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"sync"

// 	"github.com/gorilla/websocket"
// 	"github.com/pion/webrtc/v3"
// )

// type Message struct {
// 	Type string          `json:"type"`
// 	Data json.RawMessage `json:"data"`
// }

// type OfferMessage struct {
// 	Description webrtc.SessionDescription `json:"description"`
// }

// type AnswerMessage struct {
// 	Description webrtc.SessionDescription `json:"description"`
// }

// type CandidateMessage struct {
// 	Candidate string `json:"candidate"`
// }

// type Connection struct {
// 	sync.Mutex
// 	peerConnection *webrtc.PeerConnection
// }

// var connections = make(map[string]*Connection)
// var upgrader = websocket.Upgrader{}

// func main() {
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		// Upgrade the HTTP connection to a WebSocket connection
// 		conn, err := upgrader.Upgrade(w, r, nil)
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		// Create a new connection object for the WebSocket connection
// 		connection := &Connection{}

// 		// Start a new Goroutine to handle incoming WebSocket messages
// 		go handleMessages(conn, connection)
// 	})

// 	// Start the HTTP server
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

// func handleMessages(conn *websocket.Conn, connection *Connection) {
// 	defer conn.Close()

// 	for {
// 		// Read a message from the WebSocket connection
// 		_, data, err := conn.ReadMessage()
// 		if err != nil {
// 			log.Println(err)
// 			return
// 		}

// 		// Parse the message type and data
// 		var message Message
// 		err = json.Unmarshal(data, &message)
// 		if err != nil {
// 			log.Println(err)
// 			continue
// 		}

// 		// Handle the message based on its type
// 		switch message.Type {
// 		case "offer":
// 			var offer OfferMessage
// 			err = json.Unmarshal(message.Data, &offer)
// 			if err != nil {
// 				log.Println(err)
// 				continue
// 			}
// 			handleOffer(conn, connection, offer.Description)
// 		case "answer":
// 			var answer AnswerMessage
// 			err = json.Unmarshal(message.Data, &answer)
// 			if err != nil {
// 				log.Println(err)
// 				continue
// 			}
// 			handleAnswer(conn, connection, answer.Description)
// 		case "candidate":
// 			var candidate CandidateMessage
// 			err = json.Unmarshal(message.Data, &candidate)
// 			if err != nil {
// 				log.Println(err)
// 				continue
// 			}
// 			handleCandidate(conn, connection, candidate.Candidate)
// 		default:
// 			log.Println("unknown message type:", message.Type)
// 		}
// 	}
// }

// func handleOffer(conn *websocket.Conn, connection *Connection, offer webrtc.SessionDescription) {
// 	connection.Lock()
// 	defer connection.Unlock()

// 	// Create a new WebRTC API object
// 	api := webrtc.NewAPI(webrtc.WithMediaEngine(&webrtc.MediaEngine{}))

// 	// Create a new WebRTC peer connection with the given configuration
// 	peerConnection, err := api.NewPeerConnection(webrtc.Configuration{})
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Set up a handler for when a remote track is added
// 	peerConnection.OnTrack(func(track *webrtc.TrackRemote, receiver *webrtc.RTPReceiver) {
// 		fmt.Printf("Remote track added: %s\n", track.ID())
// 	})

// 	// Set the remote description to the offer
// 	err = peerConnection.SetRemoteDescription(offer)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Generate an answer to the offer
// 	answer, err := peerConnection.CreateAnswer(nil)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Set the local description to the answer
// 	err = peerConnection.SetLocalDescription(answer)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Send the answer back to the client
// 	answerMessage := AnswerMessage{Description: answer}
// 	data, err := json.Marshal(answerMessage)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// 	message := Message{Type: "answer", Data: data}
// 	err = conn.WriteJSON(message)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Add the connection to the list of connections
// 	connections[conn.RemoteAddr().String()] = connection
// 	connection.peerConnection = peerConnection
// }

// func handleAnswer(conn *websocket.Conn, connection *Connection, answer webrtc.SessionDescription) {
// 	connection.Lock()
// 	defer connection.Unlock()
// 	// Set the remote description to the answer
// 	err := connection.peerConnection.SetRemoteDescription(answer)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// }

// func handleCandidate(conn *websocket.Conn, connection *Connection, candidate string) {
// 	connection.Lock()
// 	defer connection.Unlock()
// 	// Parse the candidate data
// 	var iceCandidate webrtc.ICECandidateInit
// 	err := json.Unmarshal([]byte(candidate), &iceCandidate)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}

// 	// Add the candidate to the peer connection
// 	err = connection.peerConnection.AddICECandidate(iceCandidate)
// 	if err != nil {
// 		log.Println(err)
// 		return
// 	}
// }
