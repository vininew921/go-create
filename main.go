package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"strings"
)

func main() {
	project_name, dir := create_dir()
	create_module(project_name, dir)
	init_main(dir)

	fmt.Printf("Done! Use `cd %s` to go to the newly created module's path", project_name)
}

func create_module(project_name string, dir string) {
	cmd := exec.Command("go", "mod", "init", project_name)
	cmd.Dir = dir

	if err := cmd.Run(); err != nil {
		fmt.Println("Error running `go mod init` command: ", err)
		os.Exit(1)
	}

	fmt.Println("Created module", project_name)
}

func init_main(dir string) {
	main_content := main_content()
	main_path := dir + string(os.PathSeparator) + "main.go"

	if err := os.WriteFile(main_path, main_content, fs.ModePerm); err != nil {
		fmt.Println("Error initializing `main.go` file: ", err)
		os.Exit(1)
	}

	fmt.Println("Initialized main.go")
}

func create_dir() (string, string) {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	project_name := strings.Trim(os.Args[1], " ")
	if len(project_name) == 0 {
		usage()
		os.Exit(1)
	}

	full_path := dir + string(os.PathSeparator) + project_name

	os.Mkdir(full_path, fs.ModePerm)

	return project_name, full_path
}

func usage() {
	fmt.Println("Project name wasn't provided")
	fmt.Println("Usage: go-create <package_name>")
}

func main_content() []byte {
	content :=
		`package main

func main() {
    println("Hello World")
}
`

	return []byte(content)
}
