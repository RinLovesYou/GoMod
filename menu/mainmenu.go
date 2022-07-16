package menu

import (
	"GoMod/keybinds"
	"GoMod/reflect/unity"
	"GoMod/utils"
	"fmt"

	"github.com/RinLovesYou/imgui-go"
	"github.com/lucasb-eyer/go-colorful"
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

var rainbowMenu StringBool

type StringBool bool

func (s StringBool) String() string {
	if s {
		return "ON"
	}

	return "OFF"
}

var tabSys = NewTabSystem()

func (m *MainMenu) Init() {
	tabSys.tabs = append(tabSys.tabs, NewGoButton("Home", m.homePageRender))

	tabSys.tabs = append(tabSys.tabs, NewGoButton("About", func() {
		imgui.Text("GoMod v0.1.0 by RinLovesYou :)")
	}))
}

var rainbow = colorful.Hsv(0, 1, 1)
var green = getColor(0, 1, 0, 1)
var colorFloat float64 = 0

func (m *MainMenu) Render() {
	colorFloat++
	if colorFloat > 360 {
		colorFloat = 0
	}
	rainbow = colorful.Hsv(colorFloat, 1, 1)

	var colorToUse imgui.Vec4
	if rainbowMenu {
		colorToUse = imgui.Vec4{X: float32(rainbow.Clamped().R), Y: float32(rainbow.Clamped().G), Z: float32(rainbow.Clamped().B), W: 1}
	} else {
		colorToUse = green
	}
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
			imgui.PushStyleColor(imgui.StyleColorText, colorToUse)
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
			if utils.JustLogged {
				imgui.SetScrollHereY(1)
			}
		}
		if utils.JustLogged {
			imgui.SetScrollHereY(1)
			utils.JustLogged = false
		}
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

	if rainbowMenu {
		colorToUse = imgui.Vec4{X: float32(rainbow.Clamped().R), Y: float32(rainbow.Clamped().G), Z: float32(rainbow.Clamped().B), W: 1}
	} else {
		colorToUse = Cyan
	}

	// Logo
	imgui.PushFont(ImpactBeeg)
	imgui.PushStyleColor(imgui.StyleColorText, colorToUse)
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
	imgui.BeginChildV("HomePage", imgui.Vec2{X: 0, Y: 0}, false, imgui.WindowFlagsHorizontalScrollbar)
	if imgui.CollapsingHeader("General") {
		//being: console button
		if imgui.Selectable("Console") {
			m.ShowConsole = !m.ShowConsole
		}

		imgui.SameLine()
		color := getColor(1, 0, 0, 1)
		if m.ShowConsole {
			color = getColor(0, 1, 0, 1)
			if rainbowMenu {
				color = imgui.Vec4{X: float32(rainbow.Clamped().R), Y: float32(rainbow.Clamped().G), Z: float32(rainbow.Clamped().B), W: 1}
			}
		}
		imgui.PushStyleColor(imgui.StyleColorText, color)
		imgui.Text(m.ShowConsole.String())
		imgui.PopStyleColor()

		//end: console button

		//begin: rainbow button
		if imgui.Selectable("Rainbow") {
			rainbowMenu = !rainbowMenu
		}

		imgui.SameLine()
		color = getColor(1, 0, 0, 1)
		if rainbowMenu {
			color = imgui.Vec4{X: float32(rainbow.Clamped().R), Y: float32(rainbow.Clamped().G), Z: float32(rainbow.Clamped().B), W: 1}
		}
		imgui.PushStyleColor(imgui.StyleColorText, color)
		imgui.Text(rainbowMenu.String())
		imgui.PopStyleColor()
		//end: rainbow button

		//begin: framerate slider
		imgui.DragIntV("Target Framerate", &targetFramerate, 1, 10, 240, "%.0f", imgui.SliderFlagsNone)

		imgui.SameLine()
		if imgui.Button("Set") {
			unity.SetTargetFramerate(int(targetFramerate + 2))
			utils.Log("Target Framerate set to %d", targetFramerate)
		}
		//end: framerate slider

		//begin: fly speed slider
		imgui.DragFloatV("Fly Speed", &keybinds.FlySpeed, 0.1, 0.1, 10, "%.1f", imgui.SliderFlagsNone)
		//end: fly speed slider
	}

	imgui.CollapsingHeader("Placeholder")
	imgui.CollapsingHeader("Placeholder")
	imgui.EndChild()
}
