package AsmCs

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
	"GoMod/reflect/sdkbase"
	"GoMod/reflect/unity"
	"unsafe"
)

var (
	PlayerClass, _ = reflect.AssemblyCSharp.GetClassWhere(func(c il2cpp.Class) bool {
		return c.HasMethod("Awake") &&
			c.HasMethod("OnDestroy") &&
			c.HasMethod("OnNetworkReady") &&
			c.HasMethod("Start") &&
			c.HasMethod("ToString") &&
			c.HasMethod("Update")
	})

	apiPlayerProp, _ = PlayerClass.GetPropertyWhere(func(p il2cpp.Property) bool {
		return p.GetGet().GetReturnType().GetName() == "VRC.SDKBase.VRCPlayerApi"
	})
)

type Player struct {
	unity.Component
}

func NewPlayer(o *il2cpp.Object) *Player {
	player := Player{}
	player.Il2CppObject = o

	return &player
}

func PlayerFrom(o unsafe.Pointer) *Player {
	object := il2cpp.NewObject(o)
	return NewPlayer(object)
}

func (p *Player) ApiPlayer() *sdkbase.VRCPlayerApi {
	obj, err := apiPlayerProp.GetGet().InvokeObject(p.Il2CppObject)
	if err != nil {
		return nil
	}

	return sdkbase.NewPlayerApi(obj)
}
