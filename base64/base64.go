/*------------------------
name        base64
describe    wmntec base64 library
author      ailn(z.ailn@wmntec.com)
create      2016-05-09
version     1.0
------------------------*/
package base64

import (
	"encoding/base64"
)

func Encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Decode(baseStr string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(baseStr)
}
