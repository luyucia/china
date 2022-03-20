package china

import (
	"strconv"
	"time"
)

var year_tiangan map[int]string = map[int]string{
	4: "甲",
	5: "乙",
	6: "丙",
	7: "丁",
	8: "戊",
	9: "己",
	0: "庚",
	1: "辛",
	2: "壬",
	3: "癸",
}

var year_dizhi map[int]string = map[int]string{
	4:  "子",
	5:  "丑",
	6:  "寅",
	7:  "卯",
	8:  "辰",
	9:  "巳",
	10: "午",
	11: "未",
	0:  "申",
	1:  "酉",
	2:  "戌",
	3:  "亥",
}

// 获取天干地支
func GetGanzhi(year string) (string, string, error) {
	tianganIndex, err := strconv.Atoi(string(year[3]))
	if err != nil {
		return "", "", err
	}
	yearint, err := strconv.Atoi(year)
	if err != nil {
		return "", "", err
	}

	return year_tiangan[tianganIndex], year_dizhi[yearint%12], nil
}

var shichen map[int]string = map[int]string{
	23: "子",
	0:  "子",
	1:  "丑",
	2:  "丑",
	3:  "寅",
	4:  "寅",
	5:  "卯",
	6:  "卯",
	7:  "辰",
	8:  "辰",
	9:  "巳",
	10: "巳",
	11: "午",
	12: "午",
	13: "未",
	14: "未",
	15: "申",
	16: "申",
	17: "酉",
	18: "酉",
	19: "戌",
	20: "戌",
	21: "亥",
	22: "亥",
}

// 获取时辰
func GetShichen(time time.Time) string {
	return shichen[time.Hour()]
}

// 获取当前时辰
func GetNowShichen() string {
	return GetShichen(time.Now())
}

var tiangan map[int]string = map[int]string{
	1:  "甲",
	2:  "乙",
	3:  "丙",
	4:  "丁",
	5:  "戊",
	6:  "己",
	7:  "庚",
	8:  "辛",
	9:  "壬",
	10: "癸",
}
var dizhi map[int]string = map[int]string{
	1:  "子",
	2:  "丑",
	3:  "寅",
	4:  "卯",
	5:  "辰",
	6:  "巳",
	7:  "午",
	8:  "未",
	9:  "申",
	10: "酉",
	11: "戌",
	12: "亥",
}

var yuandanri_dizhi map[int]string = map[int]string{
	0: "酉",
	1: "午",
	2: "卯",
	3: "子",
	4: "酉",
}

// https://baijiahao.baidu.com/s?id=1686942526629345194&wfr=spider&for=pc
func GetYuandanGanzhi(time time.Time) (string, string) {
	// 末尾2位数
	year_last_two := time.Year() % 100

	t := 0

	if time.Year() >= 2000 {
		t = (year_last_two + 100) / 4
	} else {
		t = year_last_two / 4
	}

	tg := t
	for true {
		if tg <= 10 {
			break
		}
		tg -= 10
	}

	dz := t % 4
	if dz == 0 {
		dz = 10
	}

	return tiangan[tg], yuandanri_dizhi[dz]
}
