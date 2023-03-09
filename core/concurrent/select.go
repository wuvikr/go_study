package main

func main() {

	select {
		case x := <-ch1:     // 从channel ch1接收数据
		  ... ...
		
		case y, ok := <-ch2: // 从channel ch2接收数据，并根据ok值判断ch2是否已经关闭
		  ... ...
		
		case ch3 <- z:       // 将z值发送到channel ch3中:
		  ... ...
		
		default:             // 当上面case中的channel通信均无法实施时，执行该默认分支
		}
}
