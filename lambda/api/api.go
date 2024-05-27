package api

import (
	"fmt"
	"lambda-func/database"
	"lambda-func/types"
)

type ApiHandler struct {
	dbStore database.DynamoDBClient
}

func NewApiHandler(dbStore database.DynamoDBClient) ApiHandler {
	return ApiHandler{
		dbStore: dbStore,
	}
}

func (api ApiHandler) RegisterUserHandler(event types.RegisterUser) error {
	if event.Username == "" || event.Password == "" {
		return fmt.Errorf("username or password is empty")
	}

	exists, err := api.dbStore.DoesUserExist(event.Username)
	if err != nil {
		return fmt.Errorf("username check returned an error %w", err)
	}

	if exists {
		return fmt.Errorf("username already exists")
	}

	err = api.dbStore.InsertUser(event)
	if err != nil {
		return fmt.Errorf("error registering the user %w", err)
	}

	return nil
}
