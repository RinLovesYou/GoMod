package keybinds

import (
	"GoMod/il2cpp"
	"GoMod/reflect/AsmCs"
	"GoMod/reflect/unity"

	"github.com/RinLovesYou/imgui-go"
)

var (
	flying          bool
	originalGravity *il2cpp.Vector3

	vectorZero = il2cpp.NewVector3(0, 0, 0)
)

func ToggleFlight(state bool) {
	if originalGravity == nil {
		originalGravity = unity.GetGravity()
	}

	if state {
		flying = true
		unity.SetGravity(vectorZero)
		ToggleColliders(false)
	} else {
		flying = false
		unity.SetGravity(originalGravity)
		ToggleColliders(true)
	}
}

func ToggleColliders(state bool) {
	il2cpp.GetDomain().AttachThread()

	localPlayer = AsmCs.LocalPlayer()
	localPlayerObject = localPlayer.GameObject()
	localPlayerTransform = localPlayerObject.Transform()
	camera = localPlayerObject.Find("Camera (eye)")

	components := unity.GameObjectGetComponents[unity.Collider](localPlayerObject)
	if components == nil {
		return
	}

	length := components.Length()
	for i := 0; i < length; i++ {
		component := components.Get(i)
		if component == nil || component.Object.Il2CppObject.IsNull() {
			continue
		}
		defer component.Il2CppObject.Free()

		component.SetEnabled(state)
	}
}

var (
	localPlayer          *AsmCs.Player
	localPlayerObject    *unity.GameObject
	localPlayerTransform *unity.Transform

	FlySpeed float32 = 4.0

	camera *unity.GameObject
)

func UpdateFly(param uint64) {
	switch param {
	case 'F', 'f':
		if imgui.CurrentIO().KeyShiftPressed() {
			ToggleFlight(!flying)
		}
	}
}

func UpdateFlyOnUpdate() {
	if flying {
		currentSpeed := FlySpeed

		if unity.GetKey(unity.KeyCode_LeftShift) {
			currentSpeed *= 2
		}

		pos := localPlayerTransform.Position()
		delta := unity.DeltaTime()

		if unity.GetKey(unity.KeyCode_W) {
			pos = pos.Plus(camera.Transform().Forward().Times(currentSpeed).Times(delta))
		}

		if unity.GetKey(unity.KeyCode_S) {
			pos = pos.Minus(camera.Transform().Forward().Times(currentSpeed).Times(delta))
		}

		if unity.GetKey(unity.KeyCode_A) {
			pos = pos.Minus(camera.Transform().Right().Times(currentSpeed).Times(delta))
		}

		if unity.GetKey(unity.KeyCode_D) {
			pos = pos.Plus(camera.Transform().Right().Times(currentSpeed).Times(delta))
		}

		if unity.GetKey(unity.KeyCode_Q) {
			pos = pos.Minus(camera.Transform().Up().Times(currentSpeed).Times(delta))
		}

		if unity.GetKey(unity.KeyCode_E) {
			pos = pos.Plus(camera.Transform().Up().Times(currentSpeed).Times(delta))
		}

		localPlayerTransform.SetPosition(pos)
	}
}
