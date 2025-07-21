package automation_test

import "testing"

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

func Add(a, b int) int {
	return a + b
}
