package mongo

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/go-test/deep"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"golang.org/x/tools/imports"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	generatedCommon bool
)

// Collection is a struct that generates a collection file.
type Collection struct {
	File    *File
	Message *Message
	Fields  *Fields
	Queries *Queries
	Opts    *annotations.DalOptions

	plugin   *protogen.Plugin
	filename string
}

// NewCollection returns a new collection renderer.
func NewCollection(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) *Collection {
	base := path.Base(file.GeneratedFilenamePrefix)
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/mongo/%s.genms.dal.%s.go", base, strings.ToLower(msg.GoIdent.GoName)))
	outfile := plugin.NewGeneratedFile(filename, ".")

	cfile := NewFile(outfile, file)
	cmsg := NewMessage(cfile, msg)
	cfields := NewFields(cmsg)
	cqueries := NewQueries(cfile, cfields, opts)

	return &Collection{
		File:     cfile,
		Message:  cmsg,
		Fields:   cfields,
		Queries:  cqueries,
		Opts:     opts,
		plugin:   plugin,
		filename: filename,
	}
}

func (c *Collection) render() error {
	if !generatedCommon {
		if err := c.defineCommon(); err != nil {
			return err
		}
		generatedCommon = true
	}

	steps := []func() error{
		c.definePackage,
		c.defineConfig,
		c.defineInternalStructs,
		c.defineCollection,
		c.defineService,
		c.defineDefaultQueries,
		c.defineQueries,
		c.defineNewCollection,
	}

	for _, s := range steps {
		if err := s(); err != nil {
			return err
		}
	}

	outfile := c.File.Outfile()
	original, err := outfile.Content()
	if err != nil {
		return err
	}
	formatted, err := imports.Process(c.filename, original, nil)

	if diff := deep.Equal(original, formatted); diff != nil {
		formattedOutfile := c.plugin.NewGeneratedFile(c.filename, ".")
		if _, err := formattedOutfile.Write(formatted); err != nil {
			return err
		}
		outfile.Skip()
	}

	return nil
}

func (c *Collection) defineCommon() error {
	file := c.File.Proto()
	plugin := c.plugin
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/mongo/mongo.genms.go"))
	outfile := plugin.NewGeneratedFile(filename, ".")

	tmplSrc := `// Package {{ $.C.File.MongoPackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ $.C.File.MongoPackageName }}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

var (
	logs = {{ $.P.Mongo }}.Logs()
)

`

	tmpl, err := template.New("defineCommon").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	p := map[string]string{
		"Mongo": c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := outfile.Write(buf.Bytes()); err != nil {
		return err
	}

	original, err := outfile.Content()
	if err != nil {
		return err
	}
	formatted, err := imports.Process(filename, original, nil)

	if diff := deep.Equal(original, formatted); diff != nil {
		formattedOutfile := plugin.NewGeneratedFile(filename, ".")
		if _, err := formattedOutfile.Write(formatted); err != nil {
			return err
		}
		outfile.Skip()
		outfile = formattedOutfile
	}

	return nil
}

func (c *Collection) definePackage() error {
	tmplSrc := `// Package {{ .File.MongoPackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ .File.MongoPackageName }}
`

	tmpl, err := template.New("defineMongoPackage").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, c); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineCollection() error {
	tmplSrc := `// {{ $.C.Message.Name }}Collection is an autogenerated implementation of {{ $.C.Message.QualifiedDalKind }}Collection.
type {{ $.C.Message.Name }}Collection struct {
	{{ $.P.Collection }}.Unimplemented{{ $.C.Message.Name }}Collection

	name string
	dialer {{ $.P.Mongo }}.Dialer
	config *{{ $.C.Message.Name }}Config
	mutators []{{ $.P.Collection }}.{{ $.C.Message.Name }}Mutator
	defaultFilter {{ $.P.Bson }}.M
}

// WithMutators adds {{ $.P.Collection }}.{{ $.C.Message.Name }}Mutators to the collection. These will be applied to all values after they are read from mongo.
func (x *{{ $.C.Message.Name }}Collection) WithMutators(muts ...{{ $.P.Collection }}.{{ $.C.Message.Name }}Mutator) {
	x.mutators = append(x.mutators, muts...)
}

// WithDefaultFilter sets the default the default mongo filter to apply to all find queries. This is applied after any query args.
func (x *{{ $.C.Message.Name }}Collection) WithDefaultFilter(f {{ $.P.Bson }}.M) {
	x.defaultFilter = f
}
`

	tmpl, err := template.New("defineMongoCollection").
		Funcs(template.FuncMap{
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Bson":       c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo/bson"),
		"Collection": c.File.QualifiedPackageName(c.File.DalPackagePath()),
		"Mongo":      c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineService() error {
	tmplSrc := `// Initialize initializes and starts the service. Initialize should panic in case of
// any errors. It is intended that Initialize be called only once during the service life-cycle.
func (x *{{ $.C.Message.Name }}Collection) Initialize(_ {{ $.P.Context }}.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *{{ $.C.Message.Name }}Collection) Shutdown(_ {{ $.P.Context }}.Context) error {
	return nil
}

// String returns a string identifier for the service.
func (x *{{ $.C.Message.Name }}Collection) String() string {
	{{- $pkg := $.C.File.MongoPackageName -}}
	if x.name != "" {
		return "{{ ToDashCase $pkg }}-{{ ToDashCase .C.Message.Name }}-" + x.name
	}
	return "{{ ToDashCase $pkg }}-{{ ToDashCase .C.Message.Name }}"
}

// NameOf returns the name of a service. This must be unique if there are multiple instances of the same
// service.
func (x *{{ $.C.Message.Name }}Collection) NameOf() string {
	return x.String()
}
`

	tmpl, err := template.New("defineMongoService").
		Funcs(template.FuncMap{
			"ToDashCase": protocgenlib.ToDashCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Bson":    c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo/bson"),
		"Context": c.File.QualifiedPackageName("context"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineDefaultQueries() error {
	tmplSrc := `{{- $id := $.C.Fields.ID -}}	
// Find scans the collection for records matching the filter. 
func (x *{{ $.C.Message.Name }}Collection) Find(ctx {{ $.P.Context }}.Context, label string, filter {{ $.P.Bson }}.M, opts ...*{{ $.P.Mongo }}.FindOptions) ([]*{{ $.C.Message.QualifiedKind }}, error) {	
	ctx, cancel := {{ $.P.Context }}.WithTimeout(ctx, x.config.Timeout)
	defer cancel()

	client, err := x.dialer.Dial(ctx)
	if err != nil {
		logs.Error("could not dial:", err)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
		return nil, err
	}
	defer client.Close(ctx)

	ctx, _ = {{ $.P.Tag }}.New(ctx,
		{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagCollection, "{{ ToSnakeCase $.C.Message.Name }}"),
		{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagInstance, x.name),
		{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagMethod, label),
	)
	{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureInflight.M(1))
	start := {{ $.P.Time }}.Now()
	defer func(ctx {{ $.P.Context }}.Context) {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ $.P.Time }}.Millisecond)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureLatency.M(dur), {{ $.P.Mongo }}.MeasureInflight.M(-1))
	}(ctx)

	for k, v := range x.defaultFilter {
		filter[k] = v
	}	

	cur, err := client.
		Database(x.config.Database).
		Collection(x.config.Collection).
		Find(ctx, filter, {{ ToLower $.C.Message.Name }}Projection, opts...)
	if err != nil {
		logs.Error("could not execute rest request:", err)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
		return nil, err
	}

	vals := []*{{ $.C.Message.QualifiedKind }}{}
	for cur.Next(ctx) {
		obj := &{{ $.C.Message.Name }}Mongo{}
		if err = cur.Decode(obj); err != nil {
			logs.Errorf("could not parse %s - %v", label, err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, err
		}

		val, err := obj.{{ $.C.Message.Name }}()
		if err != nil {
			logs.Error("could not convert from mongo to internal:", err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, err
		}

		for _, m := range x.mutators {
			val, err = m(val)
			if err != nil {
				logs.Error("could not mutate value:", val, err)
				{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
				return nil, err
			}
		}

		vals = append(vals, val)
	}

	return vals, nil
}

// All implements {{ $.C.Message.QualifiedDalKind }}Collection.All
func (x *{{ $.C.Message.Name }}Collection) All(ctx {{ $.P.Context }}.Context) ([]*{{ $.C.Message.QualifiedKind }}, error) {	
	return x.Find(ctx, "all", {{ $.P.Bson }}.M{})
}

// Filter implements {{ $.C.Message.Name }}CollectionReader
func (x *{{ $.C.Message.Name }}Collection) Filter(ctx {{ $.P.Context }}.Context, fvs *{{ $.P.Collection }}.{{ $.C.Message.Name }}FieldValues) ([]*{{ $.C.Message.QualifiedKind }}, error) {
	filter := {{ $.P.Bson }}.M{}

	{{ range $fn := $.C.Fields.Names -}}
		{{- $f := ($.C.Fields.ByName $fn) -}}
		{{- if not $f.Ignore -}}
			{{- $conv := $f.ToMongo -}}
			if fvs.{{ ToTitleCase $f.Name }} != nil {
				{{ if eq $conv "ObjectID" -}}
					conv{{ ToTitleCase $f.Name }}, err := {{ $.P.Bson }}.ToObjectID({{ if not $f.IsRef }}*{{ end }}fvs.{{ ToTitleCase $f.Name }})
					if err != nil {
						logs.Error("could not convert value to ObjectID:", err)
						return nil, err
					}
					filter["{{ $f.QueryName }}"] = conv{{ ToTitleCase $f.Name }}
				{{- else -}}
					filter["{{ $f.QueryName }}"] = {{ if not $f.IsRef }}*{{ end }}fvs.{{ ToTitleCase $f.Name }}
				{{- end }}
			}
		{{- end }}
	{{ end -}}

	return x.Find(ctx, "filter", filter)
}

// Insert implements {{ $.C.Message.Name }}CollectionWriter
func (x *{{ $.C.Message.Name }}Collection) Insert(ctx {{ $.P.Context }}.Context, obj *{{ $.C.Message.QualifiedKind }}) (*{{ $.C.Message.QualifiedKind }}, error) {
	ctx, cancel := {{ $.P.Context }}.WithTimeout(ctx, x.config.Timeout)
	defer cancel()

	ctx, _ = {{ $.P.Tag }}.New(ctx,
		{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagCollection, "{{ ToSnakeCase $.C.Message.Name }}"),
		{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagInstance, x.name),
		{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagMethod, "insert"),
	)
	{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureInflight.M(1))
	start := {{ $.P.Time }}.Now()
	defer func(ctx {{ $.P.Context }}.Context) {
		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ $.P.Time }}.Millisecond)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureLatency.M(dur), {{ $.P.Mongo }}.MeasureInflight.M(-1))
	}(ctx)

	client, err := x.dialer.Dial(ctx)
	if err != nil {
		logs.Error("could not dial:", err)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
		return nil, err
	}
	defer client.Close(ctx)
	
	mObj, err := To{{ $.C.Message.Name }}Mongo(obj)
	if err != nil {
		logs.Error("could not convert internal to mongo:", obj, err)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
		return nil, err
	}

	{{ if $id -}}res{{- else -}}_{{- end }}, err := client.
		Database(x.config.Database).
		Collection(x.config.Collection).
		InsertOne(ctx, mObj)

	if err != nil {
		logs.Error("could not execute insert:", err)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
		return nil, err
	}

	{{ if $id }}
		oid, ok := res.InsertedID.({{ $.P.Bson }}.ObjectID)
		if !ok {
			logs.Error("could not convert returned upsert id:", oid, err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, {{ $.P.Mongo }}.ErrBadObjID
		}

		{{ if eq $id.Kind "string" }}
			obj.{{ ToTitleCase $id.Name }} = oid.Hex()
		{{ else if eq $id.Kind "[]byte" }}
			obj.{{ ToTitleCase $id.Name }} = []byte(oid[:])
		{{ end }}
	{{ end }}

	return obj, nil
}

// Upsert implements {{ $.C.Message.Name }}CollectionWriter
func (x *{{ $.C.Message.Name }}Collection) Upsert(ctx {{ $.P.Context }}.Context, obj *{{ $.C.Message.QualifiedKind }}) (*{{ $.C.Message.QualifiedKind }}, error) {
	{{ if not $id -}}
		return nil, {{ $.P.Mongo }}.ErrNoID
	{{- else -}}
		ctx, cancel := {{ $.P.Context }}.WithTimeout(ctx, x.config.Timeout)
		defer cancel()

		ctx, _ = {{ $.P.Tag }}.New(ctx,
			{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagCollection, "{{ ToSnakeCase $.C.Message.Name }}"),
			{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagInstance, x.name),
			{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagMethod, "upsert"),
		)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureInflight.M(1))
		start := {{ $.P.Time }}.Now()
		defer func(ctx {{ $.P.Context }}.Context) {
			stop := time.Now()
			dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ $.P.Time }}.Millisecond)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureLatency.M(dur), {{ $.P.Mongo }}.MeasureInflight.M(-1))
		}(ctx)

		client, err := x.dialer.Dial(ctx)
		if err != nil {
			logs.Error("could not dial:", err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, err
		}
		defer client.Close(ctx)

		mObj, err := To{{ $.C.Message.Name }}Mongo(obj)
		if err != nil {
			logs.Error("could not convert internal to mongo:", obj, err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, err
		}

		opts := &{{ $.P.Mongo }}.UpdateOptions{}
		opts.SetUpsert(true)
		
		filter := {{ $.P.Bson }}.M{"{{ $id.QueryName }}": mObj.{{ ToTitleCase $id.Name }} }

		res, err := client.
			Database(x.config.Database).
			Collection(x.config.Collection).
			UpdateOne(ctx, filter, {{ $.P.Bson }}.M{"$set": mObj}, opts)

		if err != nil {
			logs.Error("could not execute upsert:", err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, err
		}

		oid, ok := res.UpsertedID.({{ $.P.Bson }}.ObjectID)
		if !ok {
			logs.Error("could not convert returned upsert id:", oid, err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, {{ $.P.Mongo }}.ErrBadObjID
		}

		{{ if eq $id.Kind "string" }}
			obj.{{ ToTitleCase $id.Name }} = oid.Hex()
		{{ else if eq $id.Kind "[]byte" }}
			obj.{{ ToTitleCase $id.Name }} = []byte(oid[:])
		{{ end }}

		return obj, nil
	{{- end }}
}

// Update implements {{ $.C.Message.Name }}CollectionWriter
func (x *{{ $.C.Message.Name }}Collection) Update(ctx {{ $.P.Context }}.Context, obj *{{ $.C.Message.QualifiedKind }}, fvs *{{ $.P.Collection }}.{{ $.C.Message.Name }}FieldValues) (*{{ $.C.Message.QualifiedKind }}, error){
	{{ if not $id -}}
		return nil, {{ $.P.Mongo }}.ErrNoID
	{{- else -}}
		ctx, cancel := {{ $.P.Context }}.WithTimeout(ctx, x.config.Timeout)
		defer cancel()

		ctx, _ = {{ $.P.Tag }}.New(ctx,
			{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagCollection, "{{ ToSnakeCase $.C.Message.Name }}"),
			{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagInstance, x.name),
			{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagMethod, "update"),
		)
		{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureInflight.M(1))
		start := {{ $.P.Time }}.Now()
		defer func(ctx {{ $.P.Context }}.Context) {
			stop := time.Now()
			dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ $.P.Time }}.Millisecond)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureLatency.M(dur), {{ $.P.Mongo }}.MeasureInflight.M(-1))
		}(ctx)

		client, err := x.dialer.Dial(ctx)
		if err != nil {
			return nil, err
		}
		defer client.Close(ctx)

		objID, err := {{ $.P.Bson }}.ToObjectID(obj.{{ ToTitleCase $id.Name }})
		if err != nil {
			logs.Error("could not convert to ObjectID:", obj.{{ ToTitleCase $id.Name }}, err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, err
		}
		
		upd := {{ $.P.Bson }}.M{}
		{{ range $fn := $.C.Fields.Names -}}
			{{- $f := ($.C.Fields.ByName $fn) -}}
			{{- if not $f.Ignore -}}
				{{- $conv := $f.ToMongo -}}
				if fvs.{{ ToTitleCase $f.Name }} != nil {
					{{ if eq $conv "ObjectID" -}}
						conv{{ ToTitleCase $f.Name }}, err := {{ $.P.Bson }}.ToObjectID({{- if not $f.IsRef -}}*{{- end -}}fvs.{{ ToTitleCase $f.Name }})
						if err != nil {
							logs.Error("could not convert to ObjectID:", {{- if not $f.IsRef -}}*{{- end -}}fvs.{{ ToTitleCase $f.Name }}, err)
							{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
							return nil, err
						}
						upd["{{ $f.QueryName }}"] = conv{{ ToTitleCase $f.Name }}
					{{- else -}}
						upd["{{ $f.QueryName }}"] = {{- if not $f.IsRef -}}*{{- end -}}fvs.{{ ToTitleCase $f.Name }}
					{{- end }}
				}
			{{- end }}
		{{ end -}}

		filter := {{ $.P.Bson }}.M{"{{ $id.QueryName }}": objID}

		_, err = client.
			Database(x.config.Database).
			Collection(x.config.Collection).
			UpdateOne(ctx, filter, {{ $.P.Bson }}.M{"$set": upd})

		if err != nil {
			logs.Error("could not update:", err)
			{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
			return nil, err
		}

		{{ range $fn := $.C.Fields.Names -}}
			{{- $f := ($.C.Fields.ByName $fn) -}}
			{{- if not $f.Ignore -}}
				if fvs.{{ ToTitleCase $f.Name }} != nil {
					obj.{{ ToTitleCase $f.Name }} = {{- if not $f.IsRef -}}*{{- end -}}fvs.{{ ToTitleCase $f.Name }}
				}
			{{- end }}
		{{ end -}}

		return obj, nil
	{{- end }}
}

`

	tmpl, err := template.New("defineMongoDefaultQueries").
		Funcs(template.FuncMap{
			"Arg":         NewArg,
			"ToCamelCase": protocgenlib.ToCamelCase,
			"ToLower":     strings.ToLower,
			"ToSnakeCase": protocgenlib.ToSnakeCase,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Bson":       c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo/bson"),
		"Collection": c.File.QualifiedPackageName(c.File.DalPackagePath()),
		"Context":    c.File.QualifiedPackageName("context"),
		"Mongo":      c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo"),
		"Stats":      c.File.QualifiedPackageName("go.opencensus.io/stats"),
		"Tag":        c.File.QualifiedPackageName("go.opencensus.io/tag"),
		"Time":       c.File.QualifiedPackageName("time"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineQueries() error {
	tmplSrc := `{{ range $qn := $.C.Queries.Names }}
	{{- $q := ($.C.Queries.ByName $qn) -}}
	{{- if $q.QueryProvided }}
		// {{ ToTitleCase $q.Name }} implements {{ $.C.Message.QualifiedDalKind }}Collection.{{ ToTitleCase $q.Name }}
		func (x *{{ $.C.Message.Name }}Collection){{ ToTitleCase $q.Name }}(ctx {{ $.P.Context }}.Context
			{{- range $a := $q.Args -}}
				{{- $arg := (Arg $.C.File $.C.Fields $a) -}}
				, {{ ToSnakeCase $arg.Name }} {{ $arg.QualifiedKind }}
			{{- end -}}
		) ([]*{{ $.C.Message.QualifiedKind }}, error) {
			ctx, _ = {{ $.P.Tag }}.New(ctx,
				{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagCollection, "{{ ToSnakeCase $.C.Message.Name }}"),
				{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagInstance, x.name),
				{{ $.P.Tag }}.Upsert({{ $.P.Mongo }}.TagMethod, "{{ ToSnakeCase $q.Name }}"),
			)

			filter := {{ $.P.Bson }}.M{}
			
			{{ range $a := $q.Args }}
				{{ $arg := (Arg $.C.File $.C.Fields $a) }}
				{{ $conv := $arg.ToMongo }}
				{{ if eq $conv "ObjectID" -}}
					conv{{ ToTitleCase $arg.Name }}, err := {{ $.P.Bson }}.ToObjectID({{ $arg.Name }})
					if err != nil {
						logs.Error("could not convert to ObjectID:", {{ $arg.Name }}, err)
						{{ $.P.Stats }}.Record(ctx, {{ $.P.Mongo }}.MeasureError.M(1))
						return nil, err
					}
					filter["{{ $arg.QueryName }}"] = {{ $.P.Bson }}.M{ "{{ $arg.Comparison }}": conv{{ ToTitleCase $arg.Name }} }
				{{- else -}}
					filter["{{ $arg.QueryName }}"] = {{ $.P.Bson }}.M{ "{{ $arg.Comparison }}": {{ ToSnakeCase $arg.Name }} }
				{{- end }}
			{{- end }}
						
			return x.Find(ctx, "{{ ToSnakeCase $q.Name }}", filter)
		}
	{{- end -}}
{{- end }}
`

	tmpl, err := template.New("defineMongoQueries").
		Funcs(template.FuncMap{
			"Arg":         NewArg,
			"ToCamelCase": protocgenlib.ToCamelCase,
			"ToSnakeCase": protocgenlib.ToSnakeCase,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Bson":    c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo/bson"),
		"Mongo":   c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo"),
		"Context": c.File.QualifiedPackageName("context"),
		"Stats":   c.File.QualifiedPackageName("go.opencensus.io/stats"),
		"Tag":     c.File.QualifiedPackageName("go.opencensus.io/tag"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineNewCollection() error {
	tmplSrc := `// New{{ $.C.Message.Name }}Collection returns a new {{ $.C.Message.Name }}Collection.
func New{{ $.C.Message.Name }}Collection(instance string, dialer {{ $.P.Mongo }}.Dialer, config *{{ $.C.Message.Name }}Config) (*{{ $.C.Message.Name }}Collection, error) {
	coll := &{{ $.C.Message.Name }}Collection{
		name: instance,
		dialer: dialer,
		config: config,
	}

	return coll, nil
}
`

	tmpl, err := template.New("defineMongoNewCollection").
		Funcs(template.FuncMap{
			"ToSnakeCase": protocgenlib.ToSnakeCase,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Mongo": c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineInternalStructs() error {
	tmplSrc := `var (
	{{ ToLower $.C.Message.Name }}Projection = {{ $.P.Bson }}.M{
		{{ range $n := $.C.Fields.Names -}}
			{{- $f := ($.C.Fields.ByName $n) -}}
			{{- if not $f.Ignore -}}
				"{{ $f.QueryName }}": 1,
			{{- end }}
		{{ end -}}
	}
)

// {{ $.C.Message.Name }}Mongo is an autogenerated struct that
// is used to parse query results.
type {{ $.C.Message.Name }}Mongo struct {
	{{ range $n := $.C.Fields.Names -}}
		{{- $f := ($.C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			{{ ToTitleCase $f.Name }} {{ $f.QualifiedQueryKind }} ` + "`" + `bson:"{{ $f.QueryName }},omitempty"` + "`" + `
		{{- end }}
	{{ end -}}
}

// {{ $.C.Message.Name }} returns a new {{ $.C.Message.QualifiedKind }} populated with scanned values.
func (x *{{ $.C.Message.Name }}Mongo) {{ $.C.Message.Name }}() (*{{ $.C.Message.QualifiedKind }}, error) {
	y := &{{ $.C.Message.QualifiedKind }}{}

	{{ range $n := $.C.Fields.Names -}}
		{{- $f := ($.C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			{{- $convFrom := $f.ToMongo -}}
			{{ if eq $convFrom "ObjectID" -}}
				{{- $convTo := $f.ToGo -}}
				{{- if eq $convTo "string" -}}
					y.{{ ToTitleCase $f.Name }} = x.{{ ToTitleCase $f.Name }}.Hex()
				{{- else if eq $convTo "[]byte" -}}
					y.{{ ToTitleCase $f.Name }} = (x.{{ ToTitleCase $f.Name }})[:]
				{{- end -}}
			{{- else -}}
				y.{{ ToTitleCase $f.Name }} = x.{{ ToTitleCase $f.Name }}
			{{- end }}
		{{- end }}
	{{ end -}}
	return y, nil
}

// To{{ $.C.Message.Name }}Mongo converts the given {{ $.C.Message.Name }} into the internal mongo equivalent.
func To{{ $.C.Message.Name }}Mongo(obj *{{ $.C.Message.QualifiedKind }}) (*{{ $.C.Message.Name }}Mongo, error) {
	mObj := &{{ $.C.Message.Name }}Mongo{}

	{{ range $fn := $.C.Fields.Names -}}
		{{- $f := ($.C.Fields.ByName $fn) -}}
		{{- if not $f.Ignore -}}
			{{- $convTo := $f.ToMongo -}}
			{{- $convFrom := $f.ToGo -}}
			{{ if eq $convTo "ObjectID" -}}
				{{ if eq $convFrom "string" }}
					if obj.{{ ToTitleCase $f.Name }} != "" {
						conv{{ ToTitleCase $f.Name }}, err := {{ $.P.Bson }}.ToObjectID(obj.{{ ToTitleCase $f.Name }})
						if err != nil {
							return nil, err
						}
						mObj.{{ ToTitleCase $f.Name }} = conv{{ ToTitleCase $f.Name }}
					}
				{{ else if eq $convFrom "[]byte" }}
					if len(obj.{{ ToTitleCase $f.Name }}) != 0 {
						conv{{ ToTitleCase $f.Name }}, err := {{ $.P.Bson }}.ToObjectID(obj.{{ ToTitleCase $f.Name }})
						if err != nil {
							return nil, err
						}
						mObj.{{ ToTitleCase $f.Name }} = conv{{ ToTitleCase $f.Name }}
					}
				{{ end }}				
			{{- else -}}
				mObj.{{ ToTitleCase $f.Name }} = obj.{{ ToTitleCase $f.Name }}
			{{- end }}
		{{- end }}
	{{ end -}}

	return mObj, nil
}

`

	tmpl, err := template.New("defineMongoStructs").
		Funcs(template.FuncMap{
			"ToLower":     strings.ToLower,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Bson": c.File.QualifiedPackageName("github.com/rleszilm/genms/mongo/bson"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineConfig() error {
	tmplSrc := `// {{ $.C.Message.Name }}Config is a struct that can be used to configure a {{ $.C.Message.Name }}Collection
type {{ $.C.Message.Name }}Config struct {
	Name string ` + "`" + `envconfig:"name"` + "`" + `
	Database string ` + "`" + `envconfig:"database"` + "`" + `
	Collection string ` + "`" + `envconfig:"collection"` + "`" + `
	Timeout {{ $.P.Time }}.Duration ` + "`" + `envconfig:"timeout" default:"5s"` + "`" + `
}
`

	tmpl, err := template.New("defineMongoConfig").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Time": c.File.QualifiedPackageName("time"),
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

// GenerateCollection generates the dal interface for the collection
func GenerateCollection(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	c := NewCollection(plugin, file, msg, opts)
	return c.render()
}
