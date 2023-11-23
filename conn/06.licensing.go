package conn

import (
	"github.com/Hypdncy/gordp/proto/pdu/licPdu"
)

func (c *Client) readLicensing() {
	licensing := licPdu.ServerLicensingPDU{}
	licensing.Read(c.stream)
}
