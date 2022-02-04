package tutorial

import "reflect"

// 了解反射的相关知识点

func getInterfaceType() reflect.Type {
    str := "Gopher"
    return reflect.TypeOf(str)
}
