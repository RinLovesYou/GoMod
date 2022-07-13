package AsmCs

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
	"GoMod/reflect/unity"
)

var VRCApplicationClass, _ = reflect.AssemblyCSharp.GetClassWhere(func(c il2cpp.Class) bool {
	return c.HasMethod("Awake") &&
		c.HasMethod("OnApplicationQuit") &&
		c.HasMethod("OnApplicationFocus") &&
		c.HasMethod("OnApplicationPause") &&
		c.HasMethod("OnDestroy")
})

var UpdateMethod, _ = VRCApplicationClass.GetMethod("Update")

type VRCApplication struct {
	unity.Component
}
