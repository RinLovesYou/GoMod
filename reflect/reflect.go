package reflect

import (
	"GoMod/il2cpp"
)

var (
	AssemblyCSharp, _               = il2cpp.GetDomain().GetImage("Assembly-CSharp")
	VRCSDKBase, _                   = il2cpp.GetDomain().GetImage("VRCSDKBase")
	UnityEngineCoreModule, _        = il2cpp.GetDomain().GetImage("UnityEngine.CoreModule")
	UnityEngineInputLegacyModule, _ = il2cpp.GetDomain().GetImage("UnityEngine.InputLegacyModule")
	UnityEnginePhysicsModule, _     = il2cpp.GetDomain().GetImage("UnityEngine.PhysicsModule")
	Mscorlib, _                     = il2cpp.GetDomain().GetImage("mscorlib")
	TextMeshPro, _                  = il2cpp.GetDomain().GetImage("Unity.TextMeshPro")
)

type Generic interface {
	GetType() *il2cpp.Object
	Construct(*il2cpp.Object) any
	GetClass() il2cpp.Class
}
