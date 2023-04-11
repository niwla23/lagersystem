package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/database"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
)

type scanBoxIdResponse struct {
	Status   string    `json:"status"`
	BoxId    uuid.UUID `json:"boxId"`
	Duration float64   `json:"duration"`
}

func GetBoxFromScanner() (*ent.Box, error) {
	ctx := context.Background()
	// get boxId from operator service
	resp, err := http.Get("http://localhost:3000/scanBoxId")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(scanBoxIdResponse)
	json.Unmarshal(body, data)

	// get boxX from db
	boxX, err := database.Client.Box.Query().Where(box.BoxId(data.BoxId)).WithPosition().WithSections().Only(ctx)
	if err != nil {
		return nil, err
	}

	return boxX, err
}

type deliverBoxResponse struct {
	Status     string  `json:"status"`
	PositionId int     `json:"positionId"`
	Duration   float64 `json:"duration"`
}

func DeliverBoxByPositionId(positionId int) (*deliverBoxResponse, error) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/deliver/%v", positionId))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(deliverBoxResponse)
	json.Unmarshal(body, data)
	return data, nil
}
