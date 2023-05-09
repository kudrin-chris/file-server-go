package FileParse

import (
	"encoding/csv"
	"fileServer/internal/app/model"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

var files []string

func findFile() {
	root := "./inputfiles"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			ext := filepath.Ext(path)
			if ext == ".tsv" {
				files = append(files, path)
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("no find files")
		panic(err)
	}
}

func parseFiles() {
	for x := range files {
		parse(files[x])
	}
}
func parse(pathstr string) {
	rec := model.Record{}
	dev := model.Device{}
	file_ := model.File_name{}
	//file_device := model.File_and_divice{}

	f, err := os.Open(pathstr)
	if err != nil {
		fmt.Println("error open file")
		log.Fatal(err)
	}
	r := csv.NewReader(f)
	r.Comma = '\t'              // Запятой считать табуляцию
	records, err := r.ReadAll() // прочитать всё содержимое файла за один раз
	if err != nil {
		file_.Flag_error = "false"
	}
	for lnum, line := range records {
		if lnum != 0 {
			N, err := strconv.Atoi(line[0])
			if err != nil {
				fmt.Println("error convert")
				log.Fatal(err)
			}
			rec.N = N
			rec.Mqtt = line[1]
			rec.Invid = line[2]

			Unig_guid := line[3]
			dev.Unig_guid = Unig_guid

			rec.Msg_id = line[4]
			rec.Text = line[5]
			rec.Context = line[6]
			rec.Class = line[7]
			L, err := strconv.Atoi(line[8])
			if err != nil {
				rec.Level = -1
			}
			rec.Level = L
			rec.Area = line[9]
			rec.Addr = line[10]
			rec.Block = line[11]
			rec.Type_ = line[12]
			rec.Bit_ = line[13]
			rec.Invert_bit = line[14]

			str := "./outputfiles/" + Unig_guid + ".rtf"
			f, err := os.OpenFile(str, os.O_APPEND|os.O_WRONLY, 0600) //открыть файл, если существует
			if err != nil {                                           //создать файл, если не существует
				f, er := os.Create(str)
				if er != nil {
					fmt.Println("error create to file")
					panic(er)
				}
				defer f.Close()
				if _, err = f.WriteString(model.ConvertString(rec)); err != nil {
					fmt.Println("error write to file")
					panic(err)
				}
			} else {
				defer f.Close()
				if _, err = f.WriteString(model.ConvertString(rec)); err != nil {
					fmt.Println("error write to file")
					panic(err)
				}
			}
		}
	}

}
