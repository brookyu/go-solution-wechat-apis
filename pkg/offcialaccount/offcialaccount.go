package offcialaccount

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/officialaccount"
	offConfig "github.com/silenceper/wechat/v2/officialaccount/config"
	"os"
)

var instance *OfficialAccountWorker

type OfficialAccountWorker struct {
	wc              *wechat.Wechat
	officialAccount *officialaccount.OfficialAccount
}

func NewOfficialAccountWorker(wc *wechat.Wechat) *OfficialAccountWorker {
	e := godotenv.Load()
	if e != nil {
		print(e)
	}

	cache := cache.NewMemcache()
	offCfg := &offConfig.Config{
		AppID:     os.Getenv("app_id"),
		AppSecret: os.Getenv("app_secret"),
		Token:     os.Getenv("token"),
		Cache:     cache,
	}
	account := wc.GetOfficialAccount(offCfg)
	return &OfficialAccountWorker{wc, account}
}
func (oaw *OfficialAccountWorker) Serve(c *gin.Context) {

}
func GetWorkerInstance() *OfficialAccountWorker {
	if instance == nil {
		wc := wechat.NewWechat()
		instance = NewOfficialAccountWorker(wc)
	}
	return instance

}
