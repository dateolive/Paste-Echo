package paste

import (
	"math/rand"
	"time"
)

func init(){
	rand.Seed(time.Now().UnixNano()) //纳秒数级别生成随机种子
}

var(
	charset = []rune("qazwsxedcrfvtgbyhnujmikolp0123456789")
)

func getOneRune(r []rune)rune{
	return r[rand.Intn(len(r))]
}

func _generate(length int)string{

	ans:=make([]rune,length)
	for i:=0;i<length;i++{
		ans[i]= getOneRune(charset)
	}
	return string(ans)
}

func generate(length int,model interface{})string{
	str:= _generate(length)
	for exist(str,model){
		str= _generate(length)
	}
	return str
}
