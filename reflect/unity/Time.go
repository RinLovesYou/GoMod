package unity

import "GoMod/reflect"

var (
	TimeClass, _ = reflect.UnityEngineCoreModule.GetClass("Time")

	deltaTimeProperty, _ = TimeClass.GetProperty("deltaTime")
)

func DeltaTime() float32 {
	res, err := deltaTimeProperty.GetGet().Invoke()
	if err != nil {
		return 0
	}

	return *(*float32)(res.Unbox())
}
