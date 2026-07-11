package osunative

/*
#include "native/bin/cabinet.h"
#include <stdlib.h>
*/
import "C"
import (
	"errors"
	"runtime"
	"unsafe"
)

var errClosedBeatmap = errors.New("beatmap is closed")

type Beatmap struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewBeatmapFromFile(path string) (*Beatmap, error) {
	cpath := C.CString(path)
	defer C.free(unsafe.Pointer(cpath))

	var native C.NativeBeatmap
	result := C.Beatmap_CreateFromFile((*C.uchar)(unsafe.Pointer(cpath)), &native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}

	b := &Beatmap{handle: native.handle}
	runtime.SetFinalizer(b, (*Beatmap).Close)
	return b, nil
}

func NewBeatmapFromText(text string) (*Beatmap, error) {
	ctext := C.CString(text)
	defer C.free(unsafe.Pointer(ctext))

	var native C.NativeBeatmap
	result := C.Beatmap_CreateFromText((*C.uchar)(unsafe.Pointer(ctext)), &native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}

	b := &Beatmap{handle: native.handle}
	runtime.SetFinalizer(b, (*Beatmap).Close)
	return b, nil
}

func (b *Beatmap) Close() {
	if !b.closed {
		C.Beatmap_Destroy(b.handle)
		b.closed = true
		runtime.SetFinalizer(b, nil)
	}
}

func (b *Beatmap) Title() (string, error) {
	if b.closed {
		return "", errClosedBeatmap
	}
	return getString(b.handle, func(h C.ManagedObjectHandle, buf *C.uchar, sz *C.int32_t) C.int32_t {
		return C.int32_t(C.Beatmap_GetTitle(h, buf, sz))
	})
}

func (b *Beatmap) Artist() (string, error) {
	if b.closed {
		return "", errClosedBeatmap
	}
	return getString(b.handle, func(h C.ManagedObjectHandle, buf *C.uchar, sz *C.int32_t) C.int32_t {
		return C.int32_t(C.Beatmap_GetArtist(h, buf, sz))
	})
}

func (b *Beatmap) Version() (string, error) {
	if b.closed {
		return "", errClosedBeatmap
	}
	return getString(b.handle, func(h C.ManagedObjectHandle, buf *C.uchar, sz *C.int32_t) C.int32_t {
		return C.int32_t(C.Beatmap_GetVersion(h, buf, sz))
	})
}
