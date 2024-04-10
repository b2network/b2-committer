package handler

import (
	"fmt"
	"testing"
	"time"
)

func TestReadFile(t *testing.T) {
	//content, _ := os.ReadFile("../../batchFile/56908-56949.json")
	////fmt.Println(content)
	//seq := make(map[uint64][]byte)
	//json.Unmarshal(content, &seq)
	//for k, v := range seq {
	//	fmt.Println(k)
	//	fmt.Println(string(v))
	//}
	fmt.Println(time.Now().Format("2006-01-02"))
}
