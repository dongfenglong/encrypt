package ea

import (
	"fmt"
	"strings"
	"github.com/tjfoc/gmsm/sm3"
)

func GM3(plainText string) (cipherText string) {
	h := sm3.New()
	h.Write([]byte(plainText))
	sum := h.Sum(nil)
	return strings.ToUpper(fmt.Sprintf("%x", sum))
}
