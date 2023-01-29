package gensql

import (
	"context"
	"testing"

	"github.com/bozkayasalih01x/proj/util"
	"github.com/stretchr/testify/require"
)

func randomEntry(t *testing.T) Entry {
	randAccount := randomAccount(t)
	randAmount := util.RandomInt(500, 1000)
	entry, err := testQueries.CreateEntity(context.Background(), CreateEntityParams{
		AccountID: randAccount.ID,
		Amount:    randAmount,
	})
	require.NoError(t, err)
	require.NotEmpty(t, entry)
	require.Equal(t, randAccount.ID, entry.AccountID)
	require.Equal(t, randAmount, entry.Amount)

	return entry
}

func TestCreateEntry(t *testing.T) {
	randomEntry(t)
}

func TestGetEntry(t *testing.T) {
	ent := randomEntry(t)

	entity, err := testQueries.GetEntry(context.Background(), ent.ID)
	require.NoError(t, err)
	require.NotEmpty(t, entity)

	require.Equal(t, entity.ID, ent.ID)
	require.Equal(t, ent.Amount, entity.Amount)
}

func TestListEntries(t *testing.T) {
	count := 10
	for i := 0; i < count; i++ {
		randomEntry(t)
	}

	entities, err := testQueries.ListEntries(context.Background(), ListEntriesParams{
		Limit:  5,
		Offset: 5,
	})
	require.NoError(t, err)
	require.Len(t, entities, 5)
}

func TestDeleteEntity(t *testing.T) {
	entity := randomEntry(t)
	err := testQueries.DeleteEntry(context.Background(), entity.ID)
	require.NoError(t, err)
}
