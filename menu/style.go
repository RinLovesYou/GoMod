package menu

import (
	_ "embed"

	"github.com/RinLovesYou/imgui-go"
)

var mainColor = getColor(0.960, 0.498, 0.952, 1)
var mainColorInactive = getColor(0.960, 0.498, 0.952, 0.5)
var mainColorTransparent = getColor(0.960, 0.498, 0.952, 0.35)

var White = getColor(0.80, 0.80, 0.83, 1.00)
var Gray = getColor(0.34, 0.33, 0.39, 1.00)
var Background = getColor(0.06, 0.05, 0.07, 1.00)
var ChildBackground = getColor(0.07, 0.07, 0.09, 1.00)
var FrameBg = getColor(0.10, 0.09, 0.12, 1.00)
var FrameBgActive = getColor(0.56, 0.56, 0.58, 1.00)
var ScrollGrab = getColor(0.80, 0.80, 0.83, 0.31)
var PlotLines = getColor(0.40, 0.39, 0.38, 0.63)
var PlotLinesHover = getColor(0.25, 1.00, 0.00, 1.00)
var Cyan = getColor(0.003, 0.976, 0.917, 1)

//go:embed fonts/DroidSans.ttf
var font []byte

var DroidSans imgui.Font
var DroidSansMedium imgui.Font
var DroidSansBeeg imgui.Font

var Impact imgui.Font
var ImpactMedium imgui.Font
var ImpactBeeg imgui.Font

func SetStyle() {
	style := imgui.CurrentStyle()
	io := imgui.CurrentIO()

	DroidSans = io.Fonts().AddFontFromMemoryTTF(font, 14)
	io.Fonts().AddFontDefault()

	Impact = io.Fonts().AddFontFromFileTTF("C:/Windows/Fonts/impact.ttf", 14)
	ImpactMedium = io.Fonts().AddFontFromFileTTF("C:/Windows/Fonts/impact.ttf", 16)
	ImpactBeeg = io.Fonts().AddFontFromFileTTF("C:/Windows/Fonts/impact.ttf", 24)
	DroidSansMedium = io.Fonts().AddFontFromMemoryTTF(font, 16)
	DroidSansBeeg = io.Fonts().AddFontFromMemoryTTF(font, 20)

	style.SetColor(imgui.StyleColorText, White)
	style.SetColor(imgui.StyleColorTextDisabled, Gray)
	style.SetColor(imgui.StyleColorWindowBg, Background)
	style.SetColor(imgui.StyleColorChildBg, ChildBackground)
	style.SetColor(imgui.StyleColorPopupBg, ChildBackground)
	style.SetColor(imgui.StyleColorBorder, getColor(0.80, 0.80, 0.83, 0))
	style.SetColor(imgui.StyleColorBorderShadow, getColor(0, 0, 0, 0))
	style.SetColor(imgui.StyleColorFrameBg, FrameBg)
	style.SetColor(imgui.StyleColorFrameBgHovered, Gray)
	style.SetColor(imgui.StyleColorFrameBgActive, FrameBgActive)
	style.SetColor(imgui.StyleColorTitleBg, FrameBg)
	style.SetColor(imgui.StyleColorTitleBgActive, ChildBackground)
	style.SetColor(imgui.StyleColorTitleBgCollapsed, getColor(1.00, 0.98, 0.95, 0.75))
	style.SetColor(imgui.StyleColorMenuBarBg, FrameBg)
	style.SetColor(imgui.StyleColorScrollbarBg, FrameBg)
	style.SetColor(imgui.StyleColorScrollbarGrab, ScrollGrab)
	style.SetColor(imgui.StyleColorScrollbarGrabHovered, FrameBgActive)
	style.SetColor(imgui.StyleColorScrollbarGrabActive, Background)
	style.SetColor(imgui.StyleColorCheckMark, ScrollGrab)
	style.SetColor(imgui.StyleColorSliderGrab, ScrollGrab)
	style.SetColor(imgui.StyleColorSliderGrabActive, Background)
	style.SetColor(imgui.StyleColorButton, Background)
	style.SetColor(imgui.StyleColorButtonHovered, Gray)
	style.SetColor(imgui.StyleColorButtonActive, Gray)
	style.SetColor(imgui.StyleColorHeader, FrameBg)
	style.SetColor(imgui.StyleColorHeaderHovered, FrameBgActive)
	style.SetColor(imgui.StyleColorHeaderActive, Background)
	style.SetColor(imgui.StyleColorSeparator, getColor(0.43, 0.43, 0.50, 0.50))
	style.SetColor(imgui.StyleColorSeparatorHovered, getColor(0.72, 0.72, 0.72, 0.78))
	style.SetColor(imgui.StyleColorSeparatorActive, getColor(0.51, 0.51, 0.51, 1.00))
	style.SetColor(imgui.StyleColorResizeGrip, getColor(0, 0, 0, 0))
	style.SetColor(imgui.StyleColorResizeGripHovered, FrameBgActive)
	style.SetColor(imgui.StyleColorResizeGripActive, Background)
	style.SetColor(imgui.StyleColorPlotLines, PlotLines)
	style.SetColor(imgui.StyleColorPlotLinesHovered, PlotLinesHover)
	style.SetColor(imgui.StyleColorPlotHistogram, PlotLines)
	style.SetColor(imgui.StyleColorPlotHistogramHovered, PlotLinesHover)
	style.SetColor(imgui.StyleColorTextSelectedBg, getColor(0.87, 0.87, 0.87, 0.35))
	style.SetColor(imgui.StyleColorModalWindowDarkening, getColor(0.80, 0.80, 0.80, 0.35))
	style.SetColor(imgui.StyleColorDragDropTarget, getColor(1.00, 1.00, 0.00, 0.90))
	style.SetColor(imgui.StyleColorNavHighlight, getColor(0.60, 0.60, 0.60, 1.00))
	style.SetColor(imgui.StyleColorNavWindowingHighlight, getColor(1.00, 1.00, 1.00, 0.70))
	style.SetColor(imgui.StyleColorTabActive, mainColor)
	style.SetColor(imgui.StyleColorTabUnfocused, mainColorTransparent)
	style.SetColor(imgui.StyleColorTabHovered, mainColorInactive)
	style.SetColor(imgui.StyleColorTabUnfocusedActive, mainColor)
	style.SetColor(imgui.StyleColorTab, mainColorTransparent)

	style.SetWindowPadding(imgui.Vec2{X: 15, Y: 15})
	style.SetWindowRounding(5)
	style.SetFramePadding(imgui.Vec2{X: 5, Y: 5})
	style.SetFrameRounding(4)
	style.SetItemSpacing(imgui.Vec2{X: 12, Y: 8})
	style.SetItemInnerSpacing(imgui.Vec2{X: 8, Y: 6})
	style.SetIndentSpacing(25)
	style.SetScrollbarSize(15)
	style.SetScrollbarRounding(9)
	style.SetGrabMinSize(5)
	style.SetGrabRounding(3)
	style.SetWindowTitleAlign(imgui.Vec2{X: 0.5, Y: 0.5})
}

func getColor(r, g, b, a float32) imgui.Vec4 {
	return imgui.Vec4{X: r, Y: g, Z: b, W: a}
}
