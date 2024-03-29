// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"db/ent/tourproduct"
	"db/ent/user"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TourProductCreate is the builder for creating a TourProduct entity.
type TourProductCreate struct {
	config
	mutation *TourProductMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (tpc *TourProductCreate) SetName(s string) *TourProductCreate {
	tpc.mutation.SetName(s)
	return tpc
}

// SetPrice sets the "price" field.
func (tpc *TourProductCreate) SetPrice(i int) *TourProductCreate {
	tpc.mutation.SetPrice(i)
	return tpc
}

// SetForSale sets the "forSale" field.
func (tpc *TourProductCreate) SetForSale(b bool) *TourProductCreate {
	tpc.mutation.SetForSale(b)
	return tpc
}

// SetNillableForSale sets the "forSale" field if the given value is not nil.
func (tpc *TourProductCreate) SetNillableForSale(b *bool) *TourProductCreate {
	if b != nil {
		tpc.SetForSale(*b)
	}
	return tpc
}

// SetManagerID sets the "manager" edge to the User entity by ID.
func (tpc *TourProductCreate) SetManagerID(id string) *TourProductCreate {
	tpc.mutation.SetManagerID(id)
	return tpc
}

// SetManager sets the "manager" edge to the User entity.
func (tpc *TourProductCreate) SetManager(u *User) *TourProductCreate {
	return tpc.SetManagerID(u.ID)
}

// Mutation returns the TourProductMutation object of the builder.
func (tpc *TourProductCreate) Mutation() *TourProductMutation {
	return tpc.mutation
}

// Save creates the TourProduct in the database.
func (tpc *TourProductCreate) Save(ctx context.Context) (*TourProduct, error) {
	var (
		err  error
		node *TourProduct
	)
	tpc.defaults()
	if len(tpc.hooks) == 0 {
		if err = tpc.check(); err != nil {
			return nil, err
		}
		node, err = tpc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TourProductMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tpc.check(); err != nil {
				return nil, err
			}
			tpc.mutation = mutation
			node, err = tpc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tpc.hooks) - 1; i >= 0; i-- {
			mut = tpc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tpc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (tpc *TourProductCreate) SaveX(ctx context.Context) *TourProduct {
	v, err := tpc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (tpc *TourProductCreate) defaults() {
	if _, ok := tpc.mutation.ForSale(); !ok {
		v := tourproduct.DefaultForSale
		tpc.mutation.SetForSale(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tpc *TourProductCreate) check() error {
	if _, ok := tpc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if _, ok := tpc.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if _, ok := tpc.mutation.ForSale(); !ok {
		return &ValidationError{Name: "forSale", err: errors.New("ent: missing required field \"forSale\"")}
	}
	if _, ok := tpc.mutation.ManagerID(); !ok {
		return &ValidationError{Name: "manager", err: errors.New("ent: missing required edge \"manager\"")}
	}
	return nil
}

func (tpc *TourProductCreate) sqlSave(ctx context.Context) (*TourProduct, error) {
	_node, _spec := tpc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tpc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tpc *TourProductCreate) createSpec() (*TourProduct, *sqlgraph.CreateSpec) {
	var (
		_node = &TourProduct{config: tpc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: tourproduct.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tourproduct.FieldID,
			},
		}
	)
	if value, ok := tpc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tourproduct.FieldName,
		})
		_node.Name = value
	}
	if value, ok := tpc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tourproduct.FieldPrice,
		})
		_node.Price = value
	}
	if value, ok := tpc.mutation.ForSale(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: tourproduct.FieldForSale,
		})
		_node.ForSale = value
	}
	if nodes := tpc.mutation.ManagerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tourproduct.ManagerTable,
			Columns: []string{tourproduct.ManagerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_products = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TourProductCreateBulk is the builder for creating many TourProduct entities in bulk.
type TourProductCreateBulk struct {
	config
	builders []*TourProductCreate
}

// Save creates the TourProduct entities in the database.
func (tpcb *TourProductCreateBulk) Save(ctx context.Context) ([]*TourProduct, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tpcb.builders))
	nodes := make([]*TourProduct, len(tpcb.builders))
	mutators := make([]Mutator, len(tpcb.builders))
	for i := range tpcb.builders {
		func(i int, root context.Context) {
			builder := tpcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TourProductMutation)
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
					_, err = mutators[i+1].Mutate(root, tpcb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tpcb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tpcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tpcb *TourProductCreateBulk) SaveX(ctx context.Context) []*TourProduct {
	v, err := tpcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
