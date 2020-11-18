package offcialaccount

import (
	officialaccount "github.com/silenceper/wechat/v2/officialaccount"
	"github.com/silenceper/wechat/v2/officialaccount/user"
)

func (oaw *OfficialAccountWorker) GetAccount() *officialaccount.OfficialAccount {
	return oaw.officialAccount
}
func (oaw OfficialAccountWorker) GetUserInfo(openId string) (*user.Info, error) {
	return oaw.officialAccount.GetUser().GetUserInfo(openId)
}
