// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"web/ent/pet"
	"web/ent/predicate"
	"web/ent/token"
	"web/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUsername sets the "username" field.
func (uu *UserUpdate) SetUsername(s string) *UserUpdate {
	uu.mutation.SetUsername(s)
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(i)
	return uu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAge(i *int) *UserUpdate {
	if i != nil {
		uu.SetAge(*i)
	}
	return uu
}

// AddAge adds i to the "age" field.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.mutation.AddAge(i)
	return uu
}

// ClearAge clears the value of the "age" field.
func (uu *UserUpdate) ClearAge() *UserUpdate {
	uu.mutation.ClearAge()
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of the "phone" field.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.mutation.ClearPhone()
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// ClearEmail clears the value of the "email" field.
func (uu *UserUpdate) ClearEmail() *UserUpdate {
	uu.mutation.ClearEmail()
	return uu
}

// SetAddr sets the "addr" field.
func (uu *UserUpdate) SetAddr(s string) *UserUpdate {
	uu.mutation.SetAddr(s)
	return uu
}

// SetNillableAddr sets the "addr" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddr(s *string) *UserUpdate {
	if s != nil {
		uu.SetAddr(*s)
	}
	return uu
}

// ClearAddr clears the value of the "addr" field.
func (uu *UserUpdate) ClearAddr() *UserUpdate {
	uu.mutation.ClearAddr()
	return uu
}

// SetAvatar sets the "avatar" field.
func (uu *UserUpdate) SetAvatar(s string) *UserUpdate {
	uu.mutation.SetAvatar(s)
	return uu
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAvatar(s *string) *UserUpdate {
	if s != nil {
		uu.SetAvatar(*s)
	}
	return uu
}

// SetCareer sets the "career" field.
func (uu *UserUpdate) SetCareer(s string) *UserUpdate {
	uu.mutation.SetCareer(s)
	return uu
}

// SetNillableCareer sets the "career" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCareer(s *string) *UserUpdate {
	if s != nil {
		uu.SetCareer(*s)
	}
	return uu
}

// AddSentIDs adds the "sent" edge to the Pet entity by IDs.
func (uu *UserUpdate) AddSentIDs(ids ...int) *UserUpdate {
	uu.mutation.AddSentIDs(ids...)
	return uu
}

// AddSent adds the "sent" edges to the Pet entity.
func (uu *UserUpdate) AddSent(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddSentIDs(ids...)
}

// SetTokenID sets the "token" edge to the Token entity by ID.
func (uu *UserUpdate) SetTokenID(id int) *UserUpdate {
	uu.mutation.SetTokenID(id)
	return uu
}

// SetNillableTokenID sets the "token" edge to the Token entity by ID if the given value is not nil.
func (uu *UserUpdate) SetNillableTokenID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetTokenID(*id)
	}
	return uu
}

// SetToken sets the "token" edge to the Token entity.
func (uu *UserUpdate) SetToken(t *Token) *UserUpdate {
	return uu.SetTokenID(t.ID)
}

// AddAdoptedIDs adds the "adopted" edge to the Pet entity by IDs.
func (uu *UserUpdate) AddAdoptedIDs(ids ...int) *UserUpdate {
	uu.mutation.AddAdoptedIDs(ids...)
	return uu
}

// AddAdopted adds the "adopted" edges to the Pet entity.
func (uu *UserUpdate) AddAdopted(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddAdoptedIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearSent clears all "sent" edges to the Pet entity.
func (uu *UserUpdate) ClearSent() *UserUpdate {
	uu.mutation.ClearSent()
	return uu
}

// RemoveSentIDs removes the "sent" edge to Pet entities by IDs.
func (uu *UserUpdate) RemoveSentIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveSentIDs(ids...)
	return uu
}

// RemoveSent removes "sent" edges to Pet entities.
func (uu *UserUpdate) RemoveSent(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveSentIDs(ids...)
}

// ClearToken clears the "token" edge to the Token entity.
func (uu *UserUpdate) ClearToken() *UserUpdate {
	uu.mutation.ClearToken()
	return uu
}

// ClearAdopted clears all "adopted" edges to the Pet entity.
func (uu *UserUpdate) ClearAdopted() *UserUpdate {
	uu.mutation.ClearAdopted()
	return uu
}

// RemoveAdoptedIDs removes the "adopted" edge to Pet entities by IDs.
func (uu *UserUpdate) RemoveAdoptedIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveAdoptedIDs(ids...)
	return uu
}

// RemoveAdopted removes "adopted" edges to Pet entities.
func (uu *UserUpdate) RemoveAdopted(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemoveAdoptedIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "User.age": %w`, err)}
		}
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if uu.mutation.AgeCleared() {
		_spec.ClearField(user.FieldAge, field.TypeInt)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uu.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uu.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uu.mutation.Addr(); ok {
		_spec.SetField(user.FieldAddr, field.TypeString, value)
	}
	if uu.mutation.AddrCleared() {
		_spec.ClearField(user.FieldAddr, field.TypeString)
	}
	if value, ok := uu.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uu.mutation.Career(); ok {
		_spec.SetField(user.FieldCareer, field.TypeString, value)
	}
	if uu.mutation.SentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SentTable,
			Columns: []string{user.SentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedSentIDs(); len(nodes) > 0 && !uu.mutation.SentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SentTable,
			Columns: []string{user.SentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.SentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SentTable,
			Columns: []string{user.SentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.AdoptedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AdoptedTable,
			Columns: []string{user.AdoptedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedAdoptedIDs(); len(nodes) > 0 && !uu.mutation.AdoptedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AdoptedTable,
			Columns: []string{user.AdoptedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.AdoptedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AdoptedTable,
			Columns: []string{user.AdoptedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUsername sets the "username" field.
func (uuo *UserUpdateOne) SetUsername(s string) *UserUpdateOne {
	uuo.mutation.SetUsername(s)
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetAge sets the "age" field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(i)
	return uuo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAge(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetAge(*i)
	}
	return uuo
}

// AddAge adds i to the "age" field.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.mutation.AddAge(i)
	return uuo
}

// ClearAge clears the value of the "age" field.
func (uuo *UserUpdateOne) ClearAge() *UserUpdateOne {
	uuo.mutation.ClearAge()
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of the "phone" field.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.mutation.ClearPhone()
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// ClearEmail clears the value of the "email" field.
func (uuo *UserUpdateOne) ClearEmail() *UserUpdateOne {
	uuo.mutation.ClearEmail()
	return uuo
}

// SetAddr sets the "addr" field.
func (uuo *UserUpdateOne) SetAddr(s string) *UserUpdateOne {
	uuo.mutation.SetAddr(s)
	return uuo
}

// SetNillableAddr sets the "addr" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddr(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAddr(*s)
	}
	return uuo
}

// ClearAddr clears the value of the "addr" field.
func (uuo *UserUpdateOne) ClearAddr() *UserUpdateOne {
	uuo.mutation.ClearAddr()
	return uuo
}

// SetAvatar sets the "avatar" field.
func (uuo *UserUpdateOne) SetAvatar(s string) *UserUpdateOne {
	uuo.mutation.SetAvatar(s)
	return uuo
}

// SetNillableAvatar sets the "avatar" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAvatar(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAvatar(*s)
	}
	return uuo
}

// SetCareer sets the "career" field.
func (uuo *UserUpdateOne) SetCareer(s string) *UserUpdateOne {
	uuo.mutation.SetCareer(s)
	return uuo
}

// SetNillableCareer sets the "career" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCareer(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCareer(*s)
	}
	return uuo
}

// AddSentIDs adds the "sent" edge to the Pet entity by IDs.
func (uuo *UserUpdateOne) AddSentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddSentIDs(ids...)
	return uuo
}

// AddSent adds the "sent" edges to the Pet entity.
func (uuo *UserUpdateOne) AddSent(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddSentIDs(ids...)
}

// SetTokenID sets the "token" edge to the Token entity by ID.
func (uuo *UserUpdateOne) SetTokenID(id int) *UserUpdateOne {
	uuo.mutation.SetTokenID(id)
	return uuo
}

// SetNillableTokenID sets the "token" edge to the Token entity by ID if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTokenID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetTokenID(*id)
	}
	return uuo
}

// SetToken sets the "token" edge to the Token entity.
func (uuo *UserUpdateOne) SetToken(t *Token) *UserUpdateOne {
	return uuo.SetTokenID(t.ID)
}

// AddAdoptedIDs adds the "adopted" edge to the Pet entity by IDs.
func (uuo *UserUpdateOne) AddAdoptedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddAdoptedIDs(ids...)
	return uuo
}

// AddAdopted adds the "adopted" edges to the Pet entity.
func (uuo *UserUpdateOne) AddAdopted(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddAdoptedIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearSent clears all "sent" edges to the Pet entity.
func (uuo *UserUpdateOne) ClearSent() *UserUpdateOne {
	uuo.mutation.ClearSent()
	return uuo
}

// RemoveSentIDs removes the "sent" edge to Pet entities by IDs.
func (uuo *UserUpdateOne) RemoveSentIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveSentIDs(ids...)
	return uuo
}

// RemoveSent removes "sent" edges to Pet entities.
func (uuo *UserUpdateOne) RemoveSent(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveSentIDs(ids...)
}

// ClearToken clears the "token" edge to the Token entity.
func (uuo *UserUpdateOne) ClearToken() *UserUpdateOne {
	uuo.mutation.ClearToken()
	return uuo
}

// ClearAdopted clears all "adopted" edges to the Pet entity.
func (uuo *UserUpdateOne) ClearAdopted() *UserUpdateOne {
	uuo.mutation.ClearAdopted()
	return uuo
}

// RemoveAdoptedIDs removes the "adopted" edge to Pet entities by IDs.
func (uuo *UserUpdateOne) RemoveAdoptedIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveAdoptedIDs(ids...)
	return uuo
}

// RemoveAdopted removes "adopted" edges to Pet entities.
func (uuo *UserUpdateOne) RemoveAdopted(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemoveAdoptedIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.Age(); ok {
		if err := user.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "User.age": %w`, err)}
		}
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Username(); ok {
		_spec.SetField(user.FieldUsername, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if uuo.mutation.AgeCleared() {
		_spec.ClearField(user.FieldAge, field.TypeInt)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if uuo.mutation.PhoneCleared() {
		_spec.ClearField(user.FieldPhone, field.TypeString)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if uuo.mutation.EmailCleared() {
		_spec.ClearField(user.FieldEmail, field.TypeString)
	}
	if value, ok := uuo.mutation.Addr(); ok {
		_spec.SetField(user.FieldAddr, field.TypeString, value)
	}
	if uuo.mutation.AddrCleared() {
		_spec.ClearField(user.FieldAddr, field.TypeString)
	}
	if value, ok := uuo.mutation.Avatar(); ok {
		_spec.SetField(user.FieldAvatar, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Career(); ok {
		_spec.SetField(user.FieldCareer, field.TypeString, value)
	}
	if uuo.mutation.SentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SentTable,
			Columns: []string{user.SentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedSentIDs(); len(nodes) > 0 && !uuo.mutation.SentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SentTable,
			Columns: []string{user.SentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.SentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.SentTable,
			Columns: []string{user.SentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TokenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TokenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.TokenTable,
			Columns: []string{user.TokenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(token.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.AdoptedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AdoptedTable,
			Columns: []string{user.AdoptedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedAdoptedIDs(); len(nodes) > 0 && !uuo.mutation.AdoptedCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AdoptedTable,
			Columns: []string{user.AdoptedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.AdoptedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.AdoptedTable,
			Columns: []string{user.AdoptedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
