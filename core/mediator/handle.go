package mediator

import (
	"reflect"
)

const HandleMethodName string = "Handle"

func call(handler interface{}, ctx interface{}, method reflect.Value, msg interface{}) error {
	in := []reflect.Value{reflect.ValueOf(handler), reflect.ValueOf(ctx), reflect.ValueOf(msg)}

	result := method.Call(in)

	return handleErr(result)
}

func callHandle(handler interface{}, msg interface{}) error {
	handlerType := reflect.TypeOf(handler)

	handleMethod, ok := handlerType.MethodByName(HandleMethodName)

	if !ok {
		panic("handle method does not exists for the typeOf" + handlerType.String())
	}

	in := []reflect.Value{reflect.ValueOf(handler), reflect.ValueOf(msg)}

	result := handleMethod.Func.Call(in)

	return handleErr(result)
}

func handleErr(result []reflect.Value) error {
	if result == nil {
		return nil
	}

	if v := result[0].Interface(); v != nil {
		return v.(error)
	}
	return nil
}
