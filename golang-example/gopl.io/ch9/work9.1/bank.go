package work91

// 练习 9.1： 给gopl.io/ch9/bank1程序添加一个Withdraw(amount int)取款函数。其返回结果应该要表明事务是成功了还是因为没有足够资金失败了。
// 这条消息会被发送给monitor的goroutine，且消息需要包含取款的额度和一个新的channel，
// 这个新channel会被monitor goroutine来把boolean结果发回给Withdraw。

//tip: 一个提供对一个指定的变量通过channel来请求的goroutine叫做这个变量的monitor（监控）goroutine

var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

// 定义一个取款后的结果的结构体
type draw struct {
	amount  int
	success chan bool
}

// 定义一个取款动作channel
var withdrawal = make(chan draw)

// 添加一个取款函数
func Withdrawal(amount int) bool {
	success := make(chan bool)
	withdrawal <- draw{amount, success}
	return <-success
}

// 出纳员-提款机
// 存钱和查询余额
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
			// 取款过程
		case draw := <-withdrawal:
			if draw.amount <= balance {
				balance -= draw.amount
				draw.success <- true
			} else {
				draw.success <- false
			}
		}
	}
}

func init() {
	go teller()
}
