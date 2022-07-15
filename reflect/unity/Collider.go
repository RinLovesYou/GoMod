package unity

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
	"unsafe"
)

//#include <stdbool.h>
import "C"

var (
	ColliderClass, _ = reflect.UnityEnginePhysicsModule.GetClass("Collider")

	enabledProperty, _ = ColliderClass.GetProperty("enabled")
)

type Collider struct {
	Component
}

func NewCollider(obj *il2cpp.Object) *Collider {
	col := &Collider{}
	col.Il2CppObject = obj
	return col
}

func (c *Collider) SetEnabled(state bool) {
	cState := C.bool(state)
	enabledProperty.GetSet().InvokeObject(c.Il2CppObject, uintptr(unsafe.Pointer(&cState)))
}

func (c Collider) GetType() *il2cpp.Object {
	return ColliderClass.TypeObject()
}
func (c Collider) Construct(obj *il2cpp.Object) any {
	return NewCollider(obj)
}
func (c Collider) GetClass() il2cpp.Class {
	return ColliderClass
}
