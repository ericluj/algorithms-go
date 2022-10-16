package base

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUF(t *testing.T) {
	f, err := os.Open("./tinyUF.txt")
	if err != nil {
		assert.Error(t, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// 不做错误处理，认为提供的文件一定符合格式
	// 读取触点数量
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// 初始化n个分量
	uf := NewUF(n)
	for scanner.Scan() {
		// 读取整数对
		arr := strings.Split(scanner.Text(), " ")
		p, _ := strconv.Atoi(arr[0])
		q, _ := strconv.Atoi(arr[1])

		// 如果已经连通则忽略
		if uf.Connected(p, q) {
			continue
		}

		// 归并分量
		uf.Union(p, q)

		// 打印连接
		fmt.Printf("%d %d\n", p, q)
	}
}
