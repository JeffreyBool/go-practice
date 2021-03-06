/**
* Author: JeffreyBool
* Date: 2020/6/3
* Time: 14:33
* Software: GoLand
 */

package dfa

import (
	"os"
	"strings"
	"testing"

	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
)

const filterText = `我是需要过滤的陰戶，内容为：**阴@@毛，需要过滤。陰。唇。`

//
func TestDFA(t *testing.T) {
	str := " 你好呀"
	str1 := "   你好呀   "
	str2 := " 你好 呀 "
	t.Logf("%#v\n",str2)
	t.Logf("%#v\n",strings.ToUpper(strings.Trim(str," ")))
	t.Logf("%#v\n",strings.Trim(str1," "))
	t.Logf("%#v\n",strings.Trim(str2," "))

	file, err := os.Open("./data.txt")
	if err != nil {
		t.Error(err)
		return
	}
	memStore, err := store.NewMemoryStore(store.MemoryConfig{
		Reader: file,
		//DataSource: []string{"内容"},
	})
	if err != nil {
		panic(err)
	}
	filterManage := filter.NewDirtyManager(memStore)
	result, err := filterManage.Filter().Filter(filterText, '*', '@', '。')
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(result)
}
