// +build go1.1

package docker

import (
	"bufio"
)

func writePoolGet() *bufio.Writer {
	return bufio.NewWriterSize(nil, 32*1024)
}

func writePoolPut(b *bufio.Writer) {
  // is noop the correct action here?
}
