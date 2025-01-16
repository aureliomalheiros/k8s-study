package errorHandler

import (
	"fmt"
	"os"
)

func HandlerError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}