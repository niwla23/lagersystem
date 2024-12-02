package helpers

import (
	"context"
	"errors"

	"github.com/niwla23/lagersystem/manager/database"
	ent "github.com/niwla23/lagersystem/manager/ent/generated"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
	"github.com/niwla23/lagersystem/manager/ent/generated/warehouse"
)

func FindPosition(boxX *ent.Box) (*ent.Position, error) {
	ctx := context.Background()
	// check if a position is already stored for this box
	positionX, err := boxX.QueryPosition().Only(ctx)

	target := &ent.NotFoundError{}
	if errors.As(err, &target) {
		// we have no position stored, find one
		positionX, err = database.Client.Position.Query().
			Where(position.HasWarehouseWith(warehouse.ID(1))).
			Where(position.Not(position.HasStoredBox())).
			First(ctx)
	} else if err != nil {
		// shit has gone terribly wrong
		return nil, err
	}

	return positionX, err
}
