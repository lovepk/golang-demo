package monster

import (
	"encoding/json"
	"io/ioutil"
)

type Monster struct {
	Name string
	Age int
	Phone string
}

func (m *Monster) Store() bool {
	// 序列化
	jsonStr, err := json.Marshal(m);if err != nil {
		println("序列化失败 err=", err)
		return false
	}
	//	保存到文件
	filePath := "/Users/wangjun/go/src/monster/data.json"
	err = ioutil.WriteFile(filePath, jsonStr, 0666); if err != nil {
		println("写入文件失败 err=", err)
		return false
	}
	return true
}


func (m *Monster) Restore() bool {
	// 文件读取序列化数据
	filePath := "/Users/wangjun/go/src/monster/data.json"
	data, err := ioutil.ReadFile(filePath); if err != nil {
		println("读取文件失败 err=", err)
		return false
	}
	err = json.Unmarshal(data, m);if err != nil {
		println("反序列化失败 err=", err)
		return false
	}
	return true
}

func main()  {
	
}
