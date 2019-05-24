package main

import (
	"os/exec"

	"github.com/urfave/cli"
)

const (
	_getGRPCGen = "go get github.com/gogo/protobuf/protoc-gen-gogofast"
	_grpcProtoc = "protoc --proto_path=%s --proto_path=%s --proto_path=%s --gogofast_out=plugins=grpc,Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types:."
)

func installGRPCGen() error {
	if _, err := exec.LookPath("protoc-gen-gogofast"); err != nil {
		if err := goget(_getGRPCGen); err != nil {
			return err
		}
	}
	return nil
}

func genGRPC(ctx *cli.Context) error {
	return generate(ctx, _grpcProtoc)
}
