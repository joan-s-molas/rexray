// +build !linux

package executor

import "github.com/joan-s-molas/rexray/libstorage/api/types"

func getMountedBuckets(
	ctx types.Context,
	s3fsBinName string) (map[string]string, error) {

	return nil, types.ErrNotImplemented
}
