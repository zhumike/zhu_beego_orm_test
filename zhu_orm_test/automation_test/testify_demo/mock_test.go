package testify_demo

/**
 * @Author: zhenzhongzhu
 * @Description:mock方法测试
 * @File: mock_test
 * @Date: 2025/7/21 下午9:44
 */

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

// 构建mock对象
type MyMockedObj struct {
	mock.Mock
}

// mock接口逻辑
func (m *MyMockedObj) DoSomething(number int) (bool, error) {
	args := m.Called(number)
	return args.Bool(0), args.Error(1)
}

func (m *MyMockedObj) GetOrderNumberInfo(orderId int) (int, error) {
	args := m.Called(orderId)
	return args.Int(0), args.Error(1)
}

func TestSomething(t *testing.T) {
	//实例化
	testObj := new(MyMockedObj)

	//设置mock期望方法，当入参是100，返回true
	testObj.On("DoSomething", 100).Return(true, nil)
	testObj.On("GetOrderNumberInfo", 100).Return(1001200, nil)

	//调用mock对象的方法
	something, err := testObj.DoSomething(100)
	orderNumber, _ := testObj.GetOrderNumberInfo(100)

	//断言
	assert.NoError(t, err)
	assert.True(t, something)
	assert.Equal(t, 1001200, orderNumber)
	assert.Greater(t, orderNumber, 0)
	assert.NotEqual(t, 230, orderNumber)
	testObj.AssertExpectations(t)
}
