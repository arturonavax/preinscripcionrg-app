package parishes

import (
	"net/http"

	"github.com/arthurnavah/PreInscripcionRG/databases"
	"github.com/arthurnavah/PreInscripcionRG/models"
	"github.com/arthurnavah/PreInscripcionRG/models/graphqlTypes"
	"github.com/arthurnavah/PreInscripcionRG/models/student"
	"github.com/graphql-go/graphql"
)

//QueryParishes Query GraphQL para consultar lsita de parroquias
var QueryParishes = &graphql.Field{
	Type: graphql.NewList(graphqlTypes.ParishType),
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		kind := &models.UserKind{}
		var ListParishes []student.Parish

		kindID := p.Context.Value("user").(models.User).KindID

		db := databases.GetConnectionDB()
		defer db.Close()

		db.Where("id = ?", kindID).First(&kind)

		if kind.ReadParishes {
			db.Find(&ListParishes)

			for i := 0; i < len(ListParishes); i++ {
				ListParishes[i].Message = "#QueryParishes# : Consulta exitosa."
				ListParishes[i].Code = http.StatusAccepted
			}
		} else {

			Parish := student.Parish{
				Message: "#QueryParishes# : No tienes permisos para leer parroquias.",
				Code:    http.StatusNonAuthoritativeInfo,
			}

			ListParishes = append(ListParishes, Parish)
		}
		return ListParishes, nil
	},
}
