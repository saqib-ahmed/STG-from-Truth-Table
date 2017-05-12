package main

import (

  "fmt"
  "io"
  "encoding/csv"
  "log"
  "os"
)

func main() {
  // setup reader
  csvIn, err := os.Open("truth_fsm.csv")
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
  fixedStr := "digraph g{\nrankdir=\"LR\";\nedge[splines=\"curved\"]\nnode [shape = doublecircle]; \"35_cent_S_2\\n1101\" \"35_cent_S\\n1001\" \"40_cent_S\\n1010\" \"40_cent_S_2\\n1101\" \"45_cent_S\\n1011\" \"45_cent_S_2\\n1110\" \"50_cent_S\\n1111\";\nnode [shape = circle];\n"
  csvOut.WriteString(fixedStr)
  fmt.Printf("File is: %s", r)
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
    // record is an array of string so is directly printable
    fmt.Println("Record", lineCount, "is", record, "and has", len(record), "fields")
    // and we can iterate on top of that
	//output :=fmt.Sprintf("%c%c%c%c",record[5][1],record[5][3],record[5][2],record[5][4])
    str := fmt.Sprintf("\"%s\"->\"%s\"[label=%s];\n",record[1],record[4],record[2])
    csvOut.WriteString(str)
    fmt.Println()
    lineCount += 1
  }
  csvOut.WriteString("}")
  
}
