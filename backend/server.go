package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// These structures match OpenAI's format, Vllm accepts requests in open ai format
type Message struct { 
	Role	string `json:"role"`
	Content	string `json:"content"`
}

type ChatRequest struct {
	Model	string	`json:"model"`
	Messages	[]Message	`json:"messages"`
	MaxTokens	int	`json:"max_tokens"`
}

type Choice struct {
	Message Message `json:"message"`
}


type ChatResponse struct {
	Choices []Choice `json:"choices"`
}
// holds colections of Choice structs
/*
{
  "choices": [
    {"text": "Hello!"},
    {"text": "How can I help?"}
  ]
}
*/

//http://localhost:8091/process-transcription
//visit this end point trigger the send to vllm
func processTranscriptionHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintln(w, "hello World")


	// The url in which you can make a request to your graphics card, VLLM defaults to port 8000 and uses OpenAI formatting for the API enpoints
	url := "http://localhost:8000/v1/chat/completions"

	//this is the payload the endpoint is expecting
	payload := ChatRequest{
		Model: "Qwen/Qwen2.5-14B-Instruct-AWQ",
		Messages: []Message{
			{Role: "user", Content: "Hello from Go! Are you working properly?"},
		},
		MaxTokens: 150,
	}

	//Marshal converts Go data structures into into a JSON formatted byte array []byte
	jsonData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	//Make the request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer not-needed") // vllm doesnt need this

	//Send the reuest to the HTTP client (the graphics card)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	//postpones execution until surronding function returns
	defer resp.Body.Close()


	//Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	//Unmarshal converts a byte array into a Go data structure
	var chatResp ChatResponse
	if err := json.Unmarshal(body, &chatResp); err != nil {
		panic(err)
	}

	if len(chatResp.Choices) > 0 {
		fmt.Println("Qwen response")
		fmt.Println(chatResp.Choices[0].Message.Content)
	} else {
		fmt.Println("Raw Response:", string(body))
	}

	fmt.Fprintln(w, string(chatResp.Choices[0].Message.Content))

}


//will this need to listen on a port?
func main() {

	http.HandleFunc("/process-transcription", processTranscriptionHandler)

	log.Println("Server started on port 8091")

	err := http.ListenAndServe(":8091", nil)
	if err != nil {
		log.Fatalf("Server faield to start!")
	}



	
}





/*
1. Start remote server, pick and download a small llm

2. host that llm via vllm, requires a vllm install on remote server

3. Use this backend to send a request to the remote server, contact the llm, get a response

*/