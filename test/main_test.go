package test

import (
	"encoding/json"
	"fmt"
	"go-one/util"
	"strings"
	"testing"
	"time"
	"unsafe"
)

type Order struct {
	ID     string  `json:"id"`
	User   string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
	Price  float64 `json:"price"`
	Filled bool    `json:"filled"`
}

func TestInitialization(t *testing.T) {
	OrderBook := make(map[string]*Order)
	OrderBook["1"] = &Order{ID: "1", User: "Alice", Type: "Buy", Amount: 10, Price: 100, Filled: false}
	OrderBook["2"] = &Order{ID: "2", User: "Bob", Type: "Sell", Amount: 10, Price: 90, Filled: false}
	for id, order := range OrderBook {
		fmt.Println("ID:", id, "User:", order.User, "Type:", order.Type, "Amount:", order.Amount, "Price:", order.Price, "Filled:", order.Filled)

		jsonBytes, _ := json.MarshalIndent(order, "", "  ")
		fmt.Println(string(jsonBytes))
	}

	order := Order{}
	fmt.Println(order)

	order2 := &Order{}
	fmt.Println(order2)

	// 打印指针的大小
	fmt.Printf("指针的大小: %d 字节\n", unsafe.Sizeof(&order))
	fmt.Printf("的大小: %d 字节\n", unsafe.Sizeof(order))

	fmt.Println(unsafe.Offsetof(order.ID))
	fmt.Println(unsafe.Offsetof(order.User))
	fmt.Println(unsafe.Offsetof(order.Type))
	fmt.Println("Amount:", unsafe.Offsetof(order.Amount))
	fmt.Println(unsafe.Offsetof(order.Price))
	fmt.Println(unsafe.Offsetof(order.Filled))

	// 打印不同类型指针的大小
	var intPtr *int
	var stringPtr *string
	var floatPtr *float64
	var orderPtr *Order

	fmt.Printf("int指针大小: %d 字节\n", unsafe.Sizeof(intPtr))
	fmt.Printf("string指针大小: %d 字节\n", unsafe.Sizeof(stringPtr))
	fmt.Printf("float64指针大小: %d 字节\n", unsafe.Sizeof(floatPtr))
	//fmt.Printf("Order结构体指针大小: %d 字节\n", unsafe.Sizeof((*Order)(nil)))
	fmt.Printf("Order结构体指针大小: %d 字节\n", unsafe.Sizeof(orderPtr))

	fmt.Printf("*string 指针大小: %d 字节\n", unsafe.Sizeof((*string)(nil)))

	fmt.Printf("int指针: %v \n", (*int)(nil))
	fmt.Printf("int的大小: %d 字节\n", unsafe.Sizeof((*int)(nil)))
	var i *int
	fmt.Printf("i的值: %v\n", i)
	fmt.Printf("i的大小: %d 字节\n", unsafe.Sizeof(i))

	fmt.Printf("OrderBook %#v \n", OrderBook["1"])

	// 通用格式化动词
	fmt.Printf("默认格式: %v\n", 42)
	fmt.Printf("结构体字段: %+v\n", struct{ Name string }{"Alice"})
	fmt.Printf("Go 语法表示: %#v\n", struct{ Name string }{"Alice"})
	fmt.Printf("类型: %T\n", 42)

	// 布尔值
	fmt.Printf("布尔值: %t\n", true)

	// 整数
	fmt.Printf("十进制: %d\n", 42)
	fmt.Printf("二进制: %b\n", 42)
	fmt.Printf("八进制: %o\n", 42)
	fmt.Printf("十六进制: %x\n", 42)
	fmt.Printf("十六进制（大写）: %X\n", 42)

	// 浮点数
	fmt.Printf("十进制: %f\n", 3.14)
	fmt.Printf("科学计数法: %e\n", 3.14)
	fmt.Printf("科学计数法（大写）: %E\n", 3.14)
	fmt.Printf("自动选择: %g\n", 3.14)

	// 字符串
	fmt.Printf("字符串: %s\n", "hello")
	fmt.Printf("带双引号: %q\n", "hello")
	fmt.Printf("十六进制: %x\n", "hello")
	fmt.Printf("十六进制（大写）: %X\n", "hello")

	// 指针
	var ptr *int
	fmt.Printf("指针地址: %p\n", ptr)

	// 创建客户端实例
	client := util.NewHTTPClient("http://example.com", 10*time.Second)

	// GET请求示例
	headers := map[string]string{
		"Authorization": "Bearer token123",
	}
	response, err := client.Get("/users", headers)
	fmt.Printf("GET请求响应: %s\n", response)

	// POST请求示例
	data := map[string]interface{}{
		"name": "张三",
		"age":  25,
	}
	response, err = client.Post("/users", data, headers)
	if err != nil {
		fmt.Printf("POST请求失败: %v\n", err)
		return
	}
	fmt.Printf("POST请求响应: %s\n", response)
}

// 测试 Order 结构体的基本操作
func TestOrderBasicOperations(t *testing.T) {
	// 测试创建和访问订单
	t.Run("Create and Access Order", func(t *testing.T) {
		OrderBook := make(map[string]*Order)
		OrderBook["1"] = &Order{ID: "1", User: "Alice", Type: "Buy", Amount: 10, Price: 100, Filled: false}

		order := OrderBook["1"]
		if order.ID != "1" || order.User != "Alice" || order.Type != "Buy" {
			t.Errorf("订单基本信息不匹配，got: %+v", order)
		}
	})

	// 测试空订单初始化
	t.Run("Empty Order Initialization", func(t *testing.T) {
		order := Order{}
		if order.ID != "" || order.User != "" || order.Amount != 0 {
			t.Errorf("空订单初始化失败，got: %+v", order)
		}

		order2 := &Order{}
		if order2.ID != "" || order2.User != "" || order2.Amount != 0 {
			t.Errorf("空订单指针初始化失败，got: %+v", order2)
		}
	})
}

// 测试内存相关操作
func TestMemoryOperations(t *testing.T) {
	order := Order{}

	// 测试结构体大小
	t.Run("Struct Size", func(t *testing.T) {
		orderSize := unsafe.Sizeof(order)
		orderPtrSize := unsafe.Sizeof(&order)

		if orderPtrSize != 8 { // 在64位系统上，指针大小应该是8字节
			t.Errorf("指针大小异常，expected: 8, got: %d", orderPtrSize)
		}

		// Order结构体大小验证（具体大小取决于结构体字段对齐）
		if orderSize == 0 {
			t.Error("结构体大小不应为0")
		}
	})

	// 测试字段偏移量
	t.Run("Field Offsets", func(t *testing.T) {
		// 验证字段顺序
		idOffset := unsafe.Offsetof(order.ID)
		userOffset := unsafe.Offsetof(order.User)
		typeOffset := unsafe.Offsetof(order.Type)

		if idOffset > userOffset {
			t.Error("ID字段偏移量不应大于User字段")
		}
		if userOffset > typeOffset {
			t.Error("User字段偏移量不应大于Type字段")
		}
	})

	// 测试不同类型指针的大小
	t.Run("Pointer Sizes", func(t *testing.T) {
		var intPtr *int
		var stringPtr *string
		var floatPtr *float64
		var orderPtr *Order

		ptrSize := unsafe.Sizeof(intPtr)
		if ptrSize != 8 { // 64位系统上所有指针大小应该相同
			t.Errorf("指针大小异常，expected: 8, got: %d", ptrSize)
		}

		// 验证所有指针大小一致
		if unsafe.Sizeof(stringPtr) != ptrSize ||
			unsafe.Sizeof(floatPtr) != ptrSize ||
			unsafe.Sizeof(orderPtr) != ptrSize {
			t.Error("不同类型指针大小不一致")
		}
	})
}

// 测试格式化输出
func TestFormatting(t *testing.T) {
	// 测试基本类型格式化
	t.Run("Basic Type Formatting", func(t *testing.T) {
		num := 42
		str := "hello"

		cases := []struct {
			format string
			value  interface{}
			want   string
		}{
			{"%d", num, "42"},
			{"%s", str, "hello"},
			{"%t", true, "true"},
			{"%f", 3.14, "3.140000"},
		}

		for _, c := range cases {
			got := ""
			// 使用Sprint来捕获格式化结果
			got = fmt.Sprintf(c.format, c.value)
			if got != c.want {
				t.Errorf("格式化失败 format:%s, want: %s, got: %s", c.format, c.want, got)
			}
		}
	})

	// 测试结构体格式化
	t.Run("Struct Formatting", func(t *testing.T) {
		order := &Order{ID: "1", User: "Alice", Type: "Buy"}
		formatted := fmt.Sprintf("%+v", order)
		if formatted == "" {
			t.Error("结构体格式化输出为空")
		}
		if !strings.Contains(formatted, "Alice") {
			t.Error("结构体格式化输出缺少字段值")
		}
	})
}
