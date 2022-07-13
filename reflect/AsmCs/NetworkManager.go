package AsmCs

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
)

var (
	NetworkManager, _ = reflect.AssemblyCSharp.GetClassWhere(func(c il2cpp.Class) bool {
		return c.HasMethod("OnConnectedToMaster") &&
			c.HasMethod("OnCreatedRoom") &&
			c.HasMethod("OnEnable")
	})
	onPlayerFuncs []il2cpp.Method
)

func OnVRCPlayerMethods() []il2cpp.Method {
	if len(onPlayerFuncs) > 0 {
		return onPlayerFuncs
	}

	if NetworkManager.IsNull() || PlayerClass.IsNull() {
		return onPlayerFuncs
	}

	onPlayerFuncs = NetworkManager.GetMethodsWhere(func(m il2cpp.Method) bool {
		return len(m.GetParams()) == 1 &&
			m.GetParams()[0].GetName() == PlayerClass.GetName()
	})

	return onPlayerFuncs
}
