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
	"github.com/niwla23/lagersystem/manager/ent/generated/part"
	"github.com/niwla23/lagersystem/manager/ent/generated/predicate"
	"github.com/niwla23/lagersystem/manager/ent/generated/property"
	"github.com/niwla23/lagersystem/manager/ent/generated/section"
	"github.com/niwla23/lagersystem/manager/ent/generated/tag"
)

// PartUpdate is the builder for updating Part entities.
type PartUpdate struct {
	config
	hooks    []Hook
	mutation *PartMutation
}

// Where appends a list predicates to the PartUpdate builder.
func (pu *PartUpdate) Where(ps ...predicate.Part) *PartUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetCreatedAt sets the "createdAt" field.
func (pu *PartUpdate) SetCreatedAt(t time.Time) *PartUpdate {
	pu.mutation.SetCreatedAt(t)
	return pu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (pu *PartUpdate) SetNillableCreatedAt(t *time.Time) *PartUpdate {
	if t != nil {
		pu.SetCreatedAt(*t)
	}
	return pu
}

// SetUpdatedAt sets the "updatedAt" field.
func (pu *PartUpdate) SetUpdatedAt(t time.Time) *PartUpdate {
	pu.mutation.SetUpdatedAt(t)
	return pu
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (pu *PartUpdate) SetNillableUpdatedAt(t *time.Time) *PartUpdate {
	if t != nil {
		pu.SetUpdatedAt(*t)
	}
	return pu
}

// SetDeleted sets the "deleted" field.
func (pu *PartUpdate) SetDeleted(b bool) *PartUpdate {
	pu.mutation.SetDeleted(b)
	return pu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (pu *PartUpdate) SetNillableDeleted(b *bool) *PartUpdate {
	if b != nil {
		pu.SetDeleted(*b)
	}
	return pu
}

// SetName sets the "name" field.
func (pu *PartUpdate) SetName(s string) *PartUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetDescription sets the "description" field.
func (pu *PartUpdate) SetDescription(s string) *PartUpdate {
	pu.mutation.SetDescription(s)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PartUpdate) SetAmount(i int) *PartUpdate {
	pu.mutation.ResetAmount()
	pu.mutation.SetAmount(i)
	return pu
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (pu *PartUpdate) SetNillableAmount(i *int) *PartUpdate {
	if i != nil {
		pu.SetAmount(*i)
	}
	return pu
}

// AddAmount adds i to the "amount" field.
func (pu *PartUpdate) AddAmount(i int) *PartUpdate {
	pu.mutation.AddAmount(i)
	return pu
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (pu *PartUpdate) AddTagIDs(ids ...int) *PartUpdate {
	pu.mutation.AddTagIDs(ids...)
	return pu
}

// AddTags adds the "tags" edges to the Tag entity.
func (pu *PartUpdate) AddTags(t ...*Tag) *PartUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.AddTagIDs(ids...)
}

// AddPropertyIDs adds the "properties" edge to the Property entity by IDs.
func (pu *PartUpdate) AddPropertyIDs(ids ...int) *PartUpdate {
	pu.mutation.AddPropertyIDs(ids...)
	return pu
}

// AddProperties adds the "properties" edges to the Property entity.
func (pu *PartUpdate) AddProperties(p ...*Property) *PartUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddPropertyIDs(ids...)
}

// SetSectionID sets the "section" edge to the Section entity by ID.
func (pu *PartUpdate) SetSectionID(id int) *PartUpdate {
	pu.mutation.SetSectionID(id)
	return pu
}

// SetNillableSectionID sets the "section" edge to the Section entity by ID if the given value is not nil.
func (pu *PartUpdate) SetNillableSectionID(id *int) *PartUpdate {
	if id != nil {
		pu = pu.SetSectionID(*id)
	}
	return pu
}

// SetSection sets the "section" edge to the Section entity.
func (pu *PartUpdate) SetSection(s *Section) *PartUpdate {
	return pu.SetSectionID(s.ID)
}

// Mutation returns the PartMutation object of the builder.
func (pu *PartUpdate) Mutation() *PartMutation {
	return pu.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (pu *PartUpdate) ClearTags() *PartUpdate {
	pu.mutation.ClearTags()
	return pu
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (pu *PartUpdate) RemoveTagIDs(ids ...int) *PartUpdate {
	pu.mutation.RemoveTagIDs(ids...)
	return pu
}

// RemoveTags removes "tags" edges to Tag entities.
func (pu *PartUpdate) RemoveTags(t ...*Tag) *PartUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return pu.RemoveTagIDs(ids...)
}

// ClearProperties clears all "properties" edges to the Property entity.
func (pu *PartUpdate) ClearProperties() *PartUpdate {
	pu.mutation.ClearProperties()
	return pu
}

// RemovePropertyIDs removes the "properties" edge to Property entities by IDs.
func (pu *PartUpdate) RemovePropertyIDs(ids ...int) *PartUpdate {
	pu.mutation.RemovePropertyIDs(ids...)
	return pu
}

// RemoveProperties removes "properties" edges to Property entities.
func (pu *PartUpdate) RemoveProperties(p ...*Property) *PartUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemovePropertyIDs(ids...)
}

// ClearSection clears the "section" edge to the Section entity.
func (pu *PartUpdate) ClearSection() *PartUpdate {
	pu.mutation.ClearSection()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PartUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, PartMutation](ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PartUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PartUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PartUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PartUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := part.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Part.name": %w`, err)}
		}
	}
	return nil
}

func (pu *PartUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   part.Table,
			Columns: part.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: part.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.CreatedAt(); ok {
		_spec.SetField(part.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.UpdatedAt(); ok {
		_spec.SetField(part.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := pu.mutation.Deleted(); ok {
		_spec.SetField(part.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(part.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.Description(); ok {
		_spec.SetField(part.FieldDescription, field.TypeString, value)
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(part.FieldAmount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedAmount(); ok {
		_spec.AddField(part.FieldAmount, field.TypeInt, value)
	}
	if pu.mutation.TagsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedTagsIDs(); len(nodes) > 0 && !pu.mutation.TagsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.TagsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PropertiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedPropertiesIDs(); len(nodes) > 0 && !pu.mutation.PropertiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PropertiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.SectionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SectionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{part.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PartUpdateOne is the builder for updating a single Part entity.
type PartUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PartMutation
}

// SetCreatedAt sets the "createdAt" field.
func (puo *PartUpdateOne) SetCreatedAt(t time.Time) *PartUpdateOne {
	puo.mutation.SetCreatedAt(t)
	return puo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (puo *PartUpdateOne) SetNillableCreatedAt(t *time.Time) *PartUpdateOne {
	if t != nil {
		puo.SetCreatedAt(*t)
	}
	return puo
}

// SetUpdatedAt sets the "updatedAt" field.
func (puo *PartUpdateOne) SetUpdatedAt(t time.Time) *PartUpdateOne {
	puo.mutation.SetUpdatedAt(t)
	return puo
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (puo *PartUpdateOne) SetNillableUpdatedAt(t *time.Time) *PartUpdateOne {
	if t != nil {
		puo.SetUpdatedAt(*t)
	}
	return puo
}

// SetDeleted sets the "deleted" field.
func (puo *PartUpdateOne) SetDeleted(b bool) *PartUpdateOne {
	puo.mutation.SetDeleted(b)
	return puo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (puo *PartUpdateOne) SetNillableDeleted(b *bool) *PartUpdateOne {
	if b != nil {
		puo.SetDeleted(*b)
	}
	return puo
}

// SetName sets the "name" field.
func (puo *PartUpdateOne) SetName(s string) *PartUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetDescription sets the "description" field.
func (puo *PartUpdateOne) SetDescription(s string) *PartUpdateOne {
	puo.mutation.SetDescription(s)
	return puo
}

// SetAmount sets the "amount" field.
func (puo *PartUpdateOne) SetAmount(i int) *PartUpdateOne {
	puo.mutation.ResetAmount()
	puo.mutation.SetAmount(i)
	return puo
}

// SetNillableAmount sets the "amount" field if the given value is not nil.
func (puo *PartUpdateOne) SetNillableAmount(i *int) *PartUpdateOne {
	if i != nil {
		puo.SetAmount(*i)
	}
	return puo
}

// AddAmount adds i to the "amount" field.
func (puo *PartUpdateOne) AddAmount(i int) *PartUpdateOne {
	puo.mutation.AddAmount(i)
	return puo
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (puo *PartUpdateOne) AddTagIDs(ids ...int) *PartUpdateOne {
	puo.mutation.AddTagIDs(ids...)
	return puo
}

// AddTags adds the "tags" edges to the Tag entity.
func (puo *PartUpdateOne) AddTags(t ...*Tag) *PartUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.AddTagIDs(ids...)
}

// AddPropertyIDs adds the "properties" edge to the Property entity by IDs.
func (puo *PartUpdateOne) AddPropertyIDs(ids ...int) *PartUpdateOne {
	puo.mutation.AddPropertyIDs(ids...)
	return puo
}

// AddProperties adds the "properties" edges to the Property entity.
func (puo *PartUpdateOne) AddProperties(p ...*Property) *PartUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddPropertyIDs(ids...)
}

// SetSectionID sets the "section" edge to the Section entity by ID.
func (puo *PartUpdateOne) SetSectionID(id int) *PartUpdateOne {
	puo.mutation.SetSectionID(id)
	return puo
}

// SetNillableSectionID sets the "section" edge to the Section entity by ID if the given value is not nil.
func (puo *PartUpdateOne) SetNillableSectionID(id *int) *PartUpdateOne {
	if id != nil {
		puo = puo.SetSectionID(*id)
	}
	return puo
}

// SetSection sets the "section" edge to the Section entity.
func (puo *PartUpdateOne) SetSection(s *Section) *PartUpdateOne {
	return puo.SetSectionID(s.ID)
}

// Mutation returns the PartMutation object of the builder.
func (puo *PartUpdateOne) Mutation() *PartMutation {
	return puo.mutation
}

// ClearTags clears all "tags" edges to the Tag entity.
func (puo *PartUpdateOne) ClearTags() *PartUpdateOne {
	puo.mutation.ClearTags()
	return puo
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (puo *PartUpdateOne) RemoveTagIDs(ids ...int) *PartUpdateOne {
	puo.mutation.RemoveTagIDs(ids...)
	return puo
}

// RemoveTags removes "tags" edges to Tag entities.
func (puo *PartUpdateOne) RemoveTags(t ...*Tag) *PartUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return puo.RemoveTagIDs(ids...)
}

// ClearProperties clears all "properties" edges to the Property entity.
func (puo *PartUpdateOne) ClearProperties() *PartUpdateOne {
	puo.mutation.ClearProperties()
	return puo
}

// RemovePropertyIDs removes the "properties" edge to Property entities by IDs.
func (puo *PartUpdateOne) RemovePropertyIDs(ids ...int) *PartUpdateOne {
	puo.mutation.RemovePropertyIDs(ids...)
	return puo
}

// RemoveProperties removes "properties" edges to Property entities.
func (puo *PartUpdateOne) RemoveProperties(p ...*Property) *PartUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemovePropertyIDs(ids...)
}

// ClearSection clears the "section" edge to the Section entity.
func (puo *PartUpdateOne) ClearSection() *PartUpdateOne {
	puo.mutation.ClearSection()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PartUpdateOne) Select(field string, fields ...string) *PartUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Part entity.
func (puo *PartUpdateOne) Save(ctx context.Context) (*Part, error) {
	return withHooks[*Part, PartMutation](ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PartUpdateOne) SaveX(ctx context.Context) *Part {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PartUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PartUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PartUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := part.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Part.name": %w`, err)}
		}
	}
	return nil
}

func (puo *PartUpdateOne) sqlSave(ctx context.Context) (_node *Part, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   part.Table,
			Columns: part.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: part.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Part.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, part.FieldID)
		for _, f := range fields {
			if !part.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != part.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.CreatedAt(); ok {
		_spec.SetField(part.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.UpdatedAt(); ok {
		_spec.SetField(part.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := puo.mutation.Deleted(); ok {
		_spec.SetField(part.FieldDeleted, field.TypeBool, value)
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(part.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.Description(); ok {
		_spec.SetField(part.FieldDescription, field.TypeString, value)
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(part.FieldAmount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedAmount(); ok {
		_spec.AddField(part.FieldAmount, field.TypeInt, value)
	}
	if puo.mutation.TagsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedTagsIDs(); len(nodes) > 0 && !puo.mutation.TagsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.TagsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PropertiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedPropertiesIDs(); len(nodes) > 0 && !puo.mutation.PropertiesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PropertiesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.SectionCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SectionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Part{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{part.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}