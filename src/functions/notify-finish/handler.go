package main

import (
	"context"
	"encoding/json"
	"fmt"

	"poll-to-push/src/internal/task"
)

type Payload struct {
	TaskID string      `json:"taskID"`
	Result task.Result `json:"result"`
}

type Handler func(ctx context.Context, payload Payload) error

type ConnectionRetriever interface {
	Retrieve(ctx context.Context, ID string) (*task.Connection, error)
}

type Sender interface {
	Send(ctx context.Context, ID string, data []byte) error
}

func NewHandler(retriever ConnectionRetriever, sender Sender) Handler {
	return func(ctx context.Context, payload Payload) error {
		fmt.Println("retrieving taskID", payload.TaskID)
		conn, err := retriever.Retrieve(ctx, payload.TaskID)
		if err != nil {
			return err
		}

		if conn.ID == "" || conn.TaskID == "" {
			fmt.Println("client is not waiting for the answer, returning")
			return nil
		}
		buf, err := json.Marshal(payload.Result)
		err = sender.Send(ctx, conn.ID, buf)
		if err != nil {
			return err
		}

		return nil
	}
}
