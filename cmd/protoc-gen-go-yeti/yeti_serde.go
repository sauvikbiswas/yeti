package main

import (
	"fmt"
	"strings"

	"github.com/sauvikbiswas/yeti/proto/options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	fileDescriptorProtoPackageFieldNumber = 2  // FileDescriptorProto.package field number
	fileDescriptorProtoSyntaxFieldNumber  = 12 // FileDescriptorProto.syntax field number
)

const (
	yetiPackage      = protogen.GoImportPath("github.com/sauvikbiswas/yeti")
	protojsonPackage = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	fmtPackage       = protogen.GoImportPath("fmt")
	timePackage      = protogen.GoImportPath("time")
)

func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Messages) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_yeti.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	// Attach all comments associated with the syntax field.
	genLeadingComments(g, file.Desc.SourceLocations().ByPath(protoreflect.SourcePath{fileDescriptorProtoSyntaxFieldNumber}))
	g.P("// Code generated by protoc-gen-go-yeti. DO NOT EDIT.")
	g.P("// versions:")
	g.P("// - protoc-gen-go-yeti v", version)
	g.P("// - protoc             ", protocVersion(gen))
	if file.Proto.GetOptions().GetDeprecated() {
		g.P("// ", file.Desc.Path(), " is a deprecated file.")
	} else {
		g.P("// source: ", file.Desc.Path())
	}
	g.P()
	// Attach all comments associated with the package field.
	genLeadingComments(g, file.Desc.SourceLocations().ByPath(protoreflect.SourcePath{fileDescriptorProtoPackageFieldNumber}))
	g.P("package ", file.GoPackageName)
	g.P()
	generateFileContent(gen, file, g)
	return g
}

func protocVersion(gen *protogen.Plugin) string {
	v := gen.Request.GetCompilerVersion()
	if v == nil {
		return "(unknown)"
	}
	var suffix string
	if s := v.GetSuffix(); s != "" {
		suffix = "-" + s
	}
	return fmt.Sprintf("v%d.%d.%d%s", v.GetMajor(), v.GetMinor(), v.GetPatch(), suffix)
}

func genLeadingComments(g *protogen.GeneratedFile, loc protoreflect.SourceLocation) {
	for _, s := range loc.LeadingDetachedComments {
		g.P(protogen.Comments(s))
		g.P()
	}
	if s := loc.LeadingComments; s != "" {
		g.P(protogen.Comments(s))
		g.P()
	}
}

func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Messages) == 0 {
		return
	}

	for _, message := range file.Messages {
		genMessage(gen, file, g, message)
	}
}

func genMessage(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, msg *protogen.Message) {
	g.P("func (x *", msg.Desc.FullName().Name(), ") New() ", yetiPackage.Ident("Record"), " {")
	g.P("return &", msg.Desc.FullName().Name(), "{}")
	g.P("}")
	g.P()
	g.P("func (x *", msg.Desc.FullName().Name(), ") YetiSerialize() ([]byte, error) {")
	g.P("return ", protojsonPackage.Ident("Marshal"), "(x)")
	g.P("}")
	g.P()
	g.P("func (x *", msg.Desc.FullName().Name(), ") YetiDeserialize(b []byte) error {")
	g.P("return ", protojsonPackage.Ident("Unmarshal"), "(b, x)")
	g.P("}")
	g.P()
	g.P("func (x *", msg.Desc.FullName().Name(), ") YetiType() string {")
	g.P("return \"", msg.Desc.FullName().Name(), "\"")
	g.P("}")
	g.P()
	g.P("func (x *", msg.Desc.FullName().Name(), ") YetiKey() (string, error) {")
	g.P("var err error")
	getKey(gen, file, g, msg)
	g.P("return key, err")
	g.P("}")
	g.P()
}

func getKey(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile, msg *protogen.Message) {
	primaryKeyGoFuncs := make([]string, 0)
	primaryKeyFields := make([]string, 0)
	for _, field := range msg.Fields {
		messageFieldDesc := field.Desc
		fieldOpts := messageFieldDesc.Options().(*descriptorpb.FieldOptions)
		v := proto.GetExtension(fieldOpts, options.E_YetiFieldOpts).(*options.YetiFieldOptions)
		if v.GetPrimaryKey() {
			if messageFieldDesc.Kind() == protoreflect.StringKind {
				primaryKeyGoFuncs = append(primaryKeyGoFuncs, "x.Get"+field.GoName+"()")
				primaryKeyFields = append(primaryKeyFields, string(messageFieldDesc.Name()))
			} else {
				g.P("// cannot use non-string field ", messageFieldDesc.Name(), " as part of primary key")
			}
		}
	}
	if len(primaryKeyGoFuncs) == 0 {
		g.P("t := ", timePackage.Ident("Now"), "()")
		g.P("key := t.Format(\"20060102150405\")")
	} else {
		for i, goFunc := range primaryKeyGoFuncs {
			g.P("if ", goFunc, "==\"\" { return \"\", ", fmtPackage.Ident("Errorf"), "(\"", primaryKeyFields[i], " is not set\")}")
		}
		g.P("key := ", strings.Join(primaryKeyGoFuncs, " + "))
	}
}
