// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/niwla23/lagersystem/manager/ent/box"
	"github.com/niwla23/lagersystem/manager/ent/part"
	"github.com/niwla23/lagersystem/manager/ent/predicate"
	"github.com/niwla23/lagersystem/manager/ent/section"
)

// SectionQuery is the builder for querying Section entities.
type SectionQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Section
	withBox    *BoxQuery
	withParts  *PartQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SectionQuery builder.
func (sq *SectionQuery) Where(ps ...predicate.Section) *SectionQuery {
	sq.predicates = append(sq.predicates, ps...)
	return sq
}

// Limit the number of records to be returned by this query.
func (sq *SectionQuery) Limit(limit int) *SectionQuery {
	sq.ctx.Limit = &limit
	return sq
}

// Offset to start from.
func (sq *SectionQuery) Offset(offset int) *SectionQuery {
	sq.ctx.Offset = &offset
	return sq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (sq *SectionQuery) Unique(unique bool) *SectionQuery {
	sq.ctx.Unique = &unique
	return sq
}

// Order specifies how the records should be ordered.
func (sq *SectionQuery) Order(o ...OrderFunc) *SectionQuery {
	sq.order = append(sq.order, o...)
	return sq
}

// QueryBox chains the current query on the "box" edge.
func (sq *SectionQuery) QueryBox() *BoxQuery {
	query := (&BoxClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(section.Table, section.FieldID, selector),
			sqlgraph.To(box.Table, box.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, section.BoxTable, section.BoxColumn),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryParts chains the current query on the "parts" edge.
func (sq *SectionQuery) QueryParts() *PartQuery {
	query := (&PartClient{config: sq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := sq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := sq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(section.Table, section.FieldID, selector),
			sqlgraph.To(part.Table, part.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, section.PartsTable, section.PartsPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(sq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Section entity from the query.
// Returns a *NotFoundError when no Section was found.
func (sq *SectionQuery) First(ctx context.Context) (*Section, error) {
	nodes, err := sq.Limit(1).All(setContextOp(ctx, sq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{section.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (sq *SectionQuery) FirstX(ctx context.Context) *Section {
	node, err := sq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Section ID from the query.
// Returns a *NotFoundError when no Section ID was found.
func (sq *SectionQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(1).IDs(setContextOp(ctx, sq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{section.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (sq *SectionQuery) FirstIDX(ctx context.Context) int {
	id, err := sq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Section entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Section entity is found.
// Returns a *NotFoundError when no Section entities are found.
func (sq *SectionQuery) Only(ctx context.Context) (*Section, error) {
	nodes, err := sq.Limit(2).All(setContextOp(ctx, sq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{section.Label}
	default:
		return nil, &NotSingularError{section.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (sq *SectionQuery) OnlyX(ctx context.Context) *Section {
	node, err := sq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Section ID in the query.
// Returns a *NotSingularError when more than one Section ID is found.
// Returns a *NotFoundError when no entities are found.
func (sq *SectionQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = sq.Limit(2).IDs(setContextOp(ctx, sq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{section.Label}
	default:
		err = &NotSingularError{section.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (sq *SectionQuery) OnlyIDX(ctx context.Context) int {
	id, err := sq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Sections.
func (sq *SectionQuery) All(ctx context.Context) ([]*Section, error) {
	ctx = setContextOp(ctx, sq.ctx, "All")
	if err := sq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Section, *SectionQuery]()
	return withInterceptors[[]*Section](ctx, sq, qr, sq.inters)
}

// AllX is like All, but panics if an error occurs.
func (sq *SectionQuery) AllX(ctx context.Context) []*Section {
	nodes, err := sq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Section IDs.
func (sq *SectionQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	ctx = setContextOp(ctx, sq.ctx, "IDs")
	if err := sq.Select(section.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (sq *SectionQuery) IDsX(ctx context.Context) []int {
	ids, err := sq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (sq *SectionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, sq.ctx, "Count")
	if err := sq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, sq, querierCount[*SectionQuery](), sq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (sq *SectionQuery) CountX(ctx context.Context) int {
	count, err := sq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (sq *SectionQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, sq.ctx, "Exist")
	switch _, err := sq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (sq *SectionQuery) ExistX(ctx context.Context) bool {
	exist, err := sq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SectionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (sq *SectionQuery) Clone() *SectionQuery {
	if sq == nil {
		return nil
	}
	return &SectionQuery{
		config:     sq.config,
		ctx:        sq.ctx.Clone(),
		order:      append([]OrderFunc{}, sq.order...),
		inters:     append([]Interceptor{}, sq.inters...),
		predicates: append([]predicate.Section{}, sq.predicates...),
		withBox:    sq.withBox.Clone(),
		withParts:  sq.withParts.Clone(),
		// clone intermediate query.
		sql:  sq.sql.Clone(),
		path: sq.path,
	}
}

// WithBox tells the query-builder to eager-load the nodes that are connected to
// the "box" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SectionQuery) WithBox(opts ...func(*BoxQuery)) *SectionQuery {
	query := (&BoxClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withBox = query
	return sq
}

// WithParts tells the query-builder to eager-load the nodes that are connected to
// the "parts" edge. The optional arguments are used to configure the query builder of the edge.
func (sq *SectionQuery) WithParts(opts ...func(*PartQuery)) *SectionQuery {
	query := (&PartClient{config: sq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	sq.withParts = query
	return sq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Section.Query().
//		GroupBy(section.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (sq *SectionQuery) GroupBy(field string, fields ...string) *SectionGroupBy {
	sq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SectionGroupBy{build: sq}
	grbuild.flds = &sq.ctx.Fields
	grbuild.label = section.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"createdAt,omitempty"`
//	}
//
//	client.Section.Query().
//		Select(section.FieldCreatedAt).
//		Scan(ctx, &v)
func (sq *SectionQuery) Select(fields ...string) *SectionSelect {
	sq.ctx.Fields = append(sq.ctx.Fields, fields...)
	sbuild := &SectionSelect{SectionQuery: sq}
	sbuild.label = section.Label
	sbuild.flds, sbuild.scan = &sq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SectionSelect configured with the given aggregations.
func (sq *SectionQuery) Aggregate(fns ...AggregateFunc) *SectionSelect {
	return sq.Select().Aggregate(fns...)
}

func (sq *SectionQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range sq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, sq); err != nil {
				return err
			}
		}
	}
	for _, f := range sq.ctx.Fields {
		if !section.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if sq.path != nil {
		prev, err := sq.path(ctx)
		if err != nil {
			return err
		}
		sq.sql = prev
	}
	return nil
}

func (sq *SectionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Section, error) {
	var (
		nodes       = []*Section{}
		withFKs     = sq.withFKs
		_spec       = sq.querySpec()
		loadedTypes = [2]bool{
			sq.withBox != nil,
			sq.withParts != nil,
		}
	)
	if sq.withBox != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, section.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Section).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Section{config: sq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, sq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := sq.withBox; query != nil {
		if err := sq.loadBox(ctx, query, nodes, nil,
			func(n *Section, e *Box) { n.Edges.Box = e }); err != nil {
			return nil, err
		}
	}
	if query := sq.withParts; query != nil {
		if err := sq.loadParts(ctx, query, nodes,
			func(n *Section) { n.Edges.Parts = []*Part{} },
			func(n *Section, e *Part) { n.Edges.Parts = append(n.Edges.Parts, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (sq *SectionQuery) loadBox(ctx context.Context, query *BoxQuery, nodes []*Section, init func(*Section), assign func(*Section, *Box)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Section)
	for i := range nodes {
		if nodes[i].box_sections == nil {
			continue
		}
		fk := *nodes[i].box_sections
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(box.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "box_sections" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (sq *SectionQuery) loadParts(ctx context.Context, query *PartQuery, nodes []*Section, init func(*Section), assign func(*Section, *Part)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*Section)
	nids := make(map[int]map[*Section]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(section.PartsTable)
		s.Join(joinT).On(s.C(part.FieldID), joinT.C(section.PartsPrimaryKey[0]))
		s.Where(sql.InValues(joinT.C(section.PartsPrimaryKey[1]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(section.PartsPrimaryKey[1]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*Section]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Part](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "parts" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (sq *SectionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := sq.querySpec()
	_spec.Node.Columns = sq.ctx.Fields
	if len(sq.ctx.Fields) > 0 {
		_spec.Unique = sq.ctx.Unique != nil && *sq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, sq.driver, _spec)
}

func (sq *SectionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   section.Table,
			Columns: section.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: section.FieldID,
			},
		},
		From:   sq.sql,
		Unique: true,
	}
	if unique := sq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := sq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, section.FieldID)
		for i := range fields {
			if fields[i] != section.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := sq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := sq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := sq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := sq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (sq *SectionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(sq.driver.Dialect())
	t1 := builder.Table(section.Table)
	columns := sq.ctx.Fields
	if len(columns) == 0 {
		columns = section.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if sq.sql != nil {
		selector = sq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if sq.ctx.Unique != nil && *sq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range sq.predicates {
		p(selector)
	}
	for _, p := range sq.order {
		p(selector)
	}
	if offset := sq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := sq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// SectionGroupBy is the group-by builder for Section entities.
type SectionGroupBy struct {
	selector
	build *SectionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SectionGroupBy) Aggregate(fns ...AggregateFunc) *SectionGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SectionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, "GroupBy")
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SectionQuery, *SectionGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (sgb *SectionGroupBy) sqlScan(ctx context.Context, root *SectionQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(sgb.fns))
	for _, fn := range sgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*sgb.flds)+len(sgb.fns))
		for _, f := range *sgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*sgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := sgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// SectionSelect is the builder for selecting fields of Section entities.
type SectionSelect struct {
	*SectionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SectionSelect) Aggregate(fns ...AggregateFunc) *SectionSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SectionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, "Select")
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SectionQuery, *SectionSelect](ctx, ss.SectionQuery, ss, ss.inters, v)
}

func (ss *SectionSelect) sqlScan(ctx context.Context, root *SectionQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ss.fns))
	for _, fn := range ss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
