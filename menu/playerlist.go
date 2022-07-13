package menu

import (
	"fmt"

	"github.com/RinLovesYou/imgui-go"
)

type PlayerList struct {
	Players []string
	Shown   bool
}

func NewPlayerList() PlayerList {
	return PlayerList{
		Players: []string{},
		Shown:   true,
	}
}

func (p PlayerList) Render() {
	if p.Shown {
		imgui.SetNextWindowPosV(imgui.Vec2{X: 10, Y: 15}, imgui.ConditionAlways, imgui.Vec2{})
		imgui.SetNextWindowSizeConstraints(imgui.Vec2{X: 200, Y: 0}, imgui.Vec2{X: 400, Y: 800})
		imgui.SetNextWindowBgAlpha(0.5)
		imgui.BeginV("Player List", &p.Shown, imgui.WindowFlagsNoTitleBar|imgui.WindowFlagsNoDecoration|imgui.WindowFlagsNoMove|imgui.WindowFlagsAlwaysAutoResize)

		imgui.PushFont(DroidSansBeeg)

		imgui.PushStyleColor(imgui.StyleColorText, mainColor)
		imgui.Text("PlayerList")
		imgui.PopStyleColor()
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("- %d", len(p.Players)))

		imgui.PopFont()

		imgui.Separator()

		imgui.PushFont(DroidSansMedium)
		for _, name := range p.Players {
			imgui.Text(fmt.Sprintf("- %s", name))
		}
		imgui.PopFont()

		if len(p.Players) != 0 {
			imgui.Separator()
		}

		imgui.End()
	}
}
