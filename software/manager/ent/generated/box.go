// Code generated by ent, DO NOT EDIT.

package generated

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
)

// Box is the model entity for the Box schema.
type Box struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// BoxId holds the value of the "boxId" field.
	BoxId uuid.UUID `json:"boxId,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BoxQuery when eager-loading is set.
	Edges BoxEdges `json:"-"`
}

// BoxEdges holds the relations/edges for other nodes in the graph.
type BoxEdges struct {
	// Parts holds the value of the parts edge.
	Parts []*Part `json:"parts,omitempty"`
	// Position holds the value of the position edge.
	Position *Position `json:"position"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PartsOrErr returns the Parts value or an error if the edge
// was not loaded in eager-loading.
func (e BoxEdges) PartsOrErr() ([]*Part, error) {
	if e.loadedTypes[0] {
		return e.Parts, nil
	}
	return nil, &NotLoadedError{edge: "parts"}
}

// PositionOrErr returns the Position value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BoxEdges) PositionOrErr() (*Position, error) {
	if e.loadedTypes[1] {
		if e.Position == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: position.Label}
		}
		return e.Position, nil
	}
	return nil, &NotLoadedError{edge: "position"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Box) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case box.FieldID:
			values[i] = new(sql.NullInt64)
		case box.FieldCreatedAt, box.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case box.FieldBoxId:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Box", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Box fields.
func (b *Box) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case box.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			b.ID = int(value.Int64)
		case box.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				b.CreatedAt = value.Time
			}
		case box.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case box.FieldBoxId:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field boxId", values[i])
			} else if value != nil {
				b.BoxId = *value
			}
		}
	}
	return nil
}

// QueryParts queries the "parts" edge of the Box entity.
func (b *Box) QueryParts() *PartQuery {
	return NewBoxClient(b.config).QueryParts(b)
}

// QueryPosition queries the "position" edge of the Box entity.
func (b *Box) QueryPosition() *PositionQuery {
	return NewBoxClient(b.config).QueryPosition(b)
}

// Update returns a builder for updating this Box.
// Note that you need to call Box.Unwrap() before calling this method if this Box
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Box) Update() *BoxUpdateOne {
	return NewBoxClient(b.config).UpdateOne(b)
}

// Unwrap unwraps the Box entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Box) Unwrap() *Box {
	_tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("generated: Box is not a transactional entity")
	}
	b.config.driver = _tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Box) String() string {
	var builder strings.Builder
	builder.WriteString("Box(")
	builder.WriteString(fmt.Sprintf("id=%v, ", b.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(b.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("boxId=")
	builder.WriteString(fmt.Sprintf("%v", b.BoxId))
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (b *Box) MarshalJSON() ([]byte, error) {
	type Alias Box
	return json.Marshal(&struct {
		*Alias
		BoxEdges
	}{
		Alias:    (*Alias)(b),
		BoxEdges: b.Edges,
	})
}

// Boxes is a parsable slice of Box.
type Boxes []*Box

func (b Boxes) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
