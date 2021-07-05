package rest

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
	filename := path.Join(dir, fmt.Sprintf("dal/rest/%s.genms.dal.%s.go", base, strings.ToLower(msg.GoIdent.GoName)))
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
		c.defineCollection,
		c.defineService,
		c.defineDefaultQueries,
		c.defineQueries,
		c.defineNewCollection,
		c.defineInternalStructs,
		c.defineConfig,
		c.defineTemplateProvider,
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
	filename := path.Join(dir, fmt.Sprintf("dal/rest/rest.genms.go"))
	outfile := plugin.NewGeneratedFile(filename, ".")

	tmplSrc := `// Package {{ $.C.File.RestPackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ $.C.File.RestPackageName }}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

var (
	logs = {{ $.P.Rest }}.Logs()
)

`

	tmpl, err := template.New("defineCommon").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	p := map[string]string{
		"Rest": c.File.QualifiedPackageName("github.com/rleszilm/genms/rest"),
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
	tmplSrc := `// Package {{ .File.RestPackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ .File.RestPackageName }}
`

	tmpl, err := template.New("defineRestPackage").
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
	tmplSrc := `// {{ .C.Message.Name }}Collection is an autogenerated implementation of {{ .C.Message.QualifiedDalKind }}Collection.
type {{ .C.Message.Name }}Collection struct {
	{{ .P.Collection }}.Unimplemented{{ .C.Message.Name }}Collection

	name string

	client *{{ .P.HTTP }}.Client
	config *{{ .C.Message.Name }}Config

	url *{{ .P.URL }}.URL
	urlAll string
	{{ range $qn := .C.Queries.Names -}}
		{{- $q := $.C.Queries.ByName $qn -}}
		{{- if $q.QueryProvided -}}
			urlTmpl{{ ToTitleCase $q.Name }} *{{ $.P.Text }}.Template
		{{- end }}
	{{ end -}}
}
`

	tmpl, err := template.New("defineRestCollection").
		Funcs(template.FuncMap{
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Collection": c.File.QualifiedPackageName(c.File.DalPackagePath()),
		"HTTP":       c.File.QualifiedPackageName("net/http"),
		"Text":       c.File.QualifiedPackageName("text/template"),
		"URL":        c.File.QualifiedPackageName("net/url"),
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

// String returns the name of the Collection.
func (x *{{ .C.Message.Name }}Collection) String() string {
	{{- $pkg := .C.File.RestPackageName -}}
	if x.name != "" {
		return "{{ ToDashCase $pkg }}-{{ ToDashCase .C.Message.Name }}-" + x.name
	}
	return "{{ ToDashCase $pkg }}-{{ ToDashCase .C.Message.Name }}"
}

// NameOf returns the name of the Collection.
func (x *{{ .C.Message.Name }}Collection) NameOf() string {
	return x.String()
}
`

	tmpl, err := template.New("defineRestService").
		Funcs(template.FuncMap{
			"ToDashCase": protocgenlib.ToDashCase,
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

func (c *Collection) defineDefaultQueries() error {
	tmplSrc := `// DoReq executes the given http request.
func (x *{{ .C.Message.Name }}Collection) DoReq(ctx {{ .P.Context }}.Context, label string, req *{{ .P.HTTP }}.Request) ([]*{{ .C.Message.QualifiedKind }}, error) {
	var err error
	var resp *{{ .P.HTTP }}.Response

	start := {{ .P.Time }}.Now()
	ctx, _ = {{ .P.Tag }}.New(ctx,
		{{ .P.Tag }}.Upsert({{ .P.Rest }}.TagCollection, "{{ ToSnakeCase .C.Message.Name }}"),
		{{ .P.Tag }}.Upsert({{ .P.Rest }}.TagInstance, x.name),
		{{ .P.Tag }}.Upsert({{ .P.Rest }}.TagMethod, label),
		{{ .P.Tag }}.Upsert({{ .P.Rest }}.TagRestMethod, req.Method),
	)
	{{ .P.Stats }}.Record(ctx, {{ .P.Rest }}.MeasureInflight.M(1))
	defer func(ctx {{ .P.Context }}.Context) {
		if resp != nil {
			ctx, _ = {{ .P.Tag }}.New(ctx,
				{{ .P.Tag }}.Upsert({{ .P.Rest }}.TagResponseCode, {{ .P.Strconv }}.Itoa(resp.StatusCode)),
			)
		}

		stop := time.Now()
		dur := float64(stop.Sub(start).Nanoseconds()) / float64({{ .P.Time }}.Millisecond)
		{{ .P.Stats }}.Record(ctx, {{ .P.Rest }}.MeasureLatency.M(dur), {{ .P.Rest }}.MeasureInflight.M(-1))
	}(ctx)

	ctx, cancel := {{ .P.Context }}.WithTimeout(ctx, x.config.Timeout)
	defer cancel()

	resp, err = x.client.Do(req.WithContext(ctx))
	if err != nil {
		{{ .P.Rest }}.Logs().Error("could not execute rest request:", err)
		{{ .P.Stats }}.Record(ctx, {{ .P.Rest }}.MeasureError.M(1))
		return nil, err
	}

	buff, err := {{ .P.IOUtil }}.ReadAll(resp.Body)
	if err != nil {
		{{ .P.Rest }}.Logs().Error("could not read rest response:", err)
		{{ .P.Stats }}.Record(ctx, {{ .P.Rest }}.MeasureError.M(1))
		return nil, err
	}

	{{ .C.Message.Name }}Scanners := []*{{ ToTitleCase .C.Message.Name }}Scanner{}
	if err := {{ .P.JSON }}.Unmarshal(buff, &{{ .C.Message.Name }}Scanners); err != nil {
		{{ .P.Rest }}.Logs().Error("could not unmarshal rest response:", err)
		{{ .P.Stats }}.Record(ctx, {{ .P.Rest }}.MeasureError.M(1))
		return nil, err
	}

	{{ .C.Message.Name }}s := []*{{ .C.Message.QualifiedKind }}{}
	for _, c := range {{ .C.Message.Name }}Scanners {
		{{ .C.Message.Name }}s = append({{ .C.Message.Name }}s, c.{{ .C.Message.Name }}())
	}
	return {{ .C.Message.Name }}s, nil
}

`

	tmpl, err := template.New("defineRestDefaultQueries").
		Funcs(template.FuncMap{
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
		"Context": c.File.QualifiedPackageName("context"),
		"HTTP":    c.File.QualifiedPackageName("net/http"),
		"IOUtil":  c.File.QualifiedPackageName("io/ioutil"),
		"JSON":    c.File.QualifiedPackageName("encoding/json"),
		"Rest":    c.File.QualifiedPackageName("github.com/rleszilm/genms/rest"),
		"Stats":   c.File.QualifiedPackageName("go.opencensus.io/stats"),
		"Strconv": c.File.QualifiedPackageName("strconv"),
		"Tag":     c.File.QualifiedPackageName("go.opencensus.io/tag"),
		"Time":    c.File.QualifiedPackageName("time"),
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
	tmplSrc := `// All implements {{ $.C.Message.QualifiedDalKind }}Collection.All
func (x *{{ $.C.Message.Name }}Collection) All(ctx {{ $.P.Context }}.Context) ([]*{{ $.C.Message.QualifiedKind }}, error) {
	u := &{{ $.P.URL }}.URL{}
	{{ $.P.Copier }}.Copy(u, x.url)
	u.Path = x.urlAll

	req := &{{ $.P.HTTP }}.Request{
		Method: "GET",
		Header: {{ $.P.HTTP }}.Header{},
		URL: u,
	}

	for k, v := range x.config.Headers {
		req.Header.Add(k, v)
	}
	
	return x.DoReq(ctx, "all", req)
}

{{ range $qn := .C.Queries.Names }}
	{{- $q := ($.C.Queries.ByName $qn) -}}
	{{- if $q.QueryProvided }}
		// {{ ToTitleCase $q.Name }} implements {{ $.C.Message.QualifiedDalKind }}Collection.{{ ToTitleCase $q.Name }}
		func (x *{{ $.C.Message.Name }}Collection){{ ToTitleCase $q.Name }}(ctx {{ $.P.Context }}.Context
			{{- range $a1 := $q.Args -}}
				{{- $arg := (Arg $.C.File $.C.Fields $a1) -}}
				, {{ ToSnakeCase $arg.Name }} {{ $arg.QualifiedKind }}
			{{- end -}}
		) ([]*{{ $.C.Message.QualifiedKind }}, error) {
			u := &{{ $.P.URL }}.URL{}
			{{ $.P.Copier }}.Copy(u, x.url)

			req := &{{ $.P.HTTP }}.Request{
				Method: "{{ $q.Method }}",
				Header: {{ $.P.HTTP }}.Header{},
				URL: u,
			}

			queryValues := {{ $.P.URL }}.Values{}
			{{ range $a2 := $q.Args }}
				{{- $arg := (Arg $.C.File $.C.Fields $a2) -}}
				{{- if $arg.IsQuery -}}
					queryValues.Add("{{ $arg.QueryName }}", {{ $.P.Fmt }}.Sprintf("%v", {{ ToSnakeCase $arg.Name }}))
				{{ end }}
			{{- end }}
			req.URL.RawQuery = queryValues.Encode()

			pathValues := map[string]interface{}{
				{{- range $a3 := $q.Args -}}
					{{- $arg := (Arg $.C.File $.C.Fields $a3) -}}
					{{- if $arg.IsPath -}}
						"{{ $arg.Name }}": {{ $arg.ToRef }}{{ ToSnakeCase $arg.Name }},
					{{- end }}
				{{- end }}
			}
			pathBuf := &{{ $.P.Bytes }}.Buffer{}
			if err := x.urlTmpl{{ ToTitleCase $q.Name }}.Execute(pathBuf, pathValues); err != nil {
				return nil, err
			}
			req.URL.Path = pathBuf.String()
			
			{{ if ne "GET" ($q.Method) -}}
				bodyValues := map[string]interface{}{
					{{- range $a4 := $q.Args -}}
						{{- $arg := (Arg $.C.File $.C.Fields $a4) -}}
						{{- if $arg.IsBody -}}
							"{{ $arg.QueryName }}": {{ $arg.ToRef }}{{ ToSnakeCase $arg.Name }},
						{{- end }}
					{{- end }}
				}
				bodyBytes, err := {{ $.P.JSON }}.Marshal(bodyValues)
				if err != nil {
					return nil, err
				}
				bodyRC := {{ $.P.IOUtil }}.NopCloser({{ $.P.Bytes }}.NewReader(bodyBytes))
				req.Body = bodyRC
			{{- end }}

			for k, v := range x.config.Headers {
				req.Header.Add(k, v)
			}
			{{ range $a5 := $q.Args }}
				{{- $arg := (Arg $.C.File $.C.Fields $a5) -}}
				{{- if $arg.IsHeader -}}
					req.Header.Add("{{ $arg.QueryName }}", {{ $.P.Fmt }}.Sprintf("%v", {{ ToSnakeCase $arg.Name }}))
				{{- end }}
			{{ end }}
						
			return x.DoReq(ctx, "{{ ToSnakeCase $q.Name }}", req)
		}
	{{- end -}}
{{- end }}
`

	tmpl, err := template.New("defineRestQueries").
		Funcs(template.FuncMap{
			"Arg":         NewArg,
			"ToSnakeCase": protocgenlib.ToSnakeCase,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Bytes":   c.File.QualifiedPackageName("bytes"),
		"Context": c.File.QualifiedPackageName("context"),
		"Copier":  c.File.QualifiedPackageName("github.com/jinzhu/copier"),
		"Fmt":     c.File.QualifiedPackageName("fmt"),
		"HTTP":    c.File.QualifiedPackageName("net/http"),
		"IOUtil":  c.File.QualifiedPackageName("io/ioutil"),
		"JSON":    c.File.QualifiedPackageName("encoding/json"),
		"URL":     c.File.QualifiedPackageName("net/url"),
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
	tmplSrc := `// New{{ .C.Message.Name }}Collection returns a new {{ .C.Message.Name }}Collection.
func New{{ .C.Message.Name }}Collection(instance string, client *{{ .P.HTTP }}.Client, urls {{ .C.Message.Name }}UrlTemplateProvider, config *{{ .C.Message.Name }}Config) (*{{ .C.Message.Name }}Collection, error) {
	coll := &{{ .C.Message.Name }}Collection{
		name: instance,
		client: client,
		config: config,
	}

	u, err := {{ .P.URL }}.Parse(config.URL)
	if err != nil {
		return nil, err
	}
	coll.url = u

	coll.urlAll = urls.All()
	{{ range $qn := .C.Queries.Names -}}
		{{- $q := ($.C.Queries.ByName $qn) -}}
		{{- if $q.QueryProvided -}}
			if urls.{{ ToTitleCase $qn }}() != "" {
				urlTmpl{{ ToTitleCase $qn }}, err := {{ $.P.Template }}.New("urlTmpl{{ ToTitleCase $qn }}").
					Funcs({{ $.P.Template }}.FuncMap{}).
					Parse(urls.{{ ToTitleCase $qn }}())
				if err != nil {
					return nil, err
				}
				coll.urlTmpl{{ ToTitleCase $qn }} = urlTmpl{{ ToTitleCase $qn }}
			}
		{{ end }}
	{{ end -}}

	return coll, nil
}
`

	tmpl, err := template.New("defineRestNewCollection").
		Funcs(template.FuncMap{
			"ToSnakeCase": protocgenlib.ToSnakeCase,
			"ToTitleCase": protocgenlib.ToTitleCase,
		}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Collection": c.File.QualifiedPackageName(c.File.PackagePath()),
		"HTTP":       c.File.QualifiedPackageName("net/http"),
		"Template":   c.File.QualifiedPackageName("text/template"),
		"URL":        c.File.QualifiedPackageName("net/url"),
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
	tmplSrc := `// {{ .C.Message.Name }}Scanner is an autogenerated struct that
// is used to parse query results.
type {{ .C.Message.Name }}Scanner struct {
	{{ range $n := .C.Fields.Names -}}
		{{- $f := ($.C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			{{ ToTitleCase $f.Name }} {{ $f.QualifiedKind }} ` + "`" + `json:"{{ $f.QueryName }}"` + "`" + `
		{{- end }}
	{{ end -}}
}

// {{ .C.Message.Name }} returns a new {{ .C.Message.QualifiedKind }} populated with scanned values.
func (x *{{ .C.Message.Name }}Scanner) {{ .C.Message.Name }}() *{{ .C.Message.QualifiedKind }} {
	y := &{{ .C.Message.QualifiedKind }}{}

	{{ range $n := .C.Fields.Names -}}
		{{- $f := ($.C.Fields.ByName $n) -}}
		{{- if not $f.Ignore -}}
			y.{{ ToTitleCase $f.Name }} = x.{{ ToTitleCase $f.Name }}
		{{- end }}
	{{ end -}}
	return y
}

`

	tmpl, err := template.New("defineRestStructs").
		Funcs(template.FuncMap{
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
	URL string ` + "`" + `envconfig:"url"` + "`" + `
	Name string ` + "`" + `envconfig:"name"` + "`" + `
	Timeout {{ .P.Time }}.Duration ` + "`" + `envconfig:"timeout" default:"5s"` + "`" + `
	Headers map[string]string ` + "`" + `envconfig:"headers"` + "`" + `
}
`

	tmpl, err := template.New("defineRestConfig").
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

func (c *Collection) defineTemplateProvider() error {
	tmplSrc := `// {{ .C.Message.Name }}UrlTemplateProvider is an interface that returns the query templated that should be executed
// to generate the queries that the collection will use.
//counterfeiter:generate .  {{ .C.Message.Name }}UrlTemplateProvider
type {{ .C.Message.Name }}UrlTemplateProvider interface {
	All() string
	{{ range $qn := .C.Queries.Names -}}
		{{- $q := ($.C.Queries.ByName $qn) -}}
		{{- if $q.QueryProvided -}}
			{{ ToTitleCase $qn }}() string
		{{ end -}}
	{{ end -}}
}

`

	tmpl, err := template.New("defineRestTemplateProvider").
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

// GenerateCollection generates the dal interface for the collection
func GenerateCollection(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	c := NewCollection(plugin, file, msg, opts)
	return c.render()
}
