// +build go1.1 go1.2

package docker

import "archive/tar"

func addXattrs(hdr *tar.Header, path string) *tar.Header {
	return hdr
}
