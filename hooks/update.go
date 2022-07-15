package hooks

/*
	typedef void(*update_t)();
	void onUpdate();
*/
import "C"
import (
	"GoMod/dispatcher"
	"GoMod/keybinds"
	"GoMod/reflect/AsmCs"
	"GoMod/utils"
	"errors"
	"unsafe"

	"github.com/nanitefactory/gominhook"
)

var (
	updateOriginal C.update_t
)

func initUpdateHook() error {
	method := AsmCs.UpdateMethod
	if method.IsNull() {
		return errors.New("failed to find update method")
	}

	methodPtr := method.Pointer()
	err := gominhook.CreateHook(methodPtr, uintptr(C.onUpdate), uintptr(unsafe.Pointer(&updateOriginal)))
	if err != nil {
		return err
	}

	return gominhook.EnableHook(methodPtr)
}

//export onUpdate
func onUpdate() {
	keybinds.UpdateFlyOnUpdate()

	select {
	case fn := <-dispatcher.DispatcherQueue:
		err := fn()
		if err != nil {
			utils.Error("Failed to invoke command %s\n", err.Error())
		}
	default:
		break
	}

	toRemove := make([]int, 0)

	for i, coroutine := range dispatcher.Coroutines {
		if coroutine.Method(coroutine.Args...) {
			toRemove = append(toRemove, i)
		}
	}

	for _, i := range toRemove {
		dispatcher.Coroutines = append(dispatcher.Coroutines[:i], dispatcher.Coroutines[i+1:]...)
	}
}
