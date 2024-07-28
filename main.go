package main

import (
	"bufio"
	"fmt"
	"os"

	"database.com/db/models/tree"
)

var AUTO_KEY uint64 = 1

func main() {
	db := tree.OpenDB("data.db")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("DB:>")
		input, _, _ := reader.ReadLine()
		db.Insert(AUTO_KEY, input)

		AUTO_KEY += 1

		db.Store.Walk(func(t *tree.Tree) {
			fmt.Print(t.Root.Key)
		})

		fmt.Print("\n")
	}
}
