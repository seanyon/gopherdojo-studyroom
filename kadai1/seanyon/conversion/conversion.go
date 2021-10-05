/*
Conversion は 画像の拡張子の変更を行うためのパッケージです。
*/
package conversion

import (
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"
)

// A Convert image interface
type Conv struct {
	FilePath string
	SrcExt   string
	DestExt  string
	image    image.Image
}

// 指定した拡張子が対応しているか判断
func ExtensionCheck(ext string) error {
	switch ext {
	case "jpeg", "jpg", "gif", "png":
		return nil
	default:
		return errors.New("指定できない拡張子です" + ext)
	}
}

//FileConversion
func FileConversion(imagpath, afile string) error {
	//ファイルかどうか存在するか
	f, err := os.Stat(imagpath)
	if os.IsNotExist(err) || f.IsDir() {
		return errors.New("ファイルが存在しません")
	}

	//open
	reader, err := os.Open(imagpath)
	if err != nil {
		return errors.New("ファイルが開けません")
	}
	defer reader.Close()

	//decode
	img, _, err := image.Decode(reader)
	if err != nil {
		return errors.New("Decode失敗")
	}

	//create
	output, err := os.Create(afile)
	if err != nil {
		return errors.New("Create失敗")
	}
	defer output.Close()

	//encode処理 変換後のextensionによってencode分岐
	switch strings.ToLower(filepath.Ext(afile)) { //大文字も対象に
	case ".jpeg", ".jpg":
		err = jpeg.Encode(output, img, nil)
		if err != nil {
			return errors.New("Encode失敗")
		}
		fmt.Printf("%s→%s:変換成功\n", imagpath, afile)
		return nil
	case ".gif":
		err = gif.Encode(output, img, nil)
		if err != nil {
			return errors.New("Encode失敗")
		}
		fmt.Printf("%s→%s:変換成功\n", imagpath, afile)
		return nil
	case ".png":
		err = png.Encode(output, img)
		if err != nil {
			return errors.New("Encode失敗")
		}
		fmt.Printf("%s→%s:変換成功\n", imagpath, afile)
		return nil
	}
	return nil
}
