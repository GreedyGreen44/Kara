package main

import (
	"bufio"
	"os"
)

func writeData(data []byte, fileName string) (err error) {
	fo, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer func() {
		if err = fo.Close(); err != nil {
			panic(err)
		}
	}()
	w := bufio.NewWriter(fo)
	if _, err = w.Write(data); err != nil {
		return err
	}
	if err = w.Flush(); err != nil {
		return err
	}
	return nil
}
