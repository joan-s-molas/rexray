// +build !client
// +build !controller

package cli

import "github.com/joan-s-molas/rexray/agent"

func init() {
	startFuncs = append(startFuncs, agent.Start)
}
