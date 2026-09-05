package main

import "testing"

// 测试 Add 函数：名字必须 Test 开头，参数必须是 t *testing.T
func TestAdd(t *testing.T) {
	got := Add(2, 3)
	want := 5
	if got != want {
		// 报告失败，并说明期望值和实际值
		t.Errorf("Add(2, 3) = %d; 期望 %d", got, want)
	}
}

// 可以写多个测试函数，go test 会全部运行
func TestAddNegative(t *testing.T) {
	if Add(-1, -2) != -3 {
		t.Errorf("Add(-1, -2) 结果不对")
	}
}
