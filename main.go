package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	//参数

	file, err := os.Open("banmo.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	defer file.Close()
	reader := csv.NewReader(file)
	reader.Comment = '#' //注释符号
	reader.Comma = ','   //数据分隔符号，默认为逗号

	//fout,err:=os.OpenFile("out.txt",os.O_RDWR|os.O_APPEND|os.O_CREATE,0666)
	// if err != nil {
	//     fmt.Println("Error:",err)
	//     return
	// }
	//defer fout.Close()

	k := 0
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if k == 0 {
			for i, v := range record {
				fmt.Printf("%d %s\n", i, v)
			}

		}
		if k == 1 {
			fmt.Println(record)
		}
		if k > 1 {

			// tmp := strings.Split(record[0], " ")
			// fmt.Println(tmp)
			fmt.Printf("line:%d length:%d \n", k, len(record))
		}
		k = k + 1
	}

	// fmt.Printf("\n")
	// fout.WriteString(content+"\n")
	// fmt.Printf("\n")
}
