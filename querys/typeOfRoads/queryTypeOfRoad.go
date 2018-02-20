package typeOfRoads

import (
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/graphql-go/graphql"
)

//QueryTypeOfRoad Query para consultar un tipo de via.
var QueryTypeOfRoad = &graphql.Field{
	Type: graphqlTypes.TypeOfRoadType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return nil, nil
	},
}
