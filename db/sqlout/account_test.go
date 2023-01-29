package gensql

import (
	"context"
	"testing"

	"github.com/bozkayasalih01x/proj/util"
	"github.com/stretchr/testify/require"
)

func randomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Owner:    util.RandomString(5),
		Balance:  util.RandomInt(100, 1000),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)
	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Currency, account.Currency)
	require.Equal(t, arg.Balance, account.Balance)

	require.NotZero(t, account.ID)
	require.NotEmpty(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	randomAccount(t)
}

func TestGetAccount(t *testing.T) {
	acc := randomAccount(t)
	fetchedAccount, err := testQueries.GetAccount(context.Background(), acc.ID)
	require.NoError(t, err)
	require.NotEmpty(t, fetchedAccount)

	require.Equal(t, acc.ID, fetchedAccount.ID)
	require.Equal(t, acc.Balance, fetchedAccount.Balance)
	require.Equal(t, acc.Currency, fetchedAccount.Currency)
	require.Equal(t, acc.Owner, fetchedAccount.Owner)
}

func TestListAccounts(t *testing.T) {
	//create 10 acocunts and test them ;
	for i := 0; i < 10; i++ {
		randomAccount(t)
	}

	accounts, err := testQueries.ListAccounts(context.Background(), ListAccountsParams{
		Limit:  5,
		Offset: 5,
	})

	require.NoError(t, err)
	require.Len(t, accounts, 5)
	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}

func TestUpdateAccount(t *testing.T) {
	account := randomAccount(t)
	randBalance := util.RandomInt(100, 5000)

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), UpdateAccountParams{
		ID:      account.ID,
		Balance: randBalance,
	})
	require.NoError(t, err)
	require.Equal(t, updatedAccount.Balance, randBalance)
	require.Equal(t, account.ID, updatedAccount.ID)
	require.Equal(t, account.Owner, updatedAccount.Owner)
	require.Equal(t, account.Currency, updatedAccount.Currency)
}

func TestDeleteAccount(t *testing.T) {
	account := randomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)
}
