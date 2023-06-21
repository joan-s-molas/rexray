package vbox

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/vbox"
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/vbox/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, vbox.Name)
}
