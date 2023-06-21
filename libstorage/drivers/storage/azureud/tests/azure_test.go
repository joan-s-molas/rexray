package azureud

import (
	"testing"

	apitests "github.com/joan-s-molas/rexray/libstorage/api/tests"

	// load the driver packages
	"github.com/joan-s-molas/rexray/libstorage/drivers/storage/azureud"
	_ "github.com/joan-s-molas/rexray/libstorage/drivers/storage/azureud/storage"
)

func TestSuite(t *testing.T) {
	apitests.RunSuite(t, azureud.Name)
}
