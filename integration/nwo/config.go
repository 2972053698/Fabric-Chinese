
//此源码被清华学神尹成大魔王专业翻译分析并修改
//尹成QQ77025077
//尹成微信18510341407
//尹成所在QQ群721929980
//尹成邮箱 yinc13@mails.tsinghua.edu.cn
//尹成毕业于清华大学,微软区块链领域全球最有价值专家
//https://mvp.microsoft.com/zh-cn/PublicProfile/4033620
/*
版权所有IBM公司。保留所有权利。

SPDX许可证标识符：Apache-2.0
**/


package nwo

//config保存生成所需的基本信息
//结构配置文件。
type Config struct {
	Organizations []*Organization `yaml:"organizations,omitempty"`
	Consortiums   []*Consortium   `yaml:"consortiums,omitempty"`
	SystemChannel *SystemChannel  `yaml:"system_channel,omitempty"`
	Channels      []*Channel      `yaml:"channels,omitempty"`
	Consensus     *Consensus      `yaml:"consensus,omitempty"`
	Orderers      []*Orderer      `yaml:"orderers,omitempty"`
	Peers         []*Peer         `yaml:"peers,omitempty"`
	Profiles      []*Profile      `yaml:"profiles,omitempty"`
	Templates     *Templates      `yaml:"templates,omitempty"`
}

func (c *Config) RemovePeer(orgName, peerName string) {
	peers := []*Peer{}
	for _, p := range c.Peers {
		if p.Organization != orgName || p.Name != peerName {
			peers = append(peers, p)
		}
	}
	c.Peers = peers
}
