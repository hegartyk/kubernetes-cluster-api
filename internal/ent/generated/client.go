// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.infratographer.com/kubernetes-cluster-api/internal/ent/generated/migrate"
	"go.infratographer.com/x/gidx"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"go.infratographer.com/kubernetes-cluster-api/internal/ent/generated/cluster"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Cluster is the client for interacting with the Cluster builders.
	Cluster *ClusterClient
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
	c.Cluster = NewClusterClient(c.config)
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
		return nil, errors.New("generated: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("generated: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:     ctx,
		config:  cfg,
		Cluster: NewClusterClient(cfg),
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
		ctx:     ctx,
		config:  cfg,
		Cluster: NewClusterClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Cluster.
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
	c.Cluster.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Cluster.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ClusterMutation:
		return c.Cluster.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("generated: unknown mutation type %T", m)
	}
}

// ClusterClient is a client for the Cluster schema.
type ClusterClient struct {
	config
}

// NewClusterClient returns a client for the Cluster from the given config.
func NewClusterClient(c config) *ClusterClient {
	return &ClusterClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `cluster.Hooks(f(g(h())))`.
func (c *ClusterClient) Use(hooks ...Hook) {
	c.hooks.Cluster = append(c.hooks.Cluster, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `cluster.Intercept(f(g(h())))`.
func (c *ClusterClient) Intercept(interceptors ...Interceptor) {
	c.inters.Cluster = append(c.inters.Cluster, interceptors...)
}

// Create returns a builder for creating a Cluster entity.
func (c *ClusterClient) Create() *ClusterCreate {
	mutation := newClusterMutation(c.config, OpCreate)
	return &ClusterCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Cluster entities.
func (c *ClusterClient) CreateBulk(builders ...*ClusterCreate) *ClusterCreateBulk {
	return &ClusterCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Cluster.
func (c *ClusterClient) Update() *ClusterUpdate {
	mutation := newClusterMutation(c.config, OpUpdate)
	return &ClusterUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ClusterClient) UpdateOne(cl *Cluster) *ClusterUpdateOne {
	mutation := newClusterMutation(c.config, OpUpdateOne, withCluster(cl))
	return &ClusterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ClusterClient) UpdateOneID(id gidx.PrefixedID) *ClusterUpdateOne {
	mutation := newClusterMutation(c.config, OpUpdateOne, withClusterID(id))
	return &ClusterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Cluster.
func (c *ClusterClient) Delete() *ClusterDelete {
	mutation := newClusterMutation(c.config, OpDelete)
	return &ClusterDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ClusterClient) DeleteOne(cl *Cluster) *ClusterDeleteOne {
	return c.DeleteOneID(cl.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ClusterClient) DeleteOneID(id gidx.PrefixedID) *ClusterDeleteOne {
	builder := c.Delete().Where(cluster.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ClusterDeleteOne{builder}
}

// Query returns a query builder for Cluster.
func (c *ClusterClient) Query() *ClusterQuery {
	return &ClusterQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeCluster},
		inters: c.Interceptors(),
	}
}

// Get returns a Cluster entity by its id.
func (c *ClusterClient) Get(ctx context.Context, id gidx.PrefixedID) (*Cluster, error) {
	return c.Query().Where(cluster.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ClusterClient) GetX(ctx context.Context, id gidx.PrefixedID) *Cluster {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ClusterClient) Hooks() []Hook {
	return c.hooks.Cluster
}

// Interceptors returns the client interceptors.
func (c *ClusterClient) Interceptors() []Interceptor {
	return c.inters.Cluster
}

func (c *ClusterClient) mutate(ctx context.Context, m *ClusterMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ClusterCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ClusterUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ClusterUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ClusterDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown Cluster mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Cluster []ent.Hook
	}
	inters struct {
		Cluster []ent.Interceptor
	}
)
