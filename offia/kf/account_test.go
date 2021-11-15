package kf

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/shenghui0779/gochat/offia"
	"github.com/shenghui0779/gochat/wx"
	"github.com/shenghui0779/yiigo"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/customservice/getkflist?access_token=ACCESS_TOKEN").Return([]byte(`{
		"kf_list" : [
		   {
			  "kf_account": "test1@test",
			  "kf_headimgurl": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			  "kf_id": "1001",
			  "kf_nick": "ntest1",
			  "kf_wx": "kfwx1"
		   },
		   {
			  "kf_account": "test2@test",
			  "kf_headimgurl": "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			  "kf_id": "1002",
			  "kf_nick": "ntest2",
			  "invite_wx": "kfwx2",
			  "invite_expire_time": 123456789,
			  "invite_status": "waiting"
		   }
		]
	}`), nil)

	oa := offia.New("APPID", "APPSECRET")
	oa.client = client

	result := make([]*Account, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetAccountList(&result))

	assert.Nil(t, err)
	assert.Equal(t, []*Account{
		{
			ID:         "1001",
			Account:    "test1@test",
			Nickname:   "ntest1",
			HeadImgURL: "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			Weixin:     "kfwx1",
		},
		{
			ID:               "1002",
			Account:          "test2@test",
			Nickname:         "ntest2",
			HeadImgURL:       "http://mmbiz.qpic.cn/mmbiz/4whpV1VZl2iccsvYbHvnphkyGtnvjfUS8Ym0GSaLic0FD3vN0V8PILcibEGb2fPfEOmw/0",
			InviteWeixin:     "kfwx2",
			InviteExpireTime: 123456789,
			InviteStatus:     InviteWaiting,
		},
	}, result)
}

func TestGetOnlineList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/cgi-bin/customservice/getonlinekflist?access_token=ACCESS_TOKEN").Return([]byte(`{
		"kf_online_list": [
			{
				"kf_account": "test1@test",
				"status": 1,
				"kf_id": "1001",
				"accepted_case": 1
			},
			{
				"kf_account": "test2@test",
				"status": 1,
				"kf_id": "1002",
				"accepted_case": 2
			}
		]
	}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	result := make([]*Online, 0)

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", GetOnlineList(&result))

	assert.Nil(t, err)
	assert.Equal(t, []*Online{
		{
			ID:           "1001",
			Account:      "test1@test",
			Status:       1,
			AcceptedCase: 1,
		},
		{
			ID:           "1002",
			Account:      "test2@test",
			Status:       1,
			AcceptedCase: 2,
		},
	}, result)
}

func TestAddAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/add?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","nickname":"客服1"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", AddAccount("test1@test", "客服1"))

	assert.Nil(t, err)
}

func TestUpdateAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/update?access_token=ACCESS_TOKEN", []byte(`{"kf_account":"test1@test","nickname":"客服1"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UpdateAccount("test1@test", "客服1"))

	assert.Nil(t, err)
}

func TestInviteWorker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Post(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/inviteworker?access_token=ACCESS_TOKEN", []byte(`{"invite_wx":"test_kfwx","kf_account":"test1@test"}`)).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", InviteWorker("test1@test", "test_kfwx"))

	assert.Nil(t, err)
}

func TestUploadAvatar(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Upload(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/uploadheadimg?access_token=ACCESS_TOKEN&kf_account=ACCOUNT", gomock.AssignableToTypeOf(yiigo.NewUploadForm())).Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", UploadAvatar("ACCOUNT", "../test/test.jpg"))

	assert.Nil(t, err)
}

func TestDeleteAccount(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	client := wx.NewMockClient(ctrl)

	client.EXPECT().Get(gomock.AssignableToTypeOf(context.TODO()), "https://api.weixin.qq.com/customservice/kfaccount/del?access_token=ACCESS_TOKEN&kf_account=ACCOUNT").Return([]byte(`{"errcode":0,"errmsg":"ok"}`), nil)

	oa := New("APPID", "APPSECRET")
	oa.client = client

	err := oa.Do(context.TODO(), "ACCESS_TOKEN", DeleteAccount("ACCOUNT"))

	assert.Nil(t, err)
}
