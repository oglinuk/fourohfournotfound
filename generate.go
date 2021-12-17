package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

var (
	wg = &sync.WaitGroup{}

	cmd = flag.String("c", "gen", "Command to run [gen/clean]")
)

func main() {
	flag.Parse()

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		switch strings.ToLower(*cmd) {
		case "gen":
			if !info.IsDir() && strings.HasSuffix(path, ".md") {
				// TODO: Replace with PEGN-based parser
				wg.Add(1)
				go func() {
					fmt.Printf("Rendering %s ...\n", path)
					out := fmt.Sprintf("%s/index.html", filepath.Dir(path))
					err = exec.Command("pandoc", path, "-o", out).Run()
					if err != nil {
						log.Println(err)
					}
					wg.Done()
				}()
			}
		case "clean":
			if !info.IsDir() && strings.HasSuffix(path, ".html") {
				wg.Add(1)
				go func() {
					fmt.Printf("Removing %s ...\n", path)
					os.Remove(path)
					wg.Done()
				}()
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	wg.Wait()
}
