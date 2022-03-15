package main

import "fmt"

func main() {
	//map is hashmap
	//map is no order
	//default the map
	// student id ==> student name idnames
	//var idnames map[int]string //in this time,this map is not useful,the map is empty
	//deault distribute sapce
	idnames := make(map[int]string, 10)
	idnames[0] = "bob"
	idnames[1] = "mary"
	idscore := make(map[int]float64, 10)
	//if use the map,the map is nil map,must distribute space
	//use make distribute space,suggest appoint length,it is better to function
	//idnames := make(map[int]string)
	//idnames = make(map[int]string, 10)
	//suggest way
	//遍历
	for key, value := range idnames {
		fmt.Println("key", key, "value", value)
	}
	//如何确定一个key是否存在map中
	//在map中不存在访问越界，所有key都有效，会返回此数据类型的零值
	name9 := idnames[9]
	fmt.Println("name9", name9)
	fmt.Println(idscore[100])
	//无法通过 获取value的值来判断是否存在于map中，需要一个能够进行校验的方式
	value, ok := idnames[1] //如果存在，返回TRUE，反之返回FALSE
	if ok {
		fmt.Println("1", value)
	}
	value1, ok := idnames[10]
	if ok {
		fmt.Println("exist", value1)
	} else {
		fmt.Println("not exist", value1)
	}
	//删除元素，delete
	delete(idnames, 1)
	delete(idnames, 10)
	fmt.Println("now", idnames)
	//并发任务处理时，map需要上锁//TODO
}
