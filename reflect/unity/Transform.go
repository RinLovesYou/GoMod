package unity

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
)

var (
	TransformClass, _ = reflect.UnityEngineCoreModule.GetClass("Transform")

	forwardProperty, _  = TransformClass.GetProperty("forward")
	rightProperty, _    = TransformClass.GetProperty("right")
	upProperty, _       = TransformClass.GetProperty("up")
	positionProperty, _ = TransformClass.GetProperty("position")
)

type Transform struct {
	Component
}

func NewTransform(obj *il2cpp.Object) *Transform {
	t := &Transform{}
	t.Il2CppObject = obj
	return t
}

func (t *Transform) Position() *il2cpp.Vector3 {
	res, err := positionProperty.GetGet().InvokeObject(t.Il2CppObject)
	if err != nil {
		return nil
	}

	return il2cpp.Vector3FromPointer(res.Unbox())
}

func (t *Transform) SetPosition(pos *il2cpp.Vector3) {
	positionProperty.GetSet().InvokeObject(t.Il2CppObject, uintptr(pos.Handle()))
}

func (t *Transform) Forward() *il2cpp.Vector3 {
	res, err := forwardProperty.GetGet().InvokeObject(t.Il2CppObject)
	if err != nil {
		return nil
	}

	return il2cpp.Vector3FromPointer(res.Unbox())
}

func (t *Transform) Right() *il2cpp.Vector3 {
	res, err := rightProperty.GetGet().InvokeObject(t.Il2CppObject)
	if err != nil {
		return nil
	}

	return il2cpp.Vector3FromPointer(res.Unbox())
}

func (t *Transform) Up() *il2cpp.Vector3 {
	res, err := upProperty.GetGet().InvokeObject(t.Il2CppObject)
	if err != nil {
		return nil
	}

	return il2cpp.Vector3FromPointer(res.Unbox())
}

func (c Transform) GetType() *il2cpp.Object {
	return TransformClass.TypeObject()
}
func (c Transform) Construct(obj *il2cpp.Object) any {
	return NewTransform(obj)
}
func (c Transform) GetClass() il2cpp.Class {
	return TransformClass
}
