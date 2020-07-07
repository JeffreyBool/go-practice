package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

/*
!---listen16
|    |----employee
|    |    |----employee.exe
|    |    |----main.go
|    |----empty_interface
|    |    |----empty_interface.exe
|    |    |----main.go
|    |----homework
|    |    |----tree
|    |    |    |----main.go
|    |    |    |----tree.exe
|    |----interface_nest
|    |    |----interface_nest.exe
|    |    |----main.go
|    |----interface_test
|    |    |----interface_test.exe
|    |    |----main.go
|    |----multi_interface
|    |    |----main.go
|    |    |----multi_interface.exe
|    |----pointer_interface
|    |    |----main.go
|    |    |----pointer_interface.exe
|    |----type_assert
|    |    |----main.go
|    |    |----type_assert.exe
*/

func main() {
	app := cli.NewApp()
	app.Name = "golang to tree"

	app.Usage = "list all file"
	app.Action = func(c *cli.Context) error {
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}
		if err := ListDir(dir, 1); err != nil {
			panic(err)
		}
		return nil
	}
	app.Run(os.Args)
}

func ListDir(dirPath string, deep int) (err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	if deep == 1 {
		fmt.Printf("!---%s\n", filepath.Base(dirPath))
	}

	// window的目录分隔符是 \
	// linux 的目录分隔符是 /
	sep := string(os.PathSeparator)
	for _, fi := range dir {
		// 如果是目录，继续调用ListDir进行遍历
		if fi.IsDir() {
			fmt.Printf("|")
			for i := 0; i < deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", fi.Name())
			if err = ListDir(dirPath+sep+fi.Name(), deep+1); err != nil {
				fmt.Printf("err:%s\n", err)
			}
			continue
		}

		fmt.Printf("|")
		for i := 0; i < deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
	return nil
}
