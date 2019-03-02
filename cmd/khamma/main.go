package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/ingtranet/khamma"
	"github.com/paulbellamy/ratecounter"
	"github.com/urfave/cli"
)

type analyzedResult struct {
	Input  *string              `json:"input"`
	Output []*khamma.KhaiiiWord `json:"output"`
}

var toAnalyze = make(chan *string)
var toWrite = make(chan *[]byte)

func handle(input *os.File, output *os.File) error {
	reader := bufio.NewReader(input)
	writer := bufio.NewWriter(output)
	counter := ratecounter.NewRateCounter(1 * time.Second)
	count := 0

	defer writer.Flush()

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}

		jsonResult, err := json.Marshal(khamma.Analyze(line, ""))
		if err != nil {
			return err
		}

		_, err = writer.Write(jsonResult)
		if err != nil {
			return err
		}

		counter.Incr(1)
		count++
		if 1000 == count {
			fmt.Println(counter.Rate(), count)
			count = 0
		}
	}
}


func main() {
	app := cli.NewApp()
	app.Name = "khamma"
	app.Usage = ""
	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "input, I"},
		cli.StringFlag{Name: "output, O"},
	}

	app.Action = func(c *cli.Context) error {
		khamma.InitializeWithDefault()

		var input *os.File
		var output *os.File
		var err error

		inputFlag := c.String("input")
		if inputFlag == "" {
			input = os.Stdin
		} else {
			input, err = os.Open(inputFlag)
			if err != nil {
				panic(nil)
			}
		}

		outputFlag := c.String("output")
		if outputFlag == "" {
			output = os.Stdout
		} else {
			output, err = os.Create(outputFlag)
			if err != nil {
				panic(nil)
			}
		}

		handle(input, output)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
