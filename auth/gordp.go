package auth

import (
	"time"

	"github.com/BlackSnowDot0/gordp/core"
)

type Option struct {
	Addr     string
	UserName string
	Password string

	ConnectTimeout time.Duration
}

type Client struct {
	option Option

	// conn   net.Conn // TCP连接
	stream *core.Stream

	// from negotiation
	selectProtocol uint32 // 协商RDP协议，0:rdp, 1:ssl, 2:hybrid
	userId         uint16
	shareId        uint32
	serverVersion  uint32 // 服务端RDP版本号
}

func NewClient(opt *Option) *Client {
	c := &Client{
		option: Option{
			Addr:           opt.Addr,
			UserName:       opt.UserName,
			Password:       opt.Password,
			ConnectTimeout: opt.ConnectTimeout,
		},
	}

	if c.option.ConnectTimeout == 0 {
		c.option.ConnectTimeout = 5 * time.Second
	}

	return c
}

// Connect
// https://www.cyberark.com/resources/threat-research-blog/explain-like-i-m-5-remote-desktop-protocol-rdp
func (c *Client) Connect() error {
	return core.Try(func() {
		defer func() { c.Close() }()
		c.stream = core.NewStream(c.option.Addr, c.option.ConnectTimeout)
		c.negotiation()
		c.basicSettingsExchange()
		c.channelConnect()
		c.sendClientInfo()
		c.readLicensing()
	})
}

func (c *Client) Close() {
	c.stream.Close()
}
