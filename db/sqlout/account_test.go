package gensql

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	args := CreateAccountParams{
		Owner:    "Tom",
		Balance:  100,
		Currency: "USD",
	}
	account, err := testQueries.CreateAccount(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, account)

}

func TestListAccounts(t *testing.T) {
}
