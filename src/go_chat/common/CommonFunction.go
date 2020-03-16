package common

type CommonFunction struct {
}

//数组去重
func (c *CommonFunction) UniqueArr(arr []string) []string {
	strMap := make(map[string]string)
	for _, v := range arr {
		strMap[v] = v
	}
	nArr := []string{}
	for _, value := range strMap {
		nArr = append(nArr, value)
	}
	return nArr
}

//将redis返回的数据转换为字符串。
func (c *CommonFunction) B2S(bs interface{}) string {
	ba := []byte{}
	if s,ok:=bs.([]uint8);ok{
		for _, b := range s {
			ba = append(ba, byte(b))
		}
	}
	return string(ba)
}
