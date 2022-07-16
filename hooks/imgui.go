package hooks

import (
	"GoMod/keybinds"
	"GoMod/menu"
	"GoMod/utils"
	"unsafe"

	"github.com/RinLovesYou/imgui-go"
	"github.com/RinLovesYou/kiergo"
)

func initImguiHooks() error {
	return kiergo.Hook(OnPresent, WndProc)
}

var gInitialized bool
var device unsafe.Pointer
var context unsafe.Pointer
var window unsafe.Pointer

func OnPresent(pSwapChain unsafe.Pointer, SyncInterval, FlagsT uint32) error {
	if !gInitialized {
		var err error

		device, err = kiergo.SetupValuesGetDevice(pSwapChain)
		if err != nil {
			return err
		}

		context, err = kiergo.GetContext()
		if err != nil {
			return err
		}

		window, err = kiergo.GetGameWindow()
		if err != nil {
			return err
		}

		imgui.CreateContext(nil)

		menu.SetStyle()
		utils.Benchmark(func() error {
			menu.Init()
			return nil
		}, "initializing menu")

		imgui.Win32Init(window)
		imgui.Dx11Init(device, context)

		gInitialized = true
	}

	imgui.Dx11NewFrame()
	imgui.Win32NewFrame()
	imgui.NewFrame()

	menu.Render()

	imgui.Render()
	kiergo.SetRenderTargets()
	imgui.Dx11RenderDrawData(imgui.RenderedDrawData())

	return nil
}

var param uint64

func WndProc(hwnd unsafe.Pointer, msg uint32, wparam, lparam unsafe.Pointer) error {
	imgui.Win32WndProcHandler(hwnd, msg, wparam, lparam)

	if wparam != nil {
		param = *(*uint64)(wparam)
		if msg == 0x0100 || msg == 0x0104 { //keydown/syskeydown
			if param < 256 {
				switch param {
				case 'Y', 'y':
					if imgui.CurrentIO().KeyCtrlPressed() {
						menu.GoModMenu.Shown = !menu.GoModMenu.Shown
						imgui.CurrentIO().SetMouseDrawCursor(menu.GoModMenu.Shown)
					}
				case 'P', 'p':
					if imgui.CurrentIO().KeyCtrlPressed() {
						if !menu.PlayerListMenu.Toggled {
							menu.PlayerListMenu.Toggled = true
						}
					}
				}
			}

			keybinds.UpdateFly(param)
		}
	}

	return nil
}
