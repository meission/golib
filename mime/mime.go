/*------------------------
name        mime
describe    meme type
author      ailn(z.ailn@wmntec.com)
create      2016-05-08
version     1.0
------------------------*/
package mime

type Mime interface {
	Get(string) string
}

type Text struct{}
type Image struct{}
type Audio struct{}
type Video struct{}
type Application struct{}

func (t Text) Get(key string) string {
	switch key {
	case "txt":
		return "text/plain"
	case "html":
		return "text/html"
	case "css":
		return "text/css"
	default:
		return "text/plain"
	}
}

func (i Image) Get(key string) string {
	switch key {
	case "bmp":
		return "image/bmp"
	case "gif":
		return "image/gif"
	case "jpeg":
		return "image/jpeg"
	case "jpg":
		return "image/jpeg"
	case "png":
		return "image/x-png"
	default:
		return ""
	}
}

func (a Audio) Get(key string) string {
	switch key {
	case "basic":
		return "audio/basic"
	case "mid":
		return "audio/mid"
	case "mpeg":
		return "audio/mpeg"
	case "wav":
		"audio/x-wav"
	default:
		return ""
	}
}

func (v Video) Get(key string) string {
	switch key {
	case "mpeg":
		return "video/mpeg"
	default:
		return ""
	}
}

func (a Application) Get(key string) string {
	switch key {
	case "js":
		return "application/x-javascript"
	default:
		return ""
	}
}
