# 巨量引擎MarketingAPI Golang SDK

- Oauth2 授权 (api/oauth)
  - 生成授权链接 [ Url(clt *core.SDKClient, redirectUrl string, state string) string ]
  - 获取AccessToken [ AccessToken(clt *core.SDKClient, authCode String) (*oauth.AccessTokenResponse, error) ]
  - 刷新Token [ RefreshToken(clt *core.SDKClient, refreshToken string) (*oauth.AccessTokenResponse, error)]
  - 获取已授权账户 [ Advertiser(clt *core.SDKClient, accessToken string) (*oauth.AdvertiserGetResponse, error) ]
  - 获取授权User信息 [ UserInfo(clt *core.SDKClient, accessToken string) (*oauth.UserInfoResponse, error) ]
- 账号服务
  - 广告主信息与资质管理 (api/advertiser)
    - 广告主信息 [ Info(clt *core.SDKClient, accessToken string, req *advertiser.InfoRequest) (*advertiser.InfoResponse, error) ]
    - 广告主公开信息 [ PublicInfo(clt *core.SDKClient, accessToken string, req *advertiser.PublicInfoRequest) (*advertiser.PublicInfoResponse, error) ]
    - 获取广告主头像信息 [ Avatar(clt *core.SDK, accessToken string, advertiserID uint64) (*advertiser.AvatarGetResponse, error) ] 
  - 代理商账号管理 (api/agent)
    - 广告主列表 [ AdvertiserSelect(clt *core.SDKClient, accessToken string, req *agent.AdvertiserSelectRequest) (*agent.AdvertiserSelectResponse, error) ] 
    - 修改广告主 [ AdvertiserUpdate(clt *core.SDKClient, accessToken string, req *agent.AdvertiserUpdateRequest) (*agent.AdvertiserUpdateResponse, error) ]
    - 二级代理商列表 [ ChildAgentSelect(clt *core.SDKClient, accessToken string, req *agent.ChildAgentSelectRequest) (*agent.ChildAgentSelectResponse, error) ]
    - 获取代理商信息 [ Info(clt *core.SDKClient, accessToken string, req *agent.InfoRequest) (*agent.InfoResponse, error) ] 
- 数据报表
  - 广告数据报表 (api/report)
    - 广告主数据 [ Advertiser(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponse, error) ]
    - 广告组数据 [ Campaign(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponse, error) ]
    - 广告计划数据 [ Ad(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponse, error) ]
    - 广告创意数据 [ Creative(clt *core.SDKClient, accessToken string, req *report.GetRequest) (*report.GetResponse, error) ]
    - 多合一数据报表接口 [ Integrated(clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponse, error) ]
    - 视频素材报表 [ Video(clt *core.SDKClient, accessToken string, req *report.IntegratedRequest) (*report.IntegratedResponse, error) ]
    - 视频互动流失数据 [ VideoFrame(clt *core.SDKClient, accessToken string, req *report.VideoFrameRequest) (*report.VideoFrameResponse, error) ]
    - 分级模糊数据 [ Misty(clt *core.SDKClient, accessToken string, req *report.VideoFrameRequest) (*report.VideoFrameResponse, error) ]
- 数据上报管理 (api/track)
  - 转化回传 [ Active(req *track.ActiveRequest) error ]