[![Build Status](https://travis-ci.org/xenking/tcplisten.svg)](https://travis-ci.org/xenking/tcplisten)
[![Go Reference](https://pkg.go.dev/badge/github.com/xenking/tcplisten.svg)](https://pkg.go.dev/github.com/xenking/tcplisten)
[![Go Report](https://goreportcard.com/badge/github.com/xenking/tcplisten)](https://goreportcard.com/report/github.com/xenking/tcplisten)


Package tcplisten provides customizable TCP net.Listener with various
performance-related options:

 * SO_REUSEPORT. This option allows linear scaling server performance
   on multi-CPU servers.
   See https://www.nginx.com/blog/socket-sharding-nginx-release-1-9-1/ for details.

 * TCP_DEFER_ACCEPT. This option expects the server reads from the accepted
   connection before writing to them.

 * TCP_FASTOPEN. See https://lwn.net/Articles/508865/ for details.
    
 * TCP_NODELAY. This option is intended to disable/enable segment buffering so data can be sent out to peer
   as quickly as possible, so this is typically used to improve network utilization.

 * TCP_QUICKACK. This option allows sending out acknowledgments as early as possible than delayed
   under some protocol level exchanging, and it's not stable/permanent, subsequent TCP transactions
   (which may happen under the hood) can disregard this option depending on actual protocol level processing
   or any actual disagreements between user setting and stack behavior.

[Documentation](https://pkg.go.dev/github.com/xenking/tcplisten).

The package is derived from [go_reuseport](https://github.com/kavu/go_reuseport).
