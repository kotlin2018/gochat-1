package mch

// 交易类型
const (
	TradeAPP      = "APP"      // app支付
	TradeJSAPI    = "JSAPI"    // JSAPI支付（或小程序支付）
	TradeMWEB     = "MWEB"     // H5支付
	TradeNative   = "NATIVE"   // Native支付
	TradePAP      = "PAP"      // 签约续费
	TradeMicroPay = "MICROPAY" // 付款码支付
)

// 签名类型
type SignType string

const (
	SignMD5        SignType = "MD5"
	SignHMacSHA256 SignType = "HMAC-SHA256"
)

// 返回结果
const (
	ResultSuccess = "SUCCESS"
	ResultFail    = "FAIL"
	ResultNull    = "RESULT NULL" // 查询结果为空
	NotFound      = "NOT_FOUND"   // 数据不存在
	SystemError   = "SYSTEMERROR" // 系统繁忙，请稍后再试
)

const (
	TradeStateSuccess = "SUCCESS"    // 支付成功
	TradeStateRefund  = "REFUND"     // 转入退款
	TradeStateNotpay  = "NOTPAY"     // 未支付
	TradeStateClosed  = "CLOSED"     // 已关闭
	TradeStateRevoked = "REVOKED"    // 已撤销（刷卡支付）
	TradeStatePaying  = "USERPAYING" // 用户支付中
	TradeStateAccept  = "ACCEPT"     // 已接收，等待扣款
	TradeStateError   = "PAYERROR"   // 支付失败
	TradeStatePayFail = "PAY_FAIL"   // 支付失败(其他原因，如银行返回失败)
)

const (
	NoCredit         = "no_credit" // 指定不能使用信用卡支付
	CouponTypeCash   = "CASH"      // 充值代金券
	CouponTypeNoCash = "NO_CASH"   // 非充值优惠券
)

const (
	RefundStatusSuccess    = "SUCCESS"     // 退款成功
	RefundStatusClosed     = "REFUNDCLOSE" // 退款关闭
	RefundStatusProcessing = "PROCESSING"  // 退款处理中
	RefundStatusChange     = "CHANGE"      // 退款异常
)

const (
	RefundChannelOriginal      = "ORIGINAL"       // 原路退款
	RefundChannelBalance       = "BALANCE"        // 退回到余额
	RefundChannelOtherBalance  = "OTHER_BALANCE"  // 原账户异常退到其他余额账户
	RefundChannelOtherBankCard = "OTHER_BANKCARD" // 原银行卡异常退到其他银行卡
)

const (
	OrderNotExist  = "ORDERNOTEXIST"  // 订单不存在
	RefundNotExist = "REFUNDNOTEXIST" // 退款不存在
)

const (
	ContractAdd    = "ADD"    // 签约
	ContractDelete = "DELETE" // 解约
)

const (
	ContractEntrustUndo       = "1" // 未签约
	ContractEntrustOK         = "0" // 已签约
	ContractEntrustProcessing = "9" // 签约进行中
)

const (
	ContractDeleteUndo     = "0" // 未解约
	ContractDeleteExpired  = "1" // 有效期过自动解约
	ContractDeleteUser     = "2" // 用户主动解约
	ContractDeleteAPI      = "3" // 商户API解约
	ContractDeletePlatform = "4" // 商户平台解约
	ContractDeleteLogout   = "5" // 注销
	ContractDeleteContact  = "7" // 用户联系客服发起的解约
)

const (
	TransferNoCheck    = "NO_CHECK"    // 不校验真实姓名
	TransferForceCheck = "FORCE_CHECK" // 强校验真实姓名
)

const (
	TransferStatusProcessing = "PROCESSING" // 处理中
	TransferStatusSuccess    = "SUCCESS"    // 转账成功
	TransferStatusFailed     = "FAILED"     // 转账失败
	TransferStatusBankFail   = "BANK_FAIL"  // 银行退票
)

const (
	RedpackScene1 = "PRODUCT_1" // 商品促销
	RedpackScene2 = "PRODUCT_2" // 抽奖
	RedpackScene3 = "PRODUCT_3" // 虚拟物品兑奖
	RedpackScene4 = "PRODUCT_4" // 企业内部福利
	RedpackScene5 = "PRODUCT_5" // 渠道分润
	RedpackScene6 = "PRODUCT_6" // 保险回馈
	RedpackScene7 = "PRODUCT_7" // 彩票派奖
	RedpackScene8 = "PRODUCT_8" // 税务刮奖
)

const (
	RedpackStatusSending   = "SENDING"   // 发放中
	RedpackStatusSent      = "SENT"      // 已发放待领取
	RedpackStatusFailed    = "FAILED"    // 发放失败
	RedpackStatusReceived  = "RECEIVED"  // 已领取
	RedpackStatusRefunding = "RFUND_ING" // 退款中
	RedpackStatusRefund    = "REFUND"    // 已退款
)

const (
	RedpackTypeNormal = "NORMAL" // 普通红包
	RedpackTypeGroup  = "GROUP"  // 裂变红包
)

const (
	RedpackSendTypeAPI      = "API"      // 通过API接口发放
	RedpackSendTypeUpload   = "UPLOAD"   // 通过上传文件方式发放
	RedpackSendTypeActivity = "ACTIVITY" // 通过活动方式发放
)

const (
	WorkWXNormalMsg  = "NORMAL_MSG"   // 普通付款消息
	WorkWXAprovalMsg = "APPROVAL_MSG" // 审批付款消息
)

// 账单类型
const (
	BillTypeAll            = "ALL"             // 当日所有订单信息（不含充值退款订单）
	BillTypeSuccess        = "SUCCESS"         // 当日成功支付的订单（不含充值退款订单）
	BillTypeRefund         = "REFUND"          // 当日退款订单（不含充值退款订单）
	BillTypeRechargeRefund = "RECHARGE_REFUND" // 当日充值退款订单
)

// 资金账户类型
const (
	AccountTypeBasic     = "Basic"     // 基本账户
	AccountTypeOperation = "Operation" // 运营账户
	AccountTypeFees      = "Fees"      // 手续费账户
)
