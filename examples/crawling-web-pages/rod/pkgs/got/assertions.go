package got

import (
	"errors"
	"fmt"
	"github.com/energye/energy/v2/examples/crawling-web-pages/rod/pkgs/got/lib/utils"
	"math"
	"reflect"
	"regexp"
	"runtime"
	"strings"
	"sync/atomic"
)

// Assertions helpers
type Assertions struct {
	Testable

	ErrorHandler AssertionError

	must bool

	desc string
}

// Desc returns a clone with the format description. The description will be printed before the error message.
func (as Assertions) Desc(format string, args ...interface{}) Assertions {
	n := as
	n.desc = fmt.Sprintf(format, args...)
	return n
}

// Must returns a clone with the FailNow enabled. It will exit the current goroutine if the assertion fails.
func (as Assertions) Must() Assertions {
	n := as
	n.must = true
	return n
}

// Eq asserts that x equals y when converted to the same type, such as compare float 1.0 and integer 1 .
// For strict value and type comparison use Assertions.Equal .
// For how comparison works, see [utils.SmartCompare] .
func (as Assertions) Eq(x, y interface{}) {
	as.Helper()
	if utils.SmartCompare(x, y) == 0 {
		return
	}
	as.err(AssertionEq, x, y)
}

// Neq asserts that x not equals y even when converted to the same type.
// For how comparison works, see [utils.SmartCompare] .
func (as Assertions) Neq(x, y interface{}) {
	as.Helper()
	if utils.SmartCompare(x, y) != 0 {
		return
	}

	if reflect.TypeOf(x).Kind() == reflect.TypeOf(y).Kind() {
		as.err(AssertionNeqSame, x, y)
		return
	}
	as.err(AssertionNeq, x, y)
}

// Equal asserts that x equals y.
// For loose type comparison use Assertions.Eq, such as compare float 1.0 and integer 1 .
func (as Assertions) Equal(x, y interface{}) {
	as.Helper()
	if utils.Compare(x, y) == 0 {
		return
	}
	as.err(AssertionEq, x, y)
}

// Gt asserts that x is greater than y.
// For how comparison works, see [utils.SmartCompare] .
func (as Assertions) Gt(x, y interface{}) {
	as.Helper()
	if utils.SmartCompare(x, y) > 0 {
		return
	}
	as.err(AssertionGt, x, y)
}

// Gte asserts that x is greater than or equal to y.
// For how comparison works, see [utils.SmartCompare] .
func (as Assertions) Gte(x, y interface{}) {
	as.Helper()
	if utils.SmartCompare(x, y) >= 0 {
		return
	}
	as.err(AssertionGte, x, y)
}

// Lt asserts that x is less than y.
// For how comparison works, see [utils.SmartCompare] .
func (as Assertions) Lt(x, y interface{}) {
	as.Helper()
	if utils.SmartCompare(x, y) < 0 {
		return
	}
	as.err(AssertionLt, x, y)
}

// Lte asserts that x is less than or equal to b.
// For how comparison works, see [utils.SmartCompare] .
func (as Assertions) Lte(x, y interface{}) {
	as.Helper()
	if utils.SmartCompare(x, y) <= 0 {
		return
	}
	as.err(AssertionLte, x, y)
}

// InDelta asserts that x and y are within the delta of each other.
// For how comparison works, see [utils.SmartCompare] .
func (as Assertions) InDelta(x, y interface{}, delta float64) {
	as.Helper()
	if math.Abs(utils.SmartCompare(x, y)) <= delta {
		return
	}
	as.err(AssertionInDelta, x, y, delta)
}

// True asserts that x is true.
func (as Assertions) True(x bool) {
	as.Helper()
	if x {
		return
	}
	as.err(AssertionTrue)
}

// False asserts that x is false.
func (as Assertions) False(x bool) {
	as.Helper()
	if !x {
		return
	}
	as.err(AssertionFalse)
}

// Nil asserts that the last item in args is nilable and nil
func (as Assertions) Nil(args ...interface{}) {
	as.Helper()
	if len(args) == 0 {
		as.err(AssertionNoArgs)
		return
	}
	last := args[len(args)-1]
	if _, yes := isNil(last); yes {
		return
	}
	as.err(AssertionNil, last, args)
}

// NotNil asserts that the last item in args is nilable and not nil
func (as Assertions) NotNil(args ...interface{}) {
	as.Helper()
	if len(args) == 0 {
		as.err(AssertionNoArgs)
		return
	}
	last := args[len(args)-1]

	if last == nil {
		as.err(AssertionNotNil, last, args)
		return
	}

	nilable, yes := isNil(last)
	if !nilable {
		as.err(AssertionNotNilable, last, args)
		return
	}

	if yes {
		as.err(AssertionNotNilableNil, last, args)
	}
}

// Zero asserts x is zero value for its type.
func (as Assertions) Zero(x interface{}) {
	as.Helper()
	if reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface()) {
		return
	}
	as.err(AssertionZero, x)
}

// NotZero asserts that x is not zero value for its type.
func (as Assertions) NotZero(x interface{}) {
	as.Helper()
	if reflect.DeepEqual(x, reflect.Zero(reflect.TypeOf(x)).Interface()) {
		as.err(AssertionNotZero, x)
	}
}

// Regex asserts that str matches the regex pattern
func (as Assertions) Regex(pattern, str string) {
	as.Helper()
	if regexp.MustCompile(pattern).MatchString(str) {
		return
	}
	as.err(AssertionRegex, pattern, str)
}

// Has asserts that container has item.
// The container can be a string, []byte, slice, array, or map.
// For how comparison works, see [utils.SmartCompare] .
func (as Assertions) Has(container, item interface{}) {
	as.Helper()

	if c, ok := container.(string); ok && hasStr(c, item) {
		return
	} else if c, ok := container.([]byte); ok && hasStr(string(c), item) {
		return
	}

	cv := reflect.Indirect(reflect.ValueOf(container))
	switch cv.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < cv.Len(); i++ {
			if utils.SmartCompare(cv.Index(i).Interface(), item) == 0 {
				return
			}
		}
	case reflect.Map:
		keys := cv.MapKeys()
		for _, k := range keys {
			if utils.SmartCompare(cv.MapIndex(k).Interface(), item) == 0 {
				return
			}
		}
	}

	as.err(AssertionHas, container, item)
}

// Len asserts that the length of list equals l
func (as Assertions) Len(list interface{}, l int) {
	as.Helper()
	actual := reflect.ValueOf(list).Len()
	if actual == l {
		return
	}
	as.err(AssertionLen, actual, l, list)
}

// Err asserts that the last item in args is error
func (as Assertions) Err(args ...interface{}) {
	as.Helper()
	if len(args) == 0 {
		as.err(AssertionNoArgs)
		return
	}
	last := args[len(args)-1]
	if err, _ := last.(error); err != nil {
		return
	}
	as.err(AssertionErr, last, args)
}

// E is a shortcut for Must().Nil(args...)
func (as Assertions) E(args ...interface{}) {
	as.Helper()
	as.Must().Nil(args...)
}

// Panic executes fn and asserts that fn panics
func (as Assertions) Panic(fn func()) (val interface{}) {
	as.Helper()

	defer func() {
		as.Helper()

		val = recover()
		if val == nil {
			as.err(AssertionPanic, fn)
		}
	}()

	fn()

	return
}

// Is asserts that x is kind of y, it uses reflect.Kind to compare.
// If x and y are both error type, it will use errors.Is to compare.
func (as Assertions) Is(x, y interface{}) {
	as.Helper()

	if x == nil && y == nil {
		return
	}

	if ae, ok := x.(error); ok {
		if be, ok := y.(error); ok {
			if ae == be {
				return
			}

			if errors.Is(ae, be) {
				return
			}
			as.err(AssertionIsInChain, x, y)
			return
		}
	}

	at := reflect.TypeOf(x)
	bt := reflect.TypeOf(y)
	if x != nil && y != nil && at.Kind() == bt.Kind() {
		return
	}
	as.err(AssertionIsKind, x, y)
}

// Count asserts that the returned function will be called n times
func (as Assertions) Count(n int) func() {
	as.Helper()
	count := int64(0)

	as.Cleanup(func() {
		c := int(atomic.LoadInt64(&count))
		if c != n {
			as.Helper()
			as.err(AssertionCount, n, c)
		}
	})

	return func() {
		atomic.AddInt64(&count, 1)
	}
}

func (as Assertions) err(t AssertionErrType, details ...interface{}) {
	as.Helper()

	if as.desc != "" {
		as.Logf("%s", as.desc)
	}

	// TODO: we should take advantage of the Helper function
	_, f, l, _ := runtime.Caller(2)
	c := &AssertionCtx{
		Type:    t,
		Details: details,
		File:    f,
		Line:    l,
	}

	as.Logf("%s", as.ErrorHandler.Report(c))

	if as.must {
		as.FailNow()
		return
	}

	as.Fail()
}

// the first return value is true if x is nilable
func isNil(x interface{}) (bool, bool) {
	if x == nil {
		return true, true
	}

	val := reflect.ValueOf(x)
	k := val.Kind()
	nilable := k == reflect.Chan ||
		k == reflect.Func ||
		k == reflect.Interface ||
		k == reflect.Map ||
		k == reflect.Ptr ||
		k == reflect.Slice

	if nilable {
		return true, val.IsNil()
	}

	return false, false
}

func hasStr(c string, item interface{}) bool {
	if it, ok := item.(string); ok {
		if strings.Contains(c, it) {
			return true
		}
	} else if it, ok := item.([]byte); ok {
		if strings.Contains(c, string(it)) {
			return true
		}
	} else if it, ok := item.(rune); ok {
		if strings.ContainsRune(c, it) {
			return true
		}
	}
	return false
}
