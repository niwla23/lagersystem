// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
	"github.com/niwla23/lagersystem/manager/ent/generated/section"
)

// BoxCreate is the builder for creating a Box entity.
type BoxCreate struct {
	config
	mutation *BoxMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (bc *BoxCreate) SetCreatedAt(t time.Time) *BoxCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (bc *BoxCreate) SetNillableCreatedAt(t *time.Time) *BoxCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updatedAt" field.
func (bc *BoxCreate) SetUpdatedAt(t time.Time) *BoxCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (bc *BoxCreate) SetNillableUpdatedAt(t *time.Time) *BoxCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetBoxId sets the "boxId" field.
func (bc *BoxCreate) SetBoxId(u uuid.UUID) *BoxCreate {
	bc.mutation.SetBoxId(u)
	return bc
}

// AddSectionIDs adds the "sections" edge to the Section entity by IDs.
func (bc *BoxCreate) AddSectionIDs(ids ...int) *BoxCreate {
	bc.mutation.AddSectionIDs(ids...)
	return bc
}

// AddSections adds the "sections" edges to the Section entity.
func (bc *BoxCreate) AddSections(s ...*Section) *BoxCreate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return bc.AddSectionIDs(ids...)
}

// SetPositionID sets the "position" edge to the Position entity by ID.
func (bc *BoxCreate) SetPositionID(id int) *BoxCreate {
	bc.mutation.SetPositionID(id)
	return bc
}

// SetNillablePositionID sets the "position" edge to the Position entity by ID if the given value is not nil.
func (bc *BoxCreate) SetNillablePositionID(id *int) *BoxCreate {
	if id != nil {
		bc = bc.SetPositionID(*id)
	}
	return bc
}

// SetPosition sets the "position" edge to the Position entity.
func (bc *BoxCreate) SetPosition(p *Position) *BoxCreate {
	return bc.SetPositionID(p.ID)
}

// Mutation returns the BoxMutation object of the builder.
func (bc *BoxCreate) Mutation() *BoxMutation {
	return bc.mutation
}

// Save creates the Box in the database.
func (bc *BoxCreate) Save(ctx context.Context) (*Box, error) {
	if err := bc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Box, BoxMutation](ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BoxCreate) SaveX(ctx context.Context) *Box {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BoxCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BoxCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BoxCreate) defaults() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		if box.DefaultCreatedAt == nil {
			return fmt.Errorf("generated: uninitialized box.DefaultCreatedAt (forgotten import generated/runtime?)")
		}
		v := box.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		if box.DefaultUpdatedAt == nil {
			return fmt.Errorf("generated: uninitialized box.DefaultUpdatedAt (forgotten import generated/runtime?)")
		}
		v := box.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (bc *BoxCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`generated: missing required field "Box.createdAt"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`generated: missing required field "Box.updatedAt"`)}
	}
	if _, ok := bc.mutation.BoxId(); !ok {
		return &ValidationError{Name: "boxId", err: errors.New(`generated: missing required field "Box.boxId"`)}
	}
	return nil
}

func (bc *BoxCreate) sqlSave(ctx context.Context) (*Box, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BoxCreate) createSpec() (*Box, *sqlgraph.CreateSpec) {
	var (
		_node = &Box{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: box.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: box.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(box.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(box.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.BoxId(); ok {
		_spec.SetField(box.FieldBoxId, field.TypeUUID, value)
		_node.BoxId = value
	}
	if nodes := bc.mutation.SectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   box.SectionsTable,
			Columns: []string{box.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: section.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.PositionIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BoxCreateBulk is the builder for creating many Box entities in bulk.
type BoxCreateBulk struct {
	config
	builders []*BoxCreate
}

// Save creates the Box entities in the database.
func (bcb *BoxCreateBulk) Save(ctx context.Context) ([]*Box, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Box, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BoxMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BoxCreateBulk) SaveX(ctx context.Context) []*Box {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BoxCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BoxCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}