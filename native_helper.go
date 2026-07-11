package osunative

/*
#include "native/bin/cabinet.h"

static inline int nullable_double_hasValue(void* p) {
    return ((Cabinet__Nullable_double*)p)->hasValue ? 1 : 0;
}

static inline double nullable_double_value(void* p) {
    return ((Cabinet__Nullable_double*)p)->value;
}

static inline void nullable_int64_setHasValue(void* p, int v) {
    ((Cabinet__Nullable_int64_t*)p)->hasValue = v;
}

static inline void nullable_int64_setValue(void* p, long long v) {
    ((Cabinet__Nullable_int64_t*)p)->value = v;
}
*/
import "C"
import "unsafe"

type nativeStringFunc func(C.ManagedObjectHandle, *C.uchar, *C.int32_t) C.int32_t

func getString(handle C.ManagedObjectHandle, fn nativeStringFunc) (string, error) {
	var size C.int32_t
	result := fn(handle, nil, &size)
	if code := ErrorCode(result); code != BufferSizeQuery && code != Success {
		return "", code
	}
	if size <= 0 {
		return "", nil
	}
	buf := make([]byte, size)
	result = fn(handle, (*C.uchar)(unsafe.Pointer(&buf[0])), &size)
	if err := ErrorCode(result); !err.IsSuccess() {
		return "", err
	}
	s := string(buf[:size])
	for len(s) > 0 && s[len(s)-1] == 0 {
		s = s[:len(s)-1]
	}
	return s, nil
}

func readNullableDouble(p unsafe.Pointer) *float64 {
	if p == nil {
		return nil
	}
	if C.nullable_double_hasValue(p) == 0 {
		return nil
	}
	v := float64(C.nullable_double_value(p))
	return &v
}

func setNullableInt64(p unsafe.Pointer, v int64) {
	C.nullable_int64_setHasValue(p, 1)
	C.nullable_int64_setValue(p, C.longlong(v))
}
