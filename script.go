package main

import (

  "fmt"
  "io"
  "encoding/csv"
  "log"
  "os"
  "os/exec"
  "sync"
  "strings"
)

func main() {
  // setup reader
  dpi :=""
  file := ""
  if len(os.Args)<2 {
    log.Fatal("Please provide the input file.\n" +
              "go run script.go <file.csv> <dpi resolution>")
    dpi = "150"
    file = "truth_fsm.csv"
  } else if len(os.Args)<3 || os.Args[2]< "20" || os.Args[2]> "1000"{
    dpi = "150"
    file = os.Args[1]
    } else {
    dpi = os.Args[2]
    file = os.Args[1]
  }
  fmt.Println(file)
  csvIn, err := os.Open(file)
  if err != nil {
    log.Fatal(err)
  }
  r := csv.NewReader(csvIn)

  // setup writer
  csvOut, err := os.Create("out.dot")
  if err != nil {
    log.Fatal("Unable to open output")
  }

  lineCount := 0
  fixedStr := fmt.Sprintf("digraph g{\nrankdir=\"LR\";\ngraph [dpi=%s];\nedge" +
    "[splines=\"curved\"]\nnode [shape = doublecircle]; \"30_cent_S\\n1000\" \""+
    "35_cent_S_2\\n1101\" \"35_cent_S\\n1001\" \"40_cent_S\\n1010\" \"40_cent_S_2"+
    "\\n1101\" \"45_cent_S\\n1011\" \"45_cent_S_2\\n1110\" \"50_cent_S\\n1111\";"+
    "\nnode [shape = circle];\n", dpi)
  csvOut.WriteString(fixedStr)
  for {
    // read just one record, but we could ReadAll() as well
    record, err := r.Read()
    // end-of-file is fitted into err
    if err == io.EOF {
      break
    } else if err != nil {
      fmt.Println("Error:", err)
      return
    }
	if lineCount ==0 || lineCount == 1{
		lineCount +=1;
		continue;
	}

	//output :=fmt.Sprintf("%c%c%c%c",record[5][1],record[5][3],record[5][2],record[5][4])
    str := fmt.Sprintf("\"%s\"->\"%s\"[label=%s];\n",record[1],record[4],record[2])
    csvOut.WriteString(str)
    fmt.Println()
    lineCount += 1
  }
  csvOut.WriteString("}")

  // fmt.Println("BAR:", os.Getenv("PATH"))
  // cmd := exec.Command("dot -Tpng out.dot -o out.png")
  // var out bytes.Buffer
  // cmd.Stdout = &out
  // err1 := cmd.Run()
  // if err1 != nil {
  //   log.Fatal(err1)
  // }

  wg := new(sync.WaitGroup)
  dot := fmt.Sprintf("dot -Tpng out.dot -o out_%sdpi.png", dpi)
  wg.Add(1)

  x := []string{dot}
  go exe_cmd(x[0], wg)
  wg.Wait()

}

func exe_cmd(cmd string, wg *sync.WaitGroup) {
  fmt.Println("Executing: ",cmd)
  // splitting head => g++ parts => rest of the command
  parts := strings.Fields(cmd)
  head := parts[0]
  parts = parts[1:len(parts)]

  out, err := exec.Command(head,parts...).Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Printf("%s", out)
  wg.Done() // Need to signal to waitgroup that this goroutine is done
}
