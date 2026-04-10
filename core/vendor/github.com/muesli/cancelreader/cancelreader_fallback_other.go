//go:build !windows

package cancelreader

import "io"

func cancelFallbackRead(_ io.Reader) bool { return false }
