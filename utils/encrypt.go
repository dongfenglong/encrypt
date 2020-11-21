package utils

import (
	"github.com/dongfenglong/encrypt/ea"
	"github.com/lxn/walk"
)
//var wg sync.WaitGroup
func Encrypt(te, md5te, gm3te *walk.LineEdit) {
	text := te.Text()
	if len(text) == 0{
		return
	}
	//wg.Add(2)
	//go func() {
		md5te.SetText(ea.MD5(text))
		//wg.Done()
	//}()
	//go func() {
		gm3te.SetText(ea.GM3(text))
	//	wg.Done()
	//
	//}()
	//wg.Wait()
}
