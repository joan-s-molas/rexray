package dobs

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/dobs"
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/dobs/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, dobs.Name)
}
