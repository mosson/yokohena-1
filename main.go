package main

import (
	"fmt"
	"yokohena-1/ai"
)

func main() {
	b := ai.New()

	for {
		b.Wait()

		var i int
		fmt.Scanf("%d", &i)
		isContinue, message := b.Tick(i)

		if !isContinue {
			b.BoardDescription()
			if message != "" {
				fmt.Print("\n", message, "\n")
			}
			break
		}
	}

}
