package main

import (
	"fmt"

	"github.com/vinaay/mylib/utils" // replace 'your_actual_module_name' with the correct module name

	"context"

	"github.com/vinaay/mylib/db"
)

func main() {
	fmt.Println(utils.Capitalize("golang is awesome"))

	r := db.NewRedisClient("localhost:6379")

	ctx := context.Background() // Define ctx

	r.SetValue(ctx, "foo", "bar")
	val, _ := r.GetValue(ctx, "foo")
	fmt.Println("foo:", val)
}
