package demapstructure

import "reflect"

type DecodeHookFunc interface {}

//DocodeHookFuncType ia a DecodeHookFunc which has completely information about
//the source and target types.
type DecodeHookFuncType func(reflect.Type,reflect.Type,interface{})(interface{},error)

// DecodeHookFuncKind is a DecodeHookFunc which knows only the Kinds of the
// source and target types.
type DecodeHookFuncKind func(reflect.Kind,reflect.Kind,interface{})(interface{},error)


