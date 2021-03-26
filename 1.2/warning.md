- 正确关闭文件。

```go
f, err := os.Open("/tmp/file.md")
if err != nil {
	return err
}

defer func() {
	closeErr := f.Close()
	if closeErr != nil {
		if err == nil {
			err = closeErr
		} else {
			log.Println("Error occured while closing the file :", closeErr)
		}
	}
}()
return err

```

- return and defer
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x // 返回值=5,x++修改了x的值，没有修改返回值
}


func f2() (x int) {
	defer func() {
		x++
	}()
	return 5 // 返回值=x=5,x++得到6，返回x的值6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x  // 返回值=y=x=5, x++不影响y值，返回5
}

// go语言中函数的return不是原子操作，在底层是分为两步执行的。
// 第一步： 给返回值赋值
// 第二部： 执行RET指令真正的返回
// 函数中如果存在defer,那么defer执行的时机是第一步和第二步之间
func f4() (n int) {
	defer func(n int) {
		n++ // 修改的是函数中副本的值
		fmt.Printf("in defer %p\n", &n)
		fmt.Printf("in defer n=%d\n", n)
	}(n)
	fmt.Printf("in main %p\n", &n)
	return 5 // 返回值=5
}

/* f4 执行步骤
	1. 将defer入栈，此时n为零值0。将n=0拷贝了一份到匿名函数内。
	2. 将返回值n赋值为5
	3. 执行defer语句，对defer里的n进行加1操作得到1。
	4. 执行RET指令，返回n=5。
 */
