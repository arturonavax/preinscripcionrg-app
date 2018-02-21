package mentions

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryMention Query para consultar una mencion.
var QueryMention = &graphql.Field{
	Type: graphqlTypes.MentionType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idQuery, isOK := p.Args["id"].(int)
		kind := &models.UserKind{}
		mentionFound := &student.Mention{}

		kindID := p.Context.Value("user").(models.User).KindID

		if isOK {
			db := databases.GetConnectionDB()
			defer db.Close()

			db.Where("id = ?", kindID).First(&kind)

			if kind.ReadMentions {
				err := db.Where("id = ?", idQuery).First(&mentionFound).Error

				if err != nil {
					mentionFound.Message = "#QueryMention# : Ocurrio un error."
					mentionFound.Code = http.StatusInternalServerError
				} else {
					mentionFound.Message = "#QueryMention# : Consulta exitosa."
					mentionFound.Code = http.StatusAccepted
				}
			} else {
				mentionFound.Message = "#QueryMention# : No tienes permisos para leer menciones."
				mentionFound.Code = http.StatusNonAuthoritativeInfo
			}
		}

		return mentionFound, nil
	},
}
