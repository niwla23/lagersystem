// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/niwla23/lagersystem/manager/ent/generated/migrate"

	"github.com/niwla23/lagersystem/manager/ent/generated/box"
	"github.com/niwla23/lagersystem/manager/ent/generated/part"
	"github.com/niwla23/lagersystem/manager/ent/generated/position"
	"github.com/niwla23/lagersystem/manager/ent/generated/property"
	"github.com/niwla23/lagersystem/manager/ent/generated/tag"
	"github.com/niwla23/lagersystem/manager/ent/generated/warehouse"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Box is the client for interacting with the Box builders.
	Box *BoxClient
	// Part is the client for interacting with the Part builders.
	Part *PartClient
	// Position is the client for interacting with the Position builders.
	Position *PositionClient
	// Property is the client for interacting with the Property builders.
	Property *PropertyClient
	// Tag is the client for interacting with the Tag builders.
	Tag *TagClient
	// Warehouse is the client for interacting with the Warehouse builders.
	Warehouse *WarehouseClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Box = NewBoxClient(c.config)
	c.Part = NewPartClient(c.config)
	c.Position = NewPositionClient(c.config)
	c.Property = NewPropertyClient(c.config)
	c.Tag = NewTagClient(c.config)
	c.Warehouse = NewWarehouseClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("generated: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("generated: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Box:       NewBoxClient(cfg),
		Part:      NewPartClient(cfg),
		Position:  NewPositionClient(cfg),
		Property:  NewPropertyClient(cfg),
		Tag:       NewTagClient(cfg),
		Warehouse: NewWarehouseClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:       ctx,
		config:    cfg,
		Box:       NewBoxClient(cfg),
		Part:      NewPartClient(cfg),
		Position:  NewPositionClient(cfg),
		Property:  NewPropertyClient(cfg),
		Tag:       NewTagClient(cfg),
		Warehouse: NewWarehouseClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Box.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Box.Use(hooks...)
	c.Part.Use(hooks...)
	c.Position.Use(hooks...)
	c.Property.Use(hooks...)
	c.Tag.Use(hooks...)
	c.Warehouse.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Box.Intercept(interceptors...)
	c.Part.Intercept(interceptors...)
	c.Position.Intercept(interceptors...)
	c.Property.Intercept(interceptors...)
	c.Tag.Intercept(interceptors...)
	c.Warehouse.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *BoxMutation:
		return c.Box.mutate(ctx, m)
	case *PartMutation:
		return c.Part.mutate(ctx, m)
	case *PositionMutation:
		return c.Position.mutate(ctx, m)
	case *PropertyMutation:
		return c.Property.mutate(ctx, m)
	case *TagMutation:
		return c.Tag.mutate(ctx, m)
	case *WarehouseMutation:
		return c.Warehouse.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("generated: unknown mutation type %T", m)
	}
}

// BoxClient is a client for the Box schema.
type BoxClient struct {
	config
}

// NewBoxClient returns a client for the Box from the given config.
func NewBoxClient(c config) *BoxClient {
	return &BoxClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `box.Hooks(f(g(h())))`.
func (c *BoxClient) Use(hooks ...Hook) {
	c.hooks.Box = append(c.hooks.Box, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `box.Intercept(f(g(h())))`.
func (c *BoxClient) Intercept(interceptors ...Interceptor) {
	c.inters.Box = append(c.inters.Box, interceptors...)
}

// Create returns a builder for creating a Box entity.
func (c *BoxClient) Create() *BoxCreate {
	mutation := newBoxMutation(c.config, OpCreate)
	return &BoxCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Box entities.
func (c *BoxClient) CreateBulk(builders ...*BoxCreate) *BoxCreateBulk {
	return &BoxCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Box.
func (c *BoxClient) Update() *BoxUpdate {
	mutation := newBoxMutation(c.config, OpUpdate)
	return &BoxUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BoxClient) UpdateOne(b *Box) *BoxUpdateOne {
	mutation := newBoxMutation(c.config, OpUpdateOne, withBox(b))
	return &BoxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BoxClient) UpdateOneID(id uuid.UUID) *BoxUpdateOne {
	mutation := newBoxMutation(c.config, OpUpdateOne, withBoxID(id))
	return &BoxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Box.
func (c *BoxClient) Delete() *BoxDelete {
	mutation := newBoxMutation(c.config, OpDelete)
	return &BoxDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BoxClient) DeleteOne(b *Box) *BoxDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *BoxClient) DeleteOneID(id uuid.UUID) *BoxDeleteOne {
	builder := c.Delete().Where(box.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BoxDeleteOne{builder}
}

// Query returns a query builder for Box.
func (c *BoxClient) Query() *BoxQuery {
	return &BoxQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeBox},
		inters: c.Interceptors(),
	}
}

// Get returns a Box entity by its id.
func (c *BoxClient) Get(ctx context.Context, id uuid.UUID) (*Box, error) {
	return c.Query().Where(box.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BoxClient) GetX(ctx context.Context, id uuid.UUID) *Box {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParts queries the parts edge of a Box.
func (c *BoxClient) QueryParts(b *Box) *PartQuery {
	query := (&PartClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(box.Table, box.FieldID, id),
			sqlgraph.To(part.Table, part.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, box.PartsTable, box.PartsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPosition queries the position edge of a Box.
func (c *BoxClient) QueryPosition(b *Box) *PositionQuery {
	query := (&PositionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(box.Table, box.FieldID, id),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, box.PositionTable, box.PositionColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BoxClient) Hooks() []Hook {
	hooks := c.hooks.Box
	return append(hooks[:len(hooks):len(hooks)], box.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *BoxClient) Interceptors() []Interceptor {
	return c.inters.Box
}

func (c *BoxClient) mutate(ctx context.Context, m *BoxMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&BoxCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&BoxUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&BoxUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&BoxDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Box mutation op: %q", m.Op())
	}
}

// PartClient is a client for the Part schema.
type PartClient struct {
	config
}

// NewPartClient returns a client for the Part from the given config.
func NewPartClient(c config) *PartClient {
	return &PartClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `part.Hooks(f(g(h())))`.
func (c *PartClient) Use(hooks ...Hook) {
	c.hooks.Part = append(c.hooks.Part, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `part.Intercept(f(g(h())))`.
func (c *PartClient) Intercept(interceptors ...Interceptor) {
	c.inters.Part = append(c.inters.Part, interceptors...)
}

// Create returns a builder for creating a Part entity.
func (c *PartClient) Create() *PartCreate {
	mutation := newPartMutation(c.config, OpCreate)
	return &PartCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Part entities.
func (c *PartClient) CreateBulk(builders ...*PartCreate) *PartCreateBulk {
	return &PartCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Part.
func (c *PartClient) Update() *PartUpdate {
	mutation := newPartMutation(c.config, OpUpdate)
	return &PartUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PartClient) UpdateOne(pa *Part) *PartUpdateOne {
	mutation := newPartMutation(c.config, OpUpdateOne, withPart(pa))
	return &PartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PartClient) UpdateOneID(id int) *PartUpdateOne {
	mutation := newPartMutation(c.config, OpUpdateOne, withPartID(id))
	return &PartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Part.
func (c *PartClient) Delete() *PartDelete {
	mutation := newPartMutation(c.config, OpDelete)
	return &PartDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PartClient) DeleteOne(pa *Part) *PartDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PartClient) DeleteOneID(id int) *PartDeleteOne {
	builder := c.Delete().Where(part.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PartDeleteOne{builder}
}

// Query returns a query builder for Part.
func (c *PartClient) Query() *PartQuery {
	return &PartQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePart},
		inters: c.Interceptors(),
	}
}

// Get returns a Part entity by its id.
func (c *PartClient) Get(ctx context.Context, id int) (*Part, error) {
	return c.Query().Where(part.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PartClient) GetX(ctx context.Context, id int) *Part {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTags queries the tags edge of a Part.
func (c *PartClient) QueryTags(pa *Part) *TagQuery {
	query := (&TagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(part.Table, part.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, part.TagsTable, part.TagsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryProperties queries the properties edge of a Part.
func (c *PartClient) QueryProperties(pa *Part) *PropertyQuery {
	query := (&PropertyClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(part.Table, part.FieldID, id),
			sqlgraph.To(property.Table, property.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, part.PropertiesTable, part.PropertiesColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBox queries the box edge of a Part.
func (c *PartClient) QueryBox(pa *Part) *BoxQuery {
	query := (&BoxClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(part.Table, part.FieldID, id),
			sqlgraph.To(box.Table, box.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, part.BoxTable, part.BoxColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PartClient) Hooks() []Hook {
	hooks := c.hooks.Part
	return append(hooks[:len(hooks):len(hooks)], part.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *PartClient) Interceptors() []Interceptor {
	return c.inters.Part
}

func (c *PartClient) mutate(ctx context.Context, m *PartMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PartCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PartUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PartUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PartDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Part mutation op: %q", m.Op())
	}
}

// PositionClient is a client for the Position schema.
type PositionClient struct {
	config
}

// NewPositionClient returns a client for the Position from the given config.
func NewPositionClient(c config) *PositionClient {
	return &PositionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `position.Hooks(f(g(h())))`.
func (c *PositionClient) Use(hooks ...Hook) {
	c.hooks.Position = append(c.hooks.Position, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `position.Intercept(f(g(h())))`.
func (c *PositionClient) Intercept(interceptors ...Interceptor) {
	c.inters.Position = append(c.inters.Position, interceptors...)
}

// Create returns a builder for creating a Position entity.
func (c *PositionClient) Create() *PositionCreate {
	mutation := newPositionMutation(c.config, OpCreate)
	return &PositionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Position entities.
func (c *PositionClient) CreateBulk(builders ...*PositionCreate) *PositionCreateBulk {
	return &PositionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Position.
func (c *PositionClient) Update() *PositionUpdate {
	mutation := newPositionMutation(c.config, OpUpdate)
	return &PositionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PositionClient) UpdateOne(po *Position) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPosition(po))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PositionClient) UpdateOneID(id int) *PositionUpdateOne {
	mutation := newPositionMutation(c.config, OpUpdateOne, withPositionID(id))
	return &PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Position.
func (c *PositionClient) Delete() *PositionDelete {
	mutation := newPositionMutation(c.config, OpDelete)
	return &PositionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PositionClient) DeleteOne(po *Position) *PositionDeleteOne {
	return c.DeleteOneID(po.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PositionClient) DeleteOneID(id int) *PositionDeleteOne {
	builder := c.Delete().Where(position.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PositionDeleteOne{builder}
}

// Query returns a query builder for Position.
func (c *PositionClient) Query() *PositionQuery {
	return &PositionQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePosition},
		inters: c.Interceptors(),
	}
}

// Get returns a Position entity by its id.
func (c *PositionClient) Get(ctx context.Context, id int) (*Position, error) {
	return c.Query().Where(position.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PositionClient) GetX(ctx context.Context, id int) *Position {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStoredBox queries the storedBox edge of a Position.
func (c *PositionClient) QueryStoredBox(po *Position) *BoxQuery {
	query := (&BoxClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(position.Table, position.FieldID, id),
			sqlgraph.To(box.Table, box.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, position.StoredBoxTable, position.StoredBoxColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryWarehouse queries the warehouse edge of a Position.
func (c *PositionClient) QueryWarehouse(po *Position) *WarehouseQuery {
	query := (&WarehouseClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := po.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(position.Table, position.FieldID, id),
			sqlgraph.To(warehouse.Table, warehouse.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, position.WarehouseTable, position.WarehouseColumn),
		)
		fromV = sqlgraph.Neighbors(po.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PositionClient) Hooks() []Hook {
	hooks := c.hooks.Position
	return append(hooks[:len(hooks):len(hooks)], position.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *PositionClient) Interceptors() []Interceptor {
	return c.inters.Position
}

func (c *PositionClient) mutate(ctx context.Context, m *PositionMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PositionCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PositionUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PositionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PositionDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Position mutation op: %q", m.Op())
	}
}

// PropertyClient is a client for the Property schema.
type PropertyClient struct {
	config
}

// NewPropertyClient returns a client for the Property from the given config.
func NewPropertyClient(c config) *PropertyClient {
	return &PropertyClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `property.Hooks(f(g(h())))`.
func (c *PropertyClient) Use(hooks ...Hook) {
	c.hooks.Property = append(c.hooks.Property, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `property.Intercept(f(g(h())))`.
func (c *PropertyClient) Intercept(interceptors ...Interceptor) {
	c.inters.Property = append(c.inters.Property, interceptors...)
}

// Create returns a builder for creating a Property entity.
func (c *PropertyClient) Create() *PropertyCreate {
	mutation := newPropertyMutation(c.config, OpCreate)
	return &PropertyCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Property entities.
func (c *PropertyClient) CreateBulk(builders ...*PropertyCreate) *PropertyCreateBulk {
	return &PropertyCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Property.
func (c *PropertyClient) Update() *PropertyUpdate {
	mutation := newPropertyMutation(c.config, OpUpdate)
	return &PropertyUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PropertyClient) UpdateOne(pr *Property) *PropertyUpdateOne {
	mutation := newPropertyMutation(c.config, OpUpdateOne, withProperty(pr))
	return &PropertyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PropertyClient) UpdateOneID(id int) *PropertyUpdateOne {
	mutation := newPropertyMutation(c.config, OpUpdateOne, withPropertyID(id))
	return &PropertyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Property.
func (c *PropertyClient) Delete() *PropertyDelete {
	mutation := newPropertyMutation(c.config, OpDelete)
	return &PropertyDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PropertyClient) DeleteOne(pr *Property) *PropertyDeleteOne {
	return c.DeleteOneID(pr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PropertyClient) DeleteOneID(id int) *PropertyDeleteOne {
	builder := c.Delete().Where(property.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PropertyDeleteOne{builder}
}

// Query returns a query builder for Property.
func (c *PropertyClient) Query() *PropertyQuery {
	return &PropertyQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeProperty},
		inters: c.Interceptors(),
	}
}

// Get returns a Property entity by its id.
func (c *PropertyClient) Get(ctx context.Context, id int) (*Property, error) {
	return c.Query().Where(property.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PropertyClient) GetX(ctx context.Context, id int) *Property {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPart queries the part edge of a Property.
func (c *PropertyClient) QueryPart(pr *Property) *PartQuery {
	query := (&PartClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := pr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(property.Table, property.FieldID, id),
			sqlgraph.To(part.Table, part.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, property.PartTable, property.PartColumn),
		)
		fromV = sqlgraph.Neighbors(pr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *PropertyClient) Hooks() []Hook {
	hooks := c.hooks.Property
	return append(hooks[:len(hooks):len(hooks)], property.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *PropertyClient) Interceptors() []Interceptor {
	return c.inters.Property
}

func (c *PropertyClient) mutate(ctx context.Context, m *PropertyMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PropertyCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PropertyUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PropertyUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PropertyDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Property mutation op: %q", m.Op())
	}
}

// TagClient is a client for the Tag schema.
type TagClient struct {
	config
}

// NewTagClient returns a client for the Tag from the given config.
func NewTagClient(c config) *TagClient {
	return &TagClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tag.Hooks(f(g(h())))`.
func (c *TagClient) Use(hooks ...Hook) {
	c.hooks.Tag = append(c.hooks.Tag, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `tag.Intercept(f(g(h())))`.
func (c *TagClient) Intercept(interceptors ...Interceptor) {
	c.inters.Tag = append(c.inters.Tag, interceptors...)
}

// Create returns a builder for creating a Tag entity.
func (c *TagClient) Create() *TagCreate {
	mutation := newTagMutation(c.config, OpCreate)
	return &TagCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tag entities.
func (c *TagClient) CreateBulk(builders ...*TagCreate) *TagCreateBulk {
	return &TagCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tag.
func (c *TagClient) Update() *TagUpdate {
	mutation := newTagMutation(c.config, OpUpdate)
	return &TagUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TagClient) UpdateOne(t *Tag) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTag(t))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TagClient) UpdateOneID(id int) *TagUpdateOne {
	mutation := newTagMutation(c.config, OpUpdateOne, withTagID(id))
	return &TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tag.
func (c *TagClient) Delete() *TagDelete {
	mutation := newTagMutation(c.config, OpDelete)
	return &TagDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TagClient) DeleteOne(t *Tag) *TagDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TagClient) DeleteOneID(id int) *TagDeleteOne {
	builder := c.Delete().Where(tag.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TagDeleteOne{builder}
}

// Query returns a query builder for Tag.
func (c *TagClient) Query() *TagQuery {
	return &TagQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTag},
		inters: c.Interceptors(),
	}
}

// Get returns a Tag entity by its id.
func (c *TagClient) Get(ctx context.Context, id int) (*Tag, error) {
	return c.Query().Where(tag.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TagClient) GetX(ctx context.Context, id int) *Tag {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParts queries the parts edge of a Tag.
func (c *TagClient) QueryParts(t *Tag) *PartQuery {
	query := (&PartClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(part.Table, part.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tag.PartsTable, tag.PartsPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParent queries the parent edge of a Tag.
func (c *TagClient) QueryParent(t *Tag) *TagQuery {
	query := (&TagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, tag.ParentTable, tag.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a Tag.
func (c *TagClient) QueryChildren(t *Tag) *TagQuery {
	query := (&TagClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tag.Table, tag.FieldID, id),
			sqlgraph.To(tag.Table, tag.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, tag.ChildrenTable, tag.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TagClient) Hooks() []Hook {
	hooks := c.hooks.Tag
	return append(hooks[:len(hooks):len(hooks)], tag.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *TagClient) Interceptors() []Interceptor {
	return c.inters.Tag
}

func (c *TagClient) mutate(ctx context.Context, m *TagMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TagCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TagUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TagUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TagDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Tag mutation op: %q", m.Op())
	}
}

// WarehouseClient is a client for the Warehouse schema.
type WarehouseClient struct {
	config
}

// NewWarehouseClient returns a client for the Warehouse from the given config.
func NewWarehouseClient(c config) *WarehouseClient {
	return &WarehouseClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `warehouse.Hooks(f(g(h())))`.
func (c *WarehouseClient) Use(hooks ...Hook) {
	c.hooks.Warehouse = append(c.hooks.Warehouse, hooks...)
}

// Use adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `warehouse.Intercept(f(g(h())))`.
func (c *WarehouseClient) Intercept(interceptors ...Interceptor) {
	c.inters.Warehouse = append(c.inters.Warehouse, interceptors...)
}

// Create returns a builder for creating a Warehouse entity.
func (c *WarehouseClient) Create() *WarehouseCreate {
	mutation := newWarehouseMutation(c.config, OpCreate)
	return &WarehouseCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Warehouse entities.
func (c *WarehouseClient) CreateBulk(builders ...*WarehouseCreate) *WarehouseCreateBulk {
	return &WarehouseCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Warehouse.
func (c *WarehouseClient) Update() *WarehouseUpdate {
	mutation := newWarehouseMutation(c.config, OpUpdate)
	return &WarehouseUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *WarehouseClient) UpdateOne(w *Warehouse) *WarehouseUpdateOne {
	mutation := newWarehouseMutation(c.config, OpUpdateOne, withWarehouse(w))
	return &WarehouseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *WarehouseClient) UpdateOneID(id int) *WarehouseUpdateOne {
	mutation := newWarehouseMutation(c.config, OpUpdateOne, withWarehouseID(id))
	return &WarehouseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Warehouse.
func (c *WarehouseClient) Delete() *WarehouseDelete {
	mutation := newWarehouseMutation(c.config, OpDelete)
	return &WarehouseDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *WarehouseClient) DeleteOne(w *Warehouse) *WarehouseDeleteOne {
	return c.DeleteOneID(w.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *WarehouseClient) DeleteOneID(id int) *WarehouseDeleteOne {
	builder := c.Delete().Where(warehouse.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &WarehouseDeleteOne{builder}
}

// Query returns a query builder for Warehouse.
func (c *WarehouseClient) Query() *WarehouseQuery {
	return &WarehouseQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeWarehouse},
		inters: c.Interceptors(),
	}
}

// Get returns a Warehouse entity by its id.
func (c *WarehouseClient) Get(ctx context.Context, id int) (*Warehouse, error) {
	return c.Query().Where(warehouse.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *WarehouseClient) GetX(ctx context.Context, id int) *Warehouse {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPositions queries the positions edge of a Warehouse.
func (c *WarehouseClient) QueryPositions(w *Warehouse) *PositionQuery {
	query := (&PositionClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := w.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(warehouse.Table, warehouse.FieldID, id),
			sqlgraph.To(position.Table, position.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, warehouse.PositionsTable, warehouse.PositionsColumn),
		)
		fromV = sqlgraph.Neighbors(w.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *WarehouseClient) Hooks() []Hook {
	hooks := c.hooks.Warehouse
	return append(hooks[:len(hooks):len(hooks)], warehouse.Hooks[:]...)
}

// Interceptors returns the client interceptors.
func (c *WarehouseClient) Interceptors() []Interceptor {
	return c.inters.Warehouse
}

func (c *WarehouseClient) mutate(ctx context.Context, m *WarehouseMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&WarehouseCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&WarehouseUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&WarehouseUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&WarehouseDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Warehouse mutation op: %q", m.Op())
	}
}
