package map_test

import "testing"

func TestMapWithFunValue(t* testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int {return op}
	m[2] = func(op int) int {return op*op}
	m[3] = func(op int) int {return op*op*op}
	t.Log(m[1](3),m[2](4),m[3](5))
}
//map实现set
func TestMapForSet(t *testing.T){
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is exist",n)
	}else {
		t.Logf("%d not exist",n)
	}
	t.Log(len(mySet))
	delete(mySet,1)
}

