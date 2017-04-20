package api

import (
	"crypto/rand"
	"encoding/base64"
	"sync"
	"time"

	"github.com/jwowillo/trim"
	"github.com/jwowillo/trim/application"
	"github.com/jwowillo/trim/response"
	"github.com/jwowillo/trim/trimming"
)

const (
	limitAmount        = 120
	limitDuration      = time.Minute
	tokenLimitAmount   = 1
	tokenLimitDuration = time.Minute
	tokenSize          = 16
	tokenPath          = "/token"
)

type validateToken struct {
	*base
	lock sync.Mutex
}

func newValidateToken() validateToken {
	return validateToken{base: &base{}}
}

func (v validateToken) Handle(r trim.Request) trim.Response {
	vals := r.Header()["Authorization"]
	if len(vals) != 1 {
		return response.NewJSON(
			map[string]string{
				"message": "must pass token in " +
					"'Authorization' header",
			},
			trim.CodeForbidden,
		)
	}
	id := vals[0]
	t, ok := m.token(id)
	if !ok {
		return response.NewJSON(
			map[string]string{"message": "invalid token"},
			trim.CodeForbidden,
		)
	}
	if time.Since(t.Created()) > limitDuration {
		m.remove(t)
		return response.NewJSON(
			map[string]string{"message": "token is outdated"},
			trim.CodeForbidden,
		)
	}
	if t.Requests() > limitAmount {
		return response.NewJSON(
			map[string]string{
				"message": "too many requests with token",
			},
			trim.CodeTooManyRequests,
		)
	}
	v.lock.Lock()
	t.requests++
	v.lock.Unlock()
	return v.handler.Handle(r)
}

type tokenController struct {
	limiter *trimming.Limit
}

func newTokenController() *tokenController {
	return &tokenController{
		limiter: trimming.NewLimit(
			tokenLimitAmount,
			tokenLimitDuration,
			func(_ trim.Request) bool { return false },
		),
	}
}

func (c *tokenController) Trimmings() []trim.Trimming {
	return []trim.Trimming{c.limiter}
}

func (c *tokenController) Path() string {
	return tokenPath
}

func (c *tokenController) Description() *application.ControllerDescription {
	return &application.ControllerDescription{
		Get: &application.MethodDescription{
			Response: "Token to use for actions",
			Limiting: "120 requests",
		},
	}
}

func (c *tokenController) Handle(_ trim.Request) trim.Response {
	t, err := newToken()
	if err != nil {
		return response.NewJSON(
			map[string]string{"message": "couldn't make token"},
			trim.CodeInternalServerError,
		)
	}
	m.add(t)
	return response.NewJSON(map[string]string{"token": t.ID}, trim.CodeOK)
}

var m = newTokenMap()

type tokenMap struct {
	lock   sync.RWMutex
	tokens map[string]*token
}

func newTokenMap() *tokenMap {
	return &tokenMap{tokens: make(map[string]*token)}
}

func (m *tokenMap) add(t *token) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.tokens[t.ID] = t
}

func (m *tokenMap) remove(t *token) {
	m.lock.Lock()
	defer m.lock.Unlock()
	if _, ok := m.tokens[t.ID]; !ok {
		return
	}
	delete(m.tokens, t.ID)
}

func (m *tokenMap) token(id string) (*token, bool) {
	m.lock.RLock()
	defer m.lock.RUnlock()
	t, ok := m.tokens[id]
	return t, ok
}

type token struct {
	ID       string `json:"token"`
	requests int
	created  time.Time
}

func newToken() (*token, error) {
	buff := make([]byte, tokenSize)
	_, err := rand.Read(buff)
	if err != nil {
		return nil, err
	}
	id := base64.StdEncoding.EncodeToString(buff)
	return &token{ID: id, created: time.Now()}, nil
}

func (t *token) Requests() int {
	return t.requests
}

func (t *token) Created() time.Time {
	return t.created
}

func (t *token) Description() string {
	return "authentication for actions"
}
