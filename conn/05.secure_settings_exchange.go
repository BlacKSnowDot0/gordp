package conn

import (
	"github.com/Hypdncy/gordp/proto/pdu/licPdu"
)

func (c *Client) sendClientInfo() {
	clientInfo := licPdu.NewClientInfoPDU(c.userId, c.option.UserName, c.option.Password)
	clientInfo.Write(c.stream)
}
