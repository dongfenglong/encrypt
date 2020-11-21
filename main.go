package main

import (
	"github.com/dongfenglong/encrypt/utils"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

func main() {
	var dlg *walk.Dialog
	var te *walk.LineEdit
	var md5te *walk.LineEdit
	var gm3te *walk.LineEdit

	Dialog{
		AssignTo: &dlg,
		Title:    "Encrypt",
		Icon: "static/img/shell32_16782.ico",
		MinSize: Size{
			Width:  300,
			Height: 300,
		},
		Layout: VBox{},
		Children: []Widget{
			Composite{
				Layout: Grid{Columns: 2},
				Children: []Widget{
					Label{
						Text: "明文",
					},
					LineEdit{
						Text:     "",
						AssignTo: &te,
						OnMouseDown: func(x, y int, button walk.MouseButton) {
							if text, err := walk.Clipboard().Text(); err != nil {
								go func() {
									walk.MsgBox(nil, "复制失败", err.Error(), walk.MsgBoxIconError)
								}()
							} else {
								if len(text) > 0{
									te.SetText(text)
									utils.Encrypt(te, md5te, gm3te)
								}

							}
						},
					},
					Label{
						Text: "MD5",
					},
					LineEdit{
						MaxSize: Size{},
						MinSize: Size{
							Width:  120,
							Height: 30,
						},
						Text:     "",
						AssignTo: &md5te,
					},
					Label{
						Text: "GM3",
					},
					LineEdit{
						MaxSize: Size{},
						MinSize: Size{
							Width:  120,
							Height: 30,
						},
						Text:     "",
						AssignTo: &gm3te,
					},

				},
			},
		},
	}.Run(nil)
}
