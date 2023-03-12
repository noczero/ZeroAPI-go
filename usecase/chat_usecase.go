package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/noczero/ZeroAPI-go/domain/model"
	"log"
	"net/http"

	//"github.com/noczero/ZeroAPI-go/domain/web"
	//"net/http"
	"time"
)

type chatUsecase struct {
	chatRepository model.ChatRepository
	contextTimeout time.Duration
}

func NewChatUsecase(chatRepository model.ChatRepository, contextTimeout time.Duration) model.ChatUsecase {
	return &chatUsecase{chatRepository: chatRepository, contextTimeout: contextTimeout}
}

func (cu chatUsecase) Create(c context.Context, chat *model.Chat) error {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chatRepository.Create(ctx, chat)
}

func (cu chatUsecase) FetchByUserID(c context.Context, userID string) ([]model.Chat, error) {
	ctx, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	return cu.chatRepository.FetchByUserID(ctx, userID)
}

func (cu chatUsecase) GetResponseFromOpenAI(c context.Context, prompt string, userID string, token string) (string, error) {
	//var chatGPTResponse web.ChatGPTResponse

	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
	}
	log.Println(requestBody)

	// Convert the JSON body to a byte array.
	requestBodyBytes, err := json.Marshal(requestBody)
	if err != nil {
		return "", nil
	}

	fmt.Println(bytes.NewBuffer(requestBodyBytes))

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return "", err
	}

	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	log.Println(token)

	// Create an HTTP client to send the request and receive the response.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	fmt.Println(resp.Body)
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	log.Println(data)

	//log.Println(data["choices"])
	choices, ok := data["choices"].([]interface{})
	log.Println(choices)
	if ok {
		insideChoice, ok := choices[0].(map[string]interface{})
		log.Println(insideChoice)
		log.Println(ok)

		if ok {
			message, ok := insideChoice["message"].(map[string]interface{})
			log.Println(message)
			log.Println(ok)

			if ok {
				content, ok := message["content"].(string)
				fmt.Println(content)
				fmt.Println(ok)

				return content, nil
			}
		}
	}

	return "", nil
}
