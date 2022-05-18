package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	// "sync"
)

func main() {

	filename := "filename.txt"
	filename = "C:\\Users\\chris\\Desktop\\CMS GoLang\\cms.golang.teste.arquivos\\cms.golang.teste.arquivos.cnab240\\Arquivos\\ArqvRemessa\\PROC\\PAGSEGURO_20210519_001_REM.txt"
	start := time.Now()

	// "io/ioutil"
	// err := ioutil.WriteFile(filename, []byte("Hello"), 0755)
	// if err != nil {
	// 	fmt.Printf("Unable to write file: %v", err)
	// }

	// fo, err := os.Create(filename)
	// //fo, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	// //fo, err := os.OpenFile("server", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // defer fo.Close()
	// defer func() {
	// 	if err := fo.Close(); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }()

	// w := bufio.NewWriter(fo)
	// defer w.Flush()
	// for i := 0; i < 3000000; i++ {
	// 	linha := fmt.Sprint("Hello ", i, " \n")
	// 	w.WriteString(linha)
	// }
	// log.Printf("Tempo de Criação %s", time.Since(start)) // 1.1458297s

	// _, err = fo.Write([]byte("Hello 1 \n"))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// for i := 0; i < 3000000; i++ {
	// 	linha := fmt.Sprint("Hello ", i, " \n")
	// 	bytes := []byte(linha)
	// 	fo.Write(bytes)
	// }
	// log.Printf("Tempo de Criação %s", time.Since(start))

	// wg1 := sync.WaitGroup{}
	// wg1.Add(1)
	// ch1 := make(chan string)
	// ch2 := make(chan []byte)
	// go func() {
	// 	for i := 0; i < 3000000; i++ {
	// 		ch1 <- fmt.Sprint("Hello ", i, " \n")
	// 	}
	// 	close(ch1)
	// }()
	// go func() {
	// 	for v := range ch1 {
	// 		ch2 <- []byte(v)
	// 	}
	// 	close(ch2)
	// }()
	// go func() {
	// 	for v := range ch2 {
	// 		fo.Write(v)
	// 	}
	// 	wg1.Done()
	// }()
	// wg1.Wait()
	// log.Printf("Tempo de Criação %s", time.Since(start))

	//fi, err := os.Open(filename)
	fi, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// defer fi.Close()
	defer func() {
		if err := fi.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// filestat, err := fi.Stat()
	// if err != nil {
	// 	fmt.Println("Não foi possível obter a estatística do arquivo")
	// 	return
	// }
	// fileSize := filestat.Size()
	// fmt.Println("fileSize: ", fileSize)

	// buf := make([]byte, 1024) // make a buffer to keep chunks that are read
	// for {
	// 	n, err := fi.Read(buf)
	// 	if err != nil && err != io.EOF {
	// 		log.Fatal(err)
	// 	}
	// 	if n == 0 {
	// 		break
	// 	}
	// 	conteudo := string(buf[0:n])
	// 	log.Println("Linha:", string(n), "Conteudo:", conteudo)
	// }

	// rd := bufio.NewReader(fi)
	// for {
	// 	line, err := rd.ReadString('\n')
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		log.Fatalf("read file line error: %v", err)
	// 		return
	// 	}
	// 	log.Println(line)
	// }

	// rd := bufio.NewReader(fi)
	// bytes := []byte{}
	// // lines := []string{}
	// for {
	// 	line, isPrefix, err := rd.ReadLine()
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	bytes = append(bytes, line...)
	// 	if !isPrefix {
	// 		str := strings.TrimSpace(string(bytes))
	// 		// log.Println(str)
	// 		log.Println(string(bytes))
	// 		if len(str) > 0 {
	// 			// lines = append(lines, str)
	// 			bytes = []byte{}
	// 		}
	// 	}
	// }
	// if len(bytes) > 0 {
	// 	//lines = append(lines, string(bytes))
	// 	log.Println(string(bytes))
	// 	log.Println("ddddd")
	// }

	// rd := bufio.NewReader(fi)
	// for {
	// 	path, err := rd.ReadString(10) // 0x0A separator = newline
	// 	if err == io.EOF {
	// 		break
	// 	} else if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	log.Println(path)
	// }

	// reader := bufio.NewReader(fi)
	// buffer := bytes.NewBuffer(make([]byte, 0))
	// for {
	// 	part, prefix, err := reader.ReadLine()
	// 	if err != nil {
	// 		break
	// 	}
	// 	if !prefix {
	// 		buffer.Write(part)
	// 		log.Println(buffer.String())
	// 		buffer.Reset()
	// 	}
	// }
	// if err == io.EOF {
	// 	err = nil
	// }

	start = time.Now()
	seq := 0
	lineSegementoA := ""
	lineSegementoB := ""

	scanner := bufio.NewScanner(fi)
	for scanner.Scan() {

		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		lineTpReg := string(line[7:8])          // Tipo de Registro
		lineCodSegemento := string(line[13:14]) // Código Segmento do Registro Detalhe
		lineNroDocCli := ""                     // Nro. do Documento Cliente
		lineCodBcoFav := ""                     // Código do Banco Favorecido
		lineNmFav := ""                         // Nome do Favorecido

		if lineTpReg != "0" && lineTpReg != "1" && lineTpReg != "3" && lineTpReg != "5" && lineTpReg != "9" {
			continue // raise Exception.Create('Tipo de Registro invalido');
		}

		// registro.IdArqv := AObjArqvREM.Id;
		// registro.CdLegado := AObjArqvREM.CdLegado;
		// registro.Tipo := sLinhaTpReg;
		// registro.Situacao := 'PD'; // PD - Mensagem Pendente

		switch lineTpReg {
		case "0": // 0 - Header do Arquivo
			lineCodSegemento = ""
			lineSegementoA = line
			lineSegementoB = ""
		case "1": // 1 - Header de Lote
			seq++
			lineCodSegemento = ""
			lineSegementoA = line
			lineSegementoB = ""
		case "3": // 3 - Registro de Detalhe Segmento A e B

			if lineCodSegemento == "A" {
				lineSegementoA = line
				lineSegementoB = ""
				fmt.Println("3 - Registro de Detalhe Segmento A e B = ", lineCodSegemento)
			}

			if lineCodSegemento == "B" {

				lineNroDocCli := string(strings.TrimSpace(lineSegementoA[73:93])) // Nro. do Documento Cliente
				lineCodBcoFav := string(strings.TrimSpace(lineSegementoA[20:23])) // Código do Banco Favorecido
				lineNmFav := string(strings.TrimSpace(lineSegementoA[43:73]))     // Nome do Favorecido
				fmt.Println("3 - Registro de Detalhe Segmento A e B = ", "*"+lineNroDocCli+"*", "+"+lineCodBcoFav+"+", "-"+lineNmFav+"-")

				//   registro.NumCtrIF := lineNroDocCli;
				// if lineNroDocCli == "" {
				// registro.Situacao := '01'; // 01 - Mensagem não Gerada (falta de informação, formato invalido...)
				// FADConexaoThread.Monitor.Inserir(copy(BlidProc, 1, 20), '2', 'Registro não será integrado - NumCtrIF não informado.');
				// }

				// if registro.Situacao == "PD" { // PD - Mensagem Pendente
				// 	if lineNmFav == '' {
				// 	  registro.Situacao := '01'; // 01 - Mensagem não Gerada (falta de informação, formato invalido...)
				// 	  FADConexaoThread.Monitor.Inserir(copy(BlidProc, 1, 20), '2', 'Registro não será integrado - Nome do Favorecido não informado ( NumCtrIF/NroDocCli: ' + sLinhaNroDocCli + ' ).');
				// 	}
				// }

				// if registro.Situacao == "PD" { // PD - Mensagem Pendente
				// 	if registro.BuscarNumCtrlIF(QrySelNumCtrlIFReg, True) then
				// 	  registro.Situacao := '02'; // 02 - Mensagem Duplicada
				// }

				// if registro.Situacao == "PD" { // PD - Mensagem Pendente
				// 	if (lineCodBcoFav != '') And QrySelLstCodCompe.LocateEx('CDCOMPE = ' + QuotedStr(Zeros(sLinhaCodBcoFav, 3))) {
				// 	  registro.ISPBIFCred := Zeros(QrySelLstCodCompe.FieldByName('ISPB').AsString, 8);
				// 	  registro.ISPBIFCredCorretora := QrySelLstCodCompe.FieldByName('STCORRETORA').AsString;
				// 	}else{
				// 	  registro.Situacao := '01'; // 01 - Mensagem não Gerada (falta de informação, formato invalido...)
				// 	  FADConexaoThread.Monitor.Inserir(copy(BlidProc, 1, 20), '2', 'Registro não será integrado - Codigo Compe não localizado( CodCompe/CodBcoFav: ' + sLinhaCodBcoFav + ' - NumCtrIF/NroDocCli: ' + sLinhaNroDocCli + ' ).');
				// 	}
				// }
			}

		case "5": // 5 - Trailer de Lote
			lineCodSegemento = ""
			lineSegementoA = line
			lineSegementoB = ""
		case "9": // 9 - Trailer do Arquivo
			lineCodSegemento = ""
			lineSegementoA = line
			lineSegementoB = ""
		default:
			continue
		}

		seq++
		// registro.Seq := seq;
		// registro.SegementoA := lineSegementoA;
		// registro.SegementoA := lineSegementoB;
		// if not registro.Inserir(QryInsRegistro) then
		//   raise Exception.Create('Registro de Trailer do Arquivo não Inserido');
		fmt.Println("3 - Registro de Detalhe Segmento A e B", lineCodSegemento, seq, lineSegementoA, lineSegementoB, lineTpReg, lineNroDocCli, lineCodBcoFav, lineNmFav)

	}
	log.Printf("Tempo de Leitura %s", time.Since(start))

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	// start = time.Now()
	// ch := make(chan string)
	// //readerr := make(chan error)
	// //done := make(chan bool)
	// go func() {
	// 	// fi := strings.NewReader(filename)
	// 	scanner := bufio.NewScanner(fi)
	// 	for scanner.Scan() {
	// 		ch <- strings.TrimSpace(scanner.Text())
	// 	}
	// 	// if err := scanner.Err(); err != nil {
	// 	// 	readerr <- err
	// 	// }
	// 	close(ch)
	// 	// done <- true
	// }()

	// // var wg sync.WaitGroup
	// wg := sync.WaitGroup{}
	// wg.Add(1)

	// go func() {
	// 	for v := range ch {
	// 		v = v
	// 		// fmt.Println(v)
	// 		// wg.Add(1)
	// 		// go getCords(cords, wg)
	// 	}
	// 	wg.Done()
	// 	// for {
	// 	// 	select {
	// 	// 	case name := <-names:
	// 	// 		// Process each line
	// 	// 		fmt.Println(name)
	// 	// 	case err := <-readerr:
	// 	// 		log.Fatal(err)
	// 	// 	case <-done:
	// 	// 		// close(names)
	// 	// 		// close(readerr)
	// 	// }
	// 	// 		break //
	// 	// 	}
	// 	// }
	// }()

	// wg.Wait()
	// log.Printf("Tempo de Leitura %s", time.Since(start))

	fmt.Print("FIM")
}
