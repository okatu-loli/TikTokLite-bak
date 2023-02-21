package util

import (
	"fmt"
	"strconv"
	"strings"
)

var server = []string{
	"192.168.1.1",
	"192.168.2.2",
	"192.168.3.3",
	"192.168.4.4",
}

func main() {
	ring := New(server, 4)
	//hr.addNode("192.168.5.5")
	fifth := 0
	first, second, third, four := 0, 0, 0, 0
	for i := 0; i < 10; i++ {
		str := ring.getNode(strconv.Itoa(i))
		if strings.Compare(str, "192.168.1.1") == 0 {
			fmt.Printf("192.168.1.1：%v \n", i)
			first++
		} else if strings.Compare(str, "192.168.2.2") == 0 {
			fmt.Printf("192.168.2.2：%v \n", i)
			second++
		} else if strings.Compare(str, "192.168.3.3") == 0 {
			fmt.Printf("192.168.3.3：%v \n", i)
			third++
		} else if strings.Compare(str, "192.168.4.4") == 0 {
			fmt.Printf("192.168.4.4：%v \n", i)
			four++
		} else if strings.Compare(str, "192.168.5.5") == 0 {
			fmt.Printf("192.168.5.5：%v \n", i)
			fifth++
		}
	}

	fmt.Printf("%v %v %v %v %v", first, second, third, four, fifth)

}
