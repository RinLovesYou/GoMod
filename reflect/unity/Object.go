package unity

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
)

var (
	ObjectClass, _ = reflect.UnityEngineCoreModule.GetClass("Object")

	objectNameProp, _ = ObjectClass.GetProperty("name")
)

type Object struct {
	Il2CppObject *il2cpp.Object
}

func (o *Object) Name() string {
	ret, err := objectNameProp.GetGet().InvokeObject(o.Il2CppObject)
	if err != nil {
		return ""
	}

	return ret.UnboxString()
}
