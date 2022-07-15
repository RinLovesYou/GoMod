package main

import (
	"GoMod/hooks"
	"GoMod/il2cpp"
	"GoMod/utils"
	"os"
)

func init() {
	os.Setenv("GODEBUG", "cgocheck=0")

	utils.Log("initializing GoMod")
	err := utils.Benchmark(hooks.InitHooks, "intializing hooks")
	if err != nil {
		utils.Error("Failed to initialize hooks: %s\n", err.Error())
	}

	utils.Log("Welcome, GoMod's default keybind is ctrl + y")

	domain := il2cpp.GetDomain()
	domain.AttachThread()

}

func main() {

}
