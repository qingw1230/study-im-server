package captcha

import (
	"github.com/mojocn/base64Captcha"
	"github.com/qingw1230/study-im-server/pkg/common/db"
)

var Captcha *MathCaptcha

func init() {
	Captcha = NewCaptcha()
}

// MathCaptcha 验证码工具类
type MathCaptcha struct {
	captcha *base64Captcha.Captcha
}

func NewCaptcha() *MathCaptcha {
	store := base64Captcha.DefaultMemStore

	height := 60
	width := 240
	noiseCnt := 0
	showLineOptions := base64Captcha.OptionShowHollowLine

	driver := base64Captcha.NewDriverMath(height, width, noiseCnt, showLineOptions, nil, nil, nil)

	return &MathCaptcha{
		captcha: base64Captcha.NewCaptcha(driver, store),
	}
}

func (m *MathCaptcha) Generate() (string, string) {
	id, b64s, ans, _ := m.captcha.Generate()
	db.DB.SetCheckCode(id, ans)
	return id, b64s
}

func (m *MathCaptcha) Verify(id, ans string) bool {
	redisAns, err := db.DB.GetCheckCode(id)
	if err != nil {
		return false
	}
	return redisAns == ans
}
