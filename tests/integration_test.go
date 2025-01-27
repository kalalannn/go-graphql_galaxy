package tests

import (
	"context"
	"go-graphql_galaxy/internal/app"
	"go-graphql_galaxy/internal/genqlient/generated"
	"go-graphql_galaxy/internal/transformers"
	"go-graphql_galaxy/pkg/client"
	"go-graphql_galaxy/pkg/utils"
	"net/http"
	"testing"
	"time"

	"github.com/Khan/genqlient/graphql"
	"github.com/stretchr/testify/assert"
)

var Config *utils.Config
var Client graphql.Client
var Ctx = context.Background()

func TestMain(m *testing.M) {
	go func() {
		app.NewApp().Run()
	}()

	Config = utils.MustLoadConfig()
	url := client.CreateURL(Config.Server.Host, Config.Server.Port, Config.Server.PingPath)

	err := client.PingServer(url, 5*time.Second)
	if err != nil {
		panic("app run timeout")
	}

	Client = graphql.NewClient("http://localhost:8080/query", http.DefaultClient)

	m.Run()
}

const (
	AssertAverageAge                   = 40.91
	AssertAverageBeerConsumption       = 134527.91
	AssertAverageWeight                = 104.03
	AssertAverageNemesesYears          = 113.75
	AssertCharactersCount        int64 = 11
	AssertNemesesCount           int64 = 13
	AssertSecretsCount           int64 = 25
)

const (
	AssertGendersMale   int64 = 6
	AssertGendersFemale int64 = 2
	AssertGendersOther  int64 = 3
)

const (
	AssertNemesesAlive int64 = 11
	AssertNemesesDead  int64 = 2
)

func TestRoot(t *testing.T) {
	// act
	resp, err := generated.RootQuery(Ctx, Client)
	if err != nil {
		t.Fatalf("GraphQL request error: %v", err)
	}

	// assert health_check
	assert.True(t, resp.Health_check)

	// assert server_time
	parsedTime, err := time.Parse(transformers.TimeFormat, resp.Server_time)
	if err != nil {
		t.Fatalf("Time parsing error: %v", err)
	}
	assert.WithinDuration(t, parsedTime, time.Now(), time.Minute, "Server time should be within 1 minute of the current time")

	// assert average_age
	assert.Equal(t, AssertAverageAge, resp.Average_age)

	// assert average_beer_consumption
	assert.Equal(t, AssertAverageBeerConsumption, resp.Average_beer_consumption)

	// assert average_weight
	assert.Equal(t, AssertAverageWeight, resp.Average_weight)

	// assert average_nemeses_years
	assert.Equal(t, AssertAverageNemesesYears, resp.Average_nemeses_years)

	// assert characters_count
	assert.Equal(t, AssertCharactersCount, resp.Characters_count)

	// assert nemeses_count
	assert.Equal(t, AssertNemesesCount, resp.Nemeses_count)

	// assert secrets_count
	assert.Equal(t, AssertSecretsCount, resp.Secrets_count)

	// assert genders
	assert.Equal(t, AssertGendersMale, resp.Genders.Male)
	assert.Equal(t, AssertGendersFemale, resp.Genders.Female)
	assert.Equal(t, AssertGendersOther, resp.Genders.Other)

	// assert alive_nemeses
	assert.Equal(t, AssertNemesesAlive, resp.Alive_nemeses.Alive)
	assert.Equal(t, AssertNemesesDead, resp.Alive_nemeses.Dead)
}
