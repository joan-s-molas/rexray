package gcepd

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/gcepd"
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/gcepd/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, gcepd.Name)
}
