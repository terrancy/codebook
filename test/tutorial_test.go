package test

import (
    "awesome/tutorial"
    "fmt"
    "reflect"
    "testing"
)

//=======================================
//------------------反射机制--------------
//=======================================
func TestGetType(t *testing.T) {
    str := "Gopher"
    fmt.Println(reflect.TypeOf(str))
}

//=======================================
//------------------并发编程--------------
//=======================================

//
//  TestSyncDownload
//  @Description: 协程部分
//  @param t
//
func TestSyncDownload(t *testing.T) {
    tutorial.SyncDownload()
}

func TestInterface(t *testing.T) {
    var p tutorial.Person = &tutorial.Student{Name: "cai", Age: 33}
    fmt.Println(p.GetName())
}

//
//  TestMap
//  @Description:
//  @param t
//
func TestMap(t *testing.T) {
    dic := make(map[int]int)
    dic[1] = 110
    tutorial.ForMap(dic)
    fmt.Println(dic)
    fmt.Printf("%p\n", &dic)
}
