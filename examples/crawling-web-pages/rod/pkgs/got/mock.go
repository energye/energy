package got

import (
	"fmt"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/got/lib/utils"
	"reflect"
	"sync"
)

// Mock helper for interface stubbing
type Mock struct {
	lock sync.Mutex

	fallback reflect.Value

	stubs map[string]interface{}
}

// Fallback the methods that are not stubbed to fb.
func (m *Mock) Fallback(fb interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	m.fallback = reflect.ValueOf(fb)
}

// Stub the target's method with fn
func (m *Mock) Stub(method string, fn interface{}) {
	m.lock.Lock()
	defer m.lock.Unlock()

	if m.stubs == nil {
		m.stubs = map[string]interface{}{}
	}

	m.stubs[method] = fn
}

// StubFn helper
func (m *Mock) StubFn(target interface{}, method string, fn func(args []reflect.Value) []reflect.Value) {
	t := utils.MethodType(target, method)

	f := reflect.MakeFunc(t, func(args []reflect.Value) []reflect.Value {
		res := fn(args)

		out := []reflect.Value{}
		for i := 0; i < t.NumOut(); i++ {
			v := reflect.New(t.Out(i)).Elem()
			if res[i].IsValid() {
				v.Set(res[i])
			}
			out = append(out, v)
		}

		return out
	}).Interface()

	m.Stub(method, f)
}

// Stop the stub
func (m *Mock) Stop(method string) {
	m.lock.Lock()
	defer m.lock.Unlock()

	delete(m.stubs, method)
}

// Proxy a interface
func (m *Mock) Proxy(method string) interface{} {
	m.lock.Lock()
	defer m.lock.Unlock()

	if fn, has := m.stubs[method]; has {
		return fn
	}

	if !m.fallback.IsValid() {
		panic("you should specify the got.Mock.Origin")
	}

	methodVal := m.fallback.MethodByName(method)
	if !methodVal.IsValid() {
		panic(m.fallback.Type().String() + " doesn't have method: " + method)
	}

	return m.fallback.MethodByName(method).Interface()
}

// StubOn utils
type StubOn struct {
	when []*StubWhen
}

// StubWhen utils
type StubWhen struct {
	on  *StubOn
	in  []interface{}
	ret *StubReturn
}

// StubReturn utils
type StubReturn struct {
	on    *StubOn
	out   []reflect.Value
	times *StubTimes
}

// StubTimes utils
type StubTimes struct {
	count int
}

// On helper
func (m *Mock) On(target interface{}, method string) *StubOn {
	s := &StubOn{
		when: []*StubWhen{},
	}

	eq := func(in, arg []interface{}) bool {
		for i := 0; i < len(in); i++ {
			if in[i] != Any && utils.Compare(in[i], arg[i]) != 0 {
				return false
			}
		}
		return true
	}

	m.StubFn(target, method, func(args []reflect.Value) []reflect.Value {
		argsIt := utils.ToInterfaces(args)
		for _, when := range s.when {
			if eq(when.in, argsIt) {
				when.ret.times.count--
				if when.ret.times.count == 0 {
					m.Stop(method)
				}

				return when.ret.out
			}
		}
		panic(fmt.Sprintf("No got.StubOn.When matches: %#v", argsIt))
	})

	return s
}

// Any intput
var Any = struct{}{}

// When input args matches
func (s *StubOn) When(in ...interface{}) *StubWhen {
	w := &StubWhen{on: s, in: in}
	s.when = append(s.when, w)
	return w
}

// Return values for each stub
func (s *StubWhen) Return(out ...interface{}) *StubReturn {
	r := &StubReturn{on: s.on, out: utils.ToValues(out)}
	r.Times(0)
	s.ret = r
	return r
}

// Times specifies how how many stubs before stop, if n <= 0 it will never stop.
func (s *StubReturn) Times(n int) *StubOn {
	t := &StubTimes{count: n}
	s.times = t
	return s.on
}

// Once specifies stubs only once before stop
func (s *StubReturn) Once() *StubOn {
	return s.Times(1)
}
