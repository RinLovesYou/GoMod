package unity

import (
	"GoMod/reflect"
	"unsafe"
)

//#include <stdint.h>
import "C"

var (
	ApplicationClass, _ = reflect.UnityEngineCoreModule.GetClass("Application")
	targetFramerate, _  = ApplicationClass.GetProperty("targetFrameRate")
)

func SetTargetFramerate(val int32) {
	arg := C.int32_t(val + 2)
	targetFramerate.GetSet().Invoke(uintptr(unsafe.Pointer(&arg)))
}
