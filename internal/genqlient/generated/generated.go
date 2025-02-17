// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package generated

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// RootQueryAlive_nemesesAliveNemeses includes the requested fields of the GraphQL type AliveNemeses.
type RootQueryAlive_nemesesAliveNemeses struct {
	Alive int64 `json:"alive"`
	Dead  int64 `json:"dead"`
}

// GetAlive returns RootQueryAlive_nemesesAliveNemeses.Alive, and is useful for accessing the field via an interface.
func (v *RootQueryAlive_nemesesAliveNemeses) GetAlive() int64 { return v.Alive }

// GetDead returns RootQueryAlive_nemesesAliveNemeses.Dead, and is useful for accessing the field via an interface.
func (v *RootQueryAlive_nemesesAliveNemeses) GetDead() int64 { return v.Dead }

// RootQueryGenders includes the requested fields of the GraphQL type Genders.
type RootQueryGenders struct {
	Male   int64 `json:"male"`
	Female int64 `json:"female"`
	Other  int64 `json:"other"`
}

// GetMale returns RootQueryGenders.Male, and is useful for accessing the field via an interface.
func (v *RootQueryGenders) GetMale() int64 { return v.Male }

// GetFemale returns RootQueryGenders.Female, and is useful for accessing the field via an interface.
func (v *RootQueryGenders) GetFemale() int64 { return v.Female }

// GetOther returns RootQueryGenders.Other, and is useful for accessing the field via an interface.
func (v *RootQueryGenders) GetOther() int64 { return v.Other }

// RootQueryResponse is returned by RootQuery on success.
type RootQueryResponse struct {
	Health_check             bool                               `json:"health_check"`
	Server_time              string                             `json:"server_time"`
	Average_age              float64                            `json:"average_age"`
	Average_beer_consumption float64                            `json:"average_beer_consumption"`
	Average_weight           float64                            `json:"average_weight"`
	Characters_count         int64                              `json:"characters_count"`
	Genders                  RootQueryGenders                   `json:"genders"`
	Average_nemeses_years    float64                            `json:"average_nemeses_years"`
	Nemeses_count            int64                              `json:"nemeses_count"`
	Alive_nemeses            RootQueryAlive_nemesesAliveNemeses `json:"alive_nemeses"`
	Secrets_count            int64                              `json:"secrets_count"`
}

// GetHealth_check returns RootQueryResponse.Health_check, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetHealth_check() bool { return v.Health_check }

// GetServer_time returns RootQueryResponse.Server_time, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetServer_time() string { return v.Server_time }

// GetAverage_age returns RootQueryResponse.Average_age, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetAverage_age() float64 { return v.Average_age }

// GetAverage_beer_consumption returns RootQueryResponse.Average_beer_consumption, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetAverage_beer_consumption() float64 { return v.Average_beer_consumption }

// GetAverage_weight returns RootQueryResponse.Average_weight, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetAverage_weight() float64 { return v.Average_weight }

// GetCharacters_count returns RootQueryResponse.Characters_count, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetCharacters_count() int64 { return v.Characters_count }

// GetGenders returns RootQueryResponse.Genders, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetGenders() RootQueryGenders { return v.Genders }

// GetAverage_nemeses_years returns RootQueryResponse.Average_nemeses_years, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetAverage_nemeses_years() float64 { return v.Average_nemeses_years }

// GetNemeses_count returns RootQueryResponse.Nemeses_count, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetNemeses_count() int64 { return v.Nemeses_count }

// GetAlive_nemeses returns RootQueryResponse.Alive_nemeses, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetAlive_nemeses() RootQueryAlive_nemesesAliveNemeses {
	return v.Alive_nemeses
}

// GetSecrets_count returns RootQueryResponse.Secrets_count, and is useful for accessing the field via an interface.
func (v *RootQueryResponse) GetSecrets_count() int64 { return v.Secrets_count }

// The query or mutation executed by RootQuery.
const RootQuery_Operation = `
query RootQuery {
	health_check
	server_time
	average_age
	average_beer_consumption
	average_weight
	characters_count
	genders {
		male
		female
		other
	}
	average_nemeses_years
	nemeses_count
	alive_nemeses {
		alive
		dead
	}
	secrets_count
}
`

func RootQuery(
	ctx_ context.Context,
	client_ graphql.Client,
) (*RootQueryResponse, error) {
	req_ := &graphql.Request{
		OpName: "RootQuery",
		Query:  RootQuery_Operation,
	}
	var err_ error

	var data_ RootQueryResponse
	resp_ := &graphql.Response{Data: &data_}

	err_ = client_.MakeRequest(
		ctx_,
		req_,
		resp_,
	)

	return &data_, err_
}
