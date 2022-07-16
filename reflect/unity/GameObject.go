package unity

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
	"GoMod/reflect/mscorlib"
	"unsafe"
)

import "C"

var (
	GameObjectClass, _ = reflect.UnityEngineCoreModule.GetClass("GameObject")

	getComponentsMethod, _ = GameObjectClass.GetMethodWhere(func(m il2cpp.Method) bool {
		return m.GetName() == "GetComponents" && len(m.GetParams()) == 1 && m.GetReturnType().GetName() == "UnityEngine.Component[]"
	})

	getComponentsInChildrenMethod, _ = GameObjectClass.GetMethodWhere(func(m il2cpp.Method) bool {
		return m.GetName() == "GetComponentsInChildren" &&
			len(m.GetParams()) == 2 &&
			m.GetReturnType().GetName() == "UnityEngine.Component[]" &&
			m.GetParams()[1].GetName() == "System.Boolean"
	})

	transformProperty, _ = GameObjectClass.GetProperty("transform")

	findMethod, _ = GameObjectClass.GetMethod("Find")
)

type GameObject struct {
	Object
}

func NewGameObject(obj *il2cpp.Object) *GameObject {
	object := &GameObject{}
	object.Il2CppObject = obj
	return object
}

func GameObjectGetComponentsInChildren[T reflect.Generic](obj *GameObject, includeInactive bool) *mscorlib.Array[T] {
	dummy := *new(T)
	res, err := getComponentsInChildrenMethod.InvokeObject(
		obj.Il2CppObject,
		uintptr(dummy.GetType().Handle),
		uintptr(unsafe.Pointer(&includeInactive)),
	)
	if err != nil {
		return nil
	}

	arr := &mscorlib.Array[T]{}
	arr.Object = *res
	return arr
}

func GameObjectGetComponents[T reflect.Generic](obj *GameObject) *mscorlib.Array[T] {
	dummy := *new(T)
	xType := dummy.GetType()

	res, err := getComponentsMethod.InvokeObject(
		obj.Il2CppObject,
		uintptr(xType.Handle),
	)

	if err != nil {
		return nil
	}

	arr := &mscorlib.Array[T]{}
	arr.Object = *res

	return arr
}

func (g *GameObject) Transform() *Transform {
	obj, err := transformProperty.GetGet().InvokeObject(g.Il2CppObject)
	if err != nil {
		return nil
	}

	return NewTransform(obj)
}

func (g *GameObject) Find(path string) *GameObject {
	cPath := il2cpp.NewString(path)

	obj, err := findMethod.InvokeObject(g.Il2CppObject, uintptr(unsafe.Pointer(cPath.Handle)))
	if err != nil {
		return nil
	}

	return NewGameObject(obj)
}
