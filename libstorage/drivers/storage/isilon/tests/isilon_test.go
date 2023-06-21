package isilon

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/isilon"
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/isilon/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, isilon.Name)
}
