package tools

import (
	"fmt"
)

//md5加密示例代码
func ExampleMd5String() {
	plainText := "123456"
	fmt.Println(Md5String(plainText))
	// Output: e10adc3949ba59abbe56e057f20f883e
}
