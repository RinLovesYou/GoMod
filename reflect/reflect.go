package reflect

import "GoMod/il2cpp"

var (
	AssemblyCSharp, _        = il2cpp.GetDomain().GetImage("Assembly-CSharp")
	VRCSDKBase, _            = il2cpp.GetDomain().GetImage("VRCSDKBase")
	UnityEngineCoreModule, _ = il2cpp.GetDomain().GetImage("UnityEngine.CoreModule")
)
