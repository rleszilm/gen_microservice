package postgres

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
	"google.golang.org/protobuf/compiler/protogen"
)

// Collection is a struct that generates a colelction file.
type Collection struct {
	File    *File
	Message *Message
	Fields  *Fields
	Queries *Queries
	Opts    *annotations.DalOptions
}

// NewCollection returns a new collection renderer.
func NewCollection(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) *Collection {
	base := path.Base(file.GeneratedFilenamePrefix)
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/postgres/%s.genms.dal.%s.go", base, strings.ToLower(msg.GoIdent.GoName)))
	outfile := plugin.NewGeneratedFile(filename, ".")

	cfile := NewFile(outfile, file)
	cmsg := NewMessage(cfile, msg)
	cfields := NewFields(cmsg)
	cqueries := NewQueries(opts)

	return &Collection{
		File:    cfile,
		Message: cmsg,
		Fields:  cfields,
		Queries: cqueries,
		Opts:    opts,
	}
}

// GenerateCollection generates the dal interface for the collection
func GenerateCollection(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	c := NewCollection(plugin, file, msg, opts)
	return c.render()
}

func (c *Collection) render() error {
	steps := []func() error{
		c.definePackage,
		c.defineCollection,
		c.defineService,
		c.defineDefaultQueries,
		c.defineQueries,
		c.defineNewCollection,
		c.defineInternalStructs,
		c.defineConfig,
		c.defineTemplateProvider,
		c.defineMetrics,
	}

	for _, s := range steps {
		if err := s(); err != nil {
			return err
		}
	}

	return nil
}

func (c *Collection) definePackage() error {
	tmplSrc := `// Package {{ .File.PostgresPackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ .File.PostgresPackageName }}
`

	tmpl, err := template.New("definePostgresPackage").
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
	tmplSrc := `{{- $C := .C -}}
{{- $P := .P -}}
// {{ .C.Message.Name }}Collection is an autogenerated implementation of {{ .C.Message.QualifiedDalKind }}Collection.
type {{ .C.Message.Name }}Collection struct {
	{{ .P.Collection }}.Unimplemented{{ .C.Message.Name }}Collection

	db {{ .P.GenmsSQL }}.DB
	config *{{ .C.Message.Name }}Config

	execInsert string
	execUpsert string
	queryAll string

	{{ range $qn := .C.Queries.Names -}}
		{{- $q := $C.Queries.ByName $qn -}}
		{{- if $q.QueryProvided -}}
			query{{ ToTitleCase $q.Name }} string
		{{- end }}
	{{ end -}}
}
`

	tmpl, err := template.New("definePostgresCollection").
		Funcs(template.FuncMap{
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Collection": c.File.QualifiedPackageName(c.File.DalPackagePath()),
		"GenmsSQL":   c.File.QualifiedPackageName("github.com/rleszilm/genms/sql"),
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
func (x *{{ .C.Message.Name }}Collection) Initialize(_ {{ .P.Context }}.Context) error {
	return nil
}

// Shutdown closes the long-running instance, or service.
func (x *{{ .C.Message.Name }}Collection) Shutdown(_ {{ .P.Context }}.Context) error {
	return nil
}

// NameOf returns the name of a service. This must be unique if there are multiple instances of the same
// service.
func (x *{{ .C.Message.Name }}Collection) NameOf() string {
	return "{{ .C.File.PostgresPackageName }}_" + x.config.TableName
}

// String returns a string identifier for the service.
func (x *{{ .C.Message.Name }}Collection) String() string {
	return x.NameOf()
}
`

	tmpl, err := template.New("definePostgresService").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
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
	tmplSrc := `{{- $C := .C -}}
{{- $P := .P -}}
// DoInsert provides the base logic for {{ .C.Message.QualifiedDalKind }}Collection.Insert.
// The user should use this as a base for {{ .C.Message.QualifiedDalKind }}Collection.Insert, only having to add
// code that interprets the returned values.
func (x *{{ .C.Message.Name }}Collection) DoInsert(ctx {{ .P.Context }}.Context, arg interface{}) ({{ .P.SQL }}.Result, error) {
	var err error
	start := {{ .P.Time }}.Now()
	{{ .P.Stats }}.Record(ctx, {{ ToCamelCase .C.Message.Name }}Inflight.M(1))
	defer func() {
		stop := {{ .P.Time }}.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)

		if err != nil {
			ctx, err = {{ .P.Tag }}.New(ctx,
				{{ .P.Tag }}.Insert({{ ToCamelCase .C.Message.Name }}QueryError, "{{ ToSnakeCase .C.Message.Name }}_insert"),
			)
		}

		ctx, err = {{ .P.Tag }}.New(ctx,
			{{ .P.Tag }}.Insert({{ ToCamelCase .C.Message.Name }}QueryName, "{{ ToSnakeCase .C.Message.Name }}_insert"),
		)

		{{ .P.Stats }}.Record(ctx, {{ ToCamelCase .C.Message.Name }}Latency.M(dur), {{ ToCamelCase .C.Message.Name }}Inflight.M(-1))
	}()

	return x.db.ExecWithReplacements(ctx, x.execInsert, arg)
}

// DoUpsert provides the base logic for {{ .C.Message.QualifiedDalKind }}Collection.Upsert.
// The user should use this as a base for {{ .C.Message.QualifiedDalKind }}Collection.Upsert, only having to add
// code that interprets the returned values.
func (x *{{ .C.Message.Name }}Collection) DoUpsert(ctx {{ .P.Context }}.Context, arg interface{}) ({{ .P.SQL }}.Result, error) {
	var err error
	start := {{ .P.Time }}.Now()
	{{ .P.Stats }}.Record(ctx, {{ ToCamelCase .C.Message.Name }}Inflight.M(1))
	defer func() {
		stop := {{ .P.Time }}.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)

		if err != nil {
			ctx, err = {{ .P.Tag }}.New(ctx,
				{{ .P.Tag }}.Upsert({{ ToCamelCase .C.Message.Name }}QueryError, "{{ ToSnakeCase .C.Message.Name }}_upsert"),
			)
		}

		ctx, err = {{ .P.Tag }}.New(ctx,
			{{ .P.Tag }}.Upsert({{ ToCamelCase .C.Message.Name }}QueryName, "{{ ToSnakeCase .C.Message.Name }}_upsert"),
		)

		{{ .P.Stats }}.Record(ctx, {{ ToCamelCase .C.Message.Name }}Latency.M(dur), {{ ToCamelCase .C.Message.Name }}Inflight.M(-1))
	}()

	return x.db.ExecWithReplacements(ctx, x.execUpsert, arg)
}

// DoUpdate provides the base logic for {{ .C.Message.QualifiedDalKind }}Collection.Upsert.
// The user should use this as a base for {{ .C.Message.QualifiedDalKind }}Collection.Upsert, only having to add
// code that interprets the returned values.
func (x *{{ .C.Message.Name }}Collection) DoUpdate(ctx {{ .P.Context }}.Context, arg interface{}) ({{ .P.SQL }}.Result, error) {
	var err error
	start := {{ .P.Time }}.Now()
	{{ .P.Stats }}.Record(ctx, {{ ToCamelCase .C.Message.Name }}Inflight.M(1))
	defer func() {
		stop := {{ .P.Time }}.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)

		if err != nil {
			ctx, err = {{ .P.Tag }}.New(ctx,
				{{ .P.Tag }}.Upsert({{ ToCamelCase .C.Message.Name }}QueryError, "{{ ToSnakeCase .C.Message.Name }}_update"),
			)
		}

		ctx, err = {{ .P.Tag }}.New(ctx,
			{{ .P.Tag }}.Upsert({{ ToCamelCase .C.Message.Name }}QueryName, "{{ ToSnakeCase .C.Message.Name }}_update"),
		)

		{{ .P.Stats }}.Record(ctx, {{ ToCamelCase .C.Message.Name }}Latency.M(dur), {{ ToCamelCase .C.Message.Name }}Inflight.M(-1))
	}()

	return x.db.ExecWithReplacements(ctx, x.execUpdate, arg)
}

// All implements {{ .C.Message.QualifiedDalKind }}Collection.All
func (x *{{ .C.Message.Name }}Collection) All(ctx {{ .P.Context }}.Context) ([]*{{ .C.Message.QualifiedKind }}, error) {
	filter := &{{ .C.Message.QualifiedDalKind }}Filter{}
	return x.find(ctx, "all", x.queryAll, filter)
}

// Filter implements {{ .C.Message.QualifiedDalKind }}Collection.Filter
func (x *{{ .C.Message.Name }}Collection) Filter(ctx {{ .P.Context }}.Context, arg *{{ .C.Message.QualifiedDalKind }}Filter) ([]*{{ .C.Message.QualifiedKind }}, error) {
	query := "SELECT {{ .V.QueryFields }} FROM " + x.config.TableName
	
	fields := []string{}
	{{ range $fn := .C.Fields.Names -}}
		{{- $f := ($C.Fields.ByName $fn) -}}
		{{- if not $f.Ignore -}}
			if arg.{{ ToTitleCase $fn }} != nil {
				fields = append(fields, "{{ $f.QueryName }} = :{{ $f.QueryName }}")
			}
		{{- end }}
	{{ end -}}

	if len(fields) > 0 {
		query = {{ .P.Fmt }}.Sprintf("%s WHERE %s", query, {{ .P.Strings }}.Join(fields, " AND "))
	}

	return x.find(ctx, "filter", query, arg)
}

func (x *{{ .C.Message.Name }}Collection) find(ctx {{ .P.Context }}.Context, label string, query string, arg *{{ .C.Message.QualifiedDalKind }}Filter) ([]*{{ .C.Message.QualifiedKind }}, error) {
	var err error
	start := {{ .P.Time }}.Now()
	{{ .P.Stats }}.Record(ctx, {{ ToCamelCase .C.Message.Name }}Inflight.M(1))
	defer func() {
		stop := {{ .P.Time }}.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)

		if err != nil {
			ctx, err = {{ .P.Tag }}.New(ctx,
				{{ .P.Tag }}.Upsert({{ ToCamelCase .C.Message.Name }}QueryError, label),
			)
		}

		ctx, err = {{ .P.Tag }}.New(ctx,
			{{ .P.Tag }}.Upsert({{ ToCamelCase .C.Message.Name }}QueryName, label),
		)

		{{ .P.Stats }}.Record(ctx, {{ ToCamelCase .C.Message.Name }}Latency.M(dur), {{ ToCamelCase .C.Message.Name }}Inflight.M(-1))
	}()

	filter := &{{ .C.Message.Name }}Filter{}
	filter.fromGeneric(arg)

	rows, err := x.db.QueryWithReplacements(ctx, query, filter)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	{{ .C.Message.Name }}s := []*{{ .C.Message.QualifiedKind }}{}
	for rows.Next() {
		obj := &{{ .C.Message.Name }}Scanner{}
		if err = rows.StructScan(obj); err != nil {
			return nil, err
		}
		{{ .C.Message.Name }}s = append({{ .C.Message.Name }}s, obj.{{ .C.Message.Name }}())
	}
	return {{ .C.Message.Name }}s, nil
}
`

	tmpl, err := template.New("definePostgresDefaultQueries").
		Funcs(template.FuncMap{
			"ToCamelCase": protocgenlib.ToCamelCase,
			"ToSnakeCase": protocgenlib.ToSnakeCase,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Context": c.File.QualifiedPackageName("context"),
		"Fmt":     c.File.QualifiedPackageName("fmt"),
		"SQL":     c.File.QualifiedPackageName("database/sql"),
		"Strings": c.File.QualifiedPackageName("strings"),
		"Time":    c.File.QualifiedPackageName("time"),
		"Stats":   c.File.QualifiedPackageName("go.opencensus.io/stats"),
		"Tag":     c.File.QualifiedPackageName("go.opencensus.io/tag"),
	}

	fields := []string{}
	queryFields := []string{}
	for _, n := range c.Fields.Names() {
		f := c.Fields.ByName(n)
		fields = append(fields, f.Name())
		queryFields = append(queryFields, f.QueryName())
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
		"V": map[string]string{
			"Fields":      strings.Join(fields, ", "),
			"QueryFields": strings.Join(queryFields, ", "),
		},
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineQueries() error {
	tmplSrc := `{{- $C := .C -}}
{{- $P := .P -}}
{{ range $qn := .C.Queries.Names }}
	{{- $q := ($C.Queries.ByName $qn) -}}
	{{- if $q.QueryProvided }}
		// {{ ToTitleCase $q.Name }} implements {{ $C.Message.QualifiedDalKind }}Collection.{{ ToTitleCase $q.Name }}
		func (x *{{ $C.Message.Name }}Collection){{ ToTitleCase $q.Name }}(ctx {{ $P.Context }}.Context
			{{- range $a := $q.Args -}}
				{{- $f := ($C.Fields.ByName $a) -}}
				, {{ ToSnakeCase $f.Name }} {{ $f.QualifiedKind }}
			{{- end -}}
		) ([]*{{ $C.Message.QualifiedKind }}, error) {
			filter := &{{ $C.Message.QualifiedDalKind }}Filter{
				{{- range $a := $q.Args -}}
				{{- $f := ($C.Fields.ByName $a) }}
					{{ ToTitleCase $f.Name }}: {{ $f.ToRef }}{{ ToSnakeCase $f.Name }},
				{{- end }}
			}
			return x.find(ctx, "{{ ToSnakeCase $q.Name }}", x.query{{ ToTitleCase $q.Name }}, filter)
		}
	{{- end -}}
{{- end }}
`

	tmpl, err := template.New("definePostgresQueries").
		Funcs(template.FuncMap{
			"ToSnakeCase": protocgenlib.ToSnakeCase,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
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

func (c *Collection) defineNewCollection() error {
	tmplSrc := `{{- $C := .C -}}
{{- $P := .P -}}
// New{{ .C.Message.Name }}Collection returns a new {{ .C.Message.Name }}Collection.
func New{{ .C.Message.Name }}Collection(db {{ .P.GenmsSQL }}.DB, queries {{ .C.Message.Name }}QueryTemplateProvider, config *{{ .C.Message.Name }}Config) (*{{ .C.Message.Name }}Collection, error) {
	register{{ ToTitleCase .C.Message.Name }}MetricsOnce.Do(register{{ ToTitleCase .C.Message.Name }}Metrics)

	coll := &{{ .C.Message.Name }}Collection{
		db: db,
		config: config,
	}

	queryReplacements := map[string]string{
		"table": config.TableName,
		"fields": "{{ .V.QueryFields }}",
		"writeFields": "{{ .V.WriteFields }}",
	}

	// generate Insert exec
	execInsert, err := {{ .P.Dal }}.RenderQuery("{{ .C.Message.QualifiedDalKind }}-Exec-Insert", queries.Insert(), queryReplacements)
	if err != nil {
		return nil, err
	}	
	coll.execInsert = execInsert

	// generate Upsert exec
	execUpsert, err := {{ .P.Dal }}.RenderQuery("{{ .C.Message.QualifiedDalKind }}-Exec-Upsert", queries.Upsert(), queryReplacements)
	if err != nil {
		return nil, err
	}	
	coll.execUpsert = execUpsert
	
	// generate Update exec
	execUpdate, err := {{ .P.Dal }}.RenderQuery("{{ .C.Message.QualifiedDalKind }}-Exec-Update", queries.Update(), queryReplacements)
	if err != nil {
		return nil, err
	}	
	coll.execUpdate = execUpdate

	// generate All query
	queryAll, err := {{ .P.Dal }}.RenderQuery("{{ .C.Message.QualifiedDalKind }}-Query-All", queries.All(), queryReplacements)
	if err != nil {
		return nil, err
	}	
	coll.queryAll = queryAll 

	{{ range $qn := .C.Queries.Names -}}
		{{- $q := ($C.Queries.ByName $qn) -}}
		{{- if $q.QueryProvided -}}
			// generate {{ ToTitleCase $qn }} query
			query{{ ToTitleCase $qn }}, err := {{ $P.Dal }}.RenderQuery("{{ $C.Message.QualifiedDalKind }}-Query-{{ ToTitleCase $qn }}", queries.{{ ToTitleCase $qn }}(), queryReplacements)
			if err != nil {
				return nil, err
			}	
			coll.query{{ ToTitleCase $qn }} = query{{ ToTitleCase $qn }}
		{{ end }}
	{{ end -}}

	return coll, nil
}
`

	tmpl, err := template.New("definePostgresNewCollection").
		Funcs(template.FuncMap{
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Collection": c.File.QualifiedPackageName(c.File.PackagePath()),
		"Dal":        c.File.QualifiedPackageName("github.com/rleszilm/genms/dal"),
		"GenmsSQL":   c.File.QualifiedPackageName("github.com/rleszilm/genms/sql"),
	}

	fields := []string{}
	queryFields := []string{}
	writeFields := []string{}
	for _, n := range c.Fields.Names() {
		f := c.Fields.ByName(n)
		fields = append(fields, f.Name())
		queryFields = append(queryFields, f.QueryName())
		writeFields = append(writeFields, ":"+f.QueryName())
	}

	buf := &bytes.Buffer{}
	if err := tmpl.Execute(buf, map[string]interface{}{
		"C": c,
		"P": p,
		"V": map[string]string{
			"Fields":      strings.Join(fields, ", "),
			"QueryFields": strings.Join(queryFields, ", "),
			"WriteFields": strings.Join(writeFields, ", "),
		},
	}); err != nil {
		return err
	}

	if _, err := c.File.Write(buf.Bytes()); err != nil {
		return err
	}
	return nil
}

func (c *Collection) defineInternalStructs() error {
	tmplSrc := `{{- $C := .C -}}
// {{ .C.Message.Name }}Filter is an autogenerated struct that
// is used in generic {{ .C.Message.Name }} queries.
type {{ .C.Message.Name }}Filter struct {
	{{ range $n := .C.Fields.Names -}}
		{{- $f := ($C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			{{ ToTitleCase $f.Name }} {{ AsPointer $f.QualifiedKind }} ` + "`" + `db:"{{ $f.QueryName }}"` + "`" + `
		{{- end }}
	{{ end -}}
}

func (x *{{ .C.Message.Name }}Filter) fromGeneric(y *{{ .C.Message.QualifiedDalKind }}Filter) {
	{{ range $n := .C.Fields.Names -}}
		{{- $f := ($C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			if y.{{ ToTitleCase $f.Name }} != nil {
				x.{{ ToTitleCase $f.Name }} = y.{{ ToTitleCase $f.Name }}
			}
		{{- end }}
	{{ end -}}
}

// {{ .C.Message.Name }}Scanner is an autogenerated struct that
// is used to parse query results.
type {{ .C.Message.Name }}Scanner struct {
	{{ range $n := .C.Fields.Names -}}
		{{- $f := ($C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			{{ ToTitleCase $f.Name }} {{ $f.SQLNilKind }} ` + "`" + `db:"{{ $f.QueryName }}"` + "`" + `
		{{- end }}
	{{ end -}}
}

// {{ .C.Message.Name }} returns a new {{ .C.Message.QualifiedKind }} populated with scanned values.
func (x *{{ .C.Message.Name }}Scanner) {{ .C.Message.Name }}() *{{ .C.Message.QualifiedKind }} {
	y := &{{ .C.Message.QualifiedKind }}{}

	{{ range $n := .C.Fields.Names -}}
		{{- $f := ($C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			{{- if $f.HasSQLNil -}}
				if x.{{ ToTitleCase $f.Name }}.Valid {
					y.{{ ToTitleCase $f.Name }} = {{ $f.ValueFromSQLValue "x" }}
				}
			{{- else -}}
				y.{{ ToTitleCase $f.Name }} = {{ $f.ValueFromSQLValue "x" }}
			{{- end -}}
		{{- end }}
	{{ end -}}
	return y
}

// {{ .C.Message.Name }}Writer is an autogenerated struct that is used to supply values to write queries.
type {{ .C.Message.Name }}Writer struct {
	{{ range $n := .C.Fields.Names -}}
		{{- $f := ($C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			{{ ToTitleCase $f.Name }} {{ $f.QualifiedKind }} ` + "`" + `db:"{{ $f.QueryName }}"` + "`" + `
		{{- end }}
	{{ end -}}
}

// From{{ .C.Message.Name }} populates the struct with values from the base type.
func (x *{{ .C.Message.Name }}Writer) From{{ .C.Message.Name }}(y *{{ .C.Message.QualifiedKind }}) {
	{{ range $n := .C.Fields.Names -}}
		{{- $f := ($C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			x.{{ ToTitleCase $f.Name }} = y.{{ ToTitleCase $f.Name }}
		{{- end }}
	{{ end -}}
}
`

	tmpl, err := template.New("definePostgresStructs").
		Funcs(template.FuncMap{
			"AsPointer":   protocgenlib.AsPointer,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Collection": c.File.QualifiedPackageName(c.File.DalPackagePath()),
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
	tmplSrc := `// {{ .C.Message.Name }}Config is a struct that can be used to configure a {{ .C.Message.Name }}Collection
type {{ .C.Message.Name }}Config struct {
	TableName string ` + "`" + `envconfig:"table"` + "`" + `
}
`

	tmpl, err := template.New("definePostgresConfig").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{}

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

func (c *Collection) defineTemplateProvider() error {
	tmplSrc := `{{- $C := .C -}}
{{- $P := .P -}}
// {{ .C.Message.Name }}QueryTemplateProvider is an interface that returns the query templated that should be executed
// to generate the queries that the collection will use.
//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 . {{ .C.Message.Name }}QueryTemplateProvider
type {{ .C.Message.Name }}QueryTemplateProvider interface {
	Insert() string
	Upsert() string
	All() string
	{{ range $qn := .C.Queries.Names -}}
		{{- $q := ($C.Queries.ByName $qn) -}}
		{{- if $q.QueryProvided -}}
			{{ ToTitleCase $qn }}() string
		{{ end -}}
	{{ end -}}
}

// {{ .C.Message.Name }}Queries provides auto-generated queries when possible. This is not gauranteed to be a complete
// implementation of the interface. This should be used as a base for the actual query provider used.
type {{ .C.Message.Name }}Queries struct {
}

// Insert implements {{ .C.Message.Name }}QueryTemplateProvider.Insert.
func (x *{{ .C.Message.Name }}Queries) Insert() string {
	return ` + "`" + `INSERT INTO {{ "{{ .table }}" }}({{ "{{ .fields }}" }}) VALUES({{ "{{ .writeFields }}" }});` + "`" + `
}

// Upsert implements {{ .C.Message.Name }}QueryTemplateProvider.Upsert.
func (x *{{ .C.Message.Name }}Queries) Upsert() string {
	return ` + "`" + `INSERT INTO {{ "{{ .table }}" }}({{ "{{ .fields }}" }}) VALUES({{ "{{ .writeFields }}" }});` + "`" + `
}

// All implements {{ .C.Message.Name }}QueryTemplateProvider.All.
func (x *{{ .C.Message.Name }}Queries) All() string {
	return ` + "`" + `SELECT {{ "{{ .fields }}" }} FROM {{ "{{ .table }}" }};` + "`" + `
}

{{ range $qn := .C.Queries.Names -}}
	{{- $q := ($C.Queries.ByName $qn) -}}
	{{- if $q.QueryImplemented -}}
		// {{- ToTitleCase $q.Name -}} implements {{ $C.Message.Name }}QueryTemplateProvider.{{- ToTitleCase $q.Name -}}.
		func (x *{{ $C.Message.Name }}Queries) {{- ToTitleCase $q.Name -}}() string {
			return ` + "`" + `SELECT {{ "{{ .fields }}" }} FROM {{ "{{ .table }}" }} WHERE
			1 = 1
			{{- range $a := $q.Args -}}
				{{- $f := $C.Fields.ByName $a -}}	
				{{- "" }} AND
				{{ $f.QueryName }} = :{{ $f.QueryName }}
			{{- end -}};` + "`" + `
		}
	{{ end -}}
{{ end -}}
`

	tmpl, err := template.New("definePostgresTemplateProvider").
		Funcs(template.FuncMap{
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{}

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

func (c *Collection) defineMetrics() error {
	tmplSrc := `// define metrics
var (
	{{ ToCamelCase .C.Message.Name }}QueryName = {{ .P.Tag }}.MustNewKey("dal_postgres_{{ ToSnakeCase .C.Message.Name }}")
	{{ ToCamelCase .C.Message.Name }}QueryError = {{ .P.Tag }}.MustNewKey("dal_postgres_{{ ToSnakeCase .C.Message.Name }}_error")

	{{ ToCamelCase .C.Message.Name }}Latency = {{ .P.Stats }}.Float64("{{ ToSnakeCase .C.Message.Name }}_latency", "Latency of {{ .C.Message.Name }} queries", {{ .P.Stats }}.UnitMilliseconds)
	{{ ToCamelCase .C.Message.Name }}Inflight = {{ .P.Stats }}.Int64("{{ ToSnakeCase .C.Message.Name }}_inflight", "Count of {{ .C.Message.Name }} queries in flight", {{ .P.Stats }}.UnitDimensionless)

	register{{ ToTitleCase .C.Message.Name }}MetricsOnce {{ .P.Sync }}.Once
)

func register{{ ToTitleCase .C.Message.Name }}Metrics() {
	views := []*{{ .P.View }}.View{
		{
			Name:        "dal_postgres_{{ ToSnakeCase .C.Message.Name }}_latency",
			Measure:     {{ ToCamelCase .C.Message.Name }}Latency,
			Description: "The distribution of the query latencies",
			TagKeys:     []{{ .P.Tag }}.Key{ {{ ToCamelCase .C.Message.Name }}QueryName, {{ ToCamelCase .C.Message.Name }}QueryError},
			Aggregation: {{ .P.View }}.Distribution(0, 25, 100, 200, 400, 800, 10000),
		},
		{
			Name:        "dal_postgres_{{ ToSnakeCase .C.Message.Name }}_inflight",
			Measure:     {{ ToCamelCase .C.Message.Name }}Inflight,
			Description: "The number of queries being processed",
			TagKeys:     []{{ .P.Tag }}.Key{ {{ ToCamelCase .C.Message.Name }}QueryName},
			Aggregation: {{ .P.View }}.Sum(),
		},
	}

	if err := {{ .P.View }}.Register(views...); err != nil {
		{{ .P.Log }}.Fatal("Cannot register metrics:", err)
	}
}
`

	tmpl, err := template.New("definePostgresMetrics").
		Funcs(template.FuncMap{
			"ToCamelCase": protocgenlib.ToCamelCase,
			"ToSnakeCase": protocgenlib.ToSnakeCase,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Log":   c.File.QualifiedPackageName("log"),
		"Sync":  c.File.QualifiedPackageName("sync"),
		"Tag":   c.File.QualifiedPackageName("go.opencensus.io/tag"),
		"Stats": c.File.QualifiedPackageName("go.opencensus.io/stats"),
		"View":  c.File.QualifiedPackageName("go.opencensus.io/stats/view"),
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
