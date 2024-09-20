// Code generated by ent, DO NOT EDIT.

package models

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/env"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/internal"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/predicate"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/team"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/teamapikey"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/tier"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/user"
	"github.com/e2b-dev/infra/packages/shared/pkg/models/usersteams"
	"github.com/google/uuid"
)

// TeamUpdate is the builder for updating Team entities.
type TeamUpdate struct {
	config
	hooks     []Hook
	mutation  *TeamMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TeamUpdate builder.
func (tu *TeamUpdate) Where(ps ...predicate.Team) *TeamUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetIsBanned sets the "is_banned" field.
func (tu *TeamUpdate) SetIsBanned(b bool) *TeamUpdate {
	tu.mutation.SetIsBanned(b)
	return tu
}

// SetNillableIsBanned sets the "is_banned" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableIsBanned(b *bool) *TeamUpdate {
	if b != nil {
		tu.SetIsBanned(*b)
	}
	return tu
}

// SetIsBlocked sets the "is_blocked" field.
func (tu *TeamUpdate) SetIsBlocked(b bool) *TeamUpdate {
	tu.mutation.SetIsBlocked(b)
	return tu
}

// SetNillableIsBlocked sets the "is_blocked" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableIsBlocked(b *bool) *TeamUpdate {
	if b != nil {
		tu.SetIsBlocked(*b)
	}
	return tu
}

// SetBlockedReason sets the "blocked_reason" field.
func (tu *TeamUpdate) SetBlockedReason(s string) *TeamUpdate {
	tu.mutation.SetBlockedReason(s)
	return tu
}

// SetNillableBlockedReason sets the "blocked_reason" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableBlockedReason(s *string) *TeamUpdate {
	if s != nil {
		tu.SetBlockedReason(*s)
	}
	return tu
}

// ClearBlockedReason clears the value of the "blocked_reason" field.
func (tu *TeamUpdate) ClearBlockedReason() *TeamUpdate {
	tu.mutation.ClearBlockedReason()
	return tu
}

// SetName sets the "name" field.
func (tu *TeamUpdate) SetName(s string) *TeamUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableName(s *string) *TeamUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// SetTier sets the "tier" field.
func (tu *TeamUpdate) SetTier(s string) *TeamUpdate {
	tu.mutation.SetTier(s)
	return tu
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableTier(s *string) *TeamUpdate {
	if s != nil {
		tu.SetTier(*s)
	}
	return tu
}

// SetEmail sets the "email" field.
func (tu *TeamUpdate) SetEmail(s string) *TeamUpdate {
	tu.mutation.SetEmail(s)
	return tu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (tu *TeamUpdate) SetNillableEmail(s *string) *TeamUpdate {
	if s != nil {
		tu.SetEmail(*s)
	}
	return tu
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tu *TeamUpdate) AddUserIDs(ids ...uuid.UUID) *TeamUpdate {
	tu.mutation.AddUserIDs(ids...)
	return tu
}

// AddUsers adds the "users" edges to the User entity.
func (tu *TeamUpdate) AddUsers(u ...*User) *TeamUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUserIDs(ids...)
}

// AddTeamAPIKeyIDs adds the "team_api_keys" edge to the TeamAPIKey entity by IDs.
func (tu *TeamUpdate) AddTeamAPIKeyIDs(ids ...string) *TeamUpdate {
	tu.mutation.AddTeamAPIKeyIDs(ids...)
	return tu
}

// AddTeamAPIKeys adds the "team_api_keys" edges to the TeamAPIKey entity.
func (tu *TeamUpdate) AddTeamAPIKeys(t ...*TeamAPIKey) *TeamUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTeamAPIKeyIDs(ids...)
}

// SetTeamTierID sets the "team_tier" edge to the Tier entity by ID.
func (tu *TeamUpdate) SetTeamTierID(id string) *TeamUpdate {
	tu.mutation.SetTeamTierID(id)
	return tu
}

// SetTeamTier sets the "team_tier" edge to the Tier entity.
func (tu *TeamUpdate) SetTeamTier(t *Tier) *TeamUpdate {
	return tu.SetTeamTierID(t.ID)
}

// AddEnvIDs adds the "envs" edge to the Env entity by IDs.
func (tu *TeamUpdate) AddEnvIDs(ids ...string) *TeamUpdate {
	tu.mutation.AddEnvIDs(ids...)
	return tu
}

// AddEnvs adds the "envs" edges to the Env entity.
func (tu *TeamUpdate) AddEnvs(e ...*Env) *TeamUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tu.AddEnvIDs(ids...)
}

// AddUsersTeamIDs adds the "users_teams" edge to the UsersTeams entity by IDs.
func (tu *TeamUpdate) AddUsersTeamIDs(ids ...int) *TeamUpdate {
	tu.mutation.AddUsersTeamIDs(ids...)
	return tu
}

// AddUsersTeams adds the "users_teams" edges to the UsersTeams entity.
func (tu *TeamUpdate) AddUsersTeams(u ...*UsersTeams) *TeamUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.AddUsersTeamIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tu *TeamUpdate) Mutation() *TeamMutation {
	return tu.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (tu *TeamUpdate) ClearUsers() *TeamUpdate {
	tu.mutation.ClearUsers()
	return tu
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (tu *TeamUpdate) RemoveUserIDs(ids ...uuid.UUID) *TeamUpdate {
	tu.mutation.RemoveUserIDs(ids...)
	return tu
}

// RemoveUsers removes "users" edges to User entities.
func (tu *TeamUpdate) RemoveUsers(u ...*User) *TeamUpdate {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUserIDs(ids...)
}

// ClearTeamAPIKeys clears all "team_api_keys" edges to the TeamAPIKey entity.
func (tu *TeamUpdate) ClearTeamAPIKeys() *TeamUpdate {
	tu.mutation.ClearTeamAPIKeys()
	return tu
}

// RemoveTeamAPIKeyIDs removes the "team_api_keys" edge to TeamAPIKey entities by IDs.
func (tu *TeamUpdate) RemoveTeamAPIKeyIDs(ids ...string) *TeamUpdate {
	tu.mutation.RemoveTeamAPIKeyIDs(ids...)
	return tu
}

// RemoveTeamAPIKeys removes "team_api_keys" edges to TeamAPIKey entities.
func (tu *TeamUpdate) RemoveTeamAPIKeys(t ...*TeamAPIKey) *TeamUpdate {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTeamAPIKeyIDs(ids...)
}

// ClearTeamTier clears the "team_tier" edge to the Tier entity.
func (tu *TeamUpdate) ClearTeamTier() *TeamUpdate {
	tu.mutation.ClearTeamTier()
	return tu
}

// ClearEnvs clears all "envs" edges to the Env entity.
func (tu *TeamUpdate) ClearEnvs() *TeamUpdate {
	tu.mutation.ClearEnvs()
	return tu
}

// RemoveEnvIDs removes the "envs" edge to Env entities by IDs.
func (tu *TeamUpdate) RemoveEnvIDs(ids ...string) *TeamUpdate {
	tu.mutation.RemoveEnvIDs(ids...)
	return tu
}

// RemoveEnvs removes "envs" edges to Env entities.
func (tu *TeamUpdate) RemoveEnvs(e ...*Env) *TeamUpdate {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tu.RemoveEnvIDs(ids...)
}

// ClearUsersTeams clears all "users_teams" edges to the UsersTeams entity.
func (tu *TeamUpdate) ClearUsersTeams() *TeamUpdate {
	tu.mutation.ClearUsersTeams()
	return tu
}

// RemoveUsersTeamIDs removes the "users_teams" edge to UsersTeams entities by IDs.
func (tu *TeamUpdate) RemoveUsersTeamIDs(ids ...int) *TeamUpdate {
	tu.mutation.RemoveUsersTeamIDs(ids...)
	return tu
}

// RemoveUsersTeams removes "users_teams" edges to UsersTeams entities.
func (tu *TeamUpdate) RemoveUsersTeams(u ...*UsersTeams) *TeamUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tu.RemoveUsersTeamIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TeamUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TeamUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TeamUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TeamUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TeamUpdate) check() error {
	if v, ok := tu.mutation.Email(); ok {
		if err := team.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`models: validator failed for field "Team.email": %w`, err)}
		}
	}
	if _, ok := tu.mutation.TeamTierID(); tu.mutation.TeamTierCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "Team.team_tier"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TeamUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TeamUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TeamUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.IsBanned(); ok {
		_spec.SetField(team.FieldIsBanned, field.TypeBool, value)
	}
	if value, ok := tu.mutation.IsBlocked(); ok {
		_spec.SetField(team.FieldIsBlocked, field.TypeBool, value)
	}
	if value, ok := tu.mutation.BlockedReason(); ok {
		_spec.SetField(team.FieldBlockedReason, field.TypeString, value)
	}
	if tu.mutation.BlockedReasonCleared() {
		_spec.ClearField(team.FieldBlockedReason, field.TypeString)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Email(); ok {
		_spec.SetField(team.FieldEmail, field.TypeString, value)
	}
	if tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		createE := &UsersTeamsCreate{config: tu.config, mutation: newUsersTeamsMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tu.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UsersTeamsCreate{config: tu.config, mutation: newUsersTeamsMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UsersTeamsCreate{config: tu.config, mutation: newUsersTeamsMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TeamAPIKeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamAPIKeysTable,
			Columns: []string{team.TeamAPIKeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.TeamAPIKey
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTeamAPIKeysIDs(); len(nodes) > 0 && !tu.mutation.TeamAPIKeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamAPIKeysTable,
			Columns: []string{team.TeamAPIKeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.TeamAPIKey
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TeamAPIKeysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamAPIKeysTable,
			Columns: []string{team.TeamAPIKeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.TeamAPIKey
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TeamTierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   team.TeamTierTable,
			Columns: []string{team.TeamTierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.Team
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TeamTierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   team.TeamTierTable,
			Columns: []string{team.TeamTierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.Team
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.EnvsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.EnvsTable,
			Columns: []string{team.EnvsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.Env
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedEnvsIDs(); len(nodes) > 0 && !tu.mutation.EnvsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.EnvsTable,
			Columns: []string{team.EnvsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.Env
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.EnvsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.EnvsTable,
			Columns: []string{team.EnvsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = tu.schemaConfig.Env
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.UsersTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedUsersTeamsIDs(); len(nodes) > 0 && !tu.mutation.UsersTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.UsersTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tu.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tu.schemaConfig.Team
	ctx = internal.NewSchemaConfigContext(ctx, tu.schemaConfig)
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TeamUpdateOne is the builder for updating a single Team entity.
type TeamUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TeamMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetIsBanned sets the "is_banned" field.
func (tuo *TeamUpdateOne) SetIsBanned(b bool) *TeamUpdateOne {
	tuo.mutation.SetIsBanned(b)
	return tuo
}

// SetNillableIsBanned sets the "is_banned" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableIsBanned(b *bool) *TeamUpdateOne {
	if b != nil {
		tuo.SetIsBanned(*b)
	}
	return tuo
}

// SetIsBlocked sets the "is_blocked" field.
func (tuo *TeamUpdateOne) SetIsBlocked(b bool) *TeamUpdateOne {
	tuo.mutation.SetIsBlocked(b)
	return tuo
}

// SetNillableIsBlocked sets the "is_blocked" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableIsBlocked(b *bool) *TeamUpdateOne {
	if b != nil {
		tuo.SetIsBlocked(*b)
	}
	return tuo
}

// SetBlockedReason sets the "blocked_reason" field.
func (tuo *TeamUpdateOne) SetBlockedReason(s string) *TeamUpdateOne {
	tuo.mutation.SetBlockedReason(s)
	return tuo
}

// SetNillableBlockedReason sets the "blocked_reason" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableBlockedReason(s *string) *TeamUpdateOne {
	if s != nil {
		tuo.SetBlockedReason(*s)
	}
	return tuo
}

// ClearBlockedReason clears the value of the "blocked_reason" field.
func (tuo *TeamUpdateOne) ClearBlockedReason() *TeamUpdateOne {
	tuo.mutation.ClearBlockedReason()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TeamUpdateOne) SetName(s string) *TeamUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableName(s *string) *TeamUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// SetTier sets the "tier" field.
func (tuo *TeamUpdateOne) SetTier(s string) *TeamUpdateOne {
	tuo.mutation.SetTier(s)
	return tuo
}

// SetNillableTier sets the "tier" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableTier(s *string) *TeamUpdateOne {
	if s != nil {
		tuo.SetTier(*s)
	}
	return tuo
}

// SetEmail sets the "email" field.
func (tuo *TeamUpdateOne) SetEmail(s string) *TeamUpdateOne {
	tuo.mutation.SetEmail(s)
	return tuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (tuo *TeamUpdateOne) SetNillableEmail(s *string) *TeamUpdateOne {
	if s != nil {
		tuo.SetEmail(*s)
	}
	return tuo
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (tuo *TeamUpdateOne) AddUserIDs(ids ...uuid.UUID) *TeamUpdateOne {
	tuo.mutation.AddUserIDs(ids...)
	return tuo
}

// AddUsers adds the "users" edges to the User entity.
func (tuo *TeamUpdateOne) AddUsers(u ...*User) *TeamUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUserIDs(ids...)
}

// AddTeamAPIKeyIDs adds the "team_api_keys" edge to the TeamAPIKey entity by IDs.
func (tuo *TeamUpdateOne) AddTeamAPIKeyIDs(ids ...string) *TeamUpdateOne {
	tuo.mutation.AddTeamAPIKeyIDs(ids...)
	return tuo
}

// AddTeamAPIKeys adds the "team_api_keys" edges to the TeamAPIKey entity.
func (tuo *TeamUpdateOne) AddTeamAPIKeys(t ...*TeamAPIKey) *TeamUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTeamAPIKeyIDs(ids...)
}

// SetTeamTierID sets the "team_tier" edge to the Tier entity by ID.
func (tuo *TeamUpdateOne) SetTeamTierID(id string) *TeamUpdateOne {
	tuo.mutation.SetTeamTierID(id)
	return tuo
}

// SetTeamTier sets the "team_tier" edge to the Tier entity.
func (tuo *TeamUpdateOne) SetTeamTier(t *Tier) *TeamUpdateOne {
	return tuo.SetTeamTierID(t.ID)
}

// AddEnvIDs adds the "envs" edge to the Env entity by IDs.
func (tuo *TeamUpdateOne) AddEnvIDs(ids ...string) *TeamUpdateOne {
	tuo.mutation.AddEnvIDs(ids...)
	return tuo
}

// AddEnvs adds the "envs" edges to the Env entity.
func (tuo *TeamUpdateOne) AddEnvs(e ...*Env) *TeamUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tuo.AddEnvIDs(ids...)
}

// AddUsersTeamIDs adds the "users_teams" edge to the UsersTeams entity by IDs.
func (tuo *TeamUpdateOne) AddUsersTeamIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.AddUsersTeamIDs(ids...)
	return tuo
}

// AddUsersTeams adds the "users_teams" edges to the UsersTeams entity.
func (tuo *TeamUpdateOne) AddUsersTeams(u ...*UsersTeams) *TeamUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.AddUsersTeamIDs(ids...)
}

// Mutation returns the TeamMutation object of the builder.
func (tuo *TeamUpdateOne) Mutation() *TeamMutation {
	return tuo.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (tuo *TeamUpdateOne) ClearUsers() *TeamUpdateOne {
	tuo.mutation.ClearUsers()
	return tuo
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (tuo *TeamUpdateOne) RemoveUserIDs(ids ...uuid.UUID) *TeamUpdateOne {
	tuo.mutation.RemoveUserIDs(ids...)
	return tuo
}

// RemoveUsers removes "users" edges to User entities.
func (tuo *TeamUpdateOne) RemoveUsers(u ...*User) *TeamUpdateOne {
	ids := make([]uuid.UUID, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUserIDs(ids...)
}

// ClearTeamAPIKeys clears all "team_api_keys" edges to the TeamAPIKey entity.
func (tuo *TeamUpdateOne) ClearTeamAPIKeys() *TeamUpdateOne {
	tuo.mutation.ClearTeamAPIKeys()
	return tuo
}

// RemoveTeamAPIKeyIDs removes the "team_api_keys" edge to TeamAPIKey entities by IDs.
func (tuo *TeamUpdateOne) RemoveTeamAPIKeyIDs(ids ...string) *TeamUpdateOne {
	tuo.mutation.RemoveTeamAPIKeyIDs(ids...)
	return tuo
}

// RemoveTeamAPIKeys removes "team_api_keys" edges to TeamAPIKey entities.
func (tuo *TeamUpdateOne) RemoveTeamAPIKeys(t ...*TeamAPIKey) *TeamUpdateOne {
	ids := make([]string, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTeamAPIKeyIDs(ids...)
}

// ClearTeamTier clears the "team_tier" edge to the Tier entity.
func (tuo *TeamUpdateOne) ClearTeamTier() *TeamUpdateOne {
	tuo.mutation.ClearTeamTier()
	return tuo
}

// ClearEnvs clears all "envs" edges to the Env entity.
func (tuo *TeamUpdateOne) ClearEnvs() *TeamUpdateOne {
	tuo.mutation.ClearEnvs()
	return tuo
}

// RemoveEnvIDs removes the "envs" edge to Env entities by IDs.
func (tuo *TeamUpdateOne) RemoveEnvIDs(ids ...string) *TeamUpdateOne {
	tuo.mutation.RemoveEnvIDs(ids...)
	return tuo
}

// RemoveEnvs removes "envs" edges to Env entities.
func (tuo *TeamUpdateOne) RemoveEnvs(e ...*Env) *TeamUpdateOne {
	ids := make([]string, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return tuo.RemoveEnvIDs(ids...)
}

// ClearUsersTeams clears all "users_teams" edges to the UsersTeams entity.
func (tuo *TeamUpdateOne) ClearUsersTeams() *TeamUpdateOne {
	tuo.mutation.ClearUsersTeams()
	return tuo
}

// RemoveUsersTeamIDs removes the "users_teams" edge to UsersTeams entities by IDs.
func (tuo *TeamUpdateOne) RemoveUsersTeamIDs(ids ...int) *TeamUpdateOne {
	tuo.mutation.RemoveUsersTeamIDs(ids...)
	return tuo
}

// RemoveUsersTeams removes "users_teams" edges to UsersTeams entities.
func (tuo *TeamUpdateOne) RemoveUsersTeams(u ...*UsersTeams) *TeamUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return tuo.RemoveUsersTeamIDs(ids...)
}

// Where appends a list predicates to the TeamUpdate builder.
func (tuo *TeamUpdateOne) Where(ps ...predicate.Team) *TeamUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TeamUpdateOne) Select(field string, fields ...string) *TeamUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Team entity.
func (tuo *TeamUpdateOne) Save(ctx context.Context) (*Team, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TeamUpdateOne) SaveX(ctx context.Context) *Team {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TeamUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TeamUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TeamUpdateOne) check() error {
	if v, ok := tuo.mutation.Email(); ok {
		if err := team.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`models: validator failed for field "Team.email": %w`, err)}
		}
	}
	if _, ok := tuo.mutation.TeamTierID(); tuo.mutation.TeamTierCleared() && !ok {
		return errors.New(`models: clearing a required unique edge "Team.team_tier"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TeamUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TeamUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TeamUpdateOne) sqlSave(ctx context.Context) (_node *Team, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(team.Table, team.Columns, sqlgraph.NewFieldSpec(team.FieldID, field.TypeUUID))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`models: missing "Team.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, team.FieldID)
		for _, f := range fields {
			if !team.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("models: invalid field %q for query", f)}
			}
			if f != team.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.IsBanned(); ok {
		_spec.SetField(team.FieldIsBanned, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.IsBlocked(); ok {
		_spec.SetField(team.FieldIsBlocked, field.TypeBool, value)
	}
	if value, ok := tuo.mutation.BlockedReason(); ok {
		_spec.SetField(team.FieldBlockedReason, field.TypeString, value)
	}
	if tuo.mutation.BlockedReasonCleared() {
		_spec.ClearField(team.FieldBlockedReason, field.TypeString)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(team.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Email(); ok {
		_spec.SetField(team.FieldEmail, field.TypeString, value)
	}
	if tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		createE := &UsersTeamsCreate{config: tuo.config, mutation: newUsersTeamsMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUsersIDs(); len(nodes) > 0 && !tuo.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UsersTeamsCreate{config: tuo.config, mutation: newUsersTeamsMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   team.UsersTable,
			Columns: team.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUUID),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UsersTeamsCreate{config: tuo.config, mutation: newUsersTeamsMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TeamAPIKeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamAPIKeysTable,
			Columns: []string{team.TeamAPIKeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.TeamAPIKey
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTeamAPIKeysIDs(); len(nodes) > 0 && !tuo.mutation.TeamAPIKeysCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamAPIKeysTable,
			Columns: []string{team.TeamAPIKeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.TeamAPIKey
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TeamAPIKeysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.TeamAPIKeysTable,
			Columns: []string{team.TeamAPIKeysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(teamapikey.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.TeamAPIKey
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TeamTierCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   team.TeamTierTable,
			Columns: []string{team.TeamTierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.Team
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TeamTierIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   team.TeamTierTable,
			Columns: []string{team.TeamTierColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tier.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.Team
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.EnvsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.EnvsTable,
			Columns: []string{team.EnvsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.Env
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedEnvsIDs(); len(nodes) > 0 && !tuo.mutation.EnvsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.EnvsTable,
			Columns: []string{team.EnvsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.Env
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.EnvsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   team.EnvsTable,
			Columns: []string{team.EnvsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(env.FieldID, field.TypeString),
			},
		}
		edge.Schema = tuo.schemaConfig.Env
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.UsersTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedUsersTeamsIDs(); len(nodes) > 0 && !tuo.mutation.UsersTeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.UsersTeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   team.UsersTeamsTable,
			Columns: []string{team.UsersTeamsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usersteams.FieldID, field.TypeInt),
			},
		}
		edge.Schema = tuo.schemaConfig.UsersTeams
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = tuo.schemaConfig.Team
	ctx = internal.NewSchemaConfigContext(ctx, tuo.schemaConfig)
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Team{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{team.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
