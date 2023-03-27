go test -bench=.  -  run benchmark tests

go func() {//do something}(param) - start a go routine
the param is necessary to make a copy of each item in the slice 
you are iterating over, otherwise the value will change as the loop is executed

go test -race - race detector
-  race cond. happen when software output is dependent on time
- often happens when iterating over maps as two go routines can write to the same memory address at the same time

channels
- go routine data structure that can receive and send values
- prevent race conditions 
