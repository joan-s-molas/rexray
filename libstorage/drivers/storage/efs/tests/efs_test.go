package efs

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/efs"
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/efs/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, efs.Name)
}
