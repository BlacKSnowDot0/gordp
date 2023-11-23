package gordp

import (
	"github.com/Hypdncy/gordp/auth"
	"github.com/Hypdncy/gordp/core"
	"testing"
)

func TestRdpAuthConnect(t *testing.T) {
	client := auth.NewClient(&auth.Option{
		Addr:     "192.168.55.8:3389",
		UserName: "administrator",
		Password: "Admin@123",
	})
	core.ThrowError(client.Connect())
}
