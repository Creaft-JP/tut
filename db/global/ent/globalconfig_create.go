// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/Creaft-JP/tit/db/global/ent/globalconfig"
)

// GlobalConfigCreate is the builder for creating a GlobalConfig entity.
type GlobalConfigCreate struct {
	config
	mutation *GlobalConfigMutation
	hooks    []Hook
}

// SetKey sets the "key" field.
func (gcc *GlobalConfigCreate) SetKey(s string) *GlobalConfigCreate {
	gcc.mutation.SetKey(s)
	return gcc
}

// SetValue sets the "value" field.
func (gcc *GlobalConfigCreate) SetValue(s string) *GlobalConfigCreate {
	gcc.mutation.SetValue(s)
	return gcc
}

// Mutation returns the GlobalConfigMutation object of the builder.
func (gcc *GlobalConfigCreate) Mutation() *GlobalConfigMutation {
	return gcc.mutation
}

// Save creates the GlobalConfig in the database.
func (gcc *GlobalConfigCreate) Save(ctx context.Context) (*GlobalConfig, error) {
	return withHooks(ctx, gcc.sqlSave, gcc.mutation, gcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gcc *GlobalConfigCreate) SaveX(ctx context.Context) *GlobalConfig {
	v, err := gcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcc *GlobalConfigCreate) Exec(ctx context.Context) error {
	_, err := gcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcc *GlobalConfigCreate) ExecX(ctx context.Context) {
	if err := gcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gcc *GlobalConfigCreate) check() error {
	if _, ok := gcc.mutation.Key(); !ok {
		return &ValidationError{Name: "key", err: errors.New(`ent: missing required field "GlobalConfig.key"`)}
	}
	if v, ok := gcc.mutation.Key(); ok {
		if err := globalconfig.KeyValidator(v); err != nil {
			return &ValidationError{Name: "key", err: fmt.Errorf(`ent: validator failed for field "GlobalConfig.key": %w`, err)}
		}
	}
	if _, ok := gcc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`ent: missing required field "GlobalConfig.value"`)}
	}
	if v, ok := gcc.mutation.Value(); ok {
		if err := globalconfig.ValueValidator(v); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`ent: validator failed for field "GlobalConfig.value": %w`, err)}
		}
	}
	return nil
}

func (gcc *GlobalConfigCreate) sqlSave(ctx context.Context) (*GlobalConfig, error) {
	if err := gcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	gcc.mutation.id = &_node.ID
	gcc.mutation.done = true
	return _node, nil
}

func (gcc *GlobalConfigCreate) createSpec() (*GlobalConfig, *sqlgraph.CreateSpec) {
	var (
		_node = &GlobalConfig{config: gcc.config}
		_spec = sqlgraph.NewCreateSpec(globalconfig.Table, sqlgraph.NewFieldSpec(globalconfig.FieldID, field.TypeInt))
	)
	if value, ok := gcc.mutation.Key(); ok {
		_spec.SetField(globalconfig.FieldKey, field.TypeString, value)
		_node.Key = value
	}
	if value, ok := gcc.mutation.Value(); ok {
		_spec.SetField(globalconfig.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	return _node, _spec
}

// GlobalConfigCreateBulk is the builder for creating many GlobalConfig entities in bulk.
type GlobalConfigCreateBulk struct {
	config
	builders []*GlobalConfigCreate
}

// Save creates the GlobalConfig entities in the database.
func (gccb *GlobalConfigCreateBulk) Save(ctx context.Context) ([]*GlobalConfig, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gccb.builders))
	nodes := make([]*GlobalConfig, len(gccb.builders))
	mutators := make([]Mutator, len(gccb.builders))
	for i := range gccb.builders {
		func(i int, root context.Context) {
			builder := gccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GlobalConfigMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gccb *GlobalConfigCreateBulk) SaveX(ctx context.Context) []*GlobalConfig {
	v, err := gccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gccb *GlobalConfigCreateBulk) Exec(ctx context.Context) error {
	_, err := gccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gccb *GlobalConfigCreateBulk) ExecX(ctx context.Context) {
	if err := gccb.Exec(ctx); err != nil {
		panic(err)
	}
}
