// +build controller

package util

import (
	gofig "github.com/akutz/gofig/types"
	apitypes "github.com/joan-s-molas/rexray/libstorage/api/types"
)

func newClient(apitypes.Context, gofig.Config) (apitypes.Client, error) {
	return nil, nil
}
