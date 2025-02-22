main func:

```go
// 1.1 give err link, wraped error in
simpe.De(err)
// 1.2 how to use
err = add1()
if err = simpe.De(err); err != nil {
    return
}

// 2.1 show link and error source
simpe.ShowErr(err)
// 2.2 how to use
simpe.ShowErr(err)
// log:
2025/02/22 12:14:09 [Error source]: add1
2025/02/22 12:14:09 [Error link]: /main.go:30 -> : /main.go:12 -> : add1
```

use snippet - (press "ert" + "tab" to use):

```json
"simpError": {
		"prefix": "ert",
		"body": [
			"if err = simpe.De(err); err != nil {",
			"\treturn",
			"}",
			"$1"
		],
		"description": "Golang: quickly deal go err"
	},
```

code example:

```go
package main

import (
	"errors"
	"log"

	"github.com/yuan-shuo/simpe"
)

func add1() (err error) {
	err = errors.New("add1")
	if err = simpe.De(err); err != nil {
		return
	}

	return nil
}

func add2() (err error) {
	err = errors.New("add1")
	if err = simpe.De(err); err != nil {
		return
	}

	return nil
}

func other() (err error) {
	err = add1()
	if err = simpe.De(err); err != nil {
		return
	}
	err = add2()
	if err = simpe.De(err); err != nil {
		return
	}
	return
}
func main() {
	err := other()
	simpe.ShowErr(err)
	log.Println("no err")
}
```

