package base

import (
	"testing"

	fcom "github.com/hyperbench/hyperbench-common/common"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	b := NewBlockchainBase(ClientConfig{})
	assert.NotNil(t, b)

	err := b.DeployContract()
	assert.NoError(t, err)

	res := b.Confirm(&fcom.Result{})
	assert.NotNil(t, res)

	res = b.Invoke(fcom.Invoke{})
	assert.NotNil(t, res)

	res = b.Transfer(fcom.Transfer{})
	assert.NotNil(t, res)

	result := b.Query(fcom.Query{})
	assert.Nil(t, result)

	err = b.Option(fcom.Option{})
	assert.NoError(t, err)

	s, err := b.GetContext()
	assert.Equal(t, s, "")
	assert.NoError(t, err)

	err = b.SetContext("")
	assert.NoError(t, err)

	err = b.ResetContext()
	assert.NoError(t, err)

	rs, err := b.Statistic(fcom.Statistic{From: &fcom.ChainInfo{BlockHeight: 0}, To: &fcom.ChainInfo{BlockHeight: 1}})
	assert.NotNil(t, rs)
	assert.NoError(t, err)

}
