// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Creaft-JP/tit/db/local/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"github.com/Creaft-JP/tit/db/local/ent/page"
	"github.com/Creaft-JP/tit/db/local/ent/remote"
	"github.com/Creaft-JP/tit/db/local/ent/stagedfile"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Page is the client for interacting with the Page builders.
	Page *PageClient
	// Remote is the client for interacting with the Remote builders.
	Remote *RemoteClient
	// StagedFile is the client for interacting with the StagedFile builders.
	StagedFile *StagedFileClient
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
	c.Page = NewPageClient(c.config)
	c.Remote = NewRemoteClient(c.config)
	c.StagedFile = NewStagedFileClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
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
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Page:       NewPageClient(cfg),
		Remote:     NewRemoteClient(cfg),
		StagedFile: NewStagedFileClient(cfg),
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
		ctx:        ctx,
		config:     cfg,
		Page:       NewPageClient(cfg),
		Remote:     NewRemoteClient(cfg),
		StagedFile: NewStagedFileClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Page.
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
	c.Page.Use(hooks...)
	c.Remote.Use(hooks...)
	c.StagedFile.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Page.Intercept(interceptors...)
	c.Remote.Intercept(interceptors...)
	c.StagedFile.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *PageMutation:
		return c.Page.mutate(ctx, m)
	case *RemoteMutation:
		return c.Remote.mutate(ctx, m)
	case *StagedFileMutation:
		return c.StagedFile.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// PageClient is a client for the Page schema.
type PageClient struct {
	config
}

// NewPageClient returns a client for the Page from the given config.
func NewPageClient(c config) *PageClient {
	return &PageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `page.Hooks(f(g(h())))`.
func (c *PageClient) Use(hooks ...Hook) {
	c.hooks.Page = append(c.hooks.Page, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `page.Intercept(f(g(h())))`.
func (c *PageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Page = append(c.inters.Page, interceptors...)
}

// Create returns a builder for creating a Page entity.
func (c *PageClient) Create() *PageCreate {
	mutation := newPageMutation(c.config, OpCreate)
	return &PageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Page entities.
func (c *PageClient) CreateBulk(builders ...*PageCreate) *PageCreateBulk {
	return &PageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Page.
func (c *PageClient) Update() *PageUpdate {
	mutation := newPageMutation(c.config, OpUpdate)
	return &PageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PageClient) UpdateOne(pa *Page) *PageUpdateOne {
	mutation := newPageMutation(c.config, OpUpdateOne, withPage(pa))
	return &PageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PageClient) UpdateOneID(id int) *PageUpdateOne {
	mutation := newPageMutation(c.config, OpUpdateOne, withPageID(id))
	return &PageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Page.
func (c *PageClient) Delete() *PageDelete {
	mutation := newPageMutation(c.config, OpDelete)
	return &PageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *PageClient) DeleteOne(pa *Page) *PageDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *PageClient) DeleteOneID(id int) *PageDeleteOne {
	builder := c.Delete().Where(page.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PageDeleteOne{builder}
}

// Query returns a query builder for Page.
func (c *PageClient) Query() *PageQuery {
	return &PageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypePage},
		inters: c.Interceptors(),
	}
}

// Get returns a Page entity by its id.
func (c *PageClient) Get(ctx context.Context, id int) (*Page, error) {
	return c.Query().Where(page.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PageClient) GetX(ctx context.Context, id int) *Page {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PageClient) Hooks() []Hook {
	return c.hooks.Page
}

// Interceptors returns the client interceptors.
func (c *PageClient) Interceptors() []Interceptor {
	return c.inters.Page
}

func (c *PageClient) mutate(ctx context.Context, m *PageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&PageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&PageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&PageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&PageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Page mutation op: %q", m.Op())
	}
}

// RemoteClient is a client for the Remote schema.
type RemoteClient struct {
	config
}

// NewRemoteClient returns a client for the Remote from the given config.
func NewRemoteClient(c config) *RemoteClient {
	return &RemoteClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `remote.Hooks(f(g(h())))`.
func (c *RemoteClient) Use(hooks ...Hook) {
	c.hooks.Remote = append(c.hooks.Remote, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `remote.Intercept(f(g(h())))`.
func (c *RemoteClient) Intercept(interceptors ...Interceptor) {
	c.inters.Remote = append(c.inters.Remote, interceptors...)
}

// Create returns a builder for creating a Remote entity.
func (c *RemoteClient) Create() *RemoteCreate {
	mutation := newRemoteMutation(c.config, OpCreate)
	return &RemoteCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Remote entities.
func (c *RemoteClient) CreateBulk(builders ...*RemoteCreate) *RemoteCreateBulk {
	return &RemoteCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Remote.
func (c *RemoteClient) Update() *RemoteUpdate {
	mutation := newRemoteMutation(c.config, OpUpdate)
	return &RemoteUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RemoteClient) UpdateOne(r *Remote) *RemoteUpdateOne {
	mutation := newRemoteMutation(c.config, OpUpdateOne, withRemote(r))
	return &RemoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RemoteClient) UpdateOneID(id int) *RemoteUpdateOne {
	mutation := newRemoteMutation(c.config, OpUpdateOne, withRemoteID(id))
	return &RemoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Remote.
func (c *RemoteClient) Delete() *RemoteDelete {
	mutation := newRemoteMutation(c.config, OpDelete)
	return &RemoteDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RemoteClient) DeleteOne(r *Remote) *RemoteDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RemoteClient) DeleteOneID(id int) *RemoteDeleteOne {
	builder := c.Delete().Where(remote.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RemoteDeleteOne{builder}
}

// Query returns a query builder for Remote.
func (c *RemoteClient) Query() *RemoteQuery {
	return &RemoteQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRemote},
		inters: c.Interceptors(),
	}
}

// Get returns a Remote entity by its id.
func (c *RemoteClient) Get(ctx context.Context, id int) (*Remote, error) {
	return c.Query().Where(remote.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RemoteClient) GetX(ctx context.Context, id int) *Remote {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *RemoteClient) Hooks() []Hook {
	return c.hooks.Remote
}

// Interceptors returns the client interceptors.
func (c *RemoteClient) Interceptors() []Interceptor {
	return c.inters.Remote
}

func (c *RemoteClient) mutate(ctx context.Context, m *RemoteMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RemoteCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RemoteUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RemoteUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RemoteDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Remote mutation op: %q", m.Op())
	}
}

// StagedFileClient is a client for the StagedFile schema.
type StagedFileClient struct {
	config
}

// NewStagedFileClient returns a client for the StagedFile from the given config.
func NewStagedFileClient(c config) *StagedFileClient {
	return &StagedFileClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `stagedfile.Hooks(f(g(h())))`.
func (c *StagedFileClient) Use(hooks ...Hook) {
	c.hooks.StagedFile = append(c.hooks.StagedFile, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `stagedfile.Intercept(f(g(h())))`.
func (c *StagedFileClient) Intercept(interceptors ...Interceptor) {
	c.inters.StagedFile = append(c.inters.StagedFile, interceptors...)
}

// Create returns a builder for creating a StagedFile entity.
func (c *StagedFileClient) Create() *StagedFileCreate {
	mutation := newStagedFileMutation(c.config, OpCreate)
	return &StagedFileCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of StagedFile entities.
func (c *StagedFileClient) CreateBulk(builders ...*StagedFileCreate) *StagedFileCreateBulk {
	return &StagedFileCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for StagedFile.
func (c *StagedFileClient) Update() *StagedFileUpdate {
	mutation := newStagedFileMutation(c.config, OpUpdate)
	return &StagedFileUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StagedFileClient) UpdateOne(sf *StagedFile) *StagedFileUpdateOne {
	mutation := newStagedFileMutation(c.config, OpUpdateOne, withStagedFile(sf))
	return &StagedFileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StagedFileClient) UpdateOneID(id int) *StagedFileUpdateOne {
	mutation := newStagedFileMutation(c.config, OpUpdateOne, withStagedFileID(id))
	return &StagedFileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for StagedFile.
func (c *StagedFileClient) Delete() *StagedFileDelete {
	mutation := newStagedFileMutation(c.config, OpDelete)
	return &StagedFileDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StagedFileClient) DeleteOne(sf *StagedFile) *StagedFileDeleteOne {
	return c.DeleteOneID(sf.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *StagedFileClient) DeleteOneID(id int) *StagedFileDeleteOne {
	builder := c.Delete().Where(stagedfile.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StagedFileDeleteOne{builder}
}

// Query returns a query builder for StagedFile.
func (c *StagedFileClient) Query() *StagedFileQuery {
	return &StagedFileQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeStagedFile},
		inters: c.Interceptors(),
	}
}

// Get returns a StagedFile entity by its id.
func (c *StagedFileClient) Get(ctx context.Context, id int) (*StagedFile, error) {
	return c.Query().Where(stagedfile.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StagedFileClient) GetX(ctx context.Context, id int) *StagedFile {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *StagedFileClient) Hooks() []Hook {
	return c.hooks.StagedFile
}

// Interceptors returns the client interceptors.
func (c *StagedFileClient) Interceptors() []Interceptor {
	return c.inters.StagedFile
}

func (c *StagedFileClient) mutate(ctx context.Context, m *StagedFileMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&StagedFileCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&StagedFileUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&StagedFileUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&StagedFileDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown StagedFile mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Page, Remote, StagedFile []ent.Hook
	}
	inters struct {
		Page, Remote, StagedFile []ent.Interceptor
	}
)
