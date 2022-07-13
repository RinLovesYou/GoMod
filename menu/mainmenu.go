package menu

import (
	"GoMod/reflect/unity"
	"GoMod/utils"
	"fmt"

	"github.com/RinLovesYou/imgui-go"
)

type MainMenu struct {
	Shown       bool
	ShowConsole StringBool
}

func NewMainMenu() *MainMenu {
	return &MainMenu{
		Shown:       false,
		ShowConsole: StringBool(true),
	}
}

type StringBool bool

func (s StringBool) String() string {
	if s {
		return "ON"
	}

	return "OFF"
}

var tabSys = NewTabSystem()

func (m *MainMenu) Init() {
	tabSys.tabs = append(tabSys.tabs, NewGoButton("Home", func() {
		m.homePageRender()
	}))

	tabSys.tabs = append(tabSys.tabs, NewGoButton("About", func() {
		imgui.Text("GoMod v0.1.0 by RinLovesYou :)")
	}))
}

func (m *MainMenu) Render() {
	if m.ShowConsole {

		imgui.SetNextWindowSizeV(imgui.Vec2{X: 400, Y: 200}, imgui.ConditionFirstUseEver)
		imgui.SetNextWindowBgAlpha(0.75)
		imgui.BeginV("Console", &m.Shown, imgui.WindowFlagsNoTitleBar|imgui.WindowFlagsNoScrollbar)
		imgui.PushFont(DroidSansMedium)
		imgui.SetNextWindowBgAlpha(0)
		imgui.BeginChildV("ConsoleScrollingRegion", imgui.Vec2{X: 0, Y: -24}, false, imgui.WindowFlagsHorizontalScrollbar|imgui.WindowFlagsNoScrollbar)
		for _, msg := range utils.LogBuffer {
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{0, 0})
			imgui.Text("[")
			imgui.PushStyleColor(imgui.StyleColorText, imgui.Vec4{X: 0.09, Y: 0.976, Z: 0.043, W: 1.0})
			imgui.SameLine()
			imgui.Text(msg.TimeStamp)
			imgui.PopStyleColor()
			imgui.SameLine()
			imgui.Text("] [")

			imgui.PushStyleColor(imgui.StyleColorText, mainColor)
			imgui.SameLine()
			imgui.Text(msg.Prefix)
			imgui.PopStyleColor()
			imgui.SameLine()
			imgui.Text("] ")

			imgui.SameLine()
			imgui.Text(msg.Message)
			imgui.PopStyleVar()
			imgui.SetScrollHereY(1)
		}
		imgui.SetScrollHereY(1)
		imgui.EndChild()
		imgui.PopFont()

		imgui.Separator()

		framerate := imgui.CurrentIO().Framerate()
		var color imgui.Vec4

		if framerate <= 20 {
			color = getColor(1, 0, 0, 1)
		}

		if framerate > 20 && framerate < 45 {
			color = getColor(1, 1, 0, 1)
		}

		if framerate >= 45 {
			color = getColor(0, 1, 0, 1)
		}

		imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{X: 0, Y: 0})
		imgui.PushStyleColor(imgui.StyleColorText, color)
		imgui.Text(fmt.Sprintf("%.3f", 1000/framerate))
		imgui.PopStyleColor()

		imgui.SameLine()
		imgui.Text(" ms/frame ")

		imgui.SameLine()
		imgui.Text("(")
		imgui.SameLine()
		imgui.PushStyleColor(imgui.StyleColorText, color)
		imgui.Text(fmt.Sprintf("%.0f", framerate))
		imgui.PopStyleColor()
		imgui.SameLine()
		imgui.Text(")")

		imgui.SameLine()
		imgui.Text(" FPS")
		imgui.PopStyleVar()

		imgui.End()
	}

	if !m.Shown {
		return
	}
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 400, Y: 200}, imgui.ConditionFirstUseEver)
	imgui.BeginV("Main Menu", &m.Shown, imgui.WindowFlagsNoTitleBar)

	// Logo
	imgui.PushFont(ImpactBeeg)
	imgui.PushStyleColor(imgui.StyleColorText, Cyan)
	imgui.Text("GoMod")
	imgui.PopStyleColor()
	imgui.PopFont()

	// Version
	imgui.SameLine()
	imgui.PushStyleColor(imgui.StyleColorText, Gray)
	imgui.Text("v0.1.0")
	imgui.PopStyleColor()

	tabSys.Render()

	imgui.Separator()

	if len(tabSys.tabs) >= int(tabSys.selectedTab) {
		tabSys.tabs[tabSys.selectedTab].menu()
	}

	imgui.End()
}

var targetFramerate int32 = 120

func (m *MainMenu) homePageRender() {
	if imgui.CollapsingHeader("General") {
		if imgui.Selectable("Console") {
			m.ShowConsole = !m.ShowConsole
		}

		imgui.SameLine()
		color := getColor(1, 0, 0, 1)
		if m.ShowConsole {
			color = getColor(0, 1, 0, 1)
		}
		imgui.PushStyleColor(imgui.StyleColorText, color)
		imgui.Text(fmt.Sprintf("%s", m.ShowConsole))
		imgui.PopStyleColor()

		imgui.DragIntV("Target Framerate", &targetFramerate, 1, 10, 240, "%.0f", imgui.SliderFlagsNone)

		imgui.SameLine()
		if imgui.Button("Set") {
			unity.SetTargetFramerate(targetFramerate)
			utils.Log("Target Framerate set to %d", targetFramerate)
		}

	}

	imgui.CollapsingHeader("Placeholder")
	imgui.CollapsingHeader("Placeholder")
}
