package main

import (
	"fmt"
	"os"

	"github.com/gokultp/tql/pkg/fs"
)

func main() {
	query, fileName := parseArgs()
	fp, err := fs.GetFileReader(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println(query, fp)

}

func parseArgs() (*string, *string) {
	var query = new(string)
	var file = new(string)
	switch len(os.Args) {
	case 2:
		*query = os.Args[1]
	case 3:
		*query = os.Args[1]
		*file = os.Args[2]
	}
	return query, file
}
