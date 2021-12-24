package api

import (
	"github.com/gin-gonic/gin"
	"github.com/siteddv/simple-bank/api"
	"github.com/siteddv/simple-bank/internal/db/sqlc"
	util2 "github.com/siteddv/simple-bank/internal/util"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func newTestServer(t *testing.T, store db.Store) *api.Server {
	config := util2.Config{
		TokenSymmetricKey:   util2.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := api.NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
