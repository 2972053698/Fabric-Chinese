
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


package mocks

import (
	"time"

	"github.com/hyperledger/fabric/msp"
	pmsp "github.com/hyperledger/fabric/protos/msp"
	"github.com/stretchr/testify/mock"
)

type MockMSP struct {
	mock.Mock
}

func (m *MockMSP) IsWellFormed(_ *pmsp.SerializedIdentity) error {
	return nil
}

func (m *MockMSP) DeserializeIdentity(serializedIdentity []byte) (msp.Identity, error) {
	args := m.Called(serializedIdentity)
	return args.Get(0).(msp.Identity), args.Error(1)
}

func (m *MockMSP) Setup(config *pmsp.MSPConfig) error {
	args := m.Called(config)
	return args.Error(0)
}

func (m *MockMSP) GetVersion() msp.MSPVersion {
	args := m.Called()
	return args.Get(0).(msp.MSPVersion)
}

func (m *MockMSP) GetType() msp.ProviderType {
	args := m.Called()
	return args.Get(0).(msp.ProviderType)
}

func (m *MockMSP) GetIdentifier() (string, error) {
	args := m.Called()
	return args.String(0), args.Error(1)
}

func (m *MockMSP) GetSigningIdentity(identifier *msp.IdentityIdentifier) (msp.SigningIdentity, error) {
	args := m.Called(identifier)
	return args.Get(0).(msp.SigningIdentity), args.Error(1)
}

func (m *MockMSP) GetDefaultSigningIdentity() (msp.SigningIdentity, error) {
	args := m.Called()
	return args.Get(0).(msp.SigningIdentity), args.Error(1)
}

func (m *MockMSP) GetTLSRootCerts() [][]byte {
	args := m.Called()
	return args.Get(0).([][]byte)
}

func (m *MockMSP) GetTLSIntermediateCerts() [][]byte {
	args := m.Called()
	return args.Get(0).([][]byte)
}

func (m *MockMSP) Validate(id msp.Identity) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockMSP) SatisfiesPrincipal(id msp.Identity, principal *pmsp.MSPPrincipal) error {
	args := m.Called(id, principal)
	return args.Error(0)
}

type MockIdentity struct {
	mock.Mock

	ID string
}

func (m *MockIdentity) Anonymous() bool {
	panic("implement me")
}

func (m *MockIdentity) ExpiresAt() time.Time {
	panic("implement me")
}

func (m *MockIdentity) GetIdentifier() *msp.IdentityIdentifier {
	args := m.Called()
	return args.Get(0).(*msp.IdentityIdentifier)
}

func (*MockIdentity) GetMSPIdentifier() string {
	panic("implement me")
}

func (m *MockIdentity) Validate() error {
	return m.Called().Error(0)
}

func (*MockIdentity) GetOrganizationalUnits() []*msp.OUIdentifier {
	panic("implement me")
}

func (*MockIdentity) Verify(msg []byte, sig []byte) error {
	panic("implement me")
}

func (*MockIdentity) Serialize() ([]byte, error) {
	panic("implement me")
}

func (m *MockIdentity) SatisfiesPrincipal(principal *pmsp.MSPPrincipal) error {
	return m.Called(principal).Error(0)
}

type MockSigningIdentity struct {
	mock.Mock
	*MockIdentity
}

func (*MockSigningIdentity) Sign(msg []byte) ([]byte, error) {
	panic("implement me")
}

func (*MockSigningIdentity) GetPublicVersion() msp.Identity {
	panic("implement me")
}
