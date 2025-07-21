package automation_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
 * @Author: zhenzhongzhu
 * @Description:testing包下的单元测试范例
 * @File: demo_test
 */

func TestAdd(t *testing.T) {
	result := Add(1, 2)
	if result != 3 {
		t.Errorf("Add(1, 2) failed. Expected 3, got %d", result)
	}
}

func TestEquals(t *testing.T) {
	a := 2
	var b = 2
	c := 3
	assert.Equal(t, a, b, "比较相等")
	assert.NotEqual(t, a, c, "数字不相等")
	assert.Greater(t, a, c, "大于测试失败，请检查")
}

func Add(a, b int) int {
	return a + b
}
