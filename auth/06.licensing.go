package auth

import (
	"fmt"
	"github.com/BlackSnowDot0/gordp/proto/pdu/licPdu"
)

func (c *Client) readLicensing() {
	licensing := licPdu.ServerLicensingPDU{}
	licensing.Read(c.stream)
	fmt.Printf("licensing: %#v\n", licensing)
}
