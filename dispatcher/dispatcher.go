package dispatcher

var DispatcherQueue = make(chan func() error)
var Coroutines = make([]*Coroutine, 0)

func Invoke(fn func() error) {
	if fn == nil {
		return
	}

	DispatcherQueue <- fn
}

type Coroutine struct {
	Args   []any
	Method func(...any) bool
}

func StartCoroutine(method func(...any) bool, args ...any) {
	if method == nil {
		return
	}

	Coroutines = append(Coroutines, &Coroutine{
		Args:   args,
		Method: method,
	})
}
