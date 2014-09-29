// +build !go1.1

package docker

import "github.com/docker/docker/pkg/pools"

func writePoolGet() *bufio.Writer {
	return pools.BufioWriter32KPool.Get(nil)
}

func writePoolPut(b *bufio.Writer) {
	pools.BufioWriter32KPool.Put(b)
}
