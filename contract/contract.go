package contract

import "fmt"

var SuppressPanic bool = false

const EmptyDescrf = ""

func shoutAndPanic(descrf string, fmtargs ...interface{}) {
	var err error
	if descrf == EmptyDescrf {
		err = fmt.Errorf("failed to meet the contract")
	} else {
		err = fmt.Errorf(descrf, fmtargs...)
	}
	if SuppressPanic {
		fmt.Println(err)
	} else {
		panic(err)
	}
}

func shoutAndPanicIf(orly bool, descrf string, fmtargs ...interface{}) {
	if orly {
		shoutAndPanic(descrf, fmtargs...)
	}
}

func Requiref(success bool, descrf string, fmtargs ...interface{}) {
	shoutAndPanicIf(!success, descrf, fmtargs...)
}

func MustBeTruef(success bool, descrf string, fmtargs ...interface{}) {
	shoutAndPanicIf(!success, descrf, fmtargs...)
}

func MustNotBeNilf(v interface{}, descrf string, fmtargs ...interface{}) {
	shoutAndPanicIf(v == nil, descrf, fmtargs...)
}

func MustBeNilf(v interface{}, descrf string, fmtargs ...interface{}) {
	shoutAndPanicIf(v != nil, descrf, fmtargs...)
}

func RequireArgf(v interface{}, descrf string, fmtargs ...interface{}) {
	MustNotBeNilf(v, descrf, fmtargs...)
}

func MustBeFalsef(success bool, descrf string, fmtargs ...interface{}) {
	shoutAndPanicIf(success == false, descrf, fmtargs...)
}

func RequireNoErrorf(err error, descrf string, fmtargs ...interface{}) {
	shoutAndPanicIf(err != nil, descrf, fmtargs...)
}

func Guaranteef(success bool, descrf string, fmtargs ...interface{}) {
	shoutAndPanicIf(!success, descrf, fmtargs...)
}

/*
	Without formatting
*/

func Require(success bool) {
	shoutAndPanicIf(!success, EmptyDescrf)
}

func MustBeTrue(success bool) {
	shoutAndPanicIf(!success, EmptyDescrf)
}

func MustNotBeNil(v interface{}) {
	shoutAndPanicIf(v == nil, EmptyDescrf)
}

func MustBeNil(v interface{}) {
	shoutAndPanicIf(v != nil, EmptyDescrf)
}

func RequireArg(v interface{}) {
	MustNotBeNilf(v, "nil passed instead of real value")
}

func MustBeFalse(success bool) {
	shoutAndPanicIf(success == false, EmptyDescrf)
}

func RequireNoError(err error) {
	shoutAndPanicIf(err != nil, EmptyDescrf)
}

func RequireNoErrors(conds ...func() error) {
	for _, cond := range conds {
		err := cond()
		if err != nil {
			shoutAndPanic("Contract failed because of: %s", err.Error())
		}
	}
}

func Guarantee(success bool) {
	shoutAndPanicIf(!success, EmptyDescrf)
}
