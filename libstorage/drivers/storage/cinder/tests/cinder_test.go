package cinder

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/cinder"
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/cinder/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, cinder.Name)
}
