package memo1

//定义缓存结构体
type Memo struct {
	f     Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)

//定义接收结构体
type result struct {
	value interface{}
	err   error
}

//初始化函数
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

//获取过后缓存
func (memo *Memo) Get(key string) (interface{}, error) {

	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
