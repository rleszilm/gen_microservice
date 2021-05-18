// Package postgres_dal_users is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package postgres_dal_users

import (
	bytes "bytes"
	context "context"
	sql1 "database/sql"
	fmt "fmt"
	log "log"
	strings "strings"
	sync "sync"
	template "text/template"
	time "time"

	users "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/users"
	dal "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/users/dal"
	dal1 "github.com/rleszilm/genms/dal"
	sql "github.com/rleszilm/genms/sql"
	stats "go.opencensus.io/stats"
	view "go.opencensus.io/stats/view"
	tag "go.opencensus.io/tag"
)

// UserCollection is an autogenerated implementation of dal.UserCollection.
type UserCollection struct {
	dal.UnimplementedUserCollection

	db     sql.DB
	config *UserConfig

	execInsert string
	execUpsert string
	queryAll   string

	execUpdateTmpl *template.Template

	queryOneParam         string
	queryMultipleParam    string
	queryMessageParam     string
	queryWithComparator   string
	queryWithRest         string
	queryProviderStubOnly string
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
func (x *UserCollection) DoInsert(ctx context.Context, arg *users.User) (sql1.Result, error) {
	var err error
	start := time.Now()

	ctx, _ = tag.New(ctx,
		tag.Upsert(tagQueryCollection, "user"),
	)

	ctx, _ = tag.New(ctx,
		tag.Upsert(tagQueryName, "insert"),
	)

	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

		if err != nil {
			ctx, _ = tag.New(ctx,
				tag.Insert(tagQueryError, "true"),
			)
		}

		stats.Record(ctx, measureLatency.M(dur), measureInflight.M(-1))
	}()

	return x.db.ExecWithReplacements(ctx, x.execInsert, userWriterFromGeneric(arg))
}

// DoUpsert provides the base logic for dal.UserCollection.Upsert.
// The user should use this as a base for dal.UserCollection.Upsert, only having to add
// code that interprets the returned values.
func (x *UserCollection) DoUpsert(ctx context.Context, arg *users.User) (sql1.Result, error) {
	var err error
	start := time.Now()

	ctx, _ = tag.New(ctx,
		tag.Upsert(tagQueryCollection, "user"),
	)

	ctx, _ = tag.New(ctx,
		tag.Upsert(tagQueryName, "upsert"),
	)

	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

		if err != nil {
			ctx, _ = tag.New(ctx,
				tag.Insert(tagQueryError, "true"),
			)
		}

		stats.Record(ctx, measureLatency.M(dur), measureInflight.M(-1))
	}()

	return x.db.ExecWithReplacements(ctx, x.execUpsert, userWriterFromGeneric(arg))
}

// DoUpdate provides the base logic for dal.UserCollection.Upsert.
// The user should use this as a base for dal.UserCollection.Upsert, only having to add
// code that interprets the returned values.
func (x *UserCollection) DoUpdate(ctx context.Context, fvs *dal.UserFieldValues, clause string) (sql1.Result, error) {
	var err error
	start := time.Now()

	ctx, _ = tag.New(ctx,
		tag.Upsert(tagQueryCollection, "user"),
	)

	ctx, _ = tag.New(ctx,
		tag.Upsert(tagQueryName, "update"),
	)

	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

		if err != nil {
			ctx, _ = tag.New(ctx,
				tag.Insert(tagQueryError, "true"),
			)
		}

		stats.Record(ctx, measureLatency.M(dur), measureInflight.M(-1))
	}()

	updates := []string{}

	if fvs.ScalarInt32 != nil {
		updates = append(updates, "scalar_int32 = :scalar_int32")
	}

	if fvs.ScalarInt64 != nil {
		updates = append(updates, "scalar_int64 = :scalar_int64")
	}

	if fvs.ScalarFloat32 != nil {
		updates = append(updates, "scalar_float32 = :scalar_float32")
	}

	if fvs.ScalarFloat64 != nil {
		updates = append(updates, "scalar_float64 = :scalar_float64")
	}

	if fvs.ScalarString != nil {
		updates = append(updates, "scalar_string = :scalar_string")
	}

	if fvs.ScalarBool != nil {
		updates = append(updates, "scalar_bool = :scalar_bool")
	}

	if fvs.ScalarEnum != nil {
		updates = append(updates, "scalar_enum = :scalar_enum")
	}

	if fvs.ObjMessage != nil {
		updates = append(updates, "obj_message = :obj_message")
	}

	if fvs.Renamed != nil {
		updates = append(updates, "aliased = :aliased")
	}

	if fvs.RenamedPostgres != nil {
		updates = append(updates, "aliased_postgres = :aliased_postgres")
	}

	if fvs.IgnoredRest != nil {
		updates = append(updates, "ignored_rest = :ignored_rest")
	}

	if fvs.RenamedRest != nil {
		updates = append(updates, "renamed_rest = :renamed_rest")
	}
	buf := &bytes.Buffer{}
	if err := x.execUpdateTmpl.Execute(buf, map[string]interface{}{
		"clause":  clause,
		"table":   x.config.TableName,
		"updates": strings.Join(updates, ", "),
	}); err != nil {
		return nil, err
	}

	return x.db.ExecWithReplacements(ctx, string(buf.Bytes()), userFieldValuesFromGeneric(fvs))
}

// All implements dal.UserCollection.All
func (x *UserCollection) All(ctx context.Context) ([]*users.User, error) {
	return x.find(ctx, "all", x.queryAll, map[string]interface{}{})
}

// Filter implements dal.UserCollection.Filter
func (x *UserCollection) Filter(ctx context.Context, fvs *dal.UserFieldValues) ([]*users.User, error) {
	query := "SELECT scalar_int32, scalar_int64, scalar_float32, scalar_float64, scalar_string, scalar_bool, scalar_enum, obj_message, ignored, aliased, ignored_postgres, aliased_postgres, ignored_rest, renamed_rest FROM " + x.config.TableName

	fields := []string{}
	if fvs.ScalarInt32 != nil {
		fields = append(fields, "scalar_int32 = :scalar_int32")
	}
	if fvs.ScalarInt64 != nil {
		fields = append(fields, "scalar_int64 = :scalar_int64")
	}
	if fvs.ScalarFloat32 != nil {
		fields = append(fields, "scalar_float32 = :scalar_float32")
	}
	if fvs.ScalarFloat64 != nil {
		fields = append(fields, "scalar_float64 = :scalar_float64")
	}
	if fvs.ScalarString != nil {
		fields = append(fields, "scalar_string = :scalar_string")
	}
	if fvs.ScalarBool != nil {
		fields = append(fields, "scalar_bool = :scalar_bool")
	}
	if fvs.ScalarEnum != nil {
		fields = append(fields, "scalar_enum = :scalar_enum")
	}
	if fvs.ObjMessage != nil {
		fields = append(fields, "obj_message = :obj_message")
	}

	if fvs.Renamed != nil {
		fields = append(fields, "aliased = :aliased")
	}

	if fvs.RenamedPostgres != nil {
		fields = append(fields, "aliased_postgres = :aliased_postgres")
	}
	if fvs.IgnoredRest != nil {
		fields = append(fields, "ignored_rest = :ignored_rest")
	}
	if fvs.RenamedRest != nil {
		fields = append(fields, "renamed_rest = :renamed_rest")
	}
	if len(fields) > 0 {
		query = fmt.Sprintf("%s WHERE %s", query, strings.Join(fields, " AND "))
	}

	return x.find(ctx, "filter", query, userFieldValuesFromGeneric(fvs))
}

func (x *UserCollection) find(ctx context.Context, label string, query string, fvs interface{}) ([]*users.User, error) {
	var err error
	start := time.Now()

	ctx, _ = tag.New(ctx,
		tag.Upsert(tagQueryCollection, "user"),
	)

	ctx, _ = tag.New(ctx,
		tag.Upsert(tagQueryName, label),
	)

	stats.Record(ctx, measureInflight.M(1))
	defer func() {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)

		if err != nil {
			ctx, _ = tag.New(ctx,
				tag.Upsert(tagQueryError, "true"),
			)
		}

		stats.Record(ctx, measureLatency.M(dur), measureInflight.M(-1))
	}()

	rows, err := x.db.QueryWithReplacements(ctx, query, fvs)
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

// OneParam implements dal.UserCollection.OneParam
func (x *UserCollection) OneParam(ctx context.Context, scalar_int32 int32) ([]*users.User, error) {
	fvs := map[string]interface{}{"scalar_int32": &scalar_int32}
	return x.find(ctx, "one_param", x.queryOneParam, fvs)
}

// MultipleParam implements dal.UserCollection.MultipleParam
func (x *UserCollection) MultipleParam(ctx context.Context, scalar_int32 int32, scalar_int64 int64, scalar_float32 float32) ([]*users.User, error) {
	fvs := map[string]interface{}{"scalar_int32": &scalar_int32, "scalar_int64": &scalar_int64, "scalar_float32": &scalar_float32}
	return x.find(ctx, "multiple_param", x.queryMultipleParam, fvs)
}

// MessageParam implements dal.UserCollection.MessageParam
func (x *UserCollection) MessageParam(ctx context.Context, obj_message *users.User_Message) ([]*users.User, error) {
	fvs := map[string]interface{}{"obj_message": obj_message}
	return x.find(ctx, "message_param", x.queryMessageParam, fvs)
}

// WithComparator implements dal.UserCollection.WithComparator
func (x *UserCollection) WithComparator(ctx context.Context, scalar_int32 int32) ([]*users.User, error) {
	fvs := map[string]interface{}{"scalar_int32": &scalar_int32}
	return x.find(ctx, "with_comparator", x.queryWithComparator, fvs)
}

// WithRest implements dal.UserCollection.WithRest
func (x *UserCollection) WithRest(ctx context.Context, scalar_int32 int32, scalar_int64 int64, scalar_float32 float32, scalar_float64 float64) ([]*users.User, error) {
	fvs := map[string]interface{}{"scalar_int32": &scalar_int32, "scalar_int64": &scalar_int64, "scalar_float32": &scalar_float32, "scalar_float64": &scalar_float64}
	return x.find(ctx, "with_rest", x.queryWithRest, fvs)
}

// ProviderStubOnly implements dal.UserCollection.ProviderStubOnly
func (x *UserCollection) ProviderStubOnly(ctx context.Context) ([]*users.User, error) {
	fvs := map[string]interface{}{}
	return x.find(ctx, "provider_stub_only", x.queryProviderStubOnly, fvs)
}

// NewUserCollection returns a new UserCollection.
func NewUserCollection(db sql.DB, queries UserQueryTemplateProvider, config *UserConfig) (*UserCollection, error) {
	registerUserMetricsOnce.Do(registerUserMetrics)

	coll := &UserCollection{
		db:     db,
		config: config,
	}

	queryReplacements := map[string]string{
		"table":       config.TableName,
		"fields":      "scalar_int32, scalar_int64, scalar_float32, scalar_float64, scalar_string, scalar_bool, scalar_enum, obj_message, ignored, aliased, ignored_postgres, aliased_postgres, ignored_rest, renamed_rest",
		"writeFields": ":scalar_int32, :scalar_int64, :scalar_float32, :scalar_float64, :scalar_string, :scalar_bool, :scalar_enum, :obj_message, :ignored, :aliased, :ignored_postgres, :aliased_postgres, :ignored_rest, :renamed_rest",
	}

	// generate Insert exec
	execInsert, err := dal1.RenderQuery("dal.User-exec-insert", queries.Insert(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.execInsert = execInsert

	// generate Upsert exec
	execUpsert, err := dal1.RenderQuery("dal.User-exec-upsert", queries.Upsert(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.execUpsert = execUpsert

	execUpdateTmpl, err := template.New("dal.User-exec-update").
		Funcs(template.FuncMap{}).
		Parse(queries.Update())

	if err != nil {
		return nil, err
	}
	coll.execUpdateTmpl = execUpdateTmpl

	// generate All query
	queryAll, err := dal1.RenderQuery("dal.User-query-all", queries.All(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.queryAll = queryAll

	// generate OneParam query
	queryOneParam, err := dal1.RenderQuery("dal.User-query-one_param", queries.OneParam(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.queryOneParam = queryOneParam

	// generate MultipleParam query
	queryMultipleParam, err := dal1.RenderQuery("dal.User-query-multiple_param", queries.MultipleParam(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.queryMultipleParam = queryMultipleParam

	// generate MessageParam query
	queryMessageParam, err := dal1.RenderQuery("dal.User-query-message_param", queries.MessageParam(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.queryMessageParam = queryMessageParam

	// generate WithComparator query
	queryWithComparator, err := dal1.RenderQuery("dal.User-query-with_comparator", queries.WithComparator(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.queryWithComparator = queryWithComparator

	// generate WithRest query
	queryWithRest, err := dal1.RenderQuery("dal.User-query-with_rest", queries.WithRest(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.queryWithRest = queryWithRest

	// generate ProviderStubOnly query
	queryProviderStubOnly, err := dal1.RenderQuery("dal.User-query-provider_stub_only", queries.ProviderStubOnly(), queryReplacements)
	if err != nil {
		return nil, err
	}
	coll.queryProviderStubOnly = queryProviderStubOnly

	return coll, nil
}

// UserFieldValues is an autogenerated struct that is used in generic User queries.
type UserFieldValues struct {
	ScalarInt32   *int32              `db:"scalar_int32"`
	ScalarInt64   *int64              `db:"scalar_int64"`
	ScalarFloat32 *float32            `db:"scalar_float32"`
	ScalarFloat64 *float64            `db:"scalar_float64"`
	ScalarString  *string             `db:"scalar_string"`
	ScalarBool    *bool               `db:"scalar_bool"`
	ScalarEnum    *users.User_Enum    `db:"scalar_enum"`
	ObjMessage    *users.User_Message `db:"obj_message"`

	Renamed *string `db:"aliased"`

	RenamedPostgres *string `db:"aliased_postgres"`
	IgnoredRest     *string `db:"ignored_rest"`
	RenamedRest     *string `db:"renamed_rest"`
}

func userFieldValuesFromGeneric(y *dal.UserFieldValues) *UserFieldValues {
	f := &UserFieldValues{}
	if y.ScalarInt32 != nil {
		f.ScalarInt32 = y.ScalarInt32
	}
	if y.ScalarInt64 != nil {
		f.ScalarInt64 = y.ScalarInt64
	}
	if y.ScalarFloat32 != nil {
		f.ScalarFloat32 = y.ScalarFloat32
	}
	if y.ScalarFloat64 != nil {
		f.ScalarFloat64 = y.ScalarFloat64
	}
	if y.ScalarString != nil {
		f.ScalarString = y.ScalarString
	}
	if y.ScalarBool != nil {
		f.ScalarBool = y.ScalarBool
	}
	if y.ScalarEnum != nil {
		f.ScalarEnum = y.ScalarEnum
	}
	if y.ObjMessage != nil {
		f.ObjMessage = y.ObjMessage
	}

	if y.Renamed != nil {
		f.Renamed = y.Renamed
	}

	if y.RenamedPostgres != nil {
		f.RenamedPostgres = y.RenamedPostgres
	}
	if y.IgnoredRest != nil {
		f.IgnoredRest = y.IgnoredRest
	}
	if y.RenamedRest != nil {
		f.RenamedRest = y.RenamedRest
	}
	return f
}

// UserScanner is an autogenerated struct that
// is used to parse query results.
type UserScanner struct {
	ScalarInt32   sql1.NullInt32      `db:"scalar_int32"`
	ScalarInt64   sql1.NullInt64      `db:"scalar_int64"`
	ScalarFloat32 sql1.NullFloat64    `db:"scalar_float32"`
	ScalarFloat64 sql1.NullFloat64    `db:"scalar_float64"`
	ScalarString  sql1.NullString     `db:"scalar_string"`
	ScalarBool    sql1.NullBool       `db:"scalar_bool"`
	ScalarEnum    sql1.NullInt32      `db:"scalar_enum"`
	ObjMessage    *users.User_Message `db:"obj_message"`

	Renamed sql1.NullString `db:"aliased"`

	RenamedPostgres sql1.NullString `db:"aliased_postgres"`
	IgnoredRest     sql1.NullString `db:"ignored_rest"`
	RenamedRest     sql1.NullString `db:"renamed_rest"`
}

// User returns a new users.User populated with scanned values.
func (x *UserScanner) User() *users.User {
	y := &users.User{}

	if x.ScalarInt32.Valid {
		y.ScalarInt32 = x.ScalarInt32.Int32
	}
	if x.ScalarInt64.Valid {
		y.ScalarInt64 = x.ScalarInt64.Int64
	}
	if x.ScalarFloat32.Valid {
		y.ScalarFloat32 = float32(x.ScalarFloat32.Float64)
	}
	if x.ScalarFloat64.Valid {
		y.ScalarFloat64 = x.ScalarFloat64.Float64
	}
	if x.ScalarString.Valid {
		y.ScalarString = x.ScalarString.String
	}
	if x.ScalarBool.Valid {
		y.ScalarBool = x.ScalarBool.Bool
	}
	if x.ScalarEnum.Valid {
		y.ScalarEnum = users.User_Enum(x.ScalarEnum.Int32)
	}
	y.ObjMessage = x.ObjMessage

	if x.Renamed.Valid {
		y.Renamed = x.Renamed.String
	}

	if x.RenamedPostgres.Valid {
		y.RenamedPostgres = x.RenamedPostgres.String
	}
	if x.IgnoredRest.Valid {
		y.IgnoredRest = x.IgnoredRest.String
	}
	if x.RenamedRest.Valid {
		y.RenamedRest = x.RenamedRest.String
	}
	return y
}

// UserWriter is an autogenerated struct that is used to supply values to write queries.
type UserWriter struct {
	ScalarInt32   int32               `db:"scalar_int32"`
	ScalarInt64   int64               `db:"scalar_int64"`
	ScalarFloat32 float32             `db:"scalar_float32"`
	ScalarFloat64 float64             `db:"scalar_float64"`
	ScalarString  string              `db:"scalar_string"`
	ScalarBool    bool                `db:"scalar_bool"`
	ScalarEnum    users.User_Enum     `db:"scalar_enum"`
	ObjMessage    *users.User_Message `db:"obj_message"`

	Renamed string `db:"aliased"`

	RenamedPostgres string `db:"aliased_postgres"`
	IgnoredRest     string `db:"ignored_rest"`
	RenamedRest     string `db:"renamed_rest"`
}

func userWriterFromGeneric(y *users.User) *UserWriter {
	x := &UserWriter{}
	x.ScalarInt32 = y.ScalarInt32
	x.ScalarInt64 = y.ScalarInt64
	x.ScalarFloat32 = y.ScalarFloat32
	x.ScalarFloat64 = y.ScalarFloat64
	x.ScalarString = y.ScalarString
	x.ScalarBool = y.ScalarBool
	x.ScalarEnum = y.ScalarEnum
	x.ObjMessage = y.ObjMessage

	x.Renamed = y.Renamed

	x.RenamedPostgres = y.RenamedPostgres
	x.IgnoredRest = y.IgnoredRest
	x.RenamedRest = y.RenamedRest
	return x
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
	Update() string
	All() string
	OneParam() string
	MultipleParam() string
	MessageParam() string
	WithComparator() string
	WithRest() string
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

// Update implements UserQueryTemplateProvider.Update.
func (x *UserQueries) Update() string {
	return `UPDATE {{ .table }} SET {{ .updates }}{{ if .clause }} WHERE {{ .clause }}{{ end }};`
}

// All implements UserQueryTemplateProvider.All.
func (x *UserQueries) All() string {
	return `SELECT {{ .fields }} FROM {{ .table }};`
}

//OneParamimplements UserQueryTemplateProvider.OneParam.
func (x *UserQueries) OneParam() string {
	return `SELECT {{ .fields }} FROM {{ .table }} WHERE
			1 = 1 AND
				scalar_int32 = :scalar_int32;`
}

//MultipleParamimplements UserQueryTemplateProvider.MultipleParam.
func (x *UserQueries) MultipleParam() string {
	return `SELECT {{ .fields }} FROM {{ .table }} WHERE
			1 = 1 AND
				scalar_int32 = :scalar_int32 AND
				scalar_int64 = :scalar_int64 AND
				scalar_float32 = :scalar_float32;`
}

//MessageParamimplements UserQueryTemplateProvider.MessageParam.
func (x *UserQueries) MessageParam() string {
	return `SELECT {{ .fields }} FROM {{ .table }} WHERE
			1 = 1 AND
				obj_message = :obj_message;`
}

//WithComparatorimplements UserQueryTemplateProvider.WithComparator.
func (x *UserQueries) WithComparator() string {
	return `SELECT {{ .fields }} FROM {{ .table }} WHERE
			1 = 1 AND
				scalar_int32 > :scalar_int32;`
}

//WithRestimplements UserQueryTemplateProvider.WithRest.
func (x *UserQueries) WithRest() string {
	return `SELECT {{ .fields }} FROM {{ .table }} WHERE
			1 = 1 AND
				scalar_int32 = :scalar_int32 AND
				scalar_int64 = :scalar_int64 AND
				scalar_float32 = :scalar_float32 AND
				scalar_float64 = :scalar_float64;`
}

// define metrics
var (
	tagQueryName       = tag.MustNewKey("dal_postgres_query")
	tagQueryCollection = tag.MustNewKey("dal_postgres_collection")
	tagQueryError      = tag.MustNewKey("dal_postgres_error")

	measureLatency  = stats.Float64("dal_postgres_latency", "Latency of User queries", stats.UnitMilliseconds)
	measureInflight = stats.Int64("dal_postgres_inflight", "Count of User queries in flight", stats.UnitDimensionless)

	registerUserMetricsOnce sync.Once
)

func registerUserMetrics() {
	views := []*view.View{
		{
			Name:        "dal_postgres_latency",
			Measure:     measureLatency,
			Description: "The distribution of the query latencies",
			TagKeys:     []tag.Key{tagQueryName, tagQueryCollection, tagQueryError},
			Aggregation: view.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "dal_postgres_inflight",
			Measure:     measureInflight,
			Description: "The number of queries being processed",
			TagKeys:     []tag.Key{tagQueryName, tagQueryCollection},
			Aggregation: view.Sum(),
		},
	}

	if err := view.Register(views...); err != nil {
		log.Fatal("Cannot register metrics:", err)
	}
}
