package got

import (
	"reflect"
	"runtime/debug"
)

// Only run tests with it
type Only struct{}

// Skip the current test
type Skip struct{}

// Each runs each exported method Fn on type Ctx as a subtest of t.
// The iteratee can be a struct Ctx or:
//
//	iteratee(t Testable) (ctx Ctx)
//
// Each Fn will be called like:
//
//	ctx.Fn()
//
// If iteratee is Ctx, its G field will be set to New(t) for each test.
// Any Fn that has the same name with the embedded one will be ignored.
func Each(t Testable, iteratee interface{}) (count int) {
	t.Helper()

	itVal := normalizeIteratee(t, iteratee)

	ctxType := itVal.Type().Out(0)

	methods := filterMethods(ctxType)

	runVal := reflect.ValueOf(t).MethodByName("Run")
	cbType := runVal.Type().In(1)

	for _, m := range methods {
		// because the callback is in another goroutine, we create closures for each loop
		method := m

		runVal.Call([]reflect.Value{
			reflect.ValueOf(method.Name),
			reflect.MakeFunc(cbType, func(args []reflect.Value) []reflect.Value {
				t := args[0].Interface().(Testable)
				doSkip(t, method)
				count++
				res := itVal.Call(args)
				return callMethod(t, method, res[0])
			}),
		})
	}
	return
}

func normalizeIteratee(t Testable, iteratee interface{}) reflect.Value {
	t.Helper()

	if iteratee == nil {
		t.Logf("iteratee shouldn't be nil")
		t.FailNow()
	}

	itVal := reflect.ValueOf(iteratee)
	itType := itVal.Type()
	fail := true

	switch itType.Kind() {
	case reflect.Func:
		if itType.NumIn() != 1 || itType.NumOut() != 1 {
			break
		}
		try(func() {
			_ = reflect.New(itType.In(0).Elem()).Interface().(Testable)
			fail = false
		})

	case reflect.Struct:
		fnType := reflect.FuncOf([]reflect.Type{reflect.TypeOf(t)}, []reflect.Type{itType}, false)
		structVal := itVal
		itVal = reflect.MakeFunc(fnType, func(args []reflect.Value) []reflect.Value {
			sub := args[0].Interface().(Testable)
			as := reflect.ValueOf(New(sub))

			c := reflect.New(itType).Elem()
			c.Set(structVal)
			try(func() { c.FieldByName("G").Set(as) })

			return []reflect.Value{c}
		})
		fail = false
	}

	if fail {
		t.Logf("iteratee <%v> should be a struct or <func(got.Testable) Ctx>", itType)
		t.FailNow()
	}
	return itVal
}

func callMethod(t Testable, method reflect.Method, receiver reflect.Value) []reflect.Value {
	args := make([]reflect.Value, method.Type.NumIn())
	args[0] = receiver

	for i := 1; i < len(args); i++ {
		args[i] = reflect.New(method.Type.In(i)).Elem()
	}

	defer func() {
		if err := recover(); err != nil {
			t.Logf("[panic] %v\n%s", err, debug.Stack())
			t.Fail()
		}
	}()

	method.Func.Call(args)

	return []reflect.Value{}
}

func filterMethods(typ reflect.Type) []reflect.Method {
	embedded := map[string]struct{}{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Anonymous {
			for j := 0; j < field.Type.NumMethod(); j++ {
				embedded[field.Type.Method(j).Name] = struct{}{}
			}
		}
	}

	methods := []reflect.Method{}
	onlyList := []reflect.Method{}
	for i := 0; i < typ.NumMethod(); i++ {
		method := typ.Method(i)
		if _, has := embedded[method.Name]; has {
			continue
		}

		if method.Type.NumIn() > 1 && method.Type.In(1) == reflect.TypeOf(Only{}) {
			onlyList = append(onlyList, method)
		}

		methods = append(methods, method)
	}

	if len(onlyList) > 0 {
		return onlyList
	}

	return methods
}

func doSkip(t Testable, method reflect.Method) {
	if method.Type.NumIn() > 1 && method.Type.In(1) == reflect.TypeOf(Skip{}) {
		t.SkipNow()
	}
}

func try(fn func()) {
	defer func() {
		_ = recover()
	}()
	fn()
}
