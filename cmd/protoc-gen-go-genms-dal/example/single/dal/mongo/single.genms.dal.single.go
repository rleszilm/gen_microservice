// Package mongo_dal_single is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package mongo_dal_single

import (
	context "context"
	time "time"

	single "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single"
	dal "github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/example/single/dal"
	mongo "github.com/rleszilm/genms/mongo"
	bson "github.com/rleszilm/genms/mongo/bson"
	stats "go.opencensus.io/stats"
	tag "go.opencensus.io/tag"
)

// SingleConfig is a struct that can be used to configure a SingleCollection
type SingleConfig struct {
	Name       string        `envconfig:"name"`
	Database   string        `envconfig:"database"`
	Collection string        `envconfig:"collection"`
	Timeout    time.Duration `envconfig:"timeout" default:"5s"`
}

var (
	singleProjection = bson.M{
		"scalar_int32":   1,
		"scalar_int64":   1,
		"scalar_float32": 1,
		"scalar_float64": 1,
		"scalar_string":  1,
		"scalar_bytes":   1,
		"scalar_bool":    1,
		"scalar_enum":    1,
		"obj_message":    1,

		"aliased":          1,
		"ignored_postgres": 1,
		"renamed_postgres": 1,
		"ignored_rest":     1,
		"renamed_rest":     1,

		"aliased_mongo":  1,
		"_id":            1,
		"bson_bytes_oid": 1,
	}
)

// SingleMongo is an autogenerated struct that
// is used to parse query results.
type SingleMongo struct {
	ScalarInt32   int32                  `bson:"scalar_int32,omitempty"`
	ScalarInt64   int64                  `bson:"scalar_int64,omitempty"`
	ScalarFloat32 float32                `bson:"scalar_float32,omitempty"`
	ScalarFloat64 float64                `bson:"scalar_float64,omitempty"`
	ScalarString  string                 `bson:"scalar_string,omitempty"`
	ScalarBytes   []byte                 `bson:"scalar_bytes,omitempty"`
	ScalarBool    bool                   `bson:"scalar_bool,omitempty"`
	ScalarEnum    single.Single_Enum     `bson:"scalar_enum,omitempty"`
	ObjMessage    *single.Single_Message `bson:"obj_message,omitempty"`

	Renamed         string `bson:"aliased,omitempty"`
	IgnoredPostgres string `bson:"ignored_postgres,omitempty"`
	RenamedPostgres string `bson:"renamed_postgres,omitempty"`
	IgnoredRest     string `bson:"ignored_rest,omitempty"`
	RenamedRest     string `bson:"renamed_rest,omitempty"`

	RenamedMongo  string        `bson:"aliased_mongo,omitempty"`
	BsonStringOid bson.ObjectID `bson:"_id,omitempty"`
	BsonBytesOid  bson.ObjectID `bson:"bson_bytes_oid,omitempty"`
}

// Single returns a new single.Single populated with scanned values.
func (x *SingleMongo) Single() (*single.Single, error) {
	y := &single.Single{}

	y.ScalarInt32 = x.ScalarInt32
	y.ScalarInt64 = x.ScalarInt64
	y.ScalarFloat32 = x.ScalarFloat32
	y.ScalarFloat64 = x.ScalarFloat64
	y.ScalarString = x.ScalarString
	y.ScalarBytes = x.ScalarBytes
	y.ScalarBool = x.ScalarBool
	y.ScalarEnum = x.ScalarEnum
	y.ObjMessage = x.ObjMessage

	y.Renamed = x.Renamed
	y.IgnoredPostgres = x.IgnoredPostgres
	y.RenamedPostgres = x.RenamedPostgres
	y.IgnoredRest = x.IgnoredRest
	y.RenamedRest = x.RenamedRest

	y.RenamedMongo = x.RenamedMongo
	y.BsonStringOid = x.BsonStringOid.Hex()
	y.BsonBytesOid = (x.BsonBytesOid)[:]
	return y, nil
}

// ToSingleMongo converts the given Single into the internal mongo equivalent.
func ToSingleMongo(obj *single.Single) (*SingleMongo, error) {
	mObj := &SingleMongo{}

	mObj.ScalarInt32 = obj.ScalarInt32
	mObj.ScalarInt64 = obj.ScalarInt64
	mObj.ScalarFloat32 = obj.ScalarFloat32
	mObj.ScalarFloat64 = obj.ScalarFloat64
	mObj.ScalarString = obj.ScalarString
	mObj.ScalarBytes = obj.ScalarBytes
	mObj.ScalarBool = obj.ScalarBool
	mObj.ScalarEnum = obj.ScalarEnum
	mObj.ObjMessage = obj.ObjMessage

	mObj.Renamed = obj.Renamed
	mObj.IgnoredPostgres = obj.IgnoredPostgres
	mObj.RenamedPostgres = obj.RenamedPostgres
	mObj.IgnoredRest = obj.IgnoredRest
	mObj.RenamedRest = obj.RenamedRest

	mObj.RenamedMongo = obj.RenamedMongo

	if obj.BsonStringOid != "" {
		convBsonStringOid, err := bson.ToObjectID(obj.BsonStringOid)
		if err != nil {
			return nil, err
		}
		mObj.BsonStringOid = convBsonStringOid
	} else {
		mObj.BsonStringOid = nil
	}

	if len(obj.BsonBytesOid) != 0 {
		convBsonBytesOid, err := bson.ToObjectID(obj.BsonBytesOid)
		if err != nil {
			return nil, err
		}
		mObj.BsonBytesOid = &convBsonBytesOid
	} else {
		mObj.BsonBytesOid = nil
	}

	return mObj, nil
}

// SingleCollection is an autogenerated implementation of dal.SingleCollection.
type SingleCollection struct {
	dal.UnimplementedSingleCollection

	name          string
	dialer        mongo.Dialer
	config        *SingleConfig
	mutators      []dal.SingleMutator
	defaultFilter bson.M
}

// WithMutators adds dal.SingleMutators to the collection. These will be applied to all values after they are read from mongo.
func (x *SingleCollection) WithMutators(muts ...dal.SingleMutator) {
	x.mutators = append(x.mutators, muts...)
}

// WithDefaultFilter sets the default the default mongo filter to apply to all find queries. This is applied after any query args.
func (x *SingleCollection) WithDefaultFilter(f bson.M) {
	x.defaultFilter = f
}

// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *SingleCollection) Initialize(_ context.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *SingleCollection) Shutdown(_ context.Context) error {
	return nil
}

// NameOf returns the name of a service. This must be unique if there are multiple instances of the same
// service.
func (x *SingleCollection) NameOf() string {
	return "mongo_dal_single_" + x.config.Name
}

// String returns a string identifier for the service.
func (x *SingleCollection) String() string {
	return x.NameOf()
}

// Find scans the collection for records matching the filter.
func (x *SingleCollection) Find(ctx context.Context, label string, filter bson.M, opts ...*mongo.FindOptions) ([]*single.Single, error) {
	ctx, cancel := context.WithTimeout(ctx, x.config.Timeout)
	defer cancel()

	client, err := x.dialer.Dial(ctx)
	if err != nil {
		mongo.Logs().Error("could not dial:", err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}
	defer client.Close(ctx)

	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, label),
	)
	stats.Record(ctx, mongo.MeasureInflight.M(1))
	start := time.Now()
	defer func(ctx context.Context) {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, mongo.MeasureLatency.M(dur), mongo.MeasureInflight.M(-1))
	}(ctx)

	for k, v := range x.defaultFilter {
		filter[k] = v
	}

	cur, err := client.
		Database(x.config.Database).
		Collection(x.config.Collection).
		Find(ctx, filter, singleProjection, opts...)
	if err != nil {
		mongo.Logs().Error("could not execute rest request:", err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}

	vals := []*single.Single{}
	for cur.Next(ctx) {
		obj := &SingleMongo{}
		if err = cur.Decode(obj); err != nil {
			mongo.Logs().Errorf("could not parse %s - %v", label, err)
			stats.Record(ctx, mongo.MeasureError.M(1))
			return nil, err
		}

		val, err := obj.Single()
		if err != nil {
			mongo.Logs().Error("could not convert from mongo to internal:", err)
			stats.Record(ctx, mongo.MeasureError.M(1))
			return nil, err
		}

		for _, m := range x.mutators {
			val, err = m(val)
			if err != nil {
				mongo.Logs().Error("could not mutate value:", val, err)
				stats.Record(ctx, mongo.MeasureError.M(1))
				return nil, err
			}
		}

		vals = append(vals, val)
	}

	return vals, nil
}

// All implements dal.SingleCollection.All
func (x *SingleCollection) All(ctx context.Context) ([]*single.Single, error) {
	return x.Find(ctx, "all", bson.M{})
}

// Filter implements SingleCollectionReader
func (x *SingleCollection) Filter(ctx context.Context, fvs *dal.SingleFieldValues) ([]*single.Single, error) {
	filter := bson.M{}

	if fvs.ScalarInt32 != nil {
		filter["scalar_int32"] = *fvs.ScalarInt32
	}
	if fvs.ScalarInt64 != nil {
		filter["scalar_int64"] = *fvs.ScalarInt64
	}
	if fvs.ScalarFloat32 != nil {
		filter["scalar_float32"] = *fvs.ScalarFloat32
	}
	if fvs.ScalarFloat64 != nil {
		filter["scalar_float64"] = *fvs.ScalarFloat64
	}
	if fvs.ScalarString != nil {
		filter["scalar_string"] = *fvs.ScalarString
	}
	if fvs.ScalarBytes != nil {
		filter["scalar_bytes"] = fvs.ScalarBytes
	}
	if fvs.ScalarBool != nil {
		filter["scalar_bool"] = *fvs.ScalarBool
	}
	if fvs.ScalarEnum != nil {
		filter["scalar_enum"] = *fvs.ScalarEnum
	}
	if fvs.ObjMessage != nil {
		filter["obj_message"] = fvs.ObjMessage
	}

	if fvs.Renamed != nil {
		filter["aliased"] = *fvs.Renamed
	}
	if fvs.IgnoredPostgres != nil {
		filter["ignored_postgres"] = *fvs.IgnoredPostgres
	}
	if fvs.RenamedPostgres != nil {
		filter["renamed_postgres"] = *fvs.RenamedPostgres
	}
	if fvs.IgnoredRest != nil {
		filter["ignored_rest"] = *fvs.IgnoredRest
	}
	if fvs.RenamedRest != nil {
		filter["renamed_rest"] = *fvs.RenamedRest
	}

	if fvs.RenamedMongo != nil {
		filter["aliased_mongo"] = *fvs.RenamedMongo
	}
	if fvs.BsonStringOid != nil {
		convBsonStringOid, err := bson.ToObjectID(*fvs.BsonStringOid)
		if err != nil {
			mongo.Logs().Error("could not convert value to ObjectID:", err)
			return nil, err
		}
		filter["_id"] = convBsonStringOid
	}
	if fvs.BsonBytesOid != nil {
		convBsonBytesOid, err := bson.ToObjectID(fvs.BsonBytesOid)
		if err != nil {
			mongo.Logs().Error("could not convert value to ObjectID:", err)
			return nil, err
		}
		filter["bson_bytes_oid"] = convBsonBytesOid
	}
	return x.Find(ctx, "filter", filter)
}

// Insert implements SingleCollectionWriter
func (x *SingleCollection) Insert(ctx context.Context, obj *single.Single) (*single.Single, error) {
	ctx, cancel := context.WithTimeout(ctx, x.config.Timeout)
	defer cancel()

	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "insert"),
	)
	stats.Record(ctx, mongo.MeasureInflight.M(1))
	start := time.Now()
	defer func(ctx context.Context) {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, mongo.MeasureLatency.M(dur), mongo.MeasureInflight.M(-1))
	}(ctx)

	client, err := x.dialer.Dial(ctx)
	if err != nil {
		mongo.Logs().Error("could not dial:", err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}
	defer client.Close(ctx)

	res, err := client.
		Database(x.config.Database).
		Collection(x.config.Collection).
		InsertOne(ctx, obj)
	if err != nil {
		mongo.Logs().Error("could not execute insert:", err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}

	obj.BsonStringOid = res.InsertedID.(string)

	return obj, nil
}

// Upsert implements SingleCollectionWriter
func (x *SingleCollection) Upsert(ctx context.Context, obj *single.Single) (*single.Single, error) {
	ctx, cancel := context.WithTimeout(ctx, x.config.Timeout)
	defer cancel()

	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "upsert"),
	)
	stats.Record(ctx, mongo.MeasureInflight.M(1))
	start := time.Now()
	defer func(ctx context.Context) {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, mongo.MeasureLatency.M(dur), mongo.MeasureInflight.M(-1))
	}(ctx)

	client, err := x.dialer.Dial(ctx)
	if err != nil {
		mongo.Logs().Error("could not dial:", err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}
	defer client.Close(ctx)

	mObj, err := ToSingleMongo(obj)
	if err != nil {
		mongo.Logs().Error("could not convert internal to mongo:", obj, err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}

	opts := &mongo.UpdateOptions{}
	opts.SetUpsert(true)

	filter := bson.M{"_id": obj.BsonStringOid}

	res, err := client.
		Database(x.config.Database).
		Collection(x.config.Collection).
		UpdateOne(ctx, filter, mObj, opts)

	if err != nil {
		mongo.Logs().Error("could not execute upsert:", err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}

	oid, ok := res.UpsertedID.(bson.ObjectID)
	if !ok {
		mongo.Logs().Error("could not convert returned upsert id:", oid, err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, mongo.ErrBadObjID
	}

	obj.BsonStringOid = oid.Hex()

	return obj, nil
}

// Update implements SingleCollectionWriter
func (x *SingleCollection) Update(ctx context.Context, obj *single.Single, fvs *dal.SingleFieldValues) (*single.Single, error) {
	ctx, cancel := context.WithTimeout(ctx, x.config.Timeout)
	defer cancel()

	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "update"),
	)
	stats.Record(ctx, mongo.MeasureInflight.M(1))
	start := time.Now()
	defer func(ctx context.Context) {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64(time.Millisecond)
		stats.Record(ctx, mongo.MeasureLatency.M(dur), mongo.MeasureInflight.M(-1))
	}(ctx)

	client, err := x.dialer.Dial(ctx)
	if err != nil {
		return nil, err
	}
	defer client.Close(ctx)

	upd := bson.M{}
	if fvs.ScalarInt32 != nil {
		upd["scalar_int32"] = *fvs.ScalarInt32
	}
	if fvs.ScalarInt64 != nil {
		upd["scalar_int64"] = *fvs.ScalarInt64
	}
	if fvs.ScalarFloat32 != nil {
		upd["scalar_float32"] = *fvs.ScalarFloat32
	}
	if fvs.ScalarFloat64 != nil {
		upd["scalar_float64"] = *fvs.ScalarFloat64
	}
	if fvs.ScalarString != nil {
		upd["scalar_string"] = *fvs.ScalarString
	}
	if fvs.ScalarBytes != nil {
		upd["scalar_bytes"] = fvs.ScalarBytes
	}
	if fvs.ScalarBool != nil {
		upd["scalar_bool"] = *fvs.ScalarBool
	}
	if fvs.ScalarEnum != nil {
		upd["scalar_enum"] = *fvs.ScalarEnum
	}
	if fvs.ObjMessage != nil {
		upd["obj_message"] = fvs.ObjMessage
	}

	if fvs.Renamed != nil {
		upd["aliased"] = *fvs.Renamed
	}
	if fvs.IgnoredPostgres != nil {
		upd["ignored_postgres"] = *fvs.IgnoredPostgres
	}
	if fvs.RenamedPostgres != nil {
		upd["renamed_postgres"] = *fvs.RenamedPostgres
	}
	if fvs.IgnoredRest != nil {
		upd["ignored_rest"] = *fvs.IgnoredRest
	}
	if fvs.RenamedRest != nil {
		upd["renamed_rest"] = *fvs.RenamedRest
	}

	if fvs.RenamedMongo != nil {
		upd["aliased_mongo"] = *fvs.RenamedMongo
	}
	if fvs.BsonStringOid != nil {
		convBsonStringOid, err := bson.ToObjectID(*fvs.BsonStringOid)
		if err != nil {
			mongo.Logs().Error("could not convert to ObjectID:", *fvs.BsonStringOid, err)
			stats.Record(ctx, mongo.MeasureError.M(1))
			return nil, err
		}
		upd["_id"] = convBsonStringOid
	}
	if fvs.BsonBytesOid != nil {
		convBsonBytesOid, err := bson.ToObjectID(fvs.BsonBytesOid)
		if err != nil {
			mongo.Logs().Error("could not convert to ObjectID:", fvs.BsonBytesOid, err)
			stats.Record(ctx, mongo.MeasureError.M(1))
			return nil, err
		}
		upd["bson_bytes_oid"] = convBsonBytesOid
	}
	filter := bson.M{"_id": obj.BsonStringOid}

	_, err = client.
		Database(x.config.Database).
		Collection(x.config.Collection).
		UpdateOne(ctx, filter, upd)

	if err != nil {
		mongo.Logs().Error("could not update:", err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}

	if fvs.ScalarInt32 != nil {
		obj.ScalarInt32 = *fvs.ScalarInt32
	}
	if fvs.ScalarInt64 != nil {
		obj.ScalarInt64 = *fvs.ScalarInt64
	}
	if fvs.ScalarFloat32 != nil {
		obj.ScalarFloat32 = *fvs.ScalarFloat32
	}
	if fvs.ScalarFloat64 != nil {
		obj.ScalarFloat64 = *fvs.ScalarFloat64
	}
	if fvs.ScalarString != nil {
		obj.ScalarString = *fvs.ScalarString
	}
	if fvs.ScalarBytes != nil {
		obj.ScalarBytes = fvs.ScalarBytes
	}
	if fvs.ScalarBool != nil {
		obj.ScalarBool = *fvs.ScalarBool
	}
	if fvs.ScalarEnum != nil {
		obj.ScalarEnum = *fvs.ScalarEnum
	}
	if fvs.ObjMessage != nil {
		obj.ObjMessage = fvs.ObjMessage
	}

	if fvs.Renamed != nil {
		obj.Renamed = *fvs.Renamed
	}
	if fvs.IgnoredPostgres != nil {
		obj.IgnoredPostgres = *fvs.IgnoredPostgres
	}
	if fvs.RenamedPostgres != nil {
		obj.RenamedPostgres = *fvs.RenamedPostgres
	}
	if fvs.IgnoredRest != nil {
		obj.IgnoredRest = *fvs.IgnoredRest
	}
	if fvs.RenamedRest != nil {
		obj.RenamedRest = *fvs.RenamedRest
	}

	if fvs.RenamedMongo != nil {
		obj.RenamedMongo = *fvs.RenamedMongo
	}
	if fvs.BsonStringOid != nil {
		obj.BsonStringOid = *fvs.BsonStringOid
	}
	if fvs.BsonBytesOid != nil {
		obj.BsonBytesOid = fvs.BsonBytesOid
	}
	return obj, nil
}

// ById implements dal.SingleCollection.ById
func (x *SingleCollection) ById(ctx context.Context, bson_string_oid string) ([]*single.Single, error) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "by_id"),
	)

	filter := bson.M{}

	convBsonStringOid, err := bson.ToObjectID(bson_string_oid)
	if err != nil {
		mongo.Logs().Error("could not convert to ObjectID:", bson_string_oid, err)
		stats.Record(ctx, mongo.MeasureError.M(1))
		return nil, err
	}
	filter["_id"] = bson.M{"$eq": convBsonStringOid}

	return x.Find(ctx, "by_id", filter)
}

// OneParam implements dal.SingleCollection.OneParam
func (x *SingleCollection) OneParam(ctx context.Context, scalar_int32 int32) ([]*single.Single, error) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "one_param"),
	)

	filter := bson.M{}

	filter["scalar_int32"] = bson.M{"$eq": scalar_int32}

	return x.Find(ctx, "one_param", filter)
}

// MultipleParam implements dal.SingleCollection.MultipleParam
func (x *SingleCollection) MultipleParam(ctx context.Context, scalar_int32 int32, scalar_int64 int64, scalar_float32 float32) ([]*single.Single, error) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "multiple_param"),
	)

	filter := bson.M{}

	filter["scalar_int32"] = bson.M{"$eq": scalar_int32}

	filter["scalar_int64"] = bson.M{"$eq": scalar_int64}

	filter["scalar_float32"] = bson.M{"$eq": scalar_float32}

	return x.Find(ctx, "multiple_param", filter)
}

// MessageParam implements dal.SingleCollection.MessageParam
func (x *SingleCollection) MessageParam(ctx context.Context, obj_message *single.Single_Message) ([]*single.Single, error) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "message_param"),
	)

	filter := bson.M{}

	filter["obj_message"] = bson.M{"$eq": obj_message}

	return x.Find(ctx, "message_param", filter)
}

// WithComparator implements dal.SingleCollection.WithComparator
func (x *SingleCollection) WithComparator(ctx context.Context, scalar_int32 int32) ([]*single.Single, error) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "with_comparator"),
	)

	filter := bson.M{}

	filter["scalar_int32"] = bson.M{"$gt": scalar_int32}

	return x.Find(ctx, "with_comparator", filter)
}

// WithRest implements dal.SingleCollection.WithRest
func (x *SingleCollection) WithRest(ctx context.Context, scalar_int32 int32, scalar_int64 int64, scalar_float32 float32, scalar_float64 float64) ([]*single.Single, error) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "with_rest"),
	)

	filter := bson.M{}

	filter["scalar_int32"] = bson.M{"$eq": scalar_int32}

	filter["scalar_int64"] = bson.M{"$eq": scalar_int64}

	filter["scalar_float32"] = bson.M{"$eq": scalar_float32}

	filter["scalar_float64"] = bson.M{"$eq": scalar_float64}

	return x.Find(ctx, "with_rest", filter)
}

// ProviderStubOnly implements dal.SingleCollection.ProviderStubOnly
func (x *SingleCollection) ProviderStubOnly(ctx context.Context) ([]*single.Single, error) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "provider_stub_only"),
	)

	filter := bson.M{}

	return x.Find(ctx, "provider_stub_only", filter)
}

// NonFieldOnly implements dal.SingleCollection.NonFieldOnly
func (x *SingleCollection) NonFieldOnly(ctx context.Context, kind string) ([]*single.Single, error) {
	ctx, _ = tag.New(ctx,
		tag.Upsert(mongo.TagCollection, "single"),
		tag.Upsert(mongo.TagInstance, x.name),
		tag.Upsert(mongo.TagMethod, "non_field_only"),
	)

	filter := bson.M{}

	filter["kind"] = bson.M{"$eq": kind}

	return x.Find(ctx, "non_field_only", filter)
}

// NewSingleCollection returns a new SingleCollection.
func NewSingleCollection(instance string, dialer mongo.Dialer, config *SingleConfig) (*SingleCollection, error) {
	coll := &SingleCollection{
		name:   instance,
		dialer: dialer,
		config: config,
	}

	return coll, nil
}
