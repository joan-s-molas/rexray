// +build !client
// +build !agent
// +build !controller

package util

import (
	gofig "github.com/akutz/gofig/types"

	apitypes "github.com/joan-s-molas/rexray/libstorage/api/types"
	apiclient "github.com/joan-s-molas/rexray/libstorage/client"
)

func newClient(
	ctx apitypes.Context, config gofig.Config) (apitypes.Client, error) {
	return apiclient.New(ctx, config)
}
