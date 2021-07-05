package cache

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/go-test/deep"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/annotations"
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	"golang.org/x/tools/imports"
	"google.golang.org/protobuf/compiler/protogen"
)

var (
	generatedCommon bool
)

// Cache is a struct that generates an base cache file.
type Cache struct {
	File    *File
	Message *generator.Message
	Fields  *generator.Fields
	Queries *generator.Queries
	Opts    *annotations.DalOptions

	plugin   *protogen.Plugin
	filename string
}

// NewCache returns a new updater renderer.
func NewCache(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) *Cache {
	base := path.Base(file.GeneratedFilenamePrefix)
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, fmt.Sprintf("dal/cache/%s.genms.cache.%s.go", base, strings.ToLower(msg.GoIdent.GoName)))
	outfile := plugin.NewGeneratedFile(filename, ".")

	cfile := NewFile(outfile, file)
	cmsg := generator.NewMessage(cfile.Generator(), msg)
	cfields := generator.NewFields(cmsg)
	cqueries := generator.NewQueries(cfile.Generator(), cfields, opts)

	return &Cache{
		File:     cfile,
		Message:  cmsg,
		Fields:   cfields,
		Queries:  cqueries,
		Opts:     opts,
		plugin:   plugin,
		filename: filename,
	}
}

// GenerateCache generates the updater for the collection.
func GenerateCache(plugin *protogen.Plugin, file *protogen.File, msg *protogen.Message, opts *annotations.DalOptions) error {
	c := NewCache(plugin, file, msg, opts)
	return c.render()
}

func (c *Cache) render() error {
	if !generatedCommon {
		if err := c.defineCommon(); err != nil {
			return err
		}
		generatedCommon = true
	}

	steps := []func() error{
		c.definePackage,
		c.defineCache,
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

func (c *Cache) defineCommon() error {
	file := c.File.Proto()
	plugin := c.plugin
	dir := path.Dir(file.GeneratedFilenamePrefix)
	filename := path.Join(dir, "dal/cache/cache.genms.go")
	outfile := plugin.NewGeneratedFile(filename, ".")

	tmplSrc := `// Package {{ $.C.File.CachePackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ $.C.File.CachePackageName }}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

var (
	logs = {{ $.P.Cache }}.Logs()
)

`

	tmpl, err := template.New("defineCommon").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	p := map[string]string{
		"Cache": c.File.QualifiedPackageName("github.com/rleszilm/genms/cache"),
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

func (c *Cache) definePackage() error {
	tmplSrc := `// Package {{ .File.CachePackageName }} is generated by protoc-gen-go-genms-dal. *DO NOT EDIT*
package {{ .File.CachePackageName }}

`

	tmpl, err := template.New("defineCachePackage").
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

func (c *Cache) defineCache() error {
	tmplSrc := `{{- $C := .C -}}
{{- $P := .P -}}

// {{ .C.Message.Name }}Key defines a Key in the cache.
type {{ .C.Message.Name }}Key interface{}

// {{ .C.Message.Name }}Reader is defines the interface for getting values from a cache.
//counterfeiter:generate . {{ .C.Message.Name }}Reader
type {{ .C.Message.Name }}Reader interface {
	Get(ctx {{ .P.Context }}.Context, key {{ .C.Message.Name }}Key) (*{{ .C.Message.QualifiedKind }}, error)
}

// {{ .C.Message.Name }}ReadeAller is defines the interface for getting all values from a cache.
//counterfeiter:generate . {{ .C.Message.Name }}ReadAller
type {{ .C.Message.Name }}ReadAller interface {
	All(ctx {{ .P.Context }}.Context) ([]*{{ .C.Message.QualifiedKind }}, error)
}

// {{ .C.Message.Name }}Writer is defines the interface for setting values in a cache.
//counterfeiter:generate . {{ .C.Message.Name }}Writer
type {{ .C.Message.Name }}Writer interface {
	Set(ctx {{ .P.Context }}.Context, key {{ .C.Message.Name }}Key, obj *{{ .C.Message.QualifiedKind }}) (*{{ .C.Message.QualifiedKind }}, error)
}

// {{ .C.Message.Name }}ReadWriter is defines the interface for setting values in a cache.
//counterfeiter:generate . {{ .C.Message.Name }}ReadWriter
type {{ .C.Message.Name }}ReadWriter interface {
	{{ .C.Message.Name }}Reader
	{{ .C.Message.Name }}Writer
}

// {{ .C.Message.Name }}KeyFunc is a function that generates a unique deterministic key for the {{ .C.Message.QualifiedKind }}.
type {{ .C.Message.Name }}KeyFunc func(*{{ .C.Message.QualifiedKind }}) interface{}

// Unimplemented{{ .C.Message.Name }}Cache is a KV ReadWriter that takes no action on read or write.
type Unimplemented{{ .C.Message.Name }}Cache struct{
}

// GetAll implements {{ .C.Message.Name }}ReadAller.
func (x *Unimplemented{{ .C.Message.Name }}Cache) All(_ {{ .P.Context }}.Context) (*{{ .C.Message.QualifiedKind }}, error) {
	return nil, {{ .P.Cache }}.ErrUnimplemented
}


// Get implements {{ .C.Message.Name }}Reader.
func (x *Unimplemented{{ .C.Message.Name }}Cache) Get(_ {{ .P.Context }}.Context, _ {{ .C.Message.Name }}Key) (*{{ .C.Message.QualifiedKind }}, error) {
	return nil, {{ .P.Cache }}.ErrUnimplemented
}

// Set implements {{ .C.Message.Name }}Writer.
func (x *Unimplemented{{ .C.Message.Name }}Cache) Set(_ {{ .P.Context }}.Context, _ {{ .C.Message.Name }}Key, _ *{{ .C.Message.QualifiedKind }}) (*{{ .C.Message.QualifiedKind }}, error) {
	return nil, {{ .P.Cache }}.ErrUnimplemented
}

`

	tmpl, err := template.New("defineCache").
		Funcs(template.FuncMap{}).
		Parse(tmplSrc)

	if err != nil {
		return err
	}

	p := map[string]string{
		"Cache":   c.File.QualifiedPackageName("github.com/rleszilm/genms/cache"),
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
