package tool

import (
	crand "crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"math/rand"
	"strings"
	"sync"
	"time"
)

var (
	timeSource rand.Source
)

func init() {
	timeSource = rand.NewSource(time.Now().UnixNano())
}

func MakeTransferrder() string {
	return "T" + SnowflakeId()
}

func MakeRechageOrder() string {
	return fmt.Sprintf("R%s%d", time.Now().Format("20060102150405"), RandInt(1000, 9999))
}

func MakeActivityOrder() string {
	//orderNoKey := time.Now().Format("20060102")
	//orderNoPre := time.Now().Format("20060102150405")
	//orderNoAdd := global.G_REDIS.Incr(fmt.Sprintf("DEPOSIT_NO:%s", orderNoKey)).Val()
	//orderSn := orderNoPre + RandString(3) + String(orderNoAdd)
	return ""
}

func MakeCollectOrder() string {
	return fmt.Sprintf("C%s%d", time.Now().Format("20060102150405"), RandInt(1000, 9999))
}

func MakePassWord() string {
	return fmt.Sprintf("%s%d", RandString(4), RandInt(1000, 9999))
}

var mutex sync.Mutex

func MakeCollectFeeOrder() string {
	mutex.Lock()
	xxx := fmt.Sprintf("F%s%d", time.Now().Format("20060102150405"), RandInt(1000, 9999))
	mutex.Unlock()
	return xxx
}

func RandString(len int) string {

	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
		time.Sleep(time.Millisecond)
	}
	return strings.ToLower(string(bytes))
}

func RandInt(min int, max int) int {
	rand.Seed(time.Now().Unix())
	return min + rand.Intn(max-min)
}

func BigIntStringAdd(numstr string, num string) (string, error) {
	if !strings.Contains(numstr, "-") && strings.Contains(num, "-") {
		return BigIntStringReduce(numstr, strings.Replace(num, "-", "", 1))
	}

	if strings.Contains(numstr, "-") && !strings.Contains(num, "-") {
		return BigIntStringReduce(num, strings.Replace(numstr, "-", "", 1))
	}

	isFu := false
	if strings.Contains(numstr, "-") && strings.Contains(num, "-") {
		isFu = true
		numstr = strings.Replace(numstr, "-", "", 1)
		num = strings.Replace(num, "-", "", 1)
	}

	if numstr == "" {
		numstr = "0"
	}

	if num == "" {
		num = "0"
	}

	//如果有小数点 去除小数点后面的
	numstrIndex := strings.Index(numstr, ".")
	if numstrIndex != -1 {
		numstr = numstr[0:numstrIndex]
	}

	numIndex := strings.Index(num, ".")
	if numIndex != -1 {
		num = num[0:numIndex]
	}

	n, nok := new(big.Int).SetString(numstr, 10)
	if !nok {
		return "", errors.New("big.Int first param SetString fail")
	}
	m, mok := new(big.Int).SetString(num, 10)
	if !mok {
		return "", errors.New("big.Int second param SetString fail")
	}
	m.Add(n, m)

	if isFu {
		return "-" + m.String(), nil
	}
	return m.String(), nil
}

func BigIntStringReduce(numstr string, num string) (string, error) {

	if numstr == "" {
		numstr = "0"
	}

	if num == "" {
		num = "0"
	}

	//如果有小数点 去除小数点后面的
	numstrIndex := strings.Index(numstr, ".")
	if numstrIndex != -1 {
		numstr = numstr[0:numstrIndex]
	}

	numIndex := strings.Index(num, ".")
	if numIndex != -1 {
		num = num[0:numIndex]
	}

	n, nok := new(big.Int).SetString(numstr, 10)
	if !nok {
		return "", fmt.Errorf("big.Int first param SetString fail value is %v", numstr)
	}
	m, mok := new(big.Int).SetString("-"+num, 10)
	if !mok {
		return "", fmt.Errorf("big.Int second param SetString fail value is %v", num)
	}

	n.Add(n, m)
	return n.String(), nil
}

// 比较big int 字符串 0: a=b  1:a>b  -1:a<b
func BigIntStringCompare(a string, b string) (int, error) {

	if a == "" {
		a = "0"
	}

	if b == "" {
		b = "0"
	}

	//如果有小数点 去除小数点后面的
	aIndex := strings.Index(a, ".")
	if aIndex != -1 {
		a = a[0:aIndex]
	}

	bIndex := strings.Index(b, ".")
	if bIndex != -1 {
		b = b[0:bIndex]
	}
	n, nok := new(big.Int).SetString(a, 10)
	if !nok {
		return 100, errors.New("big.Int first param SetString fail")
	}
	m, mok := new(big.Int).SetString("-"+b, 10)
	if !mok {
		return 100, errors.New("big.Int second param SetString fail")
	}
	n.Add(n, m)
	res := n.Int64()
	if res > 0 {
		return 1, nil
	} else if res == 0 {
		return 0, nil
	} else {
		return -1, nil
	}
}

/*
************************************************
@ 获取随机数, 如： 00347135859b4561
@ randFlag false 伪随机, true 真随机
@ size:随机数的长度
@ strType:
@ 1: 数字,大小写字母
@ 2: 纯数字
@ 3: 纯小写字母
@ 4: 纯大写字母
@ 5: 数字+小写字母
@ 6: 数字+ 大写字母
@ 7: 小写+大写字母
@*************************************************
*/
func GetRandStr(strType int, size int, randFlag bool) string {
	var kinds [][]int
	switch strType {
	case 1:
		kinds = [][]int{{10, 48}, {26, 97}, {26, 65}}
	case 2:
		kinds = [][]int{{10, 48}}
	case 3:
		kinds = [][]int{{26, 97}}
	case 4:
		kinds = [][]int{{26, 65}}
	case 5:
		kinds = [][]int{{10, 48}, {26, 97}}
	case 6:
		kinds = [][]int{{10, 48}, {26, 65}}
	case 7:
		kinds = [][]int{{26, 97}, {26, 65}}
	default:
		return ""
	}

	kindsLen := len(kinds)
	res := make([]byte, size)

	// 如果是假随机
	if !randFlag {
		rand.New(timeSource)
	}

	var ikind int
	for i := 0; i < size; i++ {
		// random ikind
		if !randFlag {
			ikind = rand.Intn(kindsLen)
		} else {
			result, _ := crand.Int(crand.Reader, big.NewInt(int64(kindsLen)))
			ikind = int(result.Int64())
		}

		scope := kinds[ikind][0]
		base := kinds[ikind][1]

		res[i] = uint8(base + rand.Intn(scope))
	}

	return string(res)
}
