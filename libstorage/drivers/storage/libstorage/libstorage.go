package libstorage

import (
	"github.com/joan-s-molas/rexray/libstorage/api/registry"
	"github.com/joan-s-molas/rexray/libstorage/api/types"
)

const (
	// Name is the name of the driver.
	Name = types.LibStorageDriverName
)

func init() {
	registry.RegisterStorageDriver(Name, newDriver)
}
