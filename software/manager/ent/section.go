// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/niwla23/lagersystem/manager/ent/box"
	"github.com/niwla23/lagersystem/manager/ent/section"
)

// Section is the model entity for the Section schema.
type Section struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SectionQuery when eager-loading is set.
	Edges        SectionEdges `json:"edges"`
	box_sections *int
}

// SectionEdges holds the relations/edges for other nodes in the graph.
type SectionEdges struct {
	// Box holds the value of the box edge.
	Box *Box `json:"box,omitempty"`
	// Parts holds the value of the parts edge.
	Parts []*Part `json:"parts,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BoxOrErr returns the Box value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e SectionEdges) BoxOrErr() (*Box, error) {
	if e.loadedTypes[0] {
		if e.Box == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: box.Label}
		}
		return e.Box, nil
	}
	return nil, &NotLoadedError{edge: "box"}
}

// PartsOrErr returns the Parts value or an error if the edge
// was not loaded in eager-loading.
func (e SectionEdges) PartsOrErr() ([]*Part, error) {
	if e.loadedTypes[1] {
		return e.Parts, nil
	}
	return nil, &NotLoadedError{edge: "parts"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Section) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case section.FieldID:
			values[i] = new(sql.NullInt64)
		case section.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case section.ForeignKeys[0]: // box_sections
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Section", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Section fields.
func (s *Section) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case section.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case section.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				s.CreatedAt = value.Time
			}
		case section.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field box_sections", value)
			} else if value.Valid {
				s.box_sections = new(int)
				*s.box_sections = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryBox queries the "box" edge of the Section entity.
func (s *Section) QueryBox() *BoxQuery {
	return NewSectionClient(s.config).QueryBox(s)
}

// QueryParts queries the "parts" edge of the Section entity.
func (s *Section) QueryParts() *PartQuery {
	return NewSectionClient(s.config).QueryParts(s)
}

// Update returns a builder for updating this Section.
// Note that you need to call Section.Unwrap() before calling this method if this Section
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Section) Update() *SectionUpdateOne {
	return NewSectionClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Section entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Section) Unwrap() *Section {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Section is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Section) String() string {
	var builder strings.Builder
	builder.WriteString("Section(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(s.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Sections is a parsable slice of Section.
type Sections []*Section

func (s Sections) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}
