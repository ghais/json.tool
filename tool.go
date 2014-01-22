package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	d := json.NewDecoder(os.Stdin)
	for {
		obj := make(map[string]interface{})
		err := d.Decode(&obj)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		v, err := json.MarshalIndent(obj, "\n", "\t")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(string(v))

	}

}
