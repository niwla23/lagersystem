// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/part"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
	"github.com/niwla23/lagersystem/manager/ent/generated/predicate"
)

// BoxUpdate is the builder for updating Box entities.
type BoxUpdate struct {
	config
	hooks    []Hook
	mutation *BoxMutation
}

// Where appends a list predicates to the BoxUpdate builder.
func (bu *BoxUpdate) Where(ps ...predicate.Box) *BoxUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetCreatedAt sets the "createdAt" field.
func (bu *BoxUpdate) SetCreatedAt(t time.Time) *BoxUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (bu *BoxUpdate) SetNillableCreatedAt(t *time.Time) *BoxUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updatedAt" field.
func (bu *BoxUpdate) SetUpdatedAt(t time.Time) *BoxUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (bu *BoxUpdate) SetNillableUpdatedAt(t *time.Time) *BoxUpdate {
	if t != nil {
		bu.SetUpdatedAt(*t)
	}
	return bu
}

// SetBoxId sets the "boxId" field.
func (bu *BoxUpdate) SetBoxId(u uuid.UUID) *BoxUpdate {
	bu.mutation.SetBoxId(u)
	return bu
}

// AddPartIDs adds the "parts" edge to the Part entity by IDs.
func (bu *BoxUpdate) AddPartIDs(ids ...int) *BoxUpdate {
	bu.mutation.AddPartIDs(ids...)
	return bu
}

// AddParts adds the "parts" edges to the Part entity.
func (bu *BoxUpdate) AddParts(p ...*Part) *BoxUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.AddPartIDs(ids...)
}

// SetPositionID sets the "position" edge to the Position entity by ID.
func (bu *BoxUpdate) SetPositionID(id int) *BoxUpdate {
	bu.mutation.SetPositionID(id)
	return bu
}

// SetNillablePositionID sets the "position" edge to the Position entity by ID if the given value is not nil.
func (bu *BoxUpdate) SetNillablePositionID(id *int) *BoxUpdate {
	if id != nil {
		bu = bu.SetPositionID(*id)
	}
	return bu
}

// SetPosition sets the "position" edge to the Position entity.
func (bu *BoxUpdate) SetPosition(p *Position) *BoxUpdate {
	return bu.SetPositionID(p.ID)
}

// Mutation returns the BoxMutation object of the builder.
func (bu *BoxUpdate) Mutation() *BoxMutation {
	return bu.mutation
}

// ClearParts clears all "parts" edges to the Part entity.
func (bu *BoxUpdate) ClearParts() *BoxUpdate {
	bu.mutation.ClearParts()
	return bu
}

// RemovePartIDs removes the "parts" edge to Part entities by IDs.
func (bu *BoxUpdate) RemovePartIDs(ids ...int) *BoxUpdate {
	bu.mutation.RemovePartIDs(ids...)
	return bu
}

// RemoveParts removes "parts" edges to Part entities.
func (bu *BoxUpdate) RemoveParts(p ...*Part) *BoxUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return bu.RemovePartIDs(ids...)
}

// ClearPosition clears the "position" edge to the Position entity.
func (bu *BoxUpdate) ClearPosition() *BoxUpdate {
	bu.mutation.ClearPosition()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BoxUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, BoxMutation](ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BoxUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BoxUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BoxUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BoxUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   box.Table,
			Columns: box.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: box.FieldID,
			},
		},
	}
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.SetField(box.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(box.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.BoxId(); ok {
		_spec.SetField(box.FieldBoxId, field.TypeUUID, value)
	}
	if bu.mutation.PartsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   box.PartsTable,
			Columns: []string{box.PartsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: part.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.RemovedPartsIDs(); len(nodes) > 0 && !bu.mutation.PartsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   box.PartsTable,
			Columns: []string{box.PartsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: part.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.PartsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   box.PartsTable,
			Columns: []string{box.PartsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: part.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.PositionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   box.PositionTable,
			Columns: []string{box.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: position.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.PositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   box.PositionTable,
			Columns: []string{box.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: position.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{box.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BoxUpdateOne is the builder for updating a single Box entity.
type BoxUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BoxMutation
}

// SetCreatedAt sets the "createdAt" field.
func (buo *BoxUpdateOne) SetCreatedAt(t time.Time) *BoxUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (buo *BoxUpdateOne) SetNillableCreatedAt(t *time.Time) *BoxUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updatedAt" field.
func (buo *BoxUpdateOne) SetUpdatedAt(t time.Time) *BoxUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (buo *BoxUpdateOne) SetNillableUpdatedAt(t *time.Time) *BoxUpdateOne {
	if t != nil {
		buo.SetUpdatedAt(*t)
	}
	return buo
}

// SetBoxId sets the "boxId" field.
func (buo *BoxUpdateOne) SetBoxId(u uuid.UUID) *BoxUpdateOne {
	buo.mutation.SetBoxId(u)
	return buo
}

// AddPartIDs adds the "parts" edge to the Part entity by IDs.
func (buo *BoxUpdateOne) AddPartIDs(ids ...int) *BoxUpdateOne {
	buo.mutation.AddPartIDs(ids...)
	return buo
}

// AddParts adds the "parts" edges to the Part entity.
func (buo *BoxUpdateOne) AddParts(p ...*Part) *BoxUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.AddPartIDs(ids...)
}

// SetPositionID sets the "position" edge to the Position entity by ID.
func (buo *BoxUpdateOne) SetPositionID(id int) *BoxUpdateOne {
	buo.mutation.SetPositionID(id)
	return buo
}

// SetNillablePositionID sets the "position" edge to the Position entity by ID if the given value is not nil.
func (buo *BoxUpdateOne) SetNillablePositionID(id *int) *BoxUpdateOne {
	if id != nil {
		buo = buo.SetPositionID(*id)
	}
	return buo
}

// SetPosition sets the "position" edge to the Position entity.
func (buo *BoxUpdateOne) SetPosition(p *Position) *BoxUpdateOne {
	return buo.SetPositionID(p.ID)
}

// Mutation returns the BoxMutation object of the builder.
func (buo *BoxUpdateOne) Mutation() *BoxMutation {
	return buo.mutation
}

// ClearParts clears all "parts" edges to the Part entity.
func (buo *BoxUpdateOne) ClearParts() *BoxUpdateOne {
	buo.mutation.ClearParts()
	return buo
}

// RemovePartIDs removes the "parts" edge to Part entities by IDs.
func (buo *BoxUpdateOne) RemovePartIDs(ids ...int) *BoxUpdateOne {
	buo.mutation.RemovePartIDs(ids...)
	return buo
}

// RemoveParts removes "parts" edges to Part entities.
func (buo *BoxUpdateOne) RemoveParts(p ...*Part) *BoxUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return buo.RemovePartIDs(ids...)
}

// ClearPosition clears the "position" edge to the Position entity.
func (buo *BoxUpdateOne) ClearPosition() *BoxUpdateOne {
	buo.mutation.ClearPosition()
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BoxUpdateOne) Select(field string, fields ...string) *BoxUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Box entity.
func (buo *BoxUpdateOne) Save(ctx context.Context) (*Box, error) {
	return withHooks[*Box, BoxMutation](ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BoxUpdateOne) SaveX(ctx context.Context) *Box {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BoxUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BoxUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BoxUpdateOne) sqlSave(ctx context.Context) (_node *Box, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   box.Table,
			Columns: box.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: box.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Box.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, box.FieldID)
		for _, f := range fields {
			if !box.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != box.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.SetField(box.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(box.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.BoxId(); ok {
		_spec.SetField(box.FieldBoxId, field.TypeUUID, value)
	}
	if buo.mutation.PartsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   box.PartsTable,
			Columns: []string{box.PartsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: part.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.RemovedPartsIDs(); len(nodes) > 0 && !buo.mutation.PartsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   box.PartsTable,
			Columns: []string{box.PartsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: part.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.PartsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   box.PartsTable,
			Columns: []string{box.PartsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: part.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.PositionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   box.PositionTable,
			Columns: []string{box.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: position.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.PositionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   box.PositionTable,
			Columns: []string{box.PositionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: position.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Box{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{box.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
