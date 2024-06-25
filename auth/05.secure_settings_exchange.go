package auth

import (
	"github.com/BlackSnowDot0/gordp/proto/pdu/licPdu"
)

func (c *Client) sendClientInfo() {
	clientInfo := licPdu.NewClientInfoPDU(c.userId, c.option.UserName, c.option.Password)
	clientInfo.Write(c.stream)
}
