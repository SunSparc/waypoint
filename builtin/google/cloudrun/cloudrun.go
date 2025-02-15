// Package cloudrun contains components for deploying to Google Cloud Run.
package cloudrun

import (
	sdk "github.com/hashicorp/waypoint-plugin-sdk"
)

//go:generate protoc -I ../../../.. -I ../../../thirdparty/proto/opaqueany --go_opt=plugins=grpc --go_out=../../../.. waypoint/builtin/google/cloudrun/plugin.proto

// Options are the SDK options to use for instantiation for
// the Google Cloud Run plugin.
var Options = []sdk.Option{
	sdk.WithComponents(&Platform{}, &Releaser{}),
}
