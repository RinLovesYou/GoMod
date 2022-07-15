package unity

import (
	"GoMod/reflect"
	"unsafe"
)

//#include <stdint.h>
//#include <stdlib.h>
import "C"

var (
	ApplicationClass, _ = reflect.UnityEngineCoreModule.GetClass("Application")
	targetFramerate, _  = ApplicationClass.GetProperty("targetFrameRate")
)

func SetTargetFramerate(val int) {
	targetFramerate.GetSet().Invoke(
		uintptr(unsafe.Pointer(&val)),
	)
}
