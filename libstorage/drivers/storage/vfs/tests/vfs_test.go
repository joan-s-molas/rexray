package vfs

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/vfs"

	// load the vfs driver packages
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/vfs/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, vfs.Name)
}
