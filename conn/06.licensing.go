package conn

import (
	"github.com/BlackSnowDot0/gordp/proto/pdu/licPdu"
)

func (c *Client) readLicensing() {
	licensing := licPdu.ServerLicensingPDU{}
	licensing.Read(c.stream)
}
