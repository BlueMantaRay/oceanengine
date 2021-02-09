package ad

import (
	"github.com/bububa/oceanengine/marketing-api/core"
	"github.com/bububa/oceanengine/marketing-api/model/ad"
)

// 更新计划出价
// 通过此接口用于更新广告计划的出价 bid， 暂不支持更新深度出价
// 一次可以处理100个计划；
// 修改的出价不能大于当前预算；
// 修改计划出价时可以参考：【建议出价】
func UpdateBid(clt *core.SDKClient, accessToken string, req *ad.UpdateBidRequest) (*ad.UpdateBidResponse, error) {
	var resp ad.UpdateBidResponse
	err := clt.Post("2/ad/update/bid/", req, &resp, accessToken)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}