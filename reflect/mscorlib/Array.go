package mscorlib

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
	"GoMod/utils"
	"unsafe"
)

//#include <stdint.h>
import "C"

var (
	ArrayClass, _ = reflect.Mscorlib.GetClass("Array")

	createInstanceMethod, _ = ArrayClass.GetMethodWhere(func(m il2cpp.Method) bool {
		return m.GetName() == "CreateInstance" && len(m.GetParams()) == 2
	})

	lengthProperty, _ = ArrayClass.GetProperty("Length")

	getValueMethod, _ = ArrayClass.GetMethodWhere(func(m il2cpp.Method) bool {
		return m.GetName() == "GetValue" && len(m.GetParams()) == 1
	})

	setValueMethod, _ = ArrayClass.GetMethodWhere(func(m il2cpp.Method) bool {
		return m.GetName() == "SetValue" && len(m.GetParams()) == 2
	})

	resizeMethod, _ = ArrayClass.GetMethod("Resize")
)

type Array[K reflect.Generic] struct {
	il2cpp.Object
}

func NewArray[K reflect.Generic](length int) *Array[K] {
	dummy := *new(K)
	xType := dummy.GetType()
	res, err := createInstanceMethod.Invoke(
		uintptr(unsafe.Pointer(xType.Handle)),
		uintptr(unsafe.Pointer(&length)),
	)
	if err != nil {
		utils.Error(err.Error())
		return nil
	}

	return &Array[K]{
		Object: *res,
	}

}

func (a *Array[K]) Length() int {
	res, err := lengthProperty.GetGet().InvokeObject(&a.Object)
	if err != nil {
		return 0
	}

	return *(*int)(res.Unbox())
}

func (a *Array[K]) Set(value unsafe.Pointer, index int) {
	setValueMethod.InvokeObject(
		&a.Object,
		uintptr(value),
		uintptr(unsafe.Pointer(&index)),
	)
}

func (a *Array[K]) Get(index int) *K {
	res, err := getValueMethod.InvokeObject(
		&a.Object,
		uintptr(unsafe.Pointer(&index)),
	)
	if err != nil {
		utils.Error(err.Error())
		return nil
	}

	dummy := *new(K)

	return dummy.Construct(res).(*K)
}
