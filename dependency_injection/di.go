package dependency_injection

import (
	"fmt"
	"bytes"
)

func Greet(writer *bytes.Buffer ,name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
