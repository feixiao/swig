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


typedef _gostring_ swig_type_1;
extern void _wrap_Swig_free_example_7c6d0b865a722e37(uintptr_t arg1);
extern uintptr_t _wrap_Swig_malloc_example_7c6d0b865a722e37(swig_intgo arg1);
extern uintptr_t _wrap_new_Vector_example_7c6d0b865a722e37(double arg1, double arg2, double arg3);
extern void _wrap_delete_Vector_example_7c6d0b865a722e37(uintptr_t arg1);
extern swig_type_1 _wrap_Vector_print_example_7c6d0b865a722e37(uintptr_t arg1);
extern uintptr_t _wrap_addv_example_7c6d0b865a722e37(uintptr_t arg1, uintptr_t arg2);
extern uintptr_t _wrap_new_VectorArray_example_7c6d0b865a722e37(swig_intgo arg1);
extern void _wrap_delete_VectorArray_example_7c6d0b865a722e37(uintptr_t arg1);
extern swig_intgo _wrap_VectorArray_size_example_7c6d0b865a722e37(uintptr_t arg1);
extern uintptr_t _wrap_VectorArray_get_example_7c6d0b865a722e37(uintptr_t arg1, swig_intgo arg2);
extern void _wrap_VectorArray_set_example_7c6d0b865a722e37(uintptr_t arg1, swig_intgo arg2, uintptr_t arg3);
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


type swig_gostring struct { p uintptr; n int }
func swigCopyString(s string) string {
  p := *(*swig_gostring)(unsafe.Pointer(&s))
  r := string((*[0x7fffffff]byte)(unsafe.Pointer(p.p))[:p.n])
  Swig_free(p.p)
  return r
}

func Swig_free(arg1 uintptr) {
	_swig_i_0 := arg1
	C._wrap_Swig_free_example_7c6d0b865a722e37(C.uintptr_t(_swig_i_0))
}

func Swig_malloc(arg1 int) (_swig_ret uintptr) {
	var swig_r uintptr
	_swig_i_0 := arg1
	swig_r = (uintptr)(C._wrap_Swig_malloc_example_7c6d0b865a722e37(C.swig_intgo(_swig_i_0)))
	return swig_r
}

type SwigcptrVector uintptr

func (p SwigcptrVector) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVector) SwigIsVector() {
}

func NewVector(arg1 float64, arg2 float64, arg3 float64) (_swig_ret Vector) {
	var swig_r Vector
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3
	swig_r = (Vector)(SwigcptrVector(C._wrap_new_Vector_example_7c6d0b865a722e37(C.double(_swig_i_0), C.double(_swig_i_1), C.double(_swig_i_2))))
	return swig_r
}

func DeleteVector(arg1 Vector) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_Vector_example_7c6d0b865a722e37(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVector) Print() (_swig_ret string) {
	var swig_r string
	_swig_i_0 := arg1
	swig_r_p := C._wrap_Vector_print_example_7c6d0b865a722e37(C.uintptr_t(_swig_i_0))
	swig_r = *(*string)(unsafe.Pointer(&swig_r_p))
	var swig_r_1 string
 swig_r_1 = swigCopyString(swig_r) 
	return swig_r_1
}

type Vector interface {
	Swigcptr() uintptr
	SwigIsVector()
	Print() (_swig_ret string)
}

func Addv(arg1 Vector, arg2 Vector) (_swig_ret Vector) {
	var swig_r Vector
	_swig_i_0 := arg1.Swigcptr()
	_swig_i_1 := arg2.Swigcptr()
	swig_r = (Vector)(SwigcptrVector(C._wrap_addv_example_7c6d0b865a722e37(C.uintptr_t(_swig_i_0), C.uintptr_t(_swig_i_1))))
	return swig_r
}

type SwigcptrVectorArray uintptr

func (p SwigcptrVectorArray) Swigcptr() uintptr {
	return (uintptr)(p)
}

func (p SwigcptrVectorArray) SwigIsVectorArray() {
}

func NewVectorArray(arg1 int) (_swig_ret VectorArray) {
	var swig_r VectorArray
	_swig_i_0 := arg1
	swig_r = (VectorArray)(SwigcptrVectorArray(C._wrap_new_VectorArray_example_7c6d0b865a722e37(C.swig_intgo(_swig_i_0))))
	return swig_r
}

func DeleteVectorArray(arg1 VectorArray) {
	_swig_i_0 := arg1.Swigcptr()
	C._wrap_delete_VectorArray_example_7c6d0b865a722e37(C.uintptr_t(_swig_i_0))
}

func (arg1 SwigcptrVectorArray) Size() (_swig_ret int) {
	var swig_r int
	_swig_i_0 := arg1
	swig_r = (int)(C._wrap_VectorArray_size_example_7c6d0b865a722e37(C.uintptr_t(_swig_i_0)))
	return swig_r
}

func (arg1 SwigcptrVectorArray) Get(arg2 int) (_swig_ret Vector) {
	var swig_r Vector
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	swig_r = (Vector)(SwigcptrVector(C._wrap_VectorArray_get_example_7c6d0b865a722e37(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1))))
	return swig_r
}

func (arg1 SwigcptrVectorArray) Set(arg2 int, arg3 Vector) {
	_swig_i_0 := arg1
	_swig_i_1 := arg2
	_swig_i_2 := arg3.Swigcptr()
	C._wrap_VectorArray_set_example_7c6d0b865a722e37(C.uintptr_t(_swig_i_0), C.swig_intgo(_swig_i_1), C.uintptr_t(_swig_i_2))
}

type VectorArray interface {
	Swigcptr() uintptr
	SwigIsVectorArray()
	Size() (_swig_ret int)
	Get(arg2 int) (_swig_ret Vector)
	Set(arg2 int, arg3 Vector)
}


