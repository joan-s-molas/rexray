package rbd

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/rbd"
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/rbd/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, rbd.Name)
}
