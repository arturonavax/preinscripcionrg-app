package mentions

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryMentions Query GraphQL para consultar lista de menciones
var QueryMentions = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.MentionType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListMentions []student.Mention

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadMentions {
			db.Find(&ListMentions)

			for i := 0; i < len(ListMentions); i++ {
				ListMentions[i].Message = "#QueryMentions# : Consulta exitosa."
				ListMentions[i].Code = http.StatusAccepted
			}

		} else {
			Mention := student.Mention{
				Message: "#QueryMentions# : No tienes permisos para leer Menciones.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListMentions = append(ListMentions, Mention)

		}
		return ListMentions, nil
	},
}
