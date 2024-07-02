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
	"github.com/luminancetech/varso/src/services/analytics/internal/ent/build/accesslog"
	"github.com/luminancetech/varso/src/services/analytics/internal/ent/build/predicate"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeAccessLog = "AccessLog"
)

// AccessLogMutation represents an operation that mutates the AccessLog nodes in the graph.
type AccessLogMutation struct {
	config
	op               Op
	typ              string
	id               *int
	create_time      *time.Time
	ip               *string
	uri              *string
	forwarded_for    *string
	forwarded_proto  *string
	forwarded_host   *string
	forwarded_port   *string
	forwarded_server *string
	request_id       *string
	user_agent       *string
	clearedFields    map[string]struct{}
	done             bool
	oldValue         func(context.Context) (*AccessLog, error)
	predicates       []predicate.AccessLog
}

var _ ent.Mutation = (*AccessLogMutation)(nil)

// accesslogOption allows management of the mutation configuration using functional options.
type accesslogOption func(*AccessLogMutation)

// newAccessLogMutation creates new mutation for the AccessLog entity.
func newAccessLogMutation(c config, op Op, opts ...accesslogOption) *AccessLogMutation {
	m := &AccessLogMutation{
		config:        c,
		op:            op,
		typ:           TypeAccessLog,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withAccessLogID sets the ID field of the mutation.
func withAccessLogID(id int) accesslogOption {
	return func(m *AccessLogMutation) {
		var (
			err   error
			once  sync.Once
			value *AccessLog
		)
		m.oldValue = func(ctx context.Context) (*AccessLog, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().AccessLog.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withAccessLog sets the old AccessLog of the mutation.
func withAccessLog(node *AccessLog) accesslogOption {
	return func(m *AccessLogMutation) {
		m.oldValue = func(context.Context) (*AccessLog, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m AccessLogMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m AccessLogMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("build: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *AccessLogMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *AccessLogMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().AccessLog.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreateTime sets the "create_time" field.
func (m *AccessLogMutation) SetCreateTime(t time.Time) {
	m.create_time = &t
}

// CreateTime returns the value of the "create_time" field in the mutation.
func (m *AccessLogMutation) CreateTime() (r time.Time, exists bool) {
	v := m.create_time
	if v == nil {
		return
	}
	return *v, true
}

// OldCreateTime returns the old "create_time" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldCreateTime(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreateTime is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreateTime requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreateTime: %w", err)
	}
	return oldValue.CreateTime, nil
}

// ResetCreateTime resets all changes to the "create_time" field.
func (m *AccessLogMutation) ResetCreateTime() {
	m.create_time = nil
}

// SetIP sets the "ip" field.
func (m *AccessLogMutation) SetIP(s string) {
	m.ip = &s
}

// IP returns the value of the "ip" field in the mutation.
func (m *AccessLogMutation) IP() (r string, exists bool) {
	v := m.ip
	if v == nil {
		return
	}
	return *v, true
}

// OldIP returns the old "ip" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldIP(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldIP is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldIP requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldIP: %w", err)
	}
	return oldValue.IP, nil
}

// ResetIP resets all changes to the "ip" field.
func (m *AccessLogMutation) ResetIP() {
	m.ip = nil
}

// SetURI sets the "uri" field.
func (m *AccessLogMutation) SetURI(s string) {
	m.uri = &s
}

// URI returns the value of the "uri" field in the mutation.
func (m *AccessLogMutation) URI() (r string, exists bool) {
	v := m.uri
	if v == nil {
		return
	}
	return *v, true
}

// OldURI returns the old "uri" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldURI(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldURI is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldURI requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldURI: %w", err)
	}
	return oldValue.URI, nil
}

// ResetURI resets all changes to the "uri" field.
func (m *AccessLogMutation) ResetURI() {
	m.uri = nil
}

// SetForwardedFor sets the "forwarded_for" field.
func (m *AccessLogMutation) SetForwardedFor(s string) {
	m.forwarded_for = &s
}

// ForwardedFor returns the value of the "forwarded_for" field in the mutation.
func (m *AccessLogMutation) ForwardedFor() (r string, exists bool) {
	v := m.forwarded_for
	if v == nil {
		return
	}
	return *v, true
}

// OldForwardedFor returns the old "forwarded_for" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldForwardedFor(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldForwardedFor is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldForwardedFor requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldForwardedFor: %w", err)
	}
	return oldValue.ForwardedFor, nil
}

// ResetForwardedFor resets all changes to the "forwarded_for" field.
func (m *AccessLogMutation) ResetForwardedFor() {
	m.forwarded_for = nil
}

// SetForwardedProto sets the "forwarded_proto" field.
func (m *AccessLogMutation) SetForwardedProto(s string) {
	m.forwarded_proto = &s
}

// ForwardedProto returns the value of the "forwarded_proto" field in the mutation.
func (m *AccessLogMutation) ForwardedProto() (r string, exists bool) {
	v := m.forwarded_proto
	if v == nil {
		return
	}
	return *v, true
}

// OldForwardedProto returns the old "forwarded_proto" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldForwardedProto(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldForwardedProto is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldForwardedProto requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldForwardedProto: %w", err)
	}
	return oldValue.ForwardedProto, nil
}

// ResetForwardedProto resets all changes to the "forwarded_proto" field.
func (m *AccessLogMutation) ResetForwardedProto() {
	m.forwarded_proto = nil
}

// SetForwardedHost sets the "forwarded_host" field.
func (m *AccessLogMutation) SetForwardedHost(s string) {
	m.forwarded_host = &s
}

// ForwardedHost returns the value of the "forwarded_host" field in the mutation.
func (m *AccessLogMutation) ForwardedHost() (r string, exists bool) {
	v := m.forwarded_host
	if v == nil {
		return
	}
	return *v, true
}

// OldForwardedHost returns the old "forwarded_host" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldForwardedHost(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldForwardedHost is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldForwardedHost requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldForwardedHost: %w", err)
	}
	return oldValue.ForwardedHost, nil
}

// ResetForwardedHost resets all changes to the "forwarded_host" field.
func (m *AccessLogMutation) ResetForwardedHost() {
	m.forwarded_host = nil
}

// SetForwardedPort sets the "forwarded_port" field.
func (m *AccessLogMutation) SetForwardedPort(s string) {
	m.forwarded_port = &s
}

// ForwardedPort returns the value of the "forwarded_port" field in the mutation.
func (m *AccessLogMutation) ForwardedPort() (r string, exists bool) {
	v := m.forwarded_port
	if v == nil {
		return
	}
	return *v, true
}

// OldForwardedPort returns the old "forwarded_port" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldForwardedPort(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldForwardedPort is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldForwardedPort requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldForwardedPort: %w", err)
	}
	return oldValue.ForwardedPort, nil
}

// ResetForwardedPort resets all changes to the "forwarded_port" field.
func (m *AccessLogMutation) ResetForwardedPort() {
	m.forwarded_port = nil
}

// SetForwardedServer sets the "forwarded_server" field.
func (m *AccessLogMutation) SetForwardedServer(s string) {
	m.forwarded_server = &s
}

// ForwardedServer returns the value of the "forwarded_server" field in the mutation.
func (m *AccessLogMutation) ForwardedServer() (r string, exists bool) {
	v := m.forwarded_server
	if v == nil {
		return
	}
	return *v, true
}

// OldForwardedServer returns the old "forwarded_server" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldForwardedServer(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldForwardedServer is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldForwardedServer requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldForwardedServer: %w", err)
	}
	return oldValue.ForwardedServer, nil
}

// ResetForwardedServer resets all changes to the "forwarded_server" field.
func (m *AccessLogMutation) ResetForwardedServer() {
	m.forwarded_server = nil
}

// SetRequestID sets the "request_id" field.
func (m *AccessLogMutation) SetRequestID(s string) {
	m.request_id = &s
}

// RequestID returns the value of the "request_id" field in the mutation.
func (m *AccessLogMutation) RequestID() (r string, exists bool) {
	v := m.request_id
	if v == nil {
		return
	}
	return *v, true
}

// OldRequestID returns the old "request_id" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldRequestID(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldRequestID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldRequestID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRequestID: %w", err)
	}
	return oldValue.RequestID, nil
}

// ResetRequestID resets all changes to the "request_id" field.
func (m *AccessLogMutation) ResetRequestID() {
	m.request_id = nil
}

// SetUserAgent sets the "user_agent" field.
func (m *AccessLogMutation) SetUserAgent(s string) {
	m.user_agent = &s
}

// UserAgent returns the value of the "user_agent" field in the mutation.
func (m *AccessLogMutation) UserAgent() (r string, exists bool) {
	v := m.user_agent
	if v == nil {
		return
	}
	return *v, true
}

// OldUserAgent returns the old "user_agent" field's value of the AccessLog entity.
// If the AccessLog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *AccessLogMutation) OldUserAgent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUserAgent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUserAgent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUserAgent: %w", err)
	}
	return oldValue.UserAgent, nil
}

// ResetUserAgent resets all changes to the "user_agent" field.
func (m *AccessLogMutation) ResetUserAgent() {
	m.user_agent = nil
}

// Where appends a list predicates to the AccessLogMutation builder.
func (m *AccessLogMutation) Where(ps ...predicate.AccessLog) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the AccessLogMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *AccessLogMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.AccessLog, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *AccessLogMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *AccessLogMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (AccessLog).
func (m *AccessLogMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *AccessLogMutation) Fields() []string {
	fields := make([]string, 0, 10)
	if m.create_time != nil {
		fields = append(fields, accesslog.FieldCreateTime)
	}
	if m.ip != nil {
		fields = append(fields, accesslog.FieldIP)
	}
	if m.uri != nil {
		fields = append(fields, accesslog.FieldURI)
	}
	if m.forwarded_for != nil {
		fields = append(fields, accesslog.FieldForwardedFor)
	}
	if m.forwarded_proto != nil {
		fields = append(fields, accesslog.FieldForwardedProto)
	}
	if m.forwarded_host != nil {
		fields = append(fields, accesslog.FieldForwardedHost)
	}
	if m.forwarded_port != nil {
		fields = append(fields, accesslog.FieldForwardedPort)
	}
	if m.forwarded_server != nil {
		fields = append(fields, accesslog.FieldForwardedServer)
	}
	if m.request_id != nil {
		fields = append(fields, accesslog.FieldRequestID)
	}
	if m.user_agent != nil {
		fields = append(fields, accesslog.FieldUserAgent)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *AccessLogMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case accesslog.FieldCreateTime:
		return m.CreateTime()
	case accesslog.FieldIP:
		return m.IP()
	case accesslog.FieldURI:
		return m.URI()
	case accesslog.FieldForwardedFor:
		return m.ForwardedFor()
	case accesslog.FieldForwardedProto:
		return m.ForwardedProto()
	case accesslog.FieldForwardedHost:
		return m.ForwardedHost()
	case accesslog.FieldForwardedPort:
		return m.ForwardedPort()
	case accesslog.FieldForwardedServer:
		return m.ForwardedServer()
	case accesslog.FieldRequestID:
		return m.RequestID()
	case accesslog.FieldUserAgent:
		return m.UserAgent()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *AccessLogMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case accesslog.FieldCreateTime:
		return m.OldCreateTime(ctx)
	case accesslog.FieldIP:
		return m.OldIP(ctx)
	case accesslog.FieldURI:
		return m.OldURI(ctx)
	case accesslog.FieldForwardedFor:
		return m.OldForwardedFor(ctx)
	case accesslog.FieldForwardedProto:
		return m.OldForwardedProto(ctx)
	case accesslog.FieldForwardedHost:
		return m.OldForwardedHost(ctx)
	case accesslog.FieldForwardedPort:
		return m.OldForwardedPort(ctx)
	case accesslog.FieldForwardedServer:
		return m.OldForwardedServer(ctx)
	case accesslog.FieldRequestID:
		return m.OldRequestID(ctx)
	case accesslog.FieldUserAgent:
		return m.OldUserAgent(ctx)
	}
	return nil, fmt.Errorf("unknown AccessLog field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccessLogMutation) SetField(name string, value ent.Value) error {
	switch name {
	case accesslog.FieldCreateTime:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreateTime(v)
		return nil
	case accesslog.FieldIP:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetIP(v)
		return nil
	case accesslog.FieldURI:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetURI(v)
		return nil
	case accesslog.FieldForwardedFor:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetForwardedFor(v)
		return nil
	case accesslog.FieldForwardedProto:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetForwardedProto(v)
		return nil
	case accesslog.FieldForwardedHost:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetForwardedHost(v)
		return nil
	case accesslog.FieldForwardedPort:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetForwardedPort(v)
		return nil
	case accesslog.FieldForwardedServer:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetForwardedServer(v)
		return nil
	case accesslog.FieldRequestID:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRequestID(v)
		return nil
	case accesslog.FieldUserAgent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUserAgent(v)
		return nil
	}
	return fmt.Errorf("unknown AccessLog field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *AccessLogMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *AccessLogMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *AccessLogMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown AccessLog numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *AccessLogMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *AccessLogMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *AccessLogMutation) ClearField(name string) error {
	return fmt.Errorf("unknown AccessLog nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *AccessLogMutation) ResetField(name string) error {
	switch name {
	case accesslog.FieldCreateTime:
		m.ResetCreateTime()
		return nil
	case accesslog.FieldIP:
		m.ResetIP()
		return nil
	case accesslog.FieldURI:
		m.ResetURI()
		return nil
	case accesslog.FieldForwardedFor:
		m.ResetForwardedFor()
		return nil
	case accesslog.FieldForwardedProto:
		m.ResetForwardedProto()
		return nil
	case accesslog.FieldForwardedHost:
		m.ResetForwardedHost()
		return nil
	case accesslog.FieldForwardedPort:
		m.ResetForwardedPort()
		return nil
	case accesslog.FieldForwardedServer:
		m.ResetForwardedServer()
		return nil
	case accesslog.FieldRequestID:
		m.ResetRequestID()
		return nil
	case accesslog.FieldUserAgent:
		m.ResetUserAgent()
		return nil
	}
	return fmt.Errorf("unknown AccessLog field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *AccessLogMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *AccessLogMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *AccessLogMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *AccessLogMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *AccessLogMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *AccessLogMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *AccessLogMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown AccessLog unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *AccessLogMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown AccessLog edge %s", name)
}