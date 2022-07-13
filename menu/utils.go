package menu

import "github.com/RinLovesYou/imgui-go"

func inlineText(textFunc ...func()) {
	for _, f := range textFunc {
		f()
		imgui.SameLine()
	}
}
