package osunative

/*
#include "native/bin/cabinet.h"
#include <stdlib.h>
*/
import "C"
import (
	"runtime"
	"unsafe"
)

type Mod struct {
	handle C.ManagedObjectHandle
	closed bool
}

func NewMod(acronym string) (*Mod, error) {
	ca := C.CString(acronym)
	defer C.free(unsafe.Pointer(ca))

	var native C.NativeMod
	result := C.Mod_Create((*C.uchar)(unsafe.Pointer(ca)), &native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	m := &Mod{handle: native.handle}
	runtime.SetFinalizer(m, (*Mod).Close)
	return m, nil
}

func (m *Mod) Close() {
	if !m.closed {
		C.Mod_Destroy(m.handle)
		m.closed = true
		runtime.SetFinalizer(m, nil)
	}
}

func (m *Mod) SetSettingBool(key string, value bool) error {
	ck := C.CString(key)
	defer C.free(unsafe.Pointer(ck))
	result := C.Mod_SetSettingBool(m.handle, (*C.uchar)(unsafe.Pointer(ck)), C.bool(value))
	return ErrorCode(result)
}

func (m *Mod) SetSettingInt(key string, value int) error {
	ck := C.CString(key)
	defer C.free(unsafe.Pointer(ck))
	result := C.Mod_SetSettingInteger(m.handle, (*C.uchar)(unsafe.Pointer(ck)), C.int32_t(value))
	return ErrorCode(result)
}

func (m *Mod) SetSettingFloat(key string, value float64) error {
	ck := C.CString(key)
	defer C.free(unsafe.Pointer(ck))
	result := C.Mod_SetSettingFloat(m.handle, (*C.uchar)(unsafe.Pointer(ck)), C.float(value))
	return ErrorCode(result)
}
