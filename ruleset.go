package osunative

/*
#include "build/generated/cabinet.h"
*/
import "C"
import "runtime"

type Ruleset struct {
	handle    C.ManagedObjectHandle
	rulesetID int
	closed    bool
}

func NewRulesetFromID(id int) (*Ruleset, error) {
	var native C.NativeRuleset
	result := C.Ruleset_CreateFromId(C.int32_t(id), &native)
	if err := ErrorCode(result); !err.IsSuccess() {
		return nil, err
	}
	r := &Ruleset{handle: native.handle, rulesetID: id}
	runtime.SetFinalizer(r, (*Ruleset).Close)
	return r, nil
}

func (r *Ruleset) Close() {
	if !r.closed {
		C.Ruleset_Destroy(r.handle)
		r.closed = true
		runtime.SetFinalizer(r, nil)
	}
}

func (r *Ruleset) ID() int {
	return r.rulesetID
}
