package utils

import (
	"fmt"
	"log"
	"os"
)

func Output(res chan string, dst string) {
	if dst == "" {
		for r := range res {
			fmt.Println(r)
		}

		fmt.Println("\nDone")

		return
	}

	file, err := os.OpenFile(dst, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for r := range res {
		_, err := fmt.Fprintln(file, r)
		if err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("\nDone")
}
