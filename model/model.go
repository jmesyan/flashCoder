package model


type Condition struct {
    template string
    param interface{}
}



type DbModel interface {
	Select(table string,field []string,condition []Condition)
}
