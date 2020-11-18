package middlewares

import (
	"github.com/brookyu/go-solution-wechat-apis/pkg/offcialaccount"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

const SCOPE_USERINFO = "snsapi_userinfo"
const CURRENT_WECHAT_USERINFO = "currentwechatuserinfo"

func WechatAuthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		currentWechatUserInfo := session.Get(CURRENT_WECHAT_USERINFO)
		if currentWechatUserInfo == nil {
			instance := offcialaccount.GetWorkerInstance()
			account := instance.GetAccount()

			oath := account.GetOauth()
			code := c.Param("code")
			redirectUrl := c.Request.RequestURI
			if code == "" && redirectUrl != "" {
				oath.Redirect(c.Writer, c.Request, redirectUrl, SCOPE_USERINFO, "brookyu")
			}
			ak, err := oath.GetUserAccessToken(code)
			if err != nil {
				userinfo, err2 := oath.GetUserInfo(ak.AccessToken, ak.OpenID)

				if err2 != nil {
					session.Set(CURRENT_WECHAT_USERINFO, userinfo)
				}

			}

		}

	}
}
