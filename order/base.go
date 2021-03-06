package order

////////////////////////////////////////////////////////////////////////////////
// SS 订单状态
type SS uint8

const (
	SS01 SS = iota + 1 // 在线支付订单待支付；货到付款订单待确认
	SS02               // 备货中（只有此状态下，才可发货）
	SS03               // 已发货
	SS04               // 已取消
	SS05               // 已完成
	SS06               // 已发货 退货中-用户申请
	SS07               // 已发货 退货中-商家同意退货
	SS08               // 已发货 退货中-客服仲裁
	SS09               // 已发货 已关闭-退货失败
	SS10               // 已发货 退货中-客服同意
	SS11               // 已发货 退货中-用户已填物流
	SS12               // 已发货 退货成功-商户同意
	SS13               // 已发货 退货中-再次客服仲裁
	SS14               // 已发货 退货中-客服同意退款
	SS15               // 已发货 退货-用户取消
	SS16               // 未发货 退款中-用户申请
	SS17               // 未发货 退款中-商家同意
	SS21               // 未发货 退款成功-订单退款（用户申请退款，最终退款成功）
	SS22               // 已发货 退款成功-订单退款 （用户申请退货，最终退货退款成功）
	SS24               // 未发货 退货成功-商户已退款（特指货到付款订单，已发货时，用户申请退货，最终退货退款成功）
	SS25               // 未发货 退款中-用户取消
	SS26               // 未发货 退款中-商家拒绝
	SS27               // 已发货 退货中-等待买家处理（商家拒绝用户退货申请）
	SS28               // 已发货 退货失败（商家拒绝用户退货申请，客服仲裁支持商家）
	SS29               // 已发货 退货中-等待买家处理（用户填写退货物流，商家拒绝）
	SS30               // 已发货 退款中-退款申请（用户申请仅退款）
	SS31               // 已发货 退款申请取消（用户申请仅退款时，取消申请）
	SS32               // 已发货 退款中-商家同意（用户申请仅退款，商家同意申请）
	SS33               // 已发货 退款中-商家拒绝（用户申请仅退款，商家拒绝申请）
	SS34               // 已发货 退款中-客服仲裁（用户申请仅退款，商家拒绝申请，买家申请客服仲裁）
	SS35               // 已发货 退款中-客服同意（用户申请仅退款，商家拒绝申请，客服仲裁支持买家）
	SS36               // 已发货 退款中-支持商家（用户申请仅退款，商家拒绝申请，客服仲裁支持商家）
	SS37               // 已发货 已关闭-退款失败（已发货，用户申请仅退款时，退款关闭）
	SS38               // 已发货 退款成功-售后退款（特指货到付款订单，已发货，用户申请仅退款时，最终退款成功）
	SS39               // 已发货 退款成功-订单退款（已发货，用户申请仅退款时，最终退款成功）
)

////////////////////////////////////////////////////////////////////////////////
// OT 订单类型
type OT uint8

const (
	OT00 OT = 0 // 普通订单
	OT02 OT = 2 // 虚拟订单
	OT04 OT = 4 // 电子券
	OT05 OT = 5 // 三方核销
	OT06 OT = 6 // 服务市场
)

////////////////////////////////////////////////////////////////////////////////
// PT 支付方式
type PT uint8

const (
	COD    PT = iota // 货到付款
	WeChat           // 微信支付
	Alipay           // 支付宝支付
)

////////////////////////////////////////////////////////////////////////////////
// CT 优惠券类型
type CT uint8

const (
	PlatformDC CT = iota + 1  // 平台折扣券 (平台承担)
	PlatformDR                // 平台直减券 (平台承担)
	PlatformFR                // 平台满减券 (平台承担)
	CategoryDC CT = iota + 8  // 品类折扣券 (暂未开放)
	CategoryDR                // 品类直减券 (暂未开放)
	CategoryFR                // 品类满减券 (暂未开放)
	ShopDC     CT = iota + 15 // 店铺折扣券 (店铺承担)
	ShopDR                    // 店铺直减券 (店铺承担)
	ShopFR                    // 店铺满减券 (店铺承担)
	ChannelDC  CT = iota + 22 // 渠道折扣券 (平台承担)
	ChannelDR                 // 渠道直减券 (平台承担)
	ChannelFR                 // 渠道满减券 (平台承担)
	ProductDC  CT = iota + 29 // 单品折扣券 (店铺承担)
	ProductDR                 // 单品直减券 (店铺承担)
	ProductFR                 // 单品满减券 (店铺承担)
)

////////////////////////////////////////////////////////////////////////////////
// BT 订单APP渠道
type BT string

const (
	ZhanWai     BT = "0"  // 站外
	HuoShan     BT = "1"  // 火山
	DouYin      BT = "2"  // 抖音
	TouTiao     BT = "3"  // 头条
	XiGua       BT = "4"  // 西瓜
	WeiXin      BT = "5"  // 微信
	ShanGou     BT = "6"  // 闪购
	TouTiaoLite BT = "7"  // 头条lite版本
	DongCheDi   BT = "8"  // 懂车帝
	PiPiXia     BT = "9"  // 皮皮虾
	DouYinJS    BT = "10" // 抖音极速版
	TikTok      BT = "11" // 抖音海外版
)

////////////////////////////////////////////////////////////////////////////////
// SBT 订单来源类型
type SBT uint8

const (
	Unknown    SBT = iota // 未知
	App                   // app
	XiaChengXu            // 小程序
	H5                    // H5
)

////////////////////////////////////////////////////////////////////////////////
// CB 订单业务类型
type CB uint8

const (
	LuBanGuangGao CB = 1  // 鲁班广告
	LianMeng      CB = 2  // 联盟
	ShangCheng    CB = 4  // 商城
	ZiYing        CB = 8  // 自主经营
	XianSuoTong   CB = 10 // 线索通支付表单
	MenDian       CB = 12 // 抖音门店
	DouJia        CB = 14 // 抖+
	ChuanShanJia  CB = 15 // 穿山甲
)
