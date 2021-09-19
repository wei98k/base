package interview

import (
	"context"
	"sync"
	"time"
)

type Ban struct {
	VisitIPs map[string]time.Time
	Lock     sync.Mutex
}

func NewBan(ctx context.Context) *Ban {
	o := &Ban{VisitIPs: make(map[string]time.Time)}
	go func() {
		timer := time.NewTimer(time.Minute * 1)
		for {
			select {
			case <-timer.C:
				o.Lock.Lock()
				for k, v := range o.VisitIPs {
					if time.Now().Sub(v) >= time.Minute*1 {
						delete(o.VisitIPs, k)
					}
				}
				o.Lock.Unlock()
				timer.Reset(time.Minute * 1)
			case <-ctx.Done():
				return
			}
		}
	}()
	return o
}

func (o *Ban) Visit(ip string) bool {
	o.Lock.Lock()
	defer o.Lock.Unlock()
	if _, ok := o.VisitIPs[ip]; ok {
		return true
	}
	o.VisitIPs[ip] = time.Now()
	return false
}
