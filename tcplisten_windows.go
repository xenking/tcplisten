// +build windows

// Package tcplisten provides customizable TCP net.Listener with various
// performance-related options:
//
//   - SO_REUSEPORT. This option allows linear scaling server performance
//     on multi-CPU servers.
//     See https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/ for details.
//
//   - TCP_DEFER_ACCEPT. This option expects the server reads from the accepted
//     connection before writing to them.
//
//   - TCP_FASTOPEN. See https://lwn.net/Articles/508865/ for details.
//
//   - TCP_NODELAY. This option is intended to disable/enable segment buffering so data can be sent out to peer
//     as quickly as possible, so this is typically used to improve network utilization.
//
//   - TCP_QUICKACK. This option allows sending out acknowledgments as early as possible than delayed
//     under some protocol level exchanging, and it's not stable/permanent, subsequent TCP transactions
//     (which may happen under the hood) can disregard this option depending on actual protocol level processing
//     or any actual disagreements between user setting and stack behavior.
//
// The package is derived from https://github.com/kavu/go_reuseport .
package tcplisten

import (
	"net"
)

// Config provides options to enable on the returned listener.
type Config struct {
	// ReusePort enables SO_REUSEPORT.
	ReusePort bool

	// DeferAccept enables TCP_DEFER_ACCEPT.
	DeferAccept bool

	// FastOpen enables TCP_FASTOPEN.
	FastOpen bool

	// Backlog is the maximum number of pending TCP connections the listener
	// may queue before passing them to Accept.
	// See man 2 listen for details.
	//
	// By default system-level backlog value is used.
	Backlog int
}

// NewListener returns TCP listener with options set in the Config.
//
// The function may be called many times for creating distinct listeners
// with the given config.
//
// Only tcp4 and tcp6 networks are supported.
func NewListener(network, addr string, cfg Config) (net.Listener, error) {
	return net.Listen(network, addr)
}
