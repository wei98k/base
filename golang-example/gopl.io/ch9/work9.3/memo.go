package memo

// 练习 9.3：
// 扩展Func类型和(*Memo).Get方法，
// 支持调用方提供一个可选的done channel，
// 使其具备通过该channel来取消整个操作的能力（§8.9）。
// 一个被取消了的Func的调用结果不应该被缓存。

type Func func(key string, done <-chan struct{}) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

type request struct {
	key      string
	response chan<- result
	done     <-chan struct{}
}

type Memo struct {
	requests chan request
}

func New(f Func) *Memo {
	memo := &Memo{
		requests: make(chan request),
	}
	go memo.server(f)
}

var canceledKeys chan string

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response, done}
	res := <-response
	select {
	case <-done:
		canceledKeys <- key
	default:
	}
	return res.value, res.err
}

func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make()
	for {
	LOOP:
		for {
			select {
			case key := <-canceledKeys:
				delete(cache, key)
			default:
				break LOOP
			}
		}
		select {
		case req, ok := <-memo.requests:
			if !ok {
				return
			}
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done)
			}
			go e.deliver(req.response)
		default:
		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
