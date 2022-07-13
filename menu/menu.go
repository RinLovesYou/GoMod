package menu

var (
	GoModMenu      = NewMainMenu()
	PlayerListMenu = NewPlayerList()
)

func Init() {
	GoModMenu.Init()
}

func Render() {
	GoModMenu.Render()
	PlayerListMenu.Render()
}
