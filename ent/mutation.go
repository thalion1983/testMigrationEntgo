// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"testMigrationEntgo/ent/blog"
	"testMigrationEntgo/ent/predicate"
	"testMigrationEntgo/ent/user"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeBlog = "Blog"
	TypeUser = "User"
)

// BlogMutation represents an operation that mutates the Blog nodes in the graph.
type BlogMutation struct {
	config
	op            Op
	typ           string
	id            *int
	title         *string
	body          *string
	created_at    *time.Time
	clearedFields map[string]struct{}
	author        *int
	clearedauthor bool
	done          bool
	oldValue      func(context.Context) (*Blog, error)
	predicates    []predicate.Blog
}

var _ ent.Mutation = (*BlogMutation)(nil)

// blogOption allows management of the mutation configuration using functional options.
type blogOption func(*BlogMutation)

// newBlogMutation creates new mutation for the Blog entity.
func newBlogMutation(c config, op Op, opts ...blogOption) *BlogMutation {
	m := &BlogMutation{
		config:        c,
		op:            op,
		typ:           TypeBlog,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withBlogID sets the ID field of the mutation.
func withBlogID(id int) blogOption {
	return func(m *BlogMutation) {
		var (
			err   error
			once  sync.Once
			value *Blog
		)
		m.oldValue = func(ctx context.Context) (*Blog, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Blog.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withBlog sets the old Blog of the mutation.
func withBlog(node *Blog) blogOption {
	return func(m *BlogMutation) {
		m.oldValue = func(context.Context) (*Blog, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m BlogMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m BlogMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *BlogMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *BlogMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Blog.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetTitle sets the "title" field.
func (m *BlogMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *BlogMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the Blog entity.
// If the Blog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BlogMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ResetTitle resets all changes to the "title" field.
func (m *BlogMutation) ResetTitle() {
	m.title = nil
}

// SetBody sets the "body" field.
func (m *BlogMutation) SetBody(s string) {
	m.body = &s
}

// Body returns the value of the "body" field in the mutation.
func (m *BlogMutation) Body() (r string, exists bool) {
	v := m.body
	if v == nil {
		return
	}
	return *v, true
}

// OldBody returns the old "body" field's value of the Blog entity.
// If the Blog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BlogMutation) OldBody(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldBody is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldBody requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldBody: %w", err)
	}
	return oldValue.Body, nil
}

// ResetBody resets all changes to the "body" field.
func (m *BlogMutation) ResetBody() {
	m.body = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *BlogMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *BlogMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Blog entity.
// If the Blog object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *BlogMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *BlogMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetAuthorID sets the "author" edge to the User entity by id.
func (m *BlogMutation) SetAuthorID(id int) {
	m.author = &id
}

// ClearAuthor clears the "author" edge to the User entity.
func (m *BlogMutation) ClearAuthor() {
	m.clearedauthor = true
}

// AuthorCleared reports if the "author" edge to the User entity was cleared.
func (m *BlogMutation) AuthorCleared() bool {
	return m.clearedauthor
}

// AuthorID returns the "author" edge ID in the mutation.
func (m *BlogMutation) AuthorID() (id int, exists bool) {
	if m.author != nil {
		return *m.author, true
	}
	return
}

// AuthorIDs returns the "author" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// AuthorID instead. It exists only for internal usage by the builders.
func (m *BlogMutation) AuthorIDs() (ids []int) {
	if id := m.author; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetAuthor resets all changes to the "author" edge.
func (m *BlogMutation) ResetAuthor() {
	m.author = nil
	m.clearedauthor = false
}

// Where appends a list predicates to the BlogMutation builder.
func (m *BlogMutation) Where(ps ...predicate.Blog) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the BlogMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *BlogMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Blog, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *BlogMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *BlogMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Blog).
func (m *BlogMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *BlogMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.title != nil {
		fields = append(fields, blog.FieldTitle)
	}
	if m.body != nil {
		fields = append(fields, blog.FieldBody)
	}
	if m.created_at != nil {
		fields = append(fields, blog.FieldCreatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *BlogMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case blog.FieldTitle:
		return m.Title()
	case blog.FieldBody:
		return m.Body()
	case blog.FieldCreatedAt:
		return m.CreatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *BlogMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case blog.FieldTitle:
		return m.OldTitle(ctx)
	case blog.FieldBody:
		return m.OldBody(ctx)
	case blog.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Blog field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BlogMutation) SetField(name string, value ent.Value) error {
	switch name {
	case blog.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case blog.FieldBody:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetBody(v)
		return nil
	case blog.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Blog field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *BlogMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *BlogMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *BlogMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Blog numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *BlogMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *BlogMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *BlogMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Blog nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *BlogMutation) ResetField(name string) error {
	switch name {
	case blog.FieldTitle:
		m.ResetTitle()
		return nil
	case blog.FieldBody:
		m.ResetBody()
		return nil
	case blog.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	}
	return fmt.Errorf("unknown Blog field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *BlogMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.author != nil {
		edges = append(edges, blog.EdgeAuthor)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *BlogMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case blog.EdgeAuthor:
		if id := m.author; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *BlogMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *BlogMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *BlogMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedauthor {
		edges = append(edges, blog.EdgeAuthor)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *BlogMutation) EdgeCleared(name string) bool {
	switch name {
	case blog.EdgeAuthor:
		return m.clearedauthor
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *BlogMutation) ClearEdge(name string) error {
	switch name {
	case blog.EdgeAuthor:
		m.ClearAuthor()
		return nil
	}
	return fmt.Errorf("unknown Blog unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *BlogMutation) ResetEdge(name string) error {
	switch name {
	case blog.EdgeAuthor:
		m.ResetAuthor()
		return nil
	}
	return fmt.Errorf("unknown Blog edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op                Op
	typ               string
	id                *int
	name              *string
	email             *string
	title             *string
	followers         *int
	addfollowers      *int
	clearedFields     map[string]struct{}
	blog_posts        map[int]struct{}
	removedblog_posts map[int]struct{}
	clearedblog_posts bool
	done              bool
	oldValue          func(context.Context) (*User, error)
	predicates        []predicate.User
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
func withUserID(id int) userOption {
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
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
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

// SetTitle sets the "title" field.
func (m *UserMutation) SetTitle(s string) {
	m.title = &s
}

// Title returns the value of the "title" field in the mutation.
func (m *UserMutation) Title() (r string, exists bool) {
	v := m.title
	if v == nil {
		return
	}
	return *v, true
}

// OldTitle returns the old "title" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldTitle(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldTitle is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldTitle requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldTitle: %w", err)
	}
	return oldValue.Title, nil
}

// ClearTitle clears the value of the "title" field.
func (m *UserMutation) ClearTitle() {
	m.title = nil
	m.clearedFields[user.FieldTitle] = struct{}{}
}

// TitleCleared returns if the "title" field was cleared in this mutation.
func (m *UserMutation) TitleCleared() bool {
	_, ok := m.clearedFields[user.FieldTitle]
	return ok
}

// ResetTitle resets all changes to the "title" field.
func (m *UserMutation) ResetTitle() {
	m.title = nil
	delete(m.clearedFields, user.FieldTitle)
}

// SetFollowers sets the "followers" field.
func (m *UserMutation) SetFollowers(i int) {
	m.followers = &i
	m.addfollowers = nil
}

// Followers returns the value of the "followers" field in the mutation.
func (m *UserMutation) Followers() (r int, exists bool) {
	v := m.followers
	if v == nil {
		return
	}
	return *v, true
}

// OldFollowers returns the old "followers" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldFollowers(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldFollowers is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldFollowers requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFollowers: %w", err)
	}
	return oldValue.Followers, nil
}

// AddFollowers adds i to the "followers" field.
func (m *UserMutation) AddFollowers(i int) {
	if m.addfollowers != nil {
		*m.addfollowers += i
	} else {
		m.addfollowers = &i
	}
}

// AddedFollowers returns the value that was added to the "followers" field in this mutation.
func (m *UserMutation) AddedFollowers() (r int, exists bool) {
	v := m.addfollowers
	if v == nil {
		return
	}
	return *v, true
}

// ClearFollowers clears the value of the "followers" field.
func (m *UserMutation) ClearFollowers() {
	m.followers = nil
	m.addfollowers = nil
	m.clearedFields[user.FieldFollowers] = struct{}{}
}

// FollowersCleared returns if the "followers" field was cleared in this mutation.
func (m *UserMutation) FollowersCleared() bool {
	_, ok := m.clearedFields[user.FieldFollowers]
	return ok
}

// ResetFollowers resets all changes to the "followers" field.
func (m *UserMutation) ResetFollowers() {
	m.followers = nil
	m.addfollowers = nil
	delete(m.clearedFields, user.FieldFollowers)
}

// AddBlogPostIDs adds the "blog_posts" edge to the Blog entity by ids.
func (m *UserMutation) AddBlogPostIDs(ids ...int) {
	if m.blog_posts == nil {
		m.blog_posts = make(map[int]struct{})
	}
	for i := range ids {
		m.blog_posts[ids[i]] = struct{}{}
	}
}

// ClearBlogPosts clears the "blog_posts" edge to the Blog entity.
func (m *UserMutation) ClearBlogPosts() {
	m.clearedblog_posts = true
}

// BlogPostsCleared reports if the "blog_posts" edge to the Blog entity was cleared.
func (m *UserMutation) BlogPostsCleared() bool {
	return m.clearedblog_posts
}

// RemoveBlogPostIDs removes the "blog_posts" edge to the Blog entity by IDs.
func (m *UserMutation) RemoveBlogPostIDs(ids ...int) {
	if m.removedblog_posts == nil {
		m.removedblog_posts = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.blog_posts, ids[i])
		m.removedblog_posts[ids[i]] = struct{}{}
	}
}

// RemovedBlogPosts returns the removed IDs of the "blog_posts" edge to the Blog entity.
func (m *UserMutation) RemovedBlogPostsIDs() (ids []int) {
	for id := range m.removedblog_posts {
		ids = append(ids, id)
	}
	return
}

// BlogPostsIDs returns the "blog_posts" edge IDs in the mutation.
func (m *UserMutation) BlogPostsIDs() (ids []int) {
	for id := range m.blog_posts {
		ids = append(ids, id)
	}
	return
}

// ResetBlogPosts resets all changes to the "blog_posts" edge.
func (m *UserMutation) ResetBlogPosts() {
	m.blog_posts = nil
	m.clearedblog_posts = false
	m.removedblog_posts = nil
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
	fields := make([]string, 0, 4)
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.title != nil {
		fields = append(fields, user.FieldTitle)
	}
	if m.followers != nil {
		fields = append(fields, user.FieldFollowers)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldName:
		return m.Name()
	case user.FieldEmail:
		return m.Email()
	case user.FieldTitle:
		return m.Title()
	case user.FieldFollowers:
		return m.Followers()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldTitle:
		return m.OldTitle(ctx)
	case user.FieldFollowers:
		return m.OldFollowers(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldTitle:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetTitle(v)
		return nil
	case user.FieldFollowers:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFollowers(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	var fields []string
	if m.addfollowers != nil {
		fields = append(fields, user.FieldFollowers)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case user.FieldFollowers:
		return m.AddedFollowers()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	case user.FieldFollowers:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddFollowers(v)
		return nil
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldTitle) {
		fields = append(fields, user.FieldTitle)
	}
	if m.FieldCleared(user.FieldFollowers) {
		fields = append(fields, user.FieldFollowers)
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
	case user.FieldTitle:
		m.ClearTitle()
		return nil
	case user.FieldFollowers:
		m.ClearFollowers()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldTitle:
		m.ResetTitle()
		return nil
	case user.FieldFollowers:
		m.ResetFollowers()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.blog_posts != nil {
		edges = append(edges, user.EdgeBlogPosts)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeBlogPosts:
		ids := make([]ent.Value, 0, len(m.blog_posts))
		for id := range m.blog_posts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedblog_posts != nil {
		edges = append(edges, user.EdgeBlogPosts)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeBlogPosts:
		ids := make([]ent.Value, 0, len(m.removedblog_posts))
		for id := range m.removedblog_posts {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedblog_posts {
		edges = append(edges, user.EdgeBlogPosts)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeBlogPosts:
		return m.clearedblog_posts
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
	case user.EdgeBlogPosts:
		m.ResetBlogPosts()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
