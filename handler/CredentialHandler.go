package handler

import (
	"fmt"

	"github.com/google/uuid"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/store"
)

// GenerateCredential generates credential. return error, clientId, clientSecret
func GenerateCredential(clientStore *store.ClientStore) (string, string, error) {
	clientId := uuid.New().String()
	clientSecret := uuid.New().String()
	err := clientStore.Set(clientId, &models.Client{
		ID:     clientId,
		Secret: clientSecret,
		Domain: "http://localhost:9094",
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return clientId, clientSecret, err
}
