package tmpro

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
)

var TMProUGuiClass, _ = reflect.TextMeshPro.GetClass("TextMeshProUGUI")

type TMProUGui struct {
	TMP_Text
}

func NewTMProUGui(obj *il2cpp.Object) *TMProUGui {
	object := &TMProUGui{}
	object.Il2CppObject = obj
	return object
}

func (c TMProUGui) GetType() *il2cpp.Object {
	return TMProUGuiClass.TypeObject()
}
func (c TMProUGui) Construct(obj *il2cpp.Object) any {
	return NewTMProUGui(obj)
}
func (c TMProUGui) GetClass() il2cpp.Class {
	return TMProUGuiClass
}
