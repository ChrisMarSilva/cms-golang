
medir tempo de leitura delphi
medir tempo de leitura c#
medir tempo de leitura python
medir tempo de leitura golang


listar arquivos do diretorio
mover arquivo
inserir na tabela de arquivo
ler arquivo
inserir na tabela de registro
validar pk do arquivo


ler tabela de arquivo
gerar arquivo retorno
salvar arquivo


ler tabela de resposta
gerar arquivo final
salvar arquivo



package main
 
import (
	"log"
	"os"
)
 
func main() {
	oldLocation := "/var/www/html/test.txt"
	newLocation := "/var/www/html/src/test.txt"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func MoveFile(source string, destination string) {
    err := os.Rename(source, destination)
    if err != nil {
        fmt.Println(err)
    }
}



import (
    "fmt"
    "io"
    "os"
)

func MoveFile(sourcePath, destPath string) error {
    inputFile, err := os.Open(sourcePath)
    if err != nil {
        return fmt.Errorf("Couldn't open source file: %s", err)
    }
    outputFile, err := os.Create(destPath)
    if err != nil {
        inputFile.Close()
        return fmt.Errorf("Couldn't open dest file: %s", err)
    }
    defer outputFile.Close()
    _, err = io.Copy(outputFile, inputFile)
    inputFile.Close()
    if err != nil {
        return fmt.Errorf("Writing to output file failed: %s", err)
    }
    // The copy was successful, so now delete the original file
    err = os.Remove(sourcePath)
    if err != nil {
        return fmt.Errorf("Failed removing original file: %s", err)
    }
    return nil
}


func moveFile(file1, file2 string) error {
   return os.Rename(file1, file2)
}


  if !os.IsNotExist(err) {
            return
        }



package main

import (
    "time"
    "fmt"       
    "io/ioutil"
    "os"
    "bytes"
    "strings"
)

const dir = "dir00003"

func main() {

fmt.Println("Running...")

//Go into a loop forever
for {

    //Wait 60 seconds before taking any action. 
    time.Sleep(60 * time.Second)
    //Read all of the file data for all files in the directory: 
    files, err := ioutil.ReadDir(dir)
    if err != nil {fmt.Println("Failed to read transfer folder. There must be a folder named `dir00003`!"); continue}

    for _, v := range files {

        //if this is an index file, skip over it as we don't care: 
        if strings.Contains(v.Name(), "pmi") {continue}

        //if the file was created within the last 2 minutes, we should check if we need to modify it
        if time.Now().Sub(v.ModTime()) < (time.Minute * 2) {

            //open the file 
            f, err := os.Open(fmt.Sprintf("%s/%s", dir, v.Name()))
            if err != nil {fmt.Printf("\tCouldn't open file: %s\n", v.Name()); continue}

            defer f.Close()
            //read all of the bytes of the file
            bs, err := ioutil.ReadAll(f)
            if err != nil {fmt.Printf("\tCouldn't read bytes from %s\n", v.Name()); continue}

            //see if the <program_parameters/> tag is in the file
            b := bytes.Contains(bs, []byte("<program_parameters/>"))

            //if the tag is in the file, we should replace it, otherwise we move on to the next file
            if b {
                //replace the tag with nothing. Only look for the first instance and then abort the process of replacing.
                rbs := bytes.Replace(bs, []byte("<program_parameters/>"), []byte(""), 1)
                //close the file so we can delete it. 
                f.Close()
                //delete the exisint file. 
                os.Remove(fmt.Sprintf("%s/%s", dir, v.Name()))

                //create a new file with the same original name:
                nf, err := os.Create(fmt.Sprintf("%s/%s", dir, v.Name()))
                if err != nil {fmt.Printf("\tFailed to create new file for %s\n", v.Name()); continue}

                //write all of the bytes that we have in memory to our new file. 
                _, err = nf.Write(rbs)
                if err != nil {fmt.Println("Failed to write to new file %s\n", v.Name()); continue}
                //close our new file
                nf.Close()

                fmt.Printf("Modified new file: %s", v.Name())

            } else {
                continue
            }

        }

    

    }

    fmt.Printf("\nDone with round\n")

}




package main

import (
    "fmt"
    "io/ioutil"
     "log"
)

func main() {
    files, err := ioutil.ReadDir("./")
    if err != nil {
        log.Fatal(err)
    }
 
    for _, f := range files {
            fmt.Println(f.Name())
    }
}



import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
)

func main() {
    var (
        root  string
        files []string
        err   error
    )

    root := "/home/manigandan/golang/samples"
    // filepath.Walk
    files, err = FilePathWalkDir(root)
    if err != nil {
        panic(err)
    }
    // ioutil.ReadDir
    files, err = IOReadDir(root)
    if err != nil {
        panic(err)
    }
    //os.File.Readdir
    files, err = OSReadDir(root)
    if err != nil {
        panic(err)
    }

    for _, file := range files {
        fmt.Println(file)
    }
}



func IOReadDir(root string) ([]string, error) {
    var files []string
    fileInfo, err := ioutil.ReadDir(root)
    if err != nil {
        return files, err
    }

    for _, file := range fileInfo {
        files = append(files, file.Name())
    }
    return files, nil
}


package main    

import (
    "fmt"
    "log"
    "path/filepath"
)

func main() {
    files, err := filepath.Glob("*")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(files) // contains a list of all files in the current directory
}


package main

import (
    "fmt"
    "log"
    "os"
)

func main() {
    files, err := os.ReadDir(".")
    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {
        fmt.Println(file.Name())
    }
}


package main

import (
    "fmt"
    "os"
    "path/filepath"
)

func main() {
    var files []string

    root := "/some/folder/to/scan"
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        files = append(files, path)
        return nil
    })
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fmt.Println(file)
    }
}


package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	var (
    		root string
    		files []string
    		err error
	  )	

	if len(os.Args) == 1 {
		log.Fatal("No path given, Please specify path.")
		return
	}
	if root = os.Args[1]; root == "" {
		log.Fatal("No path given, Please specify path.")
		return
	}
	// filepath.Walk
	 files, err = FilePathWalkDir(root)
	 if err != nil {
	 	panic(err)
	 }
	// ioutil.ReadDir
	 files, err = IOReadDir(root)
	 if err != nil {
	 	panic(err)
	 }
	//os.File.Readdir
	files, err = OSReadDir(root)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Println(file)
	}
}



func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	return files, err
}


func IOReadDir(root string) ([]string, error) {
	var files []string
	fileInfo, err := ioutil.ReadDir(root)
	if err != nil {
		return files, err
	}

	for _, file := range fileInfo {
		files = append(files, file.Name())
	}
	return files, nil
}

port (
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	names, err := file.Readdirnames(0)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range names {
		fmt.Println(v)
	}
}

package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {

    home, err := os.UserHomeDir()

    if err != nil {

        log.Fatal(err)
    }

    files, err := ioutil.ReadDir(home)

    if err != nil {

        log.Fatal(err)
    }

    for _, f := range files {

        fmt.Println(f.Name())
    }
}


ist_files_ext.go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"
)

func main() {

    path := "/home/janbodnar/Documents/"

    files, err := ioutil.ReadDir(path)

    if err != nil {

        log.Fatal(err)
    }

    for _, f := range files {

        if filepath.Ext(f.Name()) == ".txt" {

            fmt.Println(f.Name())
        }
    }
}



package main

import (
    "fmt"
    "log"
    "path/filepath"
)

func main() {

    files, err := filepath.Glob("/root/Documents/prog/golang/**/*.go")

    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {

        fmt.Println(file)
    }
}



ZetCode
All Spring Boot Python C# Java JavaScript Subscribe
Go list directory
last modified September 15, 2020

Go list directory show how to list directory contents in Golang.

Directory definition
A directory is a unit in a computer's file system for storing and locating files. Directories are hierarchically organized into a tree. Directories have parent-child relationships. A directory is sometimes also called a folder.


 
In Go, we can list directories with ioutil.ReadDir, filepath.Walk, or filepath.Glob.

Go list directory with ioutil.ReadDir
The ioutil.ReadDir reads the directory and returns a list of directory entries sorted by filename.

func ReadDir(dirname string) ([]os.FileInfo, error)
This is the syntax of the ReadDir function.

read_homedir.go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {

    home, err := os.UserHomeDir()

    if err != nil {

        log.Fatal(err)
    }

    files, err := ioutil.ReadDir(home)

    if err != nil {

        log.Fatal(err)
    }

    for _, f := range files {

        fmt.Println(f.Name())
    }
}

The example reads the user home directory contents. The home user directory is determined with os.UserHomeDir. The listing is non-recursive.

Go list files by extension
The filepath.Ext returns the file name extension used by path.

list_files_ext.go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "path/filepath"
)

func main() {

    path := "/home/janbodnar/Documents/"

    files, err := ioutil.ReadDir(path)

    if err != nil {

        log.Fatal(err)
    }

    for _, f := range files {

        if filepath.Ext(f.Name()) == ".txt" {

            fmt.Println(f.Name())
        }
    }
}
The example shows all .txt files in the Documents directory.

Go list directories
The FileInfo's IsDir can be used to limit the listing to only files or directories.

list_files.go
package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "os"
)

func main() {

    home, err := os.UserHomeDir()

    if err != nil {

        log.Fatal(err)
    }

    files, err := ioutil.ReadDir(home)

    if err != nil {

        log.Fatal(err)
    }

    for _, f := range files {

        if !f.IsDir() {
            fmt.Println(f.Name())
        }
    }
}
The example list only files in the home directory.

Go list directory with filepath.Glob
The filepath.Glob returns the names of all files matching pattern or nil if there is no matching file.

func Glob(pattern string) (matches []string, err error)
This is the syntax of the filepath.Glob function.

globbing.go
package main

import (
    "fmt"
    "log"
    "path/filepath"
)

func main() {

    files, err := filepath.Glob("/root/Documents/prog/golang/**/*.go")

    if err != nil {
        log.Fatal(err)
    }

    for _, file := range files {

        fmt.Println(file)
    }
}
The example lists all Go files in the given directory. With the ** pattern, the listing is recursive.


 
Go list directory with filepath.Walk
For recursive directory listings, we can use the filepath.Walk function.

func Walk(root string, walkFn WalkFunc) error
The function walks the file tree rooted at root, calling walkFn for each file or directory in the tree, including root. All errors from visiting files and directories are filtered by walkFn.

walking.go
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {

    err := filepath.Walk("/home/janbodnar/Documents/prog/golang/",

        func(path string, info os.FileInfo, err error) error {

            if err != nil {
                return err
            }

            fmt.Println(path, info.Size())
            return err
        })

    if err != nil {

        log.Println(err)
    }
}
The example walks recursively the given directory. It outputs each path name and size.

Go directory size
The following example uses the filepath.Walk function to get the size of all files in the given directory. The directory sizes are not included.

dirsize.go
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {

    var size int64

    path := "/home/janbodnar/Documents/prog/golang/"

    err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {

        if err != nil {
            return err
        }

        if !info.IsDir() {

            size += info.Size()
        }

        return err
    })

    if err != nil {

        log.Println(err)
    }

    fmt.Printf("The directory size is: %d\n", size)
}
The example uses the IsDir function to tell a file from a directory. The size of a file is determined with the Size function.

Go list large files
The following example outputs large files with filepath.Walk.

large_files.go
package main

import (
    "fmt"
    "log"
    "os"
    "path/filepath"
)

func main() {

    var files []string

    var limit int64 = 1024 * 1024 * 1024

    path := "/home/janbodnar/Downloads/"

    err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {

        if err != nil {
            return err
        }

        if info.Size() > limit {

            files = append(files, path)
        }

        return err
    })

    if err != nil {
        log.Println(err)
    }

    for _, file := range files {

        fmt.Println(file)
    }
}
In the example, we list all files that are larger than 1GB in the Downloads directory.

In this tutorial, we have listed directory contents in Go.


 
List all Go tutorials.

Home Facebook Twitter Github Subscribe Privacy
© 2007 - 2021 Jan Bodnar admin(at)zetcode.com