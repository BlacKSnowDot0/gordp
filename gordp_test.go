package gordp

import (
	"fmt"
	"github.com/Hypdncy/gordp/auth"
	"os"
	"testing"

	"github.com/Hypdncy/gordp/proto/bitmap"
)

type processor struct {
	i int
}

func (p *processor) ProcessBitmap(option *bitmap.Option, bitmap *bitmap.BitMap) {
	p.i++
	_ = os.MkdirAll("./png", 0755)
	_ = os.WriteFile(fmt.Sprintf("./png/%v.png", p.i), bitmap.ToPng(), 0644)
}

func TestRdpConnect(t *testing.T) {
	client := auth.NewClient(&auth.Option{
		Addr:     "10.226.239.200:3389",
		UserName: "administrator",
		Password: "[YourPasswordHere]",
	})
	err := client.Connect()
	fmt.Println("err")
	fmt.Println(err)
	if err != nil {
		return
	}
}
