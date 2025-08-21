package proxy

import "testing"

func TestProxy(t *testing.T) {
	user := NewUser()
	proxy := NewBuyerProxy(user)
	proxy.Buy()

}
