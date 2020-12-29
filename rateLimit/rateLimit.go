package rateLimit

import (
	"fmt"
	"net/http"
)

const (
	DefaultExpireTime = 1 // 秒
	DefaultMaxThreads = 10
	DefaultPrefix     = "rlg"
)

type params struct {
	Key        string `json:"key"`
	MaxThreads int64  `json:"maxThreads"` // 最大的线程数
	ExpireTime int64  `json:"expireTime"` // 到期时间，秒
}

type Param func(*params)

func evaluateParam(param []Param) *params {
	ps := &params{}

	for _, p := range param {
		p(ps)
	}
	return ps
}

func Key(key string) Param {
	return func(o *params) {
		o.Key = key
	}
}

func MaxThreads(maxThreads int64) Param {
	return func(o *params) {
		o.MaxThreads = maxThreads
	}
}

func ExpireTime(expireTime int64) Param {
	return func(o *params) {
		o.ExpireTime = expireTime
	}
}

type LimitClient struct {
	rateLimit redis.Redis
}

func New(conf *redis.Config) *LimitClient {
	return &LimitClient{
		rateLimit: redis.New(conf),
	}
}

func (p *LimitClient) RateLimiter(param ...Param) gin.HandlerFunc {
	return func(c *gin.Context) {
		ps := evaluateParam(param)

		validAndAssignInput(c, ps)

		pipe := p.rateLimit.Pipeline(c)
		pipe.Send("INCR", ps.Key)
		pipe.Send("TTL", ps.Key)

		replies, err := pipe.Receive()
		if err != nil {
			log.Errorw("pipe filed", "message", ps.Key, "err", err)
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"ok":      false,
				"ecode":   ecode.RateLimitInvalid,
				"code":    "RATE_LIMIT",
				"message": "请稍后重试！",
			})
			return
		}

		var (
			current = replies[0].(int64)
			ttl     = replies[1].(int64)
		)

		if current == int64(1) || ttl == int64(-1) {
			p.rateLimit.Do(c, "EXPIRE", ps.Key, ps.ExpireTime)
		}

		if current > ps.MaxThreads {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{
				"ok":      false,
				"ecode":   ecode.RateLimitInvalid,
				"code":    "RATE_LIMIT",
				"message": "请稍后重试！",
			})
			return

		}
		c.Next()
	}
}

func validAndAssignInput(ctx *gin.Context, p *params) {
	keyItem := utils.ClientIP(ctx.Request)
	userID, exists := ctx.Get(zsrouter.UserKey)
	if exists && userID != "" {
		keyItem = userID.(string)
	}

	if p.ExpireTime == 0 {
		p.ExpireTime = DefaultExpireTime
	}

	if p.MaxThreads == 0 {
		p.MaxThreads = DefaultMaxThreads
	}

	if p.Key == "" {
		// 格式 rlg:60:POST:gold:/gold/issueGold:118.112.12.34
		p.Key = fmt.Sprintf("%s:%s:%s:%s", DefaultPrefix, ctx.Request.Method, ctx.Request.URL.Path, keyItem)
	}
}
