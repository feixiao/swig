/* ----------------------------------------------------------------------------
 * This file was automatically generated by SWIG (http://www.swig.org).
 * Version 3.0.12
 *
 * This file is not intended to be easily readable and contains a number of
 * coding conventions designed to improve portability and efficiency. Do not make
 * changes to this file unless you know what you are doing--modify the SWIG
 * interface file instead.
 * ----------------------------------------------------------------------------- */

// source: example.i

package example

/*
#define intgo swig_intgo
typedef void *swig_voidp;

#include <stdint.h>


typedef long long intgo;
typedef unsigned long long uintgo;



typedef struct { char *p; intgo n; } _gostring_;
typedef struct { void* array; intgo len; intgo cap; } _goslice_;


extern void _wrap_Swig_free_example_63e8d1da367391f1(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_example_63e8d1da367391f1(swig_intgo arg1);
extern swig_intgo _wrap_gcd_example_63e8d1da367391f1(swig_intgo arg1, swig_intgo arg2);
extern void _wrap_Foo_set_example_63e8d1da367391f1(double arg1);
extern double _wrap_Foo_get_example_63e8d1da367391f1(void);
#undef intgo
*/
import "C"

import "unsafe"
import _ "runtime/cgo"
import "sync"


type _ unsafe.Pointer



var Swig_escape_always_false bool
var Swig_escape_val interface{}


type _swig_fnptr *byte
type _swig_memberptr *byte


type _ sync.Mutex

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_example_63e8d1da367391f1(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_example_63e8d1da367391f1(C.swig_intgo(_swig_i_0)))
	return swig_r
}

func Gcd(arg1 int, arg2 int) (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (int)(C._wrap_gcd_example_63e8d1da367391f1(C.swig_intgo(_swig_i_0), C.swig_intgo(_swig_i_1)))
	return swig_r
}

func SetFoo(arg1 float64) {
	_swig_i_0 := arg1
	C._wrap_Foo_set_example_63e8d1da367391f1(C.double(_swig_i_0))
}

func GetFoo() (_swig_ret float64) {
	var swig_r float64
	swig_r = (float64)(C._wrap_Foo_get_example_63e8d1da367391f1())
	return swig_r
}


