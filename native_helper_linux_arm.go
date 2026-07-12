//go:build linux && arm

package osunative

/*
#cgo LDFLAGS: -L${SRCDIR}/native/bin/linux-arm -Wl,-rpath,${SRCDIR}/native/bin/linux-arm -losu.Native
*/
import "C"
