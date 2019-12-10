package main

import (
	"database/sql"
	"fmt"
)

func main() {
	fmt.Print(sql.Drivers())
}
