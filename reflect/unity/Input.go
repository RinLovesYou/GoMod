package unity

import (
	"GoMod/il2cpp"
	"GoMod/reflect"
	"unsafe"
)

var (
	InputClass, _ = reflect.UnityEngineInputLegacyModule.GetClass("Input")

	getKeyMethod, _ = InputClass.GetMethodWhere(func(method il2cpp.Method) bool {
		return method.GetName() == "GetKey" && len(method.GetParams()) == 1
	})

	getKeyDownMethod, _ = InputClass.GetMethodWhere(func(method il2cpp.Method) bool {
		return method.GetName() == "GetKeyDown" && len(method.GetParams()) == 1
	})
)

func GetKey(key KeyCode) bool {
	ret, err := getKeyMethod.Invoke(uintptr(unsafe.Pointer(&key)))
	if err != nil {
		return false
	}

	return *(*bool)(ret.Unbox())
}

func GetKeyDown(key KeyCode) bool {
	ret, err := getKeyDownMethod.Invoke(uintptr(unsafe.Pointer(&key)))
	if err != nil {
		return false
	}

	return *(*bool)(ret.Unbox())
}

type KeyCode int

const (
	KeyCode_A            KeyCode = 97
	KeyCode_B            KeyCode = 98
	KeyCode_C            KeyCode = 99
	KeyCode_D            KeyCode = 100
	KeyCode_E            KeyCode = 101
	KeyCode_F            KeyCode = 102
	KeyCode_G            KeyCode = 103
	KeyCode_H            KeyCode = 104
	KeyCode_I            KeyCode = 105
	KeyCode_J            KeyCode = 106
	KeyCode_K            KeyCode = 107
	KeyCode_L            KeyCode = 108
	KeyCode_M            KeyCode = 109
	KeyCode_N            KeyCode = 110
	KeyCode_O            KeyCode = 111
	KeyCode_P            KeyCode = 112
	KeyCode_Q            KeyCode = 113
	KeyCode_R            KeyCode = 114
	KeyCode_S            KeyCode = 115
	KeyCode_T            KeyCode = 116
	KeyCode_U            KeyCode = 117
	KeyCode_V            KeyCode = 118
	KeyCode_W            KeyCode = 119
	KeyCode_X            KeyCode = 120
	KeyCode_Y            KeyCode = 121
	KeyCode_Z            KeyCode = 122
	KeyCode_RightShift   KeyCode = 303
	KeyCode_LeftShift    KeyCode = 304
	KeyCode_RightControl KeyCode = 305
	KeyCode_LeftControl  KeyCode = 306
)
