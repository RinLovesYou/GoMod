package hooks

import "fmt"

func InitHooks() error {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic occurred:", err)
		}
	}()

	err := initImguiHooks()
	if err != nil {
		return err
	}

	err = initUpdateHook()
	if err != nil {
		return err
	}

	return initJoinHooks()
}
