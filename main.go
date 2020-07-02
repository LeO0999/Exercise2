package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	//Bai 1: Tạo 1 chương trình cứ 3s in ra dòng chữ
	// Print()

	// Bai 2: tính ra ngày hiện tại theo timestamp. Điểm mốc từ mức 0 tại 1/1/1970
	// second := time.Now().Unix()
	// GetDay(second)

	// Bài 3: chương trình với 1 vòng lặp for và 3 lần sleep mỗi lần sleep 3sec Nhưng sau 3s thì kết thúc hàm

	// ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	// defer cancel()

	// for i := 0; i < 3; i++ {
	// 	select {
	// 	case <-ctx.Done():
	// 		// fmt.Println("kt")
	// 		return

	// 	default:

	// 		time.Sleep(3 * time.Second)
	// 		fmt.Println("a")

	// 	}

	// }

	// Bai 4
	// m := int64(1592190294764144364) / int64(60*time.Second)
	// fmt.Println(m)
	// d := time.Duration(1592190294764144364)
	// fmt.Println(d)

	// Bai 5

	// time1 := time.Unix(0, 1592190385*int64(time.Second))
	// fmt.Println(time1.Weekday())

	// Bai 7: Tạo 1 đối tượng context với 1 value là timestamp hiện tại tính theo ns chạy qua 1 hàm

	// k := "key"
	// ctx := context.WithValue(context.Background(), k, time.Now().UnixNano())
	// a := X(ctx, k)
	// fmt.Println(a)

	// Bai 8: Tạo 1 interval time sao cho cứ 100ms in ra dùng chữ ${time.Now().Unix()} done
	t := time.Now().Unix()

	ctx1, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()

	for now := range time.Tick(100 * time.Millisecond) {
		select {
		case <-ctx1.Done():
			return

		default:

			fmt.Println(now, t, "done")

		}

	}

	//Bai 9: Tạo 1 đoạn code sử dụng time.AfterFunc(), sau 100ms thì in ra dòng chữ i'm study
	// f := func() {

	// 	fmt.Println("I'm study")
	// }
	// timer := time.AfterFunc(100*time.Millisecond, f)
	// defer timer.Stop()
	// time.Sleep(1 * time.Second)
	// fmt.Println("Xong rồi!")
}
