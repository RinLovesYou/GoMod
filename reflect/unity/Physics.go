package unity

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
)

var (
	PhysicsClass, _ = reflect.UnityEnginePhysicsModule.GetClass("Physics")

	gravityProperty, _ = PhysicsClass.GetProperty("gravity")
)

func GetGravity() *il2cpp.Vector3 {
	method := gravityProperty.GetGet()
	if method.IsNull() {
		return nil
	}

	ret, err := method.Invoke()
	if err != nil {
		return nil
	}

	return il2cpp.Vector3FromPointer(ret.Unbox())
}

func SetGravity(gravity *il2cpp.Vector3) {
	method := gravityProperty.GetSet()
	if method.IsNull() {
		return
	}

	method.Invoke(uintptr(gravity.Handle()))
}
