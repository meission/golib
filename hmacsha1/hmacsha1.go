/*------------------------
name        hmacsha1
describe    wmntec hmacsha1 library
author      ailn(z.ailn@wmntec.com)
create      2016-05-09
version     1.0
------------------------*/
package hmacsha1

import (
	"crypto/hmac"
	"crypto/sha1"
)

func Encrypt(secretKey string, data []byte) []byte {
	key := []byte(secretKey)
	mac := hmac.New(sha1.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}
