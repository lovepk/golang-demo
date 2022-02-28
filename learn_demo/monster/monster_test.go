package monster

import "testing"

func TestStore(t *testing.T)  {
	monster := Monster{
		Name: "王俊",
		Age: 28,
		Phone: "18021192217",
	}
	b := monster.Store();if !b {
		t.Fatalf("Store() 测试错误，希望为=%v,实际为=%v", true, b)
		return
	}
	t.Logf("Store() 测试成功")
}

// 一次测试所有 go test -v
// 单独测试某个函数命令  go test -v -test.run Restore
func TestRestore(t *testing.T)  {
	var monster Monster
	b := monster.Restore();if !b {
		t.Fatalf("Restore() 测试错误，希望为=%v,实际为=%v", true, b)
		return
	}
	if (monster.Name != "王俊") {
		t.Fatalf("Restore() 测试错误，希望为=%v,实际为=%v", "王俊", monster.Name)
	}
	t.Logf("Restore() 测试成功")
}