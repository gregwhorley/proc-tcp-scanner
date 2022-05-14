package main

import (
	"fmt"
	"io/ioutil"
)

func report(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Sprintf("%s", content)
	return nil
}
