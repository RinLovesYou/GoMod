package tmpro

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
	"GoMod/reflect/unity"
)

var (
	TMPTextClass, _ = reflect.TextMeshPro.GetClass("TMP_Text")
	textProperty, _ = TMPTextClass.GetProperty("text")
)

type TMP_Text struct {
	unity.Component
}

func (t *TMP_Text) SetText(text string) {
	il2cppString := il2cpp.NewString(text)
	textProperty.GetSet().InvokeObject(
		t.Il2CppObject,
		uintptr(il2cppString.Handle),
	)
}
