// Code generated by ent, DO NOT EDIT.

package build

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/luminancetech/varso/src/services/auth/internal/ent/build/predicate"
	"github.com/luminancetech/varso/src/services/auth/internal/ent/build/role"
	"github.com/luminancetech/varso/src/services/auth/internal/ent/build/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeRole = "Role"
	TypeUser = "User"
)

// RoleMutation represents an operation that mutates the Role nodes in the graph.
type RoleMutation struct {
	config
	op                Op
	typ               string
	id                *int
	uuid              *uuid.UUID
	key               *string
	permissions       *[]string
	appendpermissions []string
	clearedFields     map[string]struct{}
	user              map[uuid.UUID]struct{}
	removeduser       map[uuid.UUID]struct{}
	cleareduser       bool
	done              bool
	oldValue          func(context.Context) (*Role, error)
	predicates        []predicate.Role
}

var _ ent.Mutation = (*RoleMutation)(nil)

// roleOption allows management of the mutation configuration using functional options.
type roleOption func(*RoleMutation)

// newRoleMutation creates new mutation for the Role entity.
func newRoleMutation(c config, op Op, opts ...roleOption) *RoleMutation {
	m := &RoleMutation{
		config:        c,
		op:            op,
		typ:           TypeRole,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRoleID sets the ID field of the mutation.
func withRoleID(id int) roleOption {
	return func(m *RoleMutation) {
		var (
			err   error
			once  sync.Once
			value *Role
		)
		m.oldValue = func(ctx context.Context) (*Role, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Role.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRole sets the old Role of the mutation.
func withRole(node *Role) roleOption {
	return func(m *RoleMutation) {
		m.oldValue = func(context.Context) (*Role, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RoleMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RoleMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("build: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RoleMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *RoleMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Role.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetUUID sets the "uuid" field.
func (m *RoleMutation) SetUUID(u uuid.UUID) {
	m.uuid = &u
}

// UUID returns the value of the "uuid" field in the mutation.
func (m *RoleMutation) UUID() (r uuid.UUID, exists bool) {
	v := m.uuid
	if v == nil {
		return
	}
	return *v, true
}

// OldUUID returns the old "uuid" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldUUID(ctx context.Context) (v uuid.UUID, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUUID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUUID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUUID: %w", err)
	}
	return oldValue.UUID, nil
}

// ResetUUID resets all changes to the "uuid" field.
func (m *RoleMutation) ResetUUID() {
	m.uuid = nil
}

// SetKey sets the "key" field.
func (m *RoleMutation) SetKey(s string) {
	m.key = &s
}

// Key returns the value of the "key" field in the mutation.
func (m *RoleMutation) Key() (r string, exists bool) {
	v := m.key
	if v == nil {
		return
	}
	return *v, true
}

// OldKey returns the old "key" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldKey(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldKey is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldKey requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldKey: %w", err)
	}
	return oldValue.Key, nil
}

// ResetKey resets all changes to the "key" field.
func (m *RoleMutation) ResetKey() {
	m.key = nil
}

// SetPermissions sets the "permissions" field.
func (m *RoleMutation) SetPermissions(s []string) {
	m.permissions = &s
	m.appendpermissions = nil
}

// Permissions returns the value of the "permissions" field in the mutation.
func (m *RoleMutation) Permissions() (r []string, exists bool) {
	v := m.permissions
	if v == nil {
		return
	}
	return *v, true
}

// OldPermissions returns the old "permissions" field's value of the Role entity.
// If the Role object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RoleMutation) OldPermissions(ctx context.Context) (v []string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPermissions is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPermissions requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPermissions: %w", err)
	}
	return oldValue.Permissions, nil
}

// AppendPermissions adds s to the "permissions" field.
func (m *RoleMutation) AppendPermissions(s []string) {
	m.appendpermissions = append(m.appendpermissions, s...)
}

// AppendedPermissions returns the list of values that were appended to the "permissions" field in this mutation.
func (m *RoleMutation) AppendedPermissions() ([]string, bool) {
	if len(m.appendpermissions) == 0 {
		return nil, false
	}
	return m.appendpermissions, true
}

// ResetPermissions resets all changes to the "permissions" field.
func (m *RoleMutation) ResetPermissions() {
	m.permissions = nil
	m.appendpermissions = nil
}

// AddUserIDs adds the "user" edge to the User entity by ids.
func (m *RoleMutation) AddUserIDs(ids ...uuid.UUID) {
	if m.user == nil {
		m.user = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		m.user[ids[i]] = struct{}{}
	}
}

// ClearUser clears the "user" edge to the User entity.
func (m *RoleMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *RoleMutation) UserCleared() bool {
	return m.cleareduser
}

// RemoveUserIDs removes the "user" edge to the User entity by IDs.
func (m *RoleMutation) RemoveUserIDs(ids ...uuid.UUID) {
	if m.removeduser == nil {
		m.removeduser = make(map[uuid.UUID]struct{})
	}
	for i := range ids {
		delete(m.user, ids[i])
		m.removeduser[ids[i]] = struct{}{}
	}
}

// RemovedUser returns the removed IDs of the "user" edge to the User entity.
func (m *RoleMutation) RemovedUserIDs() (ids []uuid.UUID) {
	for id := range m.removeduser {
		ids = append(ids, id)
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
func (m *RoleMutation) UserIDs() (ids []uuid.UUID) {
	for id := range m.user {
		ids = append(ids, id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *RoleMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
	m.removeduser = nil
}

// Where appends a list predicates to the RoleMutation builder.
func (m *RoleMutation) Where(ps ...predicate.Role) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the RoleMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *RoleMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Role, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *RoleMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *RoleMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Role).
func (m *RoleMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RoleMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.uuid != nil {
		fields = append(fields, role.FieldUUID)
	}
	if m.key != nil {
		fields = append(fields, role.FieldKey)
	}
	if m.permissions != nil {
		fields = append(fields, role.FieldPermissions)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RoleMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case role.FieldUUID:
		return m.UUID()
	case role.FieldKey:
		return m.Key()
	case role.FieldPermissions:
		return m.Permissions()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RoleMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case role.FieldUUID:
		return m.OldUUID(ctx)
	case role.FieldKey:
		return m.OldKey(ctx)
	case role.FieldPermissions:
		return m.OldPermissions(ctx)
	}
	return nil, fmt.Errorf("unknown Role field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoleMutation) SetField(name string, value ent.Value) error {
	switch name {
	case role.FieldUUID:
		v, ok := value.(uuid.UUID)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUUID(v)
		return nil
	case role.FieldKey:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetKey(v)
		return nil
	case role.FieldPermissions:
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPermissions(v)
		return nil
	}
	return fmt.Errorf("unknown Role field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RoleMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RoleMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RoleMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Role numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RoleMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RoleMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RoleMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Role nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RoleMutation) ResetField(name string) error {
	switch name {
	case role.FieldUUID:
		m.ResetUUID()
		return nil
	case role.FieldKey:
		m.ResetKey()
		return nil
	case role.FieldPermissions:
		m.ResetPermissions()
		return nil
	}
	return fmt.Errorf("unknown Role field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RoleMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, role.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RoleMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case role.EdgeUser:
		ids := make([]ent.Value, 0, len(m.user))
		for id := range m.user {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RoleMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeduser != nil {
		edges = append(edges, role.EdgeUser)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RoleMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case role.EdgeUser:
		ids := make([]ent.Value, 0, len(m.removeduser))
		for id := range m.removeduser {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RoleMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, role.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RoleMutation) EdgeCleared(name string) bool {
	switch name {
	case role.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RoleMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Role unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RoleMutation) ResetEdge(name string) error {
	switch name {
	case role.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Role edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op            Op
	typ           string
	id            *uuid.UUID
	delete_time   *time.Time
	ban_time      *time.Time
	email         *string
	password      *string
	salt          *string
	clearedFields map[string]struct{}
	roles         map[int]struct{}
	removedroles  map[int]struct{}
	clearedroles  bool
	done          bool
	oldValue      func(context.Context) (*User, error)
	predicates    []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id uuid.UUID) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("build: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// SetID sets the value of the id field. Note that this
// operation is only accepted on creation of User entities.
func (m *UserMutation) SetID(id uuid.UUID) {
	m.id = &id
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id uuid.UUID, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]uuid.UUID, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []uuid.UUID{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetDeleteTime sets the "delete_time" field.
func (m *UserMutation) SetDeleteTime(t time.Time) {
	m.delete_time = &t
}

// DeleteTime returns the value of the "delete_time" field in the mutation.
func (m *UserMutation) DeleteTime() (r time.Time, exists bool) {
	v := m.delete_time
	if v == nil {
		return
	}
	return *v, true
}

// OldDeleteTime returns the old "delete_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldDeleteTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldDeleteTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldDeleteTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldDeleteTime: %w", err)
	}
	return oldValue.DeleteTime, nil
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (m *UserMutation) ClearDeleteTime() {
	m.delete_time = nil
	m.clearedFields[user.FieldDeleteTime] = struct{}{}
}

// DeleteTimeCleared returns if the "delete_time" field was cleared in this mutation.
func (m *UserMutation) DeleteTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldDeleteTime]
	return ok
}

// ResetDeleteTime resets all changes to the "delete_time" field.
func (m *UserMutation) ResetDeleteTime() {
	m.delete_time = nil
	delete(m.clearedFields, user.FieldDeleteTime)
}

// SetBanTime sets the "ban_time" field.
func (m *UserMutation) SetBanTime(t time.Time) {
	m.ban_time = &t
}

// BanTime returns the value of the "ban_time" field in the mutation.
func (m *UserMutation) BanTime() (r time.Time, exists bool) {
	v := m.ban_time
	if v == nil {
		return
	}
	return *v, true
}

// OldBanTime returns the old "ban_time" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldBanTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBanTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBanTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBanTime: %w", err)
	}
	return oldValue.BanTime, nil
}

// ClearBanTime clears the value of the "ban_time" field.
func (m *UserMutation) ClearBanTime() {
	m.ban_time = nil
	m.clearedFields[user.FieldBanTime] = struct{}{}
}

// BanTimeCleared returns if the "ban_time" field was cleared in this mutation.
func (m *UserMutation) BanTimeCleared() bool {
	_, ok := m.clearedFields[user.FieldBanTime]
	return ok
}

// ResetBanTime resets all changes to the "ban_time" field.
func (m *UserMutation) ResetBanTime() {
	m.ban_time = nil
	delete(m.clearedFields, user.FieldBanTime)
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetPassword sets the "password" field.
func (m *UserMutation) SetPassword(s string) {
	m.password = &s
}

// Password returns the value of the "password" field in the mutation.
func (m *UserMutation) Password() (r string, exists bool) {
	v := m.password
	if v == nil {
		return
	}
	return *v, true
}

// OldPassword returns the old "password" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldPassword(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPassword is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPassword requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPassword: %w", err)
	}
	return oldValue.Password, nil
}

// ResetPassword resets all changes to the "password" field.
func (m *UserMutation) ResetPassword() {
	m.password = nil
}

// SetSalt sets the "salt" field.
func (m *UserMutation) SetSalt(s string) {
	m.salt = &s
}

// Salt returns the value of the "salt" field in the mutation.
func (m *UserMutation) Salt() (r string, exists bool) {
	v := m.salt
	if v == nil {
		return
	}
	return *v, true
}

// OldSalt returns the old "salt" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldSalt(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldSalt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldSalt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldSalt: %w", err)
	}
	return oldValue.Salt, nil
}

// ResetSalt resets all changes to the "salt" field.
func (m *UserMutation) ResetSalt() {
	m.salt = nil
}

// AddRoleIDs adds the "roles" edge to the Role entity by ids.
func (m *UserMutation) AddRoleIDs(ids ...int) {
	if m.roles == nil {
		m.roles = make(map[int]struct{})
	}
	for i := range ids {
		m.roles[ids[i]] = struct{}{}
	}
}

// ClearRoles clears the "roles" edge to the Role entity.
func (m *UserMutation) ClearRoles() {
	m.clearedroles = true
}

// RolesCleared reports if the "roles" edge to the Role entity was cleared.
func (m *UserMutation) RolesCleared() bool {
	return m.clearedroles
}

// RemoveRoleIDs removes the "roles" edge to the Role entity by IDs.
func (m *UserMutation) RemoveRoleIDs(ids ...int) {
	if m.removedroles == nil {
		m.removedroles = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.roles, ids[i])
		m.removedroles[ids[i]] = struct{}{}
	}
}

// RemovedRoles returns the removed IDs of the "roles" edge to the Role entity.
func (m *UserMutation) RemovedRolesIDs() (ids []int) {
	for id := range m.removedroles {
		ids = append(ids, id)
	}
	return
}

// RolesIDs returns the "roles" edge IDs in the mutation.
func (m *UserMutation) RolesIDs() (ids []int) {
	for id := range m.roles {
		ids = append(ids, id)
	}
	return
}

// ResetRoles resets all changes to the "roles" edge.
func (m *UserMutation) ResetRoles() {
	m.roles = nil
	m.clearedroles = false
	m.removedroles = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.delete_time != nil {
		fields = append(fields, user.FieldDeleteTime)
	}
	if m.ban_time != nil {
		fields = append(fields, user.FieldBanTime)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.password != nil {
		fields = append(fields, user.FieldPassword)
	}
	if m.salt != nil {
		fields = append(fields, user.FieldSalt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldDeleteTime:
		return m.DeleteTime()
	case user.FieldBanTime:
		return m.BanTime()
	case user.FieldEmail:
		return m.Email()
	case user.FieldPassword:
		return m.Password()
	case user.FieldSalt:
		return m.Salt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldDeleteTime:
		return m.OldDeleteTime(ctx)
	case user.FieldBanTime:
		return m.OldBanTime(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldPassword:
		return m.OldPassword(ctx)
	case user.FieldSalt:
		return m.OldSalt(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldDeleteTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetDeleteTime(v)
		return nil
	case user.FieldBanTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBanTime(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldPassword:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPassword(v)
		return nil
	case user.FieldSalt:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetSalt(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldDeleteTime) {
		fields = append(fields, user.FieldDeleteTime)
	}
	if m.FieldCleared(user.FieldBanTime) {
		fields = append(fields, user.FieldBanTime)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldDeleteTime:
		m.ClearDeleteTime()
		return nil
	case user.FieldBanTime:
		m.ClearBanTime()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldDeleteTime:
		m.ResetDeleteTime()
		return nil
	case user.FieldBanTime:
		m.ResetBanTime()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldPassword:
		m.ResetPassword()
		return nil
	case user.FieldSalt:
		m.ResetSalt()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.roles != nil {
		edges = append(edges, user.EdgeRoles)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeRoles:
		ids := make([]ent.Value, 0, len(m.roles))
		for id := range m.roles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedroles != nil {
		edges = append(edges, user.EdgeRoles)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeRoles:
		ids := make([]ent.Value, 0, len(m.removedroles))
		for id := range m.removedroles {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedroles {
		edges = append(edges, user.EdgeRoles)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeRoles:
		return m.clearedroles
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeRoles:
		m.ResetRoles()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}