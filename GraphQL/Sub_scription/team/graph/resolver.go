package graph

import (
	"errors"
	"team/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Teams        []*model.Team
	userAddedSub []*model.User
}

func (r *Resolver) getTeamByID(id string) (*model.Team, error) {
	for _, team := range r.Teams {
		if team.ID == id {
			return team, nil
		}
	}
	return nil, errors.New("team not found")
}
