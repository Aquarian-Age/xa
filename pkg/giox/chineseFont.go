/*
 * Created by GoLand
 * User: Amrta
 * Mail: liangzi2021@yandex.com
 * Date:  2021年 7月 26日
 */

package giox

import (
	"fmt"
	"log"
	"os"

	"eliasnaur.com/font/roboto/robotoregular"
	"gioui.org/font/gofont"
	"gioui.org/font/opentype"
	"gioui.org/text"
	"gioui.org/widget/material"
	"github.com/gonoto/notosans"
)

//指定字体路径　"/path/xxx.ttf"
func FontX(fontPath string) *material.Theme {
	f, err := os.Open(fontPath)
	if err != nil {
		log.Fatal(err)
	}
	ttc, err := opentype.ParseCollectionReaderAt(f)
	if err != nil {
		log.Fatal(err)
	}
	th := material.NewTheme([]text.FontFace{{Face: ttc}})
	return th
}

//Roboto
func FontRoboto() *material.Theme {
	font := gofont.Collection()
	font = appendTTF(font, text.Font{Typeface: "Roboto"}, robotoregular.TTF)
	th := material.NewTheme(font)
	return th
}

//Noto
func FontNoto() *material.Theme {
	font := gofont.Collection()
	font = appendOTC(font, text.Font{Typeface: "Noto"}, notosans.OTC())
	th := material.NewTheme(font)
	return th
}

func appendTTF(collection []text.FontFace, fnt text.Font, ttf []byte) []text.FontFace {
	face, err := opentype.Parse(ttf)
	if err != nil {
		panic(fmt.Errorf("failed to parse font: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
}

func appendOTC(collection []text.FontFace, fnt text.Font, otc []byte) []text.FontFace {
	face, err := opentype.ParseCollection(otc)
	if err != nil {
		panic(fmt.Errorf("failed to parse font collection: %v", err))
	}
	return append(collection, text.FontFace{Font: fnt, Face: face})
}
