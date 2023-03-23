package main

import (
	"context"
	"fmt"

	"example.com/ent/ent"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:@tcp(localhost:3306)/casone_lite_local")
	if err != nil {
		fmt.Printf("ent.Open: %v", err)
	}
	defer client.Close()
	user, err := client.User.Get(context.Background(), 1)
	if err != nil {
		fmt.Printf("%v", err)
	}
	fmt.Println(user)
}
