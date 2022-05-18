package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

// go mod init github.com/ChrisMarSilva/cms.golang.teste.pastas
// go mod tidy

// go run main.go
// go build main.go

func main() {

	log.Println("INI")
	var start time.Time = time.Now()

	f, err := os.Create("arquivos.txt")
	if err != nil {
		log.Fatalln("os.Create.err", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	path := "C:\\Users\\chris\\Desktop\\Nova pasta"
	path = "C:\\Users\\chris\\Desktop\\Filmes"
	path = "C:\\"

	tot := CarregarListaArquivos(path, w)
	log.Printf("Size: '%s' - Dir: '%s'\n", ByteCountSI(tot), path)

	log.Println("FIM:", time.Since(start))
}

func CarregarListaArquivos(path string, w *bufio.Writer) int64 {

	var tot int64 = 0

	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("ReadDir.err:", err)
		return 0
	}

	for _, f := range files {
		var size int64 = 0
		if f.IsDir() {
			size := CarregarListaArquivos(path+"\\"+f.Name(), w)
			if size > 104857600 { // 100 Megabytes = 104,857,600 Bytes = 104857600 // 1 Gigabytes = 1,073,741,824 Bytes = 1073741824
				texto := fmt.Sprintf("Size: '%s' - Dir: '%s'\n", ByteCountSI(size), path+"\\"+f.Name())
				log.Printf(texto)
				w.WriteString(texto)
			}
		} else {
			size = f.Size()
			if size > 104857600 { // 100 Megabytes = 104,857,600 Bytes // 104857600
				texto := fmt.Sprintf("Size: '%s' - File: '%s'\n", ByteCountSI(size), path+"\\"+f.Name())
				log.Printf(texto)
				w.WriteString(texto)
			}
		}
		tot += size
	} // for _, f := range files {

	// fs := float64( file.Size() )
	// fmt.Println(file.Name(), HumanFileSize(fs), file.ModTime())

	return tot
}

func ByteCountSI(b int64) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func ByteCountIEC(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

var (
	suffixes [5]string
)

func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func HumanFileSize(size float64) string {
	fmt.Println(size)
	suffixes[0] = "B"
	suffixes[1] = "KB"
	suffixes[2] = "MB"
	suffixes[3] = "GB"
	suffixes[4] = "TB"

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	fmt.Println(int(math.Floor(base)))
	getSuffix := suffixes[int(math.Floor(base))]
	return strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix)
}
