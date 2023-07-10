package helpers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/database"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
)

type scanBoxIdResponse struct {
	Status   string    `json:"status"`
	BoxId    uuid.UUID `json:"boxId"`
	Duration float64   `json:"duration"`
}

var OPERATOR_BASE_URL = "http://lagersystem:5000"

// returns boxX, boxId, error
func ScanIoPos(ioPos string) (*ent.Box, *uuid.UUID, error) {
	ctx := context.Background()
	// get boxId from operator service
	resp, err := http.Get(OPERATOR_BASE_URL + "/scanBoxId/" + ioPos)
	if err != nil {
		return nil, nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}
	data := new(scanBoxIdResponse)
	json.Unmarshal(body, data)

	if data.BoxId == uuid.Nil {
		return nil, nil, errors.New("cant scan code")
	}

	// get boxX from db
	boxX, err := database.Client.Box.Query().Where(box.ID(data.BoxId)).WithPosition().WithParts().Only(ctx)

	// if we have an error
	if err != nil {
		// check if the error is a NotFoundError
		target := &ent.NotFoundError{}
		if errors.As(err, &target) {
			// if so, create the box instead
			newBox, err := database.Client.Box.Create().SetID(data.BoxId).Save(ctx)
			if err == nil {
				// that succeeded? congratulations, return it
				return newBox, &data.BoxId, nil
			}
		}
		// you fucked something up
		return nil, &data.BoxId, err
	}
	return boxX, &data.BoxId, err
}

type DeliverBoxResponse struct {
	Status     string  `json:"status"`
	PositionId int     `json:"positionId"`
	Duration   float64 `json:"duration"`
}

func PickupBox(positionId int) (*DeliverBoxResponse, error) {
	resp, err := http.Get(fmt.Sprintf(OPERATOR_BASE_URL+"/pickupBox/%v/%v", positionId, "0"))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New("something went wrong")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(DeliverBoxResponse)
	json.Unmarshal(body, data)
	return data, nil
}

func StoreBox(positionId int, ioPosId string) (*DeliverBoxResponse, error) {
	resp, err := http.Get(fmt.Sprintf(OPERATOR_BASE_URL+"/storeBox/%v/%v", positionId, ioPosId))
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(DeliverBoxResponse)
	json.Unmarshal(body, data)
	return data, nil
}

type IOState = map[string]string
type IOStateResponse struct {
	IOState IOState `json:"ioState"`
}

func GetIOState() (*IOState, error) {
	resp, err := http.Get(OPERATOR_BASE_URL + "/getIOState")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(IOStateResponse)
	json.Unmarshal(body, data)
	return &data.IOState, nil
}

func IsIoSlotFree(ioState *IOState, ioPos string) bool {
	return (*ioState)[ioPos] == "free"
}

func FindIoSlot() (string, error) {
	ioState, err := GetIOState()
	if err != nil {
		return "", err
	}
	if IsIoSlotFree(ioState, "3") {
		return "3", nil
	}
	if IsIoSlotFree(ioState, "2") {
		return "2", nil
	}
	if IsIoSlotFree(ioState, "1") {
		return "1", nil
	}
	return "", errors.New("no free io slot found")
}

type GetPositionsResponse struct {
	Positions map[string]struct {
		X     float64 `json:"x"`
		Y     float64 `json:"y"`
		BoxId string  `json:"boxId"`
	} `json:"positions"`
}

func GetPositions(client *ent.Client) (*GetPositionsResponse, error) {
	ctx := context.Background()

	resp, err := http.Get(OPERATOR_BASE_URL + "/getPositions")
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	data := new(GetPositionsResponse)
	json.Unmarshal(body, data)

	for posId := range data.Positions {
		posIdInt, err := strconv.Atoi(posId)
		if err != nil {
			continue
		}

		positionDb, err := client.Position.Query().Where(position.ID(posIdInt)).WithStoredBox().Only(ctx)
		if err != nil {
			continue
		}

		if positionDb.Edges.StoredBox != nil {

			if entry, ok := data.Positions[posId]; ok {
				entry.BoxId = positionDb.Edges.StoredBox.ID.String()
				data.Positions[posId] = entry
			}
		}
	}
	return data, nil
}
