package main

import (
	"flag"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"
)

const version = "v1.0.0"

func main() {
	// 输出版本
	showVersion := flag.Bool("version", false, "print the version and exit")

	flag.Parse()
	if *showVersion {
		fmt.Printf("protoc-gen-go-errors %v\n", version)
		return
	}


	var flags flag.FlagSet
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {		// 运行protoc插件代码
		gen.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL) // protobuf语言支持的特性
		for _, f := range gen.Files {	// A File describes a .proto source file.
			if !f.Generate {
				continue
			}
			generateFile(gen, f)
		}
		return nil
	})
}
