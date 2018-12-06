package main

import "fmt"

// go map
func main() {
	// 创建集合（其中[]内为key的类型）[]后为value的类型
	var countryCapitalMap map[string]string
	// 使用make函数初始化map 不初始化的map为nil map 不能用来存放键值对
	countryCapitalMap = make(map[string]string)
	// 上面两行等价于 countryCapitalMap := make(map[string]string)

	// 插入键值对
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["China"] = "Beijing"
	countryCapitalMap["美国"] = "华盛顿"
	countryCapitalMap["India"] = "新德里"

	// 遍历map
	for country, capital := range countryCapitalMap {
		fmt.Printf("%s的首都是：%s\n", country, capital)
	}

	// 查看key是否存在
	capital, ok := countryCapitalMap["德国"]
	if ok {
		fmt.Printf("德国的首都是：%s\n", capital)
	} else {
		fmt.Printf("找不到德国\n")
	}

	//删除元素
	fmt.Printf("删除前的元素数量%d\n", len(countryCapitalMap)) // 此处不能用cap
	delete(countryCapitalMap, "美国")
	fmt.Printf("删除后的元素数量%d\n", len(countryCapitalMap))
}
