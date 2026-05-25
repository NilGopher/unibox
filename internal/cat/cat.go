package cat

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func Run(rawArgs []string) (exitCode int) {
	args := opt.parse(rawArgs)

	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)
	defer writer.Flush()

	if len(args) == 0 {
		exitCode = process(writer, os.Stdin)
		return
	}

	for i, arg := range args {
		file, err := os.OpenFile(arg, os.O_RDONLY, 0)
		if err != nil {
			log.Print(err)
			exitCode = 1
			continue
		}

		if opt.number {
			if i > 0 {
				writer.WriteByte('\n')
			}

			writer.WriteString("       # ")
			writer.WriteString(file.Name())
			writer.WriteByte('\n')
		}

		exitCode = process(writer, file)

		file.Close()
	}

	return
}

func process(w *bufio.Writer, r io.Reader) (exitCode int) {
	reader := bufio.NewReaderSize(r, 256*1024)

	lineNum := 1
	startOfLine := true

	for {
		line, isPrefix, err := reader.ReadLine()

		if opt.number {
			if startOfLine {
				fmt.Fprintf(w, "%7d│", lineNum)
				lineNum++
				startOfLine = false
			}

			w.Write(line)

			if !isPrefix {
				w.WriteByte('\n')
				startOfLine = true
			}
		} else {
			w.Write(line)

			if !isPrefix {
				w.WriteByte('\n')
			}
		}

		if err != nil {
			if err != io.EOF {
				log.Print(err)
				exitCode = 1
			}
			break
		}
	}

	return
}
