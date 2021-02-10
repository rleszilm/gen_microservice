// Package postgres_dal_users is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package postgres_dal_users

import (
	context "context"
	sql1 "database/sql"
	fmt "fmt"
	types "github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/annotations/types"
	users "github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/example/users"
	dal "github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/example/users/dal"
	sql2 "github.com/rleszilm/gen_microservice/cmd/protoc-gen-go-genms-dal/generator/sql"
	sql "github.com/rleszilm/gen_microservice/sql"
	stats "go.opencensus.io/stats"
	view "go.opencensus.io/stats/view"
	tag "go.opencensus.io/tag"
	latlng "google.golang.org/genproto/googleapis/type/latlng"
	log "log"
	strings "strings"
	sync "sync"
	time "time"
)

// UserCollection is an autogenerated implementation of dal.UserCollection.
type UserCollection struct {
	dal.UnimplementedUserCollection

	db     sql.DB
	config *UserConfig

	execInsert string
	execUpsert string
	queryAll   string

	queryById              string
	queryByNameAndDivision string
	queryByKind            string
	queryByPhone           string
	queryProviderStubOnly  string
}

// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *UserCollection) Initialize(_ context.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *UserCollection) Shutdown(_ context.Context) error {
	return nil
}

// NameOf returns the name of a service. This must be unique if there are multiple instances of the same
// service.
func (x *UserCollection) NameOf() string {
	return "postgres_dal_users_" + x.config.TableName
}

// String returns a string identifier for the service.
func (x *UserCollection) String() string {
	return x.NameOf()
}

// DoInsert provides the base logic for dal.UserCollection.Insert.
// The user should use this as a base for dal.UserCollection.Insert, only having to add
// code that interprets the returned values.
func (x *UserCollection) DoInsert(ctx context.Context, arg interface{}) (sql1.Result, error) {
	var err error
	start := time.Now()
	stats.Record(ctx, userInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

		if err != nil {
			ctx, err = tag.New(ctx,
				tag.Insert(userQueryError, "user_insert"),
			)
		}

		ctx, err = tag.New(ctx,
			tag.Insert(userQueryName, "user_insert"),
		)

		stats.Record(ctx, userLatency.M(dur), userInflight.M(-1))
	}()

	return x.db.ExecWithReplacements(ctx, x.execInsert, arg)
}

// DoUpsert provides the base logic for dal.UserCollection.Upsert.
// The user should use this as a base for dal.UserCollection.Upsert, only having to add
// code that interprets the returned values.
func (x *UserCollection) DoUpsert(ctx context.Context, arg interface{}) (sql1.Result, error) {
	var err error
	start := time.Now()
	stats.Record(ctx, userInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

		if err != nil {
			ctx, err = tag.New(ctx,
				tag.Upsert(userQueryError, "user_upsert"),
			)
		}

		ctx, err = tag.New(ctx,
			tag.Upsert(userQueryName, "user_upsert"),
		)

		stats.Record(ctx, userLatency.M(dur), userInflight.M(-1))
	}()

	return x.db.ExecWithReplacements(ctx, x.execUpsert, arg)
}

// All implements dal.UserCollection.All
func (x *UserCollection) All(ctx context.Context) ([]*users.User, error) {
	filter := &dal.UserFilter{}
	return x.find(ctx, "all", x.queryAll, filter)
}

// Filter implements dal.UserCollection.Filter
func (x *UserCollection) Filter(ctx context.Context, arg *dal.UserFilter) ([]*users.User, error) {
	query := "SELECT id, name, division, lifetime_score, last_score, payout, point, phone, geo, type, by_backend_postgres FROM " + x.config.TableName
	fields := []string{}

	if arg.Id != nil {
		fields = append(fields, "id = :id")
	}
	if arg.Name != nil {
		fields = append(fields, "name = :name")
	}
	if arg.Division != nil {
		fields = append(fields, "division = :division")
	}
	if arg.LifetimeScore != nil {
		fields = append(fields, "lifetime_score = :lifetime_score")
	}
	if arg.LastScore != nil {
		fields = append(fields, "last_score = :last_score")
	}

	if arg.LastWinnings != nil {
		fields = append(fields, "payout = :payout")
	}
	if arg.Point != nil {
		fields = append(fields, "point = :point")
	}
	if arg.Phone != nil {
		fields = append(fields, "phone = :phone")
	}
	if arg.Geo != nil {
		fields = append(fields, "geo = :geo")
	}
	if arg.Kind != nil {
		fields = append(fields, "type = :type")
	}
	if arg.ByBackend != nil {
		fields = append(fields, "by_backend_postgres = :by_backend_postgres")
	}
	if len(fields) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(fields, " AND "))
	}

	return x.find(ctx, "filter", query, arg)
}

func (x *UserCollection) find(ctx context.Context, label string, query string, arg *dal.UserFilter) ([]*users.User, error) {
	var err error
	start := time.Now()
	stats.Record(ctx, userInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

		if err != nil {
			ctx, err = tag.New(ctx,
				tag.Upsert(userQueryError, label),
			)
		}

		ctx, err = tag.New(ctx,
			tag.Upsert(userQueryName, label),
		)

		stats.Record(ctx, userLatency.M(dur), userInflight.M(-1))
	}()

	filter := &UserFilter{}
	filter.fromGeneric(arg)

	rows, err := x.db.QueryWithReplacements(ctx, query, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	Users := []*users.User{}
	for rows.Next() {
		obj := &UserScanner{}
		if err = rows.StructScan(obj); err != nil {
			return nil, err
		}
		Users = append(Users, obj.User())
	}
	return Users, nil
}

// ById implements dal.UserCollection.ById
func (x *UserCollection) ById(ctx context.Context, id int64) ([]*users.User, error) {
	filter := &dal.UserFilter{
		Id: &id,
	}
	return x.find(ctx, "by_id", x.queryById, filter)
}

// ByNameAndDivision implements dal.UserCollection.ByNameAndDivision
func (x *UserCollection) ByNameAndDivision(ctx context.Context, name string, division string) ([]*users.User, error) {
	filter := &dal.UserFilter{
		Name:     &name,
		Division: &division,
	}
	return x.find(ctx, "by_name_and_division", x.queryByNameAndDivision, filter)
}

// ByKind implements dal.UserCollection.ByKind
func (x *UserCollection) ByKind(ctx context.Context, kind users.User_Kind) ([]*users.User, error) {
	filter := &dal.UserFilter{
		Kind: &kind,
	}
	return x.find(ctx, "by_kind", x.queryByKind, filter)
}

// ByPhone implements dal.UserCollection.ByPhone
func (x *UserCollection) ByPhone(ctx context.Context, phone *types.Phone) ([]*users.User, error) {
	filter := &dal.UserFilter{
		Phone: phone,
	}
	return x.find(ctx, "by_phone", x.queryByPhone, filter)
}

// ProviderStubOnly implements dal.UserCollection.ProviderStubOnly
func (x *UserCollection) ProviderStubOnly(ctx context.Context) ([]*users.User, error) {
	filter := &dal.UserFilter{}
	return x.find(ctx, "provider_stub_only", x.queryProviderStubOnly, filter)
}

// NewUserCollection returns a new UserCollection.
func NewUserCollection(db sql.DB, queries UserQueryTemplateProvider, config *UserConfig) (*UserCollection, error) {
	registerUserMetricsOnce.Do(registerUserMetrics)

	coll := &UserCollection{
		UnimplementedUserCollection: dal.UnimplementedUserCollection{},
		db:                          db,
		config:                      config,
	}

	queryReplacements := map[string]string{
		"table":       config.TableName,
		"fields":      "id, name, division, lifetime_score, last_score, payout, point, phone, geo, type, by_backend_postgres",
		"writeFields": ":id, :name, :division, :lifetime_score, :last_score, :payout, :point, :phone, :geo, :type, :by_backend_postgres",
	}

	// generate Upsert exec
	coll.execInsert = sql2.MustGenerateQuery("dal.User-Exec-Insert", queries.Insert(), queryReplacements)

	// generate Upsert exec
	coll.execUpsert = sql2.MustGenerateQuery("dal.User-Exec-Upsert", queries.Upsert(), queryReplacements)

	// generate All query
	coll.queryAll = sql2.MustGenerateQuery("dal.User-Query-All", queries.All(), queryReplacements)

	// generate ById query
	coll.queryById = sql2.MustGenerateQuery("dal.User-Query-ById", queries.ById(), queryReplacements)

	// generate ByNameAndDivision query
	coll.queryByNameAndDivision = sql2.MustGenerateQuery("dal.User-Query-ByNameAndDivision", queries.ByNameAndDivision(), queryReplacements)

	// generate ByKind query
	coll.queryByKind = sql2.MustGenerateQuery("dal.User-Query-ByKind", queries.ByKind(), queryReplacements)

	// generate ByPhone query
	coll.queryByPhone = sql2.MustGenerateQuery("dal.User-Query-ByPhone", queries.ByPhone(), queryReplacements)

	// generate ProviderStubOnly query
	coll.queryProviderStubOnly = sql2.MustGenerateQuery("dal.User-Query-ProviderStubOnly", queries.ProviderStubOnly(), queryReplacements)

	return coll, nil
}

// UserFilter is an autogenerated struct that
// is used in generic User queries.
type UserFilter struct {
	Id            *int64   `db:"id"`
	Name          *string  `db:"name"`
	Division      *string  `db:"division"`
	LifetimeScore *float64 `db:"lifetime_score"`
	LastScore     *float32 `db:"last_score"`

	LastWinnings *int32           `db:"payout"`
	Point        *types.Point     `db:"point"`
	Phone        *types.Phone     `db:"phone"`
	Geo          *latlng.LatLng   `db:"geo"`
	Kind         *users.User_Kind `db:"type"`
	ByBackend    *string          `db:"by_backend_postgres"`
}

func (x *UserFilter) fromGeneric(y *dal.UserFilter) {
	if y.Id != nil {
		x.Id = y.Id
	}
	if y.Name != nil {
		x.Name = y.Name
	}
	if y.Division != nil {
		x.Division = y.Division
	}
	if y.LifetimeScore != nil {
		x.LifetimeScore = y.LifetimeScore
	}
	if y.LastScore != nil {
		x.LastScore = y.LastScore
	}

	if y.LastWinnings != nil {
		x.LastWinnings = y.LastWinnings
	}
	if y.Point != nil {
		x.Point = y.Point
	}
	if y.Phone != nil {
		x.Phone = y.Phone
	}
	if y.Geo != nil {
		x.Geo = y.Geo
	}
	if y.Kind != nil {
		x.Kind = y.Kind
	}
	if y.ByBackend != nil {
		x.ByBackend = y.ByBackend
	}
}

// UserScanner is an autogenerated struct that
// is used to parse query results.
type UserScanner struct {
	Id            sql1.NullInt64   `db:"id"`
	Name          sql1.NullString  `db:"name"`
	Division      sql1.NullString  `db:"division"`
	LifetimeScore sql1.NullFloat64 `db:"lifetime_score"`
	LastScore     sql1.NullFloat64 `db:"last_score"`
	LastWinnings  sql1.NullInt32   `db:"payout"`
	Point         *types.Point     `db:"point"`
	Phone         *types.Phone     `db:"phone"`
	Geo           *latlng.LatLng   `db:"geo"`
	Kind          sql1.NullInt32   `db:"type"`
	ByBackend     sql1.NullString  `db:"by_backend_postgres"`
}

// User returns a new users.User populated with scanned values.
func (x *UserScanner) User() *users.User {
	return &users.User{
		Id:            x.Id.Int64,
		Name:          x.Name.String,
		Division:      x.Division.String,
		LifetimeScore: x.LifetimeScore.Float64,
		LastScore:     float32(x.LastScore.Float64),
		LastWinnings:  x.LastWinnings.Int32,
		Point:         x.Point,
		Phone:         x.Phone,
		Geo:           x.Geo,
		Kind:          users.User_Kind(x.Kind.Int32),
		ByBackend:     x.ByBackend.String,
	}
}

// UserWriter is an autogenerated struct that
// is used to supply values to write queries.
type UserWriter struct {
	Id            int64           `db:"id"`
	Name          string          `db:"name"`
	Division      string          `db:"division"`
	LifetimeScore float64         `db:"lifetime_score"`
	LastScore     float32         `db:"last_score"`
	LastWinnings  int32           `db:"payout"`
	Point         *types.Point    `db:"point"`
	Phone         *types.Phone    `db:"phone"`
	Geo           *latlng.LatLng  `db:"geo"`
	Kind          users.User_Kind `db:"type"`
	ByBackend     string          `db:"by_backend_postgres"`
}

// FromUser populates the struct with values from the base type.
func (x *UserWriter) FromUser(y *users.User) {
	x.Id = y.Id
	x.Name = y.Name
	x.Division = y.Division
	x.LifetimeScore = y.LifetimeScore
	x.LastScore = y.LastScore
	x.LastWinnings = y.LastWinnings
	x.Point = y.Point
	x.Phone = y.Phone
	x.Geo = y.Geo
	x.Kind = y.Kind
	x.ByBackend = y.ByBackend

}

// UserConfig is a struct that can be used to configure a UserCollection
type UserConfig struct {
	TableName string `envconfig:"table"`
}

// UserQueryTemplateProvider is an interface that returns the query templated that should be executed
// to generate the queries that the collection will use.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . UserQueryTemplateProvider
type UserQueryTemplateProvider interface {
	Insert() string
	Upsert() string
	All() string
	ById() string
	ByNameAndDivision() string
	ByKind() string
	ByPhone() string
	ProviderStubOnly() string
}

// UserQueries provides auto-generated queries when possible. This is not gauranteed to be a complete
// implementation of the interface. This should be used as a base for the actual query provider used.
type UserQueries struct {
}

// Insert implements UserQueryTemplateProvider.Insert.
func (x *UserQueries) Insert() string {
	return `INSERT INTO {{ .table }}({{ .fields }}) VALUES({{ .writeFields }});`
}

// Upsert implements UserQueryTemplateProvider.Upsert.
func (x *UserQueries) Upsert() string {
	return `INSERT INTO {{ .table }}({{ .fields }}) VALUES({{ .writeFields }});`
}

// All implements UserQueryTemplateProvider.All.
func (x *UserQueries) All() string {
	return `SELECT {{ .fields }} FROM {{ .table }};`
}

// ById implements UserQueryTemplateProvider.ById.
func (x *UserQueries) ById() string {
	return `SELECT {{ .fields }} FROM {{ .table }}
		WHERE 
			id = :id;`
}

// ByNameAndDivision implements UserQueryTemplateProvider.ByNameAndDivision.
func (x *UserQueries) ByNameAndDivision() string {
	return `SELECT {{ .fields }} FROM {{ .table }}
		WHERE 
			name = :name AND
			division = :division;`
}

// ByKind implements UserQueryTemplateProvider.ByKind.
func (x *UserQueries) ByKind() string {
	return `SELECT {{ .fields }} FROM {{ .table }}
		WHERE 
			type = :type;`
}

// ByPhone implements UserQueryTemplateProvider.ByPhone.
func (x *UserQueries) ByPhone() string {
	return `SELECT {{ .fields }} FROM {{ .table }}
		WHERE 
			phone = :phone;`
}

// define metrics
var (
	userQueryName  = tag.MustNewKey("dal_postgres_user")
	userQueryError = tag.MustNewKey("dal_postgres_user_error")

	userLatency  = stats.Float64("user_latency", "Latency of User queries", stats.UnitMilliseconds)
	userInflight = stats.Int64("user_inflight", "Count of User queries in flight", stats.UnitDimensionless)

	registerUserMetricsOnce sync.Once
)

func registerUserMetrics() {
	views := []*view.View{
		{
			Name:        "dal_postgres_user_latency",
			Measure:     userLatency,
			Description: "The distribution of the query latencies",
			TagKeys:     []tag.Key{userQueryName, userQueryError},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "dal_postgres_user_inflight",
			Measure:     userInflight,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{userQueryName},
			Aggregation: view.Sum(),
		},
	}

	if err := view.Register(views...); err != nil {
		log.Fatal("Cannot register metrics:", err)
	}
}
