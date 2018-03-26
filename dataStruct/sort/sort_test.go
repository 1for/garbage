package sort

import (
	"testing"
	"math/rand"
	"time"
	"sort"
	"reflect"
//	"fmt"
	"github.com/mohae/deepcopy"
)

var testCase [][]int
var sortedCase [][]int

func init() {
	rand.Seed(time.Now().UnixNano())
	
	//随机生成测试用例
	//100以内数字，5组
	var i, j int
	var tmpArr, tmpA, tmpB []int;
	
	for i = 1; i < 6; i++ {
		//产生5组用例
		for j = 0; j < i*5; j++ {
			//每组有 i*5 个元素
			tmpArr = append(tmpArr, rand.Intn(1000))
		}

		tmpA = deepcopy.Copy(tmpArr).([]int)
		tmpB = deepcopy.Copy(tmpArr).([]int)
		
		testCase = append(testCase, tmpA)
		
		sort.Ints(tmpB)
		sortedCase = append(sortedCase, tmpB)

		tmpArr = tmpArr[:0]
	}


}

func TestInsertion(t *testing.T) {
	var i int;
	var tmpA, tmpB []int;
	for i = 0; i< len(testCase); i++ {
		tmpA = testCase[i];

		tmpB = InsertionSort(tmpA);

		if !reflect.DeepEqual(sortedCase[i], tmpB) {
			t.Error("sort Fail");
		}

		tmpA = tmpA[:0]
		tmpB = tmpB[:0]
	}
}

func TestShell(t *testing.T) {
	var i int;
	var tmpA, tmpB []int;
	for i = 0; i< len(testCase); i++ {
		tmpA = testCase[i];

		tmpB = ShellSort(tmpA);

		if !reflect.DeepEqual(sortedCase[i], tmpB) {
			t.Error("sort Fail");
		}

		tmpA = tmpA[:0]
		tmpB = tmpB[:0]
	}
}
