package test

import (
    "awesome/tutorial"
    "fmt"
    "testing"
)

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
