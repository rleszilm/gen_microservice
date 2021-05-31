package mongo

import (
	"github.com/rleszilm/genms/cmd/protoc-gen-go-genms-dal/generator"
	protocgenlib "github.com/rleszilm/genms/internal/protoc-gen-lib"
)

// Fields is a struct that contains data about the messages fields.
type Fields struct {
	*generator.Fields
}

// NewFields returns a new Fields
func NewFields(msg *Message) *Fields {
	return AsFields(protocgenlib.NewFields(msg.ProtocGenLib()))
}

// AsFields wraps Fields.
func AsFields(fields *protocgenlib.Fields) *Fields {
	return &Fields{
		Fields: generator.AsFields(fields),
	}
}

// ByName returns the specified field.
func (f *Fields) ByName(n string) *Field {
	return AsField(f.Fields.ByName(n).ProtocGenLib())
}
