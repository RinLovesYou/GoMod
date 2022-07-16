package menu

import (
	"fmt"

	"github.com/RinLovesYou/imgui-go"
)

type PlayerList struct {
	Players []string
	Shown   bool
	Toggled bool
	speed   float32
}

func NewPlayerList() *PlayerList {
	return &PlayerList{
		Players: []string{},
		Shown:   true,
		Toggled: false,
		speed:   1,
	}
}

var nextPosY = float32(15)

func (p *PlayerList) Render() {
	var colorToUse imgui.Vec4
	if rainbowMenu {
		colorToUse = imgui.Vec4{X: float32(rainbow.Clamped().R), Y: float32(rainbow.Clamped().G), Z: float32(rainbow.Clamped().B), W: 1}
	} else {
		colorToUse = mainColor
	}
	nextPosition := imgui.Vec2{X: 10, Y: nextPosY}

	if p.Toggled {

		p.Shown = p.Toggled
		p.Toggled = false
	}

	if p.Shown {
		imgui.SetNextWindowPosV(nextPosition, imgui.ConditionAlways, imgui.Vec2{})
		imgui.SetNextWindowSizeConstraints(imgui.Vec2{X: 100, Y: 0}, imgui.Vec2{X: 400, Y: 800})
		imgui.SetNextWindowBgAlpha(0.5)
		imgui.BeginV("Player List", &p.Shown, imgui.WindowFlagsNoTitleBar|imgui.WindowFlagsNoMove|imgui.WindowFlagsNoScrollbar)

		imgui.PushFont(DroidSansBeeg)

		imgui.PushStyleColor(imgui.StyleColorText, colorToUse)
		imgui.Text("PlayerList")
		imgui.PopStyleColor()
		imgui.SameLine()
		imgui.Text(fmt.Sprintf("- %d", len(p.Players)))

		imgui.PopFont()

		imgui.Separator()

		if len(p.Players) != 0 {

			imgui.SetNextWindowBgAlpha(0)
			imgui.BeginChild("PlayerListScroll")

			imgui.PushFont(DroidSansMedium)
			for _, name := range p.Players {
				imgui.Text(fmt.Sprintf("- %s", name))
			}
			imgui.PopFont()

			imgui.EndChild()

			imgui.Separator()
		}

		imgui.End()
	}
}
