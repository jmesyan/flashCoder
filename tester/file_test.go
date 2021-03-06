package tester

import (
	// "bufio"
	// "fmt"
	// "io"
	"flashCoder/supplier/file"
	"io/ioutil"
	// "os"
	// "strconv"
	"strings"
	"testing"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func TestFile(t *testing.T) {
	//也许大部分基本的文件读取任务是将文件内容读取到内存中。
	file := new(file.FlashFile)
	path := "./test.php"
	dat, err := ioutil.ReadFile(path)
	check(err)
	data := strings.Split(string(dat), "\n")
	// lines:=[]{}
	// result := file.DeleteLines(data, []int{2, 3, 1})
	// result := file.AddContent(data, 2800, "good1\ngood2\n哈哈我陳公公了", true, true)
	result := file.AddFuncContent(data, "getSellerOrderGoodsDetail", "good1\ngood2\n哈哈我陳公公了", true, 2)
	// for k, v := range data {
	// 	line := k + 1
	// 	if line == 50 {
	// 		delete(data, v)
	// 		break
	// 	}
	// }
	// fmt.Println(len(data), result)
	if result != nil {
		newContent := strings.Join(result, "\n")
		ioutil.WriteFile(path, []byte(newContent), 0)
	}

	// fmt.Print(string(dat))
	//你经常会想对于一个文件是怎么读并且读取到哪一部分进行更多的控制。对于这个任务，从使用 os.Open打开一个文件获取一个 os.File 值开始。
	// path := "./test.php"
	// f, err := os.Open(path)
	// check(err)
	// //从文件开始位置读取一些字节。这里最多读取 5 个字节，并且这也是我们实际读取的字节数。
	// b1 := make([]byte, 5)
	// n1, err := f.Read(b1)
	// check(err)
	// fmt.Printf("%d bytes: %s\n", n1, string(b1))
	// //你也可以 Seek 到一个文件中已知的位置并从这个位置开始进行读取。
	// o2, err := f.Seek(-6, 2)
	// check(err)
	// b2 := make([]byte, 2)
	// n2, err := f.Read(b2)
	// check(err)
	// fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))
	// //io 包提供了一些可以帮助我们进行文件读取的函数。例如，上面的读取可以使用 ReadAtLeast 得到一个更健壮的实现。
	// o3, err := f.Seek(6, 0)
	// check(err)
	// b3 := make([]byte, 2)
	// n3, err := io.ReadAtLeast(f, b3, 2)
	// check(err)
	// fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))
	// //没有内置的回转支持，但是使用 Seek(0, 0) 实现。
	// _, err = f.Seek(0, 0)
	// check(err)
	// //bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的取读函数。
	// r4 := bufio.NewReader(f)
	// b4, err := r4.Peek(5)
	// check(err)
	// fmt.Printf("5 bytes: %s\n", string(b4))
	// var num int
	// num = 0
	// for {
	// 	//读出内容保存为string 每次读到以'\n'为标记的位置
	// 	line, err := r4.ReadString('\n')
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(line)

	// }
	// sql = sql[1:]
	// fmt.Println(sql, contents)
	// models.File.BatchAdd(sql, contents)
	//任务结束后要关闭这个文件（通常这个操作应该在 Open操作后立即使用 defer 来完成）。
	// f.Close()
}
