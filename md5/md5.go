/*------------------------
name        md5
describe    wmntec md5 library
author      ailn(z.ailn@wmntec.com)
create      2016-04-18
version     1.0
------------------------*/
package md5

import (
	//golang official package
	"crypto/md5"
	"encoding/hex"
)

//get md5 string form bytes
func MD5(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
