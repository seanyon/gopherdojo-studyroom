package main

import (
	"flag"
	"fmt"
	"gitlab.com/capskk/intern/gopherdojo/try53imageconversion/conversion"
	"os"
	"path/filepath"
	"strings"
)

var (
	extb string
	exta string
	dirpath string
	//afile string//変換後のファイル名
)

func main() {
	//オプション指定
	flag.StringVar(&extb, "b", "jpg", "変換前の拡張子の指定")
	flag.StringVar(&exta, "a", "png", "変換後の拡張子の指定")
	flag.StringVar(&dirpath, "d", "", "探索開始dir")
	flag.Parse()

	//-bで指定した、変換前の拡張子が対応しているか判断
	err := conversion.ExtensionCheck(extb)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	//-aで指定した、変換後の拡張子が対応しているか判断
	err = conversion.ExtensionCheck(exta)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println("変換中・・・")

	err = filepath.Walk(dirpath,
		func(path string, info os.FileInfo, err error) error {
			// フォルダならスキップ
			if info.IsDir() {
				return nil
			}

			a := strings.Split(path, ".")
			//拡張子がextbと一致していなかったら、スキップ
			if a[len(a)-1] != extb{
				fmt.Printf("%s 未変換\n", path)
				return nil
			}

			//正しいファイルなら変換
			a[len(a)-1] = exta
			afile := strings.Join(a,".")
			return conversion.FileConversion(path, afile)

		})
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

}




