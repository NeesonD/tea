package main

import (
	"bytes"
	goformat "go/format"
	"io/ioutil"
	"os"
	"path/filepath"
)

func createFileAndFormat(file string, code string) {
	createFile(file, code, true)
}

// 本质就是将字符串写到文件中
func createFile(file string, code string, format bool) error {
	dirAbs, err := filepath.Abs("./gen")
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer([]byte(code))

	if format {
		formatOutput, err := goformat.Source(buf.Bytes())
		if err != nil {
			return err
		}

		buf.Reset()
		buf.Write(formatOutput)
	}

	name := file + ".go"
	filename := filepath.Join(dirAbs, name)

	err = ioutil.WriteFile(filename, buf.Bytes(), os.ModePerm)
	if err != nil {
		return err
	}

	return err
}

func genFileName(filename string) string {
	return "gen_" + filename
}
