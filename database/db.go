package db

import (
	"github.com/go-resty/resty/v2"
)

type payload struct {
	Limit int `json:"limit"`
	With_payload bool `json:"with_payload"`
	Score_threshold float32 `json:"score_threshold"`
	Vector []float64 `json:"vector"`
}

func SearchPoint(vector []float64) (string, error) {
	load := payload{
		Limit: 1,
		With_payload: true,
		Score_threshold: 0.7,
		Vector: vector,
	}
	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(load).
		Post("http://localhost:6333/collections/faces/points/search")
	if err != nil {
		return "", err
	}
	return resp.String(), nil
}