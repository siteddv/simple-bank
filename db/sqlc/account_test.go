package db

import (
	"context"
	"database/sql"
	"github.com/siteddv/simple-bank/util"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func Test_CreateAccount(t *testing.T) {
	arg := CreateAccountParams{
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	testQueries := GetTestQueries()

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	/*err = testQueries.DeleteAccount(context.Background(), account.ID)
	if err != nil {
		log.Fatal("error handled during deleting account: ", err)
	}*/
}

func GetTestQueries() *Queries {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries := New(conn)

	return testQueries
}
