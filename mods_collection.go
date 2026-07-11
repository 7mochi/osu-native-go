package osunative

/*
#include "native/bin/cabinet.h"
*/
import "C"
import "runtime"

type ModsCollection struct {
	handle C.ManagedObjectHandle
	mods   []*Mod
	closed bool
}

func NewModsCollection() (*ModsCollection, error) {
	var native C.NativeModsCollection
	result := C.ModsCollection_Create(&native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	mc := &ModsCollection{handle: native.handle, mods: []*Mod{}}
	runtime.SetFinalizer(mc, (*ModsCollection).Close)
	return mc, nil
}

func (mc *ModsCollection) Add(mod *Mod) error {
	result := C.ModsCollection_Add(mc.handle, mod.handle)
	if e := ErrorCode(result); !e.IsSuccess() {
		return e
	}
	mc.mods = append(mc.mods, mod)
	return nil
}

func (mc *ModsCollection) Remove(mod *Mod) error {
	result := C.ModsCollection_Remove(mc.handle, mod.handle)
	if e := ErrorCode(result); !e.IsSuccess() {
		return e
	}
	for i, m := range mc.mods {
		if m == mod {
			mc.mods = append(mc.mods[:i], mc.mods[i+1:]...)
			break
		}
	}
	return nil
}

func (mc *ModsCollection) Close() {
	if !mc.closed {
		C.ModsCollection_Destroy(mc.handle)
		for _, m := range mc.mods {
			m.Close()
		}
		mc.closed = true
		runtime.SetFinalizer(mc, nil)
	}
}
