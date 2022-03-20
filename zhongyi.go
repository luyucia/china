package china

import "time"

var liuzhu map[int]string = map[int]string{
	23: "足少阳胆经",
	0:  "足少阳胆经",
	1:  "足厥阴肝经",
	2:  "足厥阴肝经",
	3:  "手太阴肺经",
	4:  "手太阴肺经",
	5:  "手阳明大肠经",
	6:  "手阳明大肠经",
	7:  "足阳明胃经",
	8:  "足阳明胃经",
	9:  "足太阴脾经",
	10: "足太阴脾经",
	11: "手少阴心经",
	12: "手少阴心经",
	13: "手太阳小肠经",
	14: "手太阳小肠经",
	15: "足太阳膀胱经",
	16: "足太阳膀胱经",
	17: "足少阴肾经",
	18: "足少阴肾经",
	19: "手厥阴心包经",
	20: "手厥阴心包经",
	21: "手少阳三焦经",
	22: "手少阳三焦经",
}

// 获取十二经络流注
func GetLiuzhu(time time.Time) string {
	return liuzhu[time.Hour()]
}

func GetNowLiuzhu() string {
	return GetLiuzhu(time.Now())
}
