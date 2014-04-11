package contract

import "fmt"

var SuppressPanic bool = false

func shoutAndPanic(descr string) {
	err := fmt.Errorf(descr)
	if SuppressPanic {
		fmt.Println(err)
	} else {
		panic(err)
	}
}

func shoutAndPanicIf(orly bool, descr string) {
	if orly {
		shoutAndPanic(descr)
	}
}

func Require(success bool, descr string) {
	if !success {
		descr := fmt.Sprintf("Input contract not met: %s", descr)
		shoutAndPanic(descr)
	}
}

func MustBeTrue(success bool, descr string) {
	shoutAndPanicIf(success == true, descr)
}

func MustNotBeNil(v interface{}, descr string) {
	shoutAndPanicIf(v == nil, descr)
}

func MustBeNil(v interface{}, descr string) {
	shoutAndPanicIf(v != nil, descr)
}

func RequireArg(v interface{}) {
	MustNotBeNil(v, "nil passed instead of real value")
}

func MustBeFalse(success bool, descr string) {
	shoutAndPanicIf(success == false, descr)
}

func RequireNoError(err error, intent string) {
	if err != nil {
		descr := fmt.Sprintf("Required no error to do: %s, but got: %s", intent, err.Error())
		shoutAndPanic(descr)
	}
}

func Guarantee(success bool, descr string) {
	if !success {
		descr := fmt.Sprintf("Output contract not met: %s", descr)
		shoutAndPanic(descr)
	}
}
