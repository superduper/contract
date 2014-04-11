package contract

import "fmt"

var SuppressPanic bool = false

func shoutAndPanic(descrf ...string) {
	var err error
	if len(descrf) == 0 {
		err = fmt.Errorf("failed to meet the contract")
	} else {
		if len(descrf) == 1 {
			err = fmt.Errorf(descrf[0])
		} else {
			fmtargs := make([]interface{}, len(descrf[1:]))
			for i, v := range descrf[1:] {
				fmtargs[i] = interface{}(v)
			}
			err = fmt.Errorf(descrf[0], fmtargs...)
		}
	}
	if SuppressPanic {
		fmt.Println(err)
	} else {
		panic(err)
	}
}

func shoutAndPanicIf(orly bool, descrf ...string) {
	if orly {
		shoutAndPanic(descrf...)
	}
}

func Require(success bool, descrf ...string) {
	shoutAndPanicIf(!success, descrf...)
}

func MustBeTrue(success bool, descrf ...string) {
	shoutAndPanicIf(success == true, descrf...)
}

func MustNotBeNil(v interface{}, descrf ...string) {
	shoutAndPanicIf(v == nil, descrf...)
}

func MustBeNil(v interface{}, descrf ...string) {
	shoutAndPanicIf(v != nil, descrf...)
}

func RequireArg(v interface{}, descrf ...string) {
	if len(descrf) == 0 {
		descrf = []string{"nil passed instead of real value"}
	}
	MustNotBeNil(v, descrf...)
}

func MustBeFalse(success bool, descrf ...string) {
	shoutAndPanicIf(success == false, descrf...)
}

func RequireNoError(err error, descrf ...string) {
	shoutAndPanicIf(err != nil, descrf...)
}

func Guarantee(success bool, descrf ...string) {
	shoutAndPanicIf(!success, descrf...)
}
