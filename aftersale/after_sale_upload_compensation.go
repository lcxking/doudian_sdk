package aftersale

import "github.com/lcxking/doudian_sdk/unit"

type ArgAfterSaleUploadCompensation struct {
	OrderID  unit.OrderID `paramName:"order_id"` // 子订单ID
	Comment  string       `paramName:"comment"`  // type = 2 时需要选择拒绝原因
	Evidence string       `paramName:"evidence"` // type = 2 凭证图片（货到付款订单，必填）
}
