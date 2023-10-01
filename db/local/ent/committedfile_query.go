// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/committedfile"
	"github.com/Creaft-JP/tit/db/local/ent/predicate"
	"github.com/Creaft-JP/tit/db/local/ent/titcommit"
)

// CommittedFileQuery is the builder for querying CommittedFile entities.
type CommittedFileQuery struct {
	config
	ctx        *QueryContext
	order      []committedfile.OrderOption
	inters     []Interceptor
	predicates []predicate.CommittedFile
	withCommit *TitCommitQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CommittedFileQuery builder.
func (cfq *CommittedFileQuery) Where(ps ...predicate.CommittedFile) *CommittedFileQuery {
	cfq.predicates = append(cfq.predicates, ps...)
	return cfq
}

// Limit the number of records to be returned by this query.
func (cfq *CommittedFileQuery) Limit(limit int) *CommittedFileQuery {
	cfq.ctx.Limit = &limit
	return cfq
}

// Offset to start from.
func (cfq *CommittedFileQuery) Offset(offset int) *CommittedFileQuery {
	cfq.ctx.Offset = &offset
	return cfq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cfq *CommittedFileQuery) Unique(unique bool) *CommittedFileQuery {
	cfq.ctx.Unique = &unique
	return cfq
}

// Order specifies how the records should be ordered.
func (cfq *CommittedFileQuery) Order(o ...committedfile.OrderOption) *CommittedFileQuery {
	cfq.order = append(cfq.order, o...)
	return cfq
}

// QueryCommit chains the current query on the "commit" edge.
func (cfq *CommittedFileQuery) QueryCommit() *TitCommitQuery {
	query := (&TitCommitClient{config: cfq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cfq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cfq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(committedfile.Table, committedfile.FieldID, selector),
			sqlgraph.To(titcommit.Table, titcommit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, committedfile.CommitTable, committedfile.CommitColumn),
		)
		fromU = sqlgraph.SetNeighbors(cfq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first CommittedFile entity from the query.
// Returns a *NotFoundError when no CommittedFile was found.
func (cfq *CommittedFileQuery) First(ctx context.Context) (*CommittedFile, error) {
	nodes, err := cfq.Limit(1).All(setContextOp(ctx, cfq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{committedfile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cfq *CommittedFileQuery) FirstX(ctx context.Context) *CommittedFile {
	node, err := cfq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first CommittedFile ID from the query.
// Returns a *NotFoundError when no CommittedFile ID was found.
func (cfq *CommittedFileQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cfq.Limit(1).IDs(setContextOp(ctx, cfq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{committedfile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cfq *CommittedFileQuery) FirstIDX(ctx context.Context) int {
	id, err := cfq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single CommittedFile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one CommittedFile entity is found.
// Returns a *NotFoundError when no CommittedFile entities are found.
func (cfq *CommittedFileQuery) Only(ctx context.Context) (*CommittedFile, error) {
	nodes, err := cfq.Limit(2).All(setContextOp(ctx, cfq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{committedfile.Label}
	default:
		return nil, &NotSingularError{committedfile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cfq *CommittedFileQuery) OnlyX(ctx context.Context) *CommittedFile {
	node, err := cfq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only CommittedFile ID in the query.
// Returns a *NotSingularError when more than one CommittedFile ID is found.
// Returns a *NotFoundError when no entities are found.
func (cfq *CommittedFileQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cfq.Limit(2).IDs(setContextOp(ctx, cfq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{committedfile.Label}
	default:
		err = &NotSingularError{committedfile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cfq *CommittedFileQuery) OnlyIDX(ctx context.Context) int {
	id, err := cfq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of CommittedFiles.
func (cfq *CommittedFileQuery) All(ctx context.Context) ([]*CommittedFile, error) {
	ctx = setContextOp(ctx, cfq.ctx, "All")
	if err := cfq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*CommittedFile, *CommittedFileQuery]()
	return withInterceptors[[]*CommittedFile](ctx, cfq, qr, cfq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cfq *CommittedFileQuery) AllX(ctx context.Context) []*CommittedFile {
	nodes, err := cfq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of CommittedFile IDs.
func (cfq *CommittedFileQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cfq.ctx.Unique == nil && cfq.path != nil {
		cfq.Unique(true)
	}
	ctx = setContextOp(ctx, cfq.ctx, "IDs")
	if err = cfq.Select(committedfile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cfq *CommittedFileQuery) IDsX(ctx context.Context) []int {
	ids, err := cfq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cfq *CommittedFileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cfq.ctx, "Count")
	if err := cfq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cfq, querierCount[*CommittedFileQuery](), cfq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cfq *CommittedFileQuery) CountX(ctx context.Context) int {
	count, err := cfq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cfq *CommittedFileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cfq.ctx, "Exist")
	switch _, err := cfq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cfq *CommittedFileQuery) ExistX(ctx context.Context) bool {
	exist, err := cfq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommittedFileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cfq *CommittedFileQuery) Clone() *CommittedFileQuery {
	if cfq == nil {
		return nil
	}
	return &CommittedFileQuery{
		config:     cfq.config,
		ctx:        cfq.ctx.Clone(),
		order:      append([]committedfile.OrderOption{}, cfq.order...),
		inters:     append([]Interceptor{}, cfq.inters...),
		predicates: append([]predicate.CommittedFile{}, cfq.predicates...),
		withCommit: cfq.withCommit.Clone(),
		// clone intermediate query.
		sql:  cfq.sql.Clone(),
		path: cfq.path,
	}
}

// WithCommit tells the query-builder to eager-load the nodes that are connected to
// the "commit" edge. The optional arguments are used to configure the query builder of the edge.
func (cfq *CommittedFileQuery) WithCommit(opts ...func(*TitCommitQuery)) *CommittedFileQuery {
	query := (&TitCommitClient{config: cfq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cfq.withCommit = query
	return cfq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Path string `json:"path,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.CommittedFile.Query().
//		GroupBy(committedfile.FieldPath).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cfq *CommittedFileQuery) GroupBy(field string, fields ...string) *CommittedFileGroupBy {
	cfq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommittedFileGroupBy{build: cfq}
	grbuild.flds = &cfq.ctx.Fields
	grbuild.label = committedfile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Path string `json:"path,omitempty"`
//	}
//
//	client.CommittedFile.Query().
//		Select(committedfile.FieldPath).
//		Scan(ctx, &v)
func (cfq *CommittedFileQuery) Select(fields ...string) *CommittedFileSelect {
	cfq.ctx.Fields = append(cfq.ctx.Fields, fields...)
	sbuild := &CommittedFileSelect{CommittedFileQuery: cfq}
	sbuild.label = committedfile.Label
	sbuild.flds, sbuild.scan = &cfq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommittedFileSelect configured with the given aggregations.
func (cfq *CommittedFileQuery) Aggregate(fns ...AggregateFunc) *CommittedFileSelect {
	return cfq.Select().Aggregate(fns...)
}

func (cfq *CommittedFileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cfq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cfq); err != nil {
				return err
			}
		}
	}
	for _, f := range cfq.ctx.Fields {
		if !committedfile.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cfq.path != nil {
		prev, err := cfq.path(ctx)
		if err != nil {
			return err
		}
		cfq.sql = prev
	}
	return nil
}

func (cfq *CommittedFileQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*CommittedFile, error) {
	var (
		nodes       = []*CommittedFile{}
		withFKs     = cfq.withFKs
		_spec       = cfq.querySpec()
		loadedTypes = [1]bool{
			cfq.withCommit != nil,
		}
	)
	if cfq.withCommit != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, committedfile.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*CommittedFile).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &CommittedFile{config: cfq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cfq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cfq.withCommit; query != nil {
		if err := cfq.loadCommit(ctx, query, nodes, nil,
			func(n *CommittedFile, e *TitCommit) { n.Edges.Commit = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cfq *CommittedFileQuery) loadCommit(ctx context.Context, query *TitCommitQuery, nodes []*CommittedFile, init func(*CommittedFile), assign func(*CommittedFile, *TitCommit)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*CommittedFile)
	for i := range nodes {
		if nodes[i].tit_commit_files == nil {
			continue
		}
		fk := *nodes[i].tit_commit_files
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(titcommit.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "tit_commit_files" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cfq *CommittedFileQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cfq.querySpec()
	_spec.Node.Columns = cfq.ctx.Fields
	if len(cfq.ctx.Fields) > 0 {
		_spec.Unique = cfq.ctx.Unique != nil && *cfq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cfq.driver, _spec)
}

func (cfq *CommittedFileQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(committedfile.Table, committedfile.Columns, sqlgraph.NewFieldSpec(committedfile.FieldID, field.TypeInt))
	_spec.From = cfq.sql
	if unique := cfq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cfq.path != nil {
		_spec.Unique = true
	}
	if fields := cfq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, committedfile.FieldID)
		for i := range fields {
			if fields[i] != committedfile.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := cfq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cfq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cfq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cfq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cfq *CommittedFileQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cfq.driver.Dialect())
	t1 := builder.Table(committedfile.Table)
	columns := cfq.ctx.Fields
	if len(columns) == 0 {
		columns = committedfile.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cfq.sql != nil {
		selector = cfq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cfq.ctx.Unique != nil && *cfq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cfq.predicates {
		p(selector)
	}
	for _, p := range cfq.order {
		p(selector)
	}
	if offset := cfq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cfq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// CommittedFileGroupBy is the group-by builder for CommittedFile entities.
type CommittedFileGroupBy struct {
	selector
	build *CommittedFileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cfgb *CommittedFileGroupBy) Aggregate(fns ...AggregateFunc) *CommittedFileGroupBy {
	cfgb.fns = append(cfgb.fns, fns...)
	return cfgb
}

// Scan applies the selector query and scans the result into the given value.
func (cfgb *CommittedFileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cfgb.build.ctx, "GroupBy")
	if err := cfgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommittedFileQuery, *CommittedFileGroupBy](ctx, cfgb.build, cfgb, cfgb.build.inters, v)
}

func (cfgb *CommittedFileGroupBy) sqlScan(ctx context.Context, root *CommittedFileQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cfgb.fns))
	for _, fn := range cfgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cfgb.flds)+len(cfgb.fns))
		for _, f := range *cfgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cfgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cfgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CommittedFileSelect is the builder for selecting fields of CommittedFile entities.
type CommittedFileSelect struct {
	*CommittedFileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cfs *CommittedFileSelect) Aggregate(fns ...AggregateFunc) *CommittedFileSelect {
	cfs.fns = append(cfs.fns, fns...)
	return cfs
}

// Scan applies the selector query and scans the result into the given value.
func (cfs *CommittedFileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cfs.ctx, "Select")
	if err := cfs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommittedFileQuery, *CommittedFileSelect](ctx, cfs.CommittedFileQuery, cfs, cfs.inters, v)
}

func (cfs *CommittedFileSelect) sqlScan(ctx context.Context, root *CommittedFileQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cfs.fns))
	for _, fn := range cfs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cfs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cfs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
