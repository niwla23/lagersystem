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
	"github.com/niwla23/lagersystem/manager/ent/generated/tag"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetCreatedAt sets the "createdAt" field.
func (tu *TagUpdate) SetCreatedAt(t time.Time) *TagUpdate {
	tu.mutation.SetCreatedAt(t)
	return tu
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (tu *TagUpdate) SetNillableCreatedAt(t *time.Time) *TagUpdate {
	if t != nil {
		tu.SetCreatedAt(*t)
	}
	return tu
}

// SetUpdatedAt sets the "updatedAt" field.
func (tu *TagUpdate) SetUpdatedAt(t time.Time) *TagUpdate {
	tu.mutation.SetUpdatedAt(t)
	return tu
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (tu *TagUpdate) SetNillableUpdatedAt(t *time.Time) *TagUpdate {
	if t != nil {
		tu.SetUpdatedAt(*t)
	}
	return tu
}

// SetName sets the "name" field.
func (tu *TagUpdate) SetName(s string) *TagUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TagUpdate) SetDescription(s string) *TagUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// AddPartIDs adds the "parts" edge to the Part entity by IDs.
func (tu *TagUpdate) AddPartIDs(ids ...int) *TagUpdate {
	tu.mutation.AddPartIDs(ids...)
	return tu
}

// AddParts adds the "parts" edges to the Part entity.
func (tu *TagUpdate) AddParts(p ...*Part) *TagUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.AddPartIDs(ids...)
}

// SetParentID sets the "parent" edge to the Tag entity by ID.
func (tu *TagUpdate) SetParentID(id int) *TagUpdate {
	tu.mutation.SetParentID(id)
	return tu
}

// SetNillableParentID sets the "parent" edge to the Tag entity by ID if the given value is not nil.
func (tu *TagUpdate) SetNillableParentID(id *int) *TagUpdate {
	if id != nil {
		tu = tu.SetParentID(*id)
	}
	return tu
}

// SetParent sets the "parent" edge to the Tag entity.
func (tu *TagUpdate) SetParent(t *Tag) *TagUpdate {
	return tu.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Tag entity by IDs.
func (tu *TagUpdate) AddChildIDs(ids ...int) *TagUpdate {
	tu.mutation.AddChildIDs(ids...)
	return tu
}

// AddChildren adds the "children" edges to the Tag entity.
func (tu *TagUpdate) AddChildren(t ...*Tag) *TagUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddChildIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearParts clears all "parts" edges to the Part entity.
func (tu *TagUpdate) ClearParts() *TagUpdate {
	tu.mutation.ClearParts()
	return tu
}

// RemovePartIDs removes the "parts" edge to Part entities by IDs.
func (tu *TagUpdate) RemovePartIDs(ids ...int) *TagUpdate {
	tu.mutation.RemovePartIDs(ids...)
	return tu
}

// RemoveParts removes "parts" edges to Part entities.
func (tu *TagUpdate) RemoveParts(p ...*Part) *TagUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tu.RemovePartIDs(ids...)
}

// ClearParent clears the "parent" edge to the Tag entity.
func (tu *TagUpdate) ClearParent() *TagUpdate {
	tu.mutation.ClearParent()
	return tu
}

// ClearChildren clears all "children" edges to the Tag entity.
func (tu *TagUpdate) ClearChildren() *TagUpdate {
	tu.mutation.ClearChildren()
	return tu
}

// RemoveChildIDs removes the "children" edge to Tag entities by IDs.
func (tu *TagUpdate) RemoveChildIDs(ids ...int) *TagUpdate {
	tu.mutation.RemoveChildIDs(ids...)
	return tu
}

// RemoveChildren removes "children" edges to Tag entities.
func (tu *TagUpdate) RemoveChildren(t ...*Tag) *TagUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TagMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TagUpdate) check() error {
	if v, ok := tu.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Tag.name": %w`, err)}
		}
	}
	return nil
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.CreatedAt(); ok {
		_spec.SetField(tag.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.UpdatedAt(); ok {
		_spec.SetField(tag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.SetField(tag.FieldDescription, field.TypeString, value)
	}
	if tu.mutation.PartsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.PartsTable,
			Columns: tag.PartsPrimaryKey,
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
	if nodes := tu.mutation.RemovedPartsIDs(); len(nodes) > 0 && !tu.mutation.PartsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.PartsTable,
			Columns: tag.PartsPrimaryKey,
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
	if nodes := tu.mutation.PartsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.PartsTable,
			Columns: tag.PartsPrimaryKey,
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
	if tu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tag.ParentTable,
			Columns: []string{tag.ParentColumn},
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
	if nodes := tu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tag.ParentTable,
			Columns: []string{tag.ParentColumn},
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
	if tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.ChildrenTable,
			Columns: []string{tag.ChildrenColumn},
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
	if nodes := tu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.ChildrenTable,
			Columns: []string{tag.ChildrenColumn},
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
	if nodes := tu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.ChildrenTable,
			Columns: []string{tag.ChildrenColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetCreatedAt sets the "createdAt" field.
func (tuo *TagUpdateOne) SetCreatedAt(t time.Time) *TagUpdateOne {
	tuo.mutation.SetCreatedAt(t)
	return tuo
}

// SetNillableCreatedAt sets the "createdAt" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableCreatedAt(t *time.Time) *TagUpdateOne {
	if t != nil {
		tuo.SetCreatedAt(*t)
	}
	return tuo
}

// SetUpdatedAt sets the "updatedAt" field.
func (tuo *TagUpdateOne) SetUpdatedAt(t time.Time) *TagUpdateOne {
	tuo.mutation.SetUpdatedAt(t)
	return tuo
}

// SetNillableUpdatedAt sets the "updatedAt" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableUpdatedAt(t *time.Time) *TagUpdateOne {
	if t != nil {
		tuo.SetUpdatedAt(*t)
	}
	return tuo
}

// SetName sets the "name" field.
func (tuo *TagUpdateOne) SetName(s string) *TagUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TagUpdateOne) SetDescription(s string) *TagUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// AddPartIDs adds the "parts" edge to the Part entity by IDs.
func (tuo *TagUpdateOne) AddPartIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.AddPartIDs(ids...)
	return tuo
}

// AddParts adds the "parts" edges to the Part entity.
func (tuo *TagUpdateOne) AddParts(p ...*Part) *TagUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.AddPartIDs(ids...)
}

// SetParentID sets the "parent" edge to the Tag entity by ID.
func (tuo *TagUpdateOne) SetParentID(id int) *TagUpdateOne {
	tuo.mutation.SetParentID(id)
	return tuo
}

// SetNillableParentID sets the "parent" edge to the Tag entity by ID if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableParentID(id *int) *TagUpdateOne {
	if id != nil {
		tuo = tuo.SetParentID(*id)
	}
	return tuo
}

// SetParent sets the "parent" edge to the Tag entity.
func (tuo *TagUpdateOne) SetParent(t *Tag) *TagUpdateOne {
	return tuo.SetParentID(t.ID)
}

// AddChildIDs adds the "children" edge to the Tag entity by IDs.
func (tuo *TagUpdateOne) AddChildIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.AddChildIDs(ids...)
	return tuo
}

// AddChildren adds the "children" edges to the Tag entity.
func (tuo *TagUpdateOne) AddChildren(t ...*Tag) *TagUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddChildIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearParts clears all "parts" edges to the Part entity.
func (tuo *TagUpdateOne) ClearParts() *TagUpdateOne {
	tuo.mutation.ClearParts()
	return tuo
}

// RemovePartIDs removes the "parts" edge to Part entities by IDs.
func (tuo *TagUpdateOne) RemovePartIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.RemovePartIDs(ids...)
	return tuo
}

// RemoveParts removes "parts" edges to Part entities.
func (tuo *TagUpdateOne) RemoveParts(p ...*Part) *TagUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return tuo.RemovePartIDs(ids...)
}

// ClearParent clears the "parent" edge to the Tag entity.
func (tuo *TagUpdateOne) ClearParent() *TagUpdateOne {
	tuo.mutation.ClearParent()
	return tuo
}

// ClearChildren clears all "children" edges to the Tag entity.
func (tuo *TagUpdateOne) ClearChildren() *TagUpdateOne {
	tuo.mutation.ClearChildren()
	return tuo
}

// RemoveChildIDs removes the "children" edge to Tag entities by IDs.
func (tuo *TagUpdateOne) RemoveChildIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.RemoveChildIDs(ids...)
	return tuo
}

// RemoveChildren removes "children" edges to Tag entities.
func (tuo *TagUpdateOne) RemoveChildren(t ...*Tag) *TagUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveChildIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	return withHooks[*Tag, TagMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TagUpdateOne) check() error {
	if v, ok := tuo.mutation.Name(); ok {
		if err := tag.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`generated: validator failed for field "Tag.name": %w`, err)}
		}
	}
	return nil
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tag.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.CreatedAt(); ok {
		_spec.SetField(tag.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.UpdatedAt(); ok {
		_spec.SetField(tag.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(tag.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.SetField(tag.FieldDescription, field.TypeString, value)
	}
	if tuo.mutation.PartsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.PartsTable,
			Columns: tag.PartsPrimaryKey,
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
	if nodes := tuo.mutation.RemovedPartsIDs(); len(nodes) > 0 && !tuo.mutation.PartsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.PartsTable,
			Columns: tag.PartsPrimaryKey,
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
	if nodes := tuo.mutation.PartsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.PartsTable,
			Columns: tag.PartsPrimaryKey,
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
	if tuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tag.ParentTable,
			Columns: []string{tag.ParentColumn},
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
	if nodes := tuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tag.ParentTable,
			Columns: []string{tag.ParentColumn},
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
	if tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.ChildrenTable,
			Columns: []string{tag.ChildrenColumn},
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
	if nodes := tuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !tuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.ChildrenTable,
			Columns: []string{tag.ChildrenColumn},
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
	if nodes := tuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   tag.ChildrenTable,
			Columns: []string{tag.ChildrenColumn},
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
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
