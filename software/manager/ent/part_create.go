// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/niwla23/lagersystem/manager/ent/part"
	"github.com/niwla23/lagersystem/manager/ent/property"
	"github.com/niwla23/lagersystem/manager/ent/section"
	"github.com/niwla23/lagersystem/manager/ent/tag"
)

// PartCreate is the builder for creating a Part entity.
type PartCreate struct {
	config
	mutation *PartMutation
	hooks    []Hook
}

// SetCreatedAt sets the "createdAt" field.
func (pc *PartCreate) SetCreatedAt(t time.Time) *PartCreate {
	pc.mutation.SetCreatedAt(t)
	return pc
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (pc *PartCreate) SetNillableCreatedAt(t *time.Time) *PartCreate {
	if t != nil {
		pc.SetCreatedAt(*t)
	}
	return pc
}

// SetUpdatedAt sets the "updatedAt" field.
func (pc *PartCreate) SetUpdatedAt(t time.Time) *PartCreate {
	pc.mutation.SetUpdatedAt(t)
	return pc
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (pc *PartCreate) SetNillableUpdatedAt(t *time.Time) *PartCreate {
	if t != nil {
		pc.SetUpdatedAt(*t)
	}
	return pc
}

// SetDeleted sets the "deleted" field.
func (pc *PartCreate) SetDeleted(b bool) *PartCreate {
	pc.mutation.SetDeleted(b)
	return pc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (pc *PartCreate) SetNillableDeleted(b *bool) *PartCreate {
	if b != nil {
		pc.SetDeleted(*b)
	}
	return pc
}

// SetName sets the "name" field.
func (pc *PartCreate) SetName(s string) *PartCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PartCreate) SetDescription(s string) *PartCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetAmount sets the "amount" field.
func (pc *PartCreate) SetAmount(i int) *PartCreate {
	pc.mutation.SetAmount(i)
	return pc
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pc *PartCreate) SetNillableAmount(i *int) *PartCreate {
	if i != nil {
		pc.SetAmount(*i)
	}
	return pc
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (pc *PartCreate) AddTagIDs(ids ...int) *PartCreate {
	pc.mutation.AddTagIDs(ids...)
	return pc
}

// AddTags adds the "tags" edges to the Tag entity.
func (pc *PartCreate) AddTags(t ...*Tag) *PartCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pc.AddTagIDs(ids...)
}

// AddPropertyIDs adds the "properties" edge to the Property entity by IDs.
func (pc *PartCreate) AddPropertyIDs(ids ...int) *PartCreate {
	pc.mutation.AddPropertyIDs(ids...)
	return pc
}

// AddProperties adds the "properties" edges to the Property entity.
func (pc *PartCreate) AddProperties(p ...*Property) *PartCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddPropertyIDs(ids...)
}

// SetSectionID sets the "section" edge to the Section entity by ID.
func (pc *PartCreate) SetSectionID(id int) *PartCreate {
	pc.mutation.SetSectionID(id)
	return pc
}

// SetNillableSectionID sets the "section" edge to the Section entity by ID if the given value is not nil.
func (pc *PartCreate) SetNillableSectionID(id *int) *PartCreate {
	if id != nil {
		pc = pc.SetSectionID(*id)
	}
	return pc
}

// SetSection sets the "section" edge to the Section entity.
func (pc *PartCreate) SetSection(s *Section) *PartCreate {
	return pc.SetSectionID(s.ID)
}

// Mutation returns the PartMutation object of the builder.
func (pc *PartCreate) Mutation() *PartMutation {
	return pc.mutation
}

// Save creates the Part in the database.
func (pc *PartCreate) Save(ctx context.Context) (*Part, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Part, PartMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PartCreate) SaveX(ctx context.Context) *Part {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PartCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PartCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PartCreate) defaults() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		if part.DefaultCreatedAt == nil {
			return fmt.Errorf("ent: uninitialized part.DefaultCreatedAt (forgotten import ent/runtime?)")
		}
		v := part.DefaultCreatedAt()
		pc.mutation.SetCreatedAt(v)
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		if part.DefaultUpdatedAt == nil {
			return fmt.Errorf("ent: uninitialized part.DefaultUpdatedAt (forgotten import ent/runtime?)")
		}
		v := part.DefaultUpdatedAt()
		pc.mutation.SetUpdatedAt(v)
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		v := part.DefaultDeleted
		pc.mutation.SetDeleted(v)
	}
	if _, ok := pc.mutation.Amount(); !ok {
		v := part.DefaultAmount
		pc.mutation.SetAmount(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *PartCreate) check() error {
	if _, ok := pc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "createdAt", err: errors.New(`ent: missing required field "Part.createdAt"`)}
	}
	if _, ok := pc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updatedAt", err: errors.New(`ent: missing required field "Part.updatedAt"`)}
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Part.deleted"`)}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Part.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := part.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Part.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`ent: missing required field "Part.description"`)}
	}
	if _, ok := pc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "Part.amount"`)}
	}
	return nil
}

func (pc *PartCreate) sqlSave(ctx context.Context) (*Part, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PartCreate) createSpec() (*Part, *sqlgraph.CreateSpec) {
	var (
		_node = &Part{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: part.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: part.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.CreatedAt(); ok {
		_spec.SetField(part.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := pc.mutation.UpdatedAt(); ok {
		_spec.SetField(part.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := pc.mutation.Deleted(); ok {
		_spec.SetField(part.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(part.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(part.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Amount(); ok {
		_spec.SetField(part.FieldAmount, field.TypeInt, value)
		_node.Amount = value
	}
	if nodes := pc.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   part.TagsTable,
			Columns: part.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PropertiesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   part.PropertiesTable,
			Columns: []string{part.PropertiesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: property.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SectionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   part.SectionTable,
			Columns: []string{part.SectionColumn},
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
		_node.part_section = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PartCreateBulk is the builder for creating many Part entities in bulk.
type PartCreateBulk struct {
	config
	builders []*PartCreate
}

// Save creates the Part entities in the database.
func (pcb *PartCreateBulk) Save(ctx context.Context) ([]*Part, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Part, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PartMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PartCreateBulk) SaveX(ctx context.Context) []*Part {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PartCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PartCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
