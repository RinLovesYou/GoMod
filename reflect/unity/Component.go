package unity

import "GoMod/reflect"

var (
	ComponentClass, _ = reflect.UnityEngineCoreModule.GetClass("Component")

	gameObjectProperty, _ = ComponentClass.GetProperty("gameObject")
)

type Component struct {
	Object
}

func (c *Component) GameObject() *GameObject {
	obj, err := gameObjectProperty.GetGet().InvokeObject(c.Il2CppObject)
	if err != nil {
		return nil
	}

	return NewGameObject(obj)
}
