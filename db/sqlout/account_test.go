package gensql

import (
	"context"
	"testing"

	"github.com/bozkayasalih01x/proj/util"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Owner:    util.RandomString(5),
		Balance:  util.RandomInt(100, 1000),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

}

func TestListAccounts(t *testing.T) {
}
