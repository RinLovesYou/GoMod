package sdkbase

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
	"GoMod/reflect/unity"
	"unsafe"
)

var (
	PlayerApiClass, _ = reflect.VRCSDKBase.GetClassWhere(func(c il2cpp.Class) bool {
		return c.GetName() == "VRCPlayerApi" && c.GetNamespace() == "VRC.SDKBase"
	})

	playerApiDisplayName, _ = PlayerApiClass.GetField("displayName")
)

type VRCPlayerApi struct {
	unity.Component
}

func NewPlayerApi(o *il2cpp.Object) *VRCPlayerApi {
	api := VRCPlayerApi{}
	api.Il2CppObject = o

	return &api
}

func PlayerApiFrom(o unsafe.Pointer) *VRCPlayerApi {
	object := il2cpp.NewObject(o)
	return NewPlayerApi(object)
}

func (api *VRCPlayerApi) DisplayName() string {
	return playerApiDisplayName.GetValueObject(api.Il2CppObject).UnboxString()
}
