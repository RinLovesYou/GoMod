package main

import (
	"GoMod/hooks"
	"GoMod/il2cpp"
	"GoMod/utils"
)

func init() {
	utils.Log("initializing GoMod")
	utils.Log("initializing hooks")
	err := hooks.InitHooks()
	if err != nil {
		utils.Error("Failed to initialize hooks: %s\n", err.Error())
	}

	domain := il2cpp.GetDomain()
	domain.AttachThread()

	utils.Log("Welcome, GoMod's default keybind is ctrl + y")
}

func main() {

}
