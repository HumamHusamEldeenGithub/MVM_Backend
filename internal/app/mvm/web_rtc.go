package mvm

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/pion/webrtc/v3"
)

// Define a struct to hold the WebRTC configuration
type Config struct {
	ICEServers []webrtc.ICEServer `json:"iceServers"`
}

// Define a struct to hold the WebRTC session data
type SessionData struct {
	Offer  webrtc.SessionDescription `json:"offer"`
	Answer webrtc.SessionDescription `json:"answer"`
}

// Define a struct to hold the WebRTC data channel message
type DataChannelMessage struct {
	Data string `json:"data"`
}

func (s *MVMServiceServer) HandleOffer(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Parse the WebRTC configuration
	var config Config
	err = json.Unmarshal(body, &config)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create a new WebRTC API instance
	api := webrtc.NewAPI(webrtc.WithMediaEngine(&webrtc.MediaEngine{}))

	// Create a new WebRTC peer connection
	peerConnection, err := api.NewPeerConnection(webrtc.Configuration{
		ICEServers: config.ICEServers,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create a data channel for sending and receiving raw data
	dataChannel, err := peerConnection.CreateDataChannel("raw-data", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Register the data channel message handler
	dataChannel.OnMessage(func(msg webrtc.DataChannelMessage) {
		var data DataChannelMessage
		err := json.Unmarshal(msg.Data, &data)
		if err != nil {
			log.Println("Failed to unmarshal data channel message:", err)
			return
		}

		log.Println("Received data channel message:", data.Data)
	})

	// Create an offer and set it as the local description
	offer, err := peerConnection.CreateOffer(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = peerConnection.SetLocalDescription(offer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the offer in the response
	sessionData := SessionData{
		Offer: offer,
	}
	response, err := json.Marshal(sessionData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

	// Wait for the answer to be received
	answer := webrtc.SessionDescription{}
	err = json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the answer as the remote description
	err = peerConnection.SetRemoteDescription(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Wait for the ICE connection state to change to connected
	connectionState := peerConnection.ConnectionState()
	for connectionState != webrtc.PeerConnectionStateConnected {
		time.Sleep(time.Second)
		connectionState = peerConnection.ConnectionState()
	}

	log.Println("Peer connection is connected")

	// Wait for the data channel to open
	dataChannelState := dataChannel.ReadyState()
	for dataChannelState != webrtc.DataChannelStateOpen {
		time.Sleep(time.Second)
		dataChannelState = dataChannel.ReadyState()
	}

	log.Println("Data channel is open")

	// Send a message on the data channel
	message := DataChannelMessage{
		Data: "Hello, world!",
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Failed to marshal data channel message:", err)
		return
	}
	err = dataChannel.Send(messageBytes)
	if err != nil {
		log.Println("Failed to send data channel message:", err)
		return
	}

	log.Println("Sent data channel message")

	// Wait for the data channel to close
	for dataChannel.ReadyState() != webrtc.DataChannelStateClosed {
		time.Sleep(time.Second)
	}

	log.Println("Data channel is closed")

	// Close the peer connection
	err = peerConnection.Close()
	if err != nil {
		log.Println("Failed to close peer connection:", err)
		return
	}

	log.Println("Peer connection is closed")
}

func (s *MVMServiceServer) HandleAnswer(w http.ResponseWriter, r *http.Request) {
	// Read the answer from the request body
	var sessionData SessionData
	err := json.NewDecoder(r.Body).Decode(&sessionData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Get the peer connection from the WebRTC API
	api := webrtc.NewAPI(webrtc.WithMediaEngine(&webrtc.MediaEngine{}))
	peerConnection, err := api.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the remote description as the offer
	err = peerConnection.SetRemoteDescription(sessionData.Offer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Create an answer and set it as the local description
	answer, err := peerConnection.CreateAnswer(nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = peerConnection.SetLocalDescription(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Send the answer in the response
	sessionData.Answer = answer
	response, err := json.Marshal(sessionData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (s *MVMServiceServer) HandleWebSocket(w http.ResponseWriter, r *http.Request) {
	// Upgrade the HTTP connection to a WebSocket connection
	upgrader := websocket.Upgrader{}
	_, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection to WebSocket:", err)
		return
	}
	// Create a new WebRTC API instance
	api := webrtc.NewAPI(webrtc.WithMediaEngine(&webrtc.MediaEngine{}))

	// Create a new WebRTC peer connection
	peerConnection, err := api.NewPeerConnection(webrtc.Configuration{})
	if err != nil {
		log.Println("Failed to create peer connection:", err)
		return
	}

	// Create a new data channel on the peer connection
	dataChannel, err := peerConnection.CreateDataChannel("data", nil)
	if err != nil {
		log.Println("Failed to create data channel:", err)
		return
	}

	// Wait for the data channel to open
	dataChannelState := dataChannel.ReadyState()
	for dataChannelState != webrtc.DataChannelStateOpen {
		time.Sleep(time.Second)
		dataChannelState = dataChannel.ReadyState()
	}

	log.Println("Data channel is open")

	// Set the handler for data channel messages
	dataChannel.OnMessage(func(message webrtc.DataChannelMessage) {
		// Unmarshal the message data into a DataChannelMessage struct
		var messageData DataChannelMessage
		err := json.Unmarshal(message.Data, &messageData)
		if err != nil {
			log.Println("Failed to unmarshal data channel message:", err)
			return
		}

		// Print the message data to the console
		log.Println("Received data channel message:", messageData.Data)

		// Send a reply on the data channel
		reply := DataChannelMessage{
			Data: "Received message: " + messageData.Data,
		}
		replyBytes, err := json.Marshal(reply)
		if err != nil {
			log.Println("Failed to marshal data channel message:", err)
			return
		}
		err = dataChannel.Send(replyBytes)
		if err != nil {
			log.Println("Failed to send data channel message:", err)
			return
		}

		log.Println("Sent data channel reply")
	})

	// Send a message on the data channel
	message := DataChannelMessage{
		Data: "Hello, world!",
	}
	messageBytes, err := json.Marshal(message)
	if err != nil {
		log.Println("Failed to marshal data channel message:", err)
		return
	}
	err = dataChannel.Send(messageBytes)
	if err != nil {
		log.Println("Failed to send data channel message:", err)
		return
	}

	log.Println("Sent data channel message")

	// Wait for the data channel to close
	for dataChannel.ReadyState() != webrtc.DataChannelStateClosed {
		time.Sleep(time.Second)
	}

	log.Println("Data channel is closed")

	// Close the peer connection
	err = peerConnection.Close()
	if err != nil {
		log.Println("Failed to close peer connection:", err)
		return
	}

	log.Println("Peer connection is closed")
}
