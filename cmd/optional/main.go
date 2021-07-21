// Optional is a tool that generates 'optional' type wrappers around a given type T.
//
// Typically this process would be run using go generate, like this:
//
//	//go:generate optional -type=Foo
//
// running this command
//
//	optional -type=Foo
//
// in the same directory will create the file optional_foo.go
// containing a definition of
//
//	type OptionalFoo struct {
//		...
//	}
//
// The default type is OptionalT or optionalT (depending on if the type is exported)
// and output file is optional_t.go. This can be overridden with the -output flag.
//
package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/build"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"time"
)

type generator struct {
	packageName    string
	outputName     string
	typeName       string
	typePackage    string
	WithJSONEncode bool
}

func (g *generator) generate() ([]byte, error) {

	var (
		t = template.Must(template.New("").Parse(tmpl))

		data = struct {
			Timestamp      time.Time
			PackageName    string
			TypeName       string
			OutputName     string
			VariableName   string
			TypePackage    string
			WithJSONEncode bool
		}{
			time.Now().UTC(),
			g.packageName,
			g.typeName,
			g.outputName,
			strings.ToLower(string(g.outputName[0])),
			g.typePackage,
			g.WithJSONEncode,
		}

		buf bytes.Buffer
		err = t.Execute(&buf, data)
	)

	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func main() {
	log.SetFlags(0)
	log.SetPrefix("optional: ")

	var (
		typeName     = flag.String("type", "", "type name; must be set")
		outputName   = flag.String("output", "", "output type and file name; default [o|O]ptional<type> and srcdir/optional_<type>.go")
		typePackage  = flag.String("type-package", "", "package in which the type is defined")
		noJSONEncode = flag.Bool("no-json", false, "do not provide json (un)marshall functions")
	)

	flag.Parse()

	if len(*typeName) == 0 {
		flag.Usage()
		os.Exit(2)
	}

	pkg, err := build.Default.ImportDir(".", 0)
	if err != nil {
		log.Fatal(err)
	}

	var (
		filename string
		g        generator
	)

	g.typeName = *typeName
	g.packageName = pkg.Name
	g.typePackage = *typePackage
	g.WithJSONEncode = !*noJSONEncode

	if len(*outputName) == 0 {
		// no output specified, use default optional_<type>

		// TODO: may not be the most reliable method
		exported := strings.Title(g.typeName) == g.typeName

		if exported {
			g.outputName = "Optional" + strings.Title(g.typeName)
		} else {
			g.outputName = "optional" + strings.Title(g.typeName)
		}
		filename = fmt.Sprintf("optional_%s.go", strings.ToLower(g.typeName))
	} else {
		g.outputName = *outputName
		filename = strings.ToLower(g.outputName + ".go")
	}

	src, err := g.generate()
	if err != nil {
		log.Fatal(err)
	}

	if err = ioutil.WriteFile(filename, src, 0644); err != nil {
		log.Fatalf("writing output: %s", err)
	}
}

const tmpl = `// Code generated by go generate
// This file was generated by robots at {{ .Timestamp }}

package {{ .PackageName }}

import (
{{- if .WithJSONEncode }}
	"encoding/json"
{{- end }}
{{- if .TypePackage }}
	{{ .TypePackage }}
{{- end }}
	"errors"
)

// {{ .OutputName }} is an optional {{ .TypeName }}.
type {{ .OutputName }} struct {
	value *{{ .TypeName }}
}

// New{{ .OutputName }} creates an optional.{{ .OutputName }} from a {{ .TypeName }}.
func New{{ .OutputName }}(v {{ .TypeName }}) {{ .OutputName }} {
	return {{ .OutputName }}{&v}
}

// Set sets the {{ .TypeName }} value.
func ({{ .VariableName }} *{{ .OutputName }}) Set(v {{ .TypeName }}) {
	{{ .VariableName }}.value = &v
}

// Get returns the {{ .TypeName }} value or an error if not present.
func ({{ .VariableName }} {{ .OutputName }}) Get() ({{ .TypeName }}, error) {
	if !{{ .VariableName }}.Present() {
		var zero {{ .TypeName }}
		return zero, errors.New("value not present")
	}
	return *{{ .VariableName }}.value, nil
}

// MustGet returns the {{ .TypeName }} value or panics if not present.
func ({{ .VariableName }} {{ .OutputName }}) MustGet() {{ .TypeName }} {
	if !{{ .VariableName }}.Present() {
		panic("value not present")
	}
	return *{{ .VariableName }}.value
}

// Present returns whether or not the value is present.
func ({{ .VariableName }} {{ .OutputName }}) Present() bool {
	return {{ .VariableName }}.value != nil
}

// OrElse returns the {{ .TypeName }} value or a default value if the value is not present.
func ({{ .VariableName }} {{ .OutputName }}) OrElse(v {{ .TypeName }}) {{ .TypeName }} {
	if {{ .VariableName }}.Present() {
		return *{{ .VariableName }}.value
	}
	return v
}

// If calls the function f with the value if the value is present.
func ({{ .VariableName }} {{ .OutputName }}) If(fn func({{ .TypeName }})) {
	if {{ .VariableName }}.Present() {
		fn(*{{ .VariableName }}.value)
	}
}

{{- if .WithJSONEncode }}

func ({{ .VariableName }} {{ .OutputName }}) MarshalJSON() ([]byte, error) {
	if {{ .VariableName }}.Present() {
		return json.Marshal({{ .VariableName }}.value)
	}
	return json.Marshal(nil)
}

func ({{ .VariableName }} *{{ .OutputName }}) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		{{ .VariableName }}.value = nil
		return nil
	}

	var value {{ .TypeName }}

	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	{{ .VariableName }}.value = &value
	return nil
}
{{- end }}
`
