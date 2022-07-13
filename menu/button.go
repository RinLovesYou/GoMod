package menu

import "github.com/RinLovesYou/imgui-go"

type GoButton struct {
	selected bool
	text     string
	menu     func()
}

func NewGoButton(text string, menu func()) *GoButton {
	return &GoButton{
		selected: false,
		text:     text,
		menu:     menu,
	}
}

func (b *GoButton) Render() bool {
	var result bool

	if imgui.Button(b.text) {
		result = true
	}

	return result
}

type TabSystem struct {
	tabs        []*GoButton
	selectedTab byte
}

func NewTabSystem() *TabSystem {
	return &TabSystem{
		tabs:        []*GoButton{},
		selectedTab: 0,
	}
}

func (t *TabSystem) Render() {
	for i, tab := range t.tabs {
		imgui.SameLine()
		if tab.Render() {
			t.selectedTab = byte(i)
			tab.selected = true
		}
	}
}
