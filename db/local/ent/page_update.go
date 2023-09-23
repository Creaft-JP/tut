// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/local/ent/page"
	"github.com/Creaft-JP/tit/db/local/ent/predicate"
	"github.com/Creaft-JP/tit/db/local/ent/section"
)

// PageUpdate is the builder for updating Page entities.
type PageUpdate struct {
	config
	hooks    []Hook
	mutation *PageMutation
}

// Where appends a list predicates to the PageUpdate builder.
func (pu *PageUpdate) Where(ps ...predicate.Page) *PageUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetPathname sets the "pathname" field.
func (pu *PageUpdate) SetPathname(s string) *PageUpdate {
	pu.mutation.SetPathname(s)
	return pu
}

// SetNumber sets the "number" field.
func (pu *PageUpdate) SetNumber(i int) *PageUpdate {
	pu.mutation.ResetNumber()
	pu.mutation.SetNumber(i)
	return pu
}

// AddNumber adds i to the "number" field.
func (pu *PageUpdate) AddNumber(i int) *PageUpdate {
	pu.mutation.AddNumber(i)
	return pu
}

// SetTitle sets the "title" field.
func (pu *PageUpdate) SetTitle(s string) *PageUpdate {
	pu.mutation.SetTitle(s)
	return pu
}

// SetOverviewSentence sets the "overview_sentence" field.
func (pu *PageUpdate) SetOverviewSentence(s string) *PageUpdate {
	pu.mutation.SetOverviewSentence(s)
	return pu
}

// AddSectionIDs adds the "sections" edge to the Section entity by IDs.
func (pu *PageUpdate) AddSectionIDs(ids ...int) *PageUpdate {
	pu.mutation.AddSectionIDs(ids...)
	return pu
}

// AddSections adds the "sections" edges to the Section entity.
func (pu *PageUpdate) AddSections(s ...*Section) *PageUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.AddSectionIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (pu *PageUpdate) Mutation() *PageMutation {
	return pu.mutation
}

// ClearSections clears all "sections" edges to the Section entity.
func (pu *PageUpdate) ClearSections() *PageUpdate {
	pu.mutation.ClearSections()
	return pu
}

// RemoveSectionIDs removes the "sections" edge to Section entities by IDs.
func (pu *PageUpdate) RemoveSectionIDs(ids ...int) *PageUpdate {
	pu.mutation.RemoveSectionIDs(ids...)
	return pu
}

// RemoveSections removes "sections" edges to Section entities.
func (pu *PageUpdate) RemoveSections(s ...*Section) *PageUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pu.RemoveSectionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PageUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PageUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PageUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PageUpdate) check() error {
	if v, ok := pu.mutation.Pathname(); ok {
		if err := page.PathnameValidator(v); err != nil {
			return &ValidationError{Name: "pathname", err: fmt.Errorf(`ent: validator failed for field "Page.pathname": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Number(); ok {
		if err := page.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Page.number": %w`, err)}
		}
	}
	return nil
}

func (pu *PageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Pathname(); ok {
		_spec.SetField(page.FieldPathname, field.TypeString, value)
	}
	if value, ok := pu.mutation.Number(); ok {
		_spec.SetField(page.FieldNumber, field.TypeInt, value)
	}
	if value, ok := pu.mutation.AddedNumber(); ok {
		_spec.AddField(page.FieldNumber, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
	}
	if value, ok := pu.mutation.OverviewSentence(); ok {
		_spec.SetField(page.FieldOverviewSentence, field.TypeString, value)
	}
	if pu.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.SectionsTable,
			Columns: []string{page.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedSectionsIDs(); len(nodes) > 0 && !pu.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.SectionsTable,
			Columns: []string{page.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.SectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.SectionsTable,
			Columns: []string{page.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PageUpdateOne is the builder for updating a single Page entity.
type PageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PageMutation
}

// SetPathname sets the "pathname" field.
func (puo *PageUpdateOne) SetPathname(s string) *PageUpdateOne {
	puo.mutation.SetPathname(s)
	return puo
}

// SetNumber sets the "number" field.
func (puo *PageUpdateOne) SetNumber(i int) *PageUpdateOne {
	puo.mutation.ResetNumber()
	puo.mutation.SetNumber(i)
	return puo
}

// AddNumber adds i to the "number" field.
func (puo *PageUpdateOne) AddNumber(i int) *PageUpdateOne {
	puo.mutation.AddNumber(i)
	return puo
}

// SetTitle sets the "title" field.
func (puo *PageUpdateOne) SetTitle(s string) *PageUpdateOne {
	puo.mutation.SetTitle(s)
	return puo
}

// SetOverviewSentence sets the "overview_sentence" field.
func (puo *PageUpdateOne) SetOverviewSentence(s string) *PageUpdateOne {
	puo.mutation.SetOverviewSentence(s)
	return puo
}

// AddSectionIDs adds the "sections" edge to the Section entity by IDs.
func (puo *PageUpdateOne) AddSectionIDs(ids ...int) *PageUpdateOne {
	puo.mutation.AddSectionIDs(ids...)
	return puo
}

// AddSections adds the "sections" edges to the Section entity.
func (puo *PageUpdateOne) AddSections(s ...*Section) *PageUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.AddSectionIDs(ids...)
}

// Mutation returns the PageMutation object of the builder.
func (puo *PageUpdateOne) Mutation() *PageMutation {
	return puo.mutation
}

// ClearSections clears all "sections" edges to the Section entity.
func (puo *PageUpdateOne) ClearSections() *PageUpdateOne {
	puo.mutation.ClearSections()
	return puo
}

// RemoveSectionIDs removes the "sections" edge to Section entities by IDs.
func (puo *PageUpdateOne) RemoveSectionIDs(ids ...int) *PageUpdateOne {
	puo.mutation.RemoveSectionIDs(ids...)
	return puo
}

// RemoveSections removes "sections" edges to Section entities.
func (puo *PageUpdateOne) RemoveSections(s ...*Section) *PageUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return puo.RemoveSectionIDs(ids...)
}

// Where appends a list predicates to the PageUpdate builder.
func (puo *PageUpdateOne) Where(ps ...predicate.Page) *PageUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PageUpdateOne) Select(field string, fields ...string) *PageUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Page entity.
func (puo *PageUpdateOne) Save(ctx context.Context) (*Page, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PageUpdateOne) SaveX(ctx context.Context) *Page {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PageUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PageUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PageUpdateOne) check() error {
	if v, ok := puo.mutation.Pathname(); ok {
		if err := page.PathnameValidator(v); err != nil {
			return &ValidationError{Name: "pathname", err: fmt.Errorf(`ent: validator failed for field "Page.pathname": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Number(); ok {
		if err := page.NumberValidator(v); err != nil {
			return &ValidationError{Name: "number", err: fmt.Errorf(`ent: validator failed for field "Page.number": %w`, err)}
		}
	}
	return nil
}

func (puo *PageUpdateOne) sqlSave(ctx context.Context) (_node *Page, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(page.Table, page.Columns, sqlgraph.NewFieldSpec(page.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Page.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, page.FieldID)
		for _, f := range fields {
			if !page.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != page.FieldID {
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
	if value, ok := puo.mutation.Pathname(); ok {
		_spec.SetField(page.FieldPathname, field.TypeString, value)
	}
	if value, ok := puo.mutation.Number(); ok {
		_spec.SetField(page.FieldNumber, field.TypeInt, value)
	}
	if value, ok := puo.mutation.AddedNumber(); ok {
		_spec.AddField(page.FieldNumber, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Title(); ok {
		_spec.SetField(page.FieldTitle, field.TypeString, value)
	}
	if value, ok := puo.mutation.OverviewSentence(); ok {
		_spec.SetField(page.FieldOverviewSentence, field.TypeString, value)
	}
	if puo.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.SectionsTable,
			Columns: []string{page.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedSectionsIDs(); len(nodes) > 0 && !puo.mutation.SectionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.SectionsTable,
			Columns: []string{page.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.SectionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   page.SectionsTable,
			Columns: []string{page.SectionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(section.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Page{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{page.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
