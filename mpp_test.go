package mpp_test

import (
	"testing"

	"github.com/lestrrat-go/mpp"
	"github.com/stretchr/testify/assert"
)

func TestMPP(t *testing.T) {
	p := mpp.New("hello").JSON()

	v, err := p.Do()
	if !assert.NoError(t, err, `p.Do() should succeed`) {
		return
	}

	t.Logf("%s", v)
}
