package main

import (
	"fmt"
	"github.com/gyamin/fcmDePushTsuchi/pkg/db"
	"github.com/gyamin/fcmDePushTsuchi/pkg/push"
)

func main() {
	push.HelloA()
	con := db.GetConnection()
	fmt.Println(con.Ping())
}
