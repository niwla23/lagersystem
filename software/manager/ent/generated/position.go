// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
	"github.com/niwla23/lagersystem/manager/ent/generated/warehouse"
)

// Position is the model entity for the Position schema.
type Position struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// PositionId holds the value of the "positionId" field.
	PositionId int `json:"positionId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PositionQuery when eager-loading is set.
	Edges               PositionEdges `json:"-"`
	box_position        *int
	warehouse_positions *int
}

// PositionEdges holds the relations/edges for other nodes in the graph.
type PositionEdges struct {
	// StoredBox holds the value of the storedBox edge.
	StoredBox *Box `json:"storedBox,omitempty"`
	// Warehouse holds the value of the warehouse edge.
	Warehouse *Warehouse `json:"warehouse,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// StoredBoxOrErr returns the StoredBox value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PositionEdges) StoredBoxOrErr() (*Box, error) {
	if e.loadedTypes[0] {
		if e.StoredBox == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: box.Label}
		}
		return e.StoredBox, nil
	}
	return nil, &NotLoadedError{edge: "storedBox"}
}

// WarehouseOrErr returns the Warehouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PositionEdges) WarehouseOrErr() (*Warehouse, error) {
	if e.loadedTypes[1] {
		if e.Warehouse == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: warehouse.Label}
		}
		return e.Warehouse, nil
	}
	return nil, &NotLoadedError{edge: "warehouse"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Position) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case position.FieldID, position.FieldPositionId:
			values[i] = new(sql.NullInt64)
		case position.FieldCreatedAt, position.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case position.ForeignKeys[0]: // box_position
			values[i] = new(sql.NullInt64)
		case position.ForeignKeys[1]: // warehouse_positions
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Position", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Position fields.
func (po *Position) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case position.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case position.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				po.CreatedAt = value.Time
			}
		case position.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				po.UpdatedAt = value.Time
			}
		case position.FieldPositionId:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field positionId", values[i])
			} else if value.Valid {
				po.PositionId = int(value.Int64)
			}
		case position.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field box_position", value)
			} else if value.Valid {
				po.box_position = new(int)
				*po.box_position = int(value.Int64)
			}
		case position.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field warehouse_positions", value)
			} else if value.Valid {
				po.warehouse_positions = new(int)
				*po.warehouse_positions = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryStoredBox queries the "storedBox" edge of the Position entity.
func (po *Position) QueryStoredBox() *BoxQuery {
	return NewPositionClient(po.config).QueryStoredBox(po)
}

// QueryWarehouse queries the "warehouse" edge of the Position entity.
func (po *Position) QueryWarehouse() *WarehouseQuery {
	return NewPositionClient(po.config).QueryWarehouse(po)
}

// Update returns a builder for updating this Position.
// Note that you need to call Position.Unwrap() before calling this method if this Position
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Position) Update() *PositionUpdateOne {
	return NewPositionClient(po.config).UpdateOne(po)
}

// Unwrap unwraps the Position entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Position) Unwrap() *Position {
	_tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("generated: Position is not a transactional entity")
	}
	po.config.driver = _tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Position) String() string {
	var builder strings.Builder
	builder.WriteString("Position(")
	builder.WriteString(fmt.Sprintf("id=%v, ", po.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(po.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(po.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("positionId=")
	builder.WriteString(fmt.Sprintf("%v", po.PositionId))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (po *Position) MarshalJSON() ([]byte, error) {
	type Alias Position
	return json.Marshal(&struct {
		*Alias
		PositionEdges
	}{
		Alias:         (*Alias)(po),
		PositionEdges: po.Edges,
	})
}

// Positions is a parsable slice of Position.
type Positions []*Position

func (po Positions) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}