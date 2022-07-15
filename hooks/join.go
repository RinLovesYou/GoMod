package hooks

/*
	#include <stdint.h>

	typedef void(*on_join)(void*, void*);

	__attribute__((weak))
	void InvokePlayerJoin(on_join callback, void* managerPtr, void* player) {
		callback(managerPtr, player);
	}

	void onJoin(void* managerPtr, void* playerPtr);
	void onLeave(void* managerPtr, void* playerPtr);
*/
import "C"
import (
	"GoMod/menu"
	"GoMod/reflect/AsmCs"
	"GoMod/utils"
	"unsafe"

	"github.com/nanitefactory/gominhook"
)

var (
	onJoinOriginal  C.on_join
	onLeaveOriginal C.on_join
	playerCache     = make(map[uintptr]string)
)

func initJoinHooks() error {
	methods := AsmCs.OnVRCPlayerMethods()

	//onJoin
	method := methods[1]
	methodPtr := method.Pointer()

	err := gominhook.CreateHook(methodPtr, uintptr(C.onJoin), uintptr(unsafe.Pointer(&onJoinOriginal)))
	if err != nil {
		return err
	}
	err = gominhook.EnableHook(methodPtr)
	if err != nil {
		return err
	}

	//onLeave
	method = methods[0]
	methodPtr = method.Pointer()
	if err != nil {
		return err
	}
	err = gominhook.CreateHook(methodPtr, uintptr(C.onLeave), uintptr(unsafe.Pointer(&onLeaveOriginal)))
	if err != nil {
		return err
	}

	return gominhook.EnableHook(methodPtr)
}

//export onJoin
func onJoin(managerPtrC, playerPtrC unsafe.Pointer) {
	playerPtr := uintptr(playerPtrC)
	player := AsmCs.PlayerFrom(playerPtrC)
	if player.Il2CppObject.IsNull() {
		utils.Error("Unknown Player joined.")
		return
	}

	name, err := utils.ToAscii(player.ApiPlayer().DisplayName())
	if err != nil {
		name = player.ApiPlayer().DisplayName()
		utils.Error("There was an error converting the player's name")
	}

	playerCache[playerPtr] = name

	utils.Log("Player joined: %s", playerCache[playerPtr])

	refreshNameCache()

	C.InvokePlayerJoin(onJoinOriginal, managerPtrC, playerPtrC)
}

//export onLeave
func onLeave(managerPtrC, playerPtrC unsafe.Pointer) {
	playerPtr := uintptr(playerPtrC)

	if player, ok := playerCache[playerPtr]; ok {
		utils.Log("Player left: %s", player)
		delete(playerCache, playerPtr)
		refreshNameCache()

	}
	C.InvokePlayerJoin(onLeaveOriginal, managerPtrC, playerPtrC)
}

func refreshNameCache() {
	menu.PlayerListMenu.Players = make([]string, 0)
	for _, player := range playerCache {
		menu.PlayerListMenu.Players = append(menu.PlayerListMenu.Players, player)
	}
}
