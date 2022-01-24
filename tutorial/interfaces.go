package tutorial

type Person interface {
    GetName() string
}

type Student struct {
    Name string
    Age  int
}

func (stu *Student) GetName() string {
    return stu.Name
}
