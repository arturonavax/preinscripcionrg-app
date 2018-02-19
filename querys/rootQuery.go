package querys

import (
	"github.com/arthurnavah/PreInscripcionRG/querys/asignatures"
	"github.com/arthurnavah/PreInscripcionRG/querys/countries"
	"github.com/arthurnavah/PreInscripcionRG/querys/me"
	"github.com/arthurnavah/PreInscripcionRG/querys/mentions"
	"github.com/arthurnavah/PreInscripcionRG/querys/municipalities"
	"github.com/arthurnavah/PreInscripcionRG/querys/parishes"
	"github.com/arthurnavah/PreInscripcionRG/querys/sections"
	"github.com/arthurnavah/PreInscripcionRG/querys/sectors"
	"github.com/arthurnavah/PreInscripcionRG/querys/states"
	"github.com/arthurnavah/PreInscripcionRG/querys/studentConditions"
	"github.com/arthurnavah/PreInscripcionRG/querys/students"
	"github.com/arthurnavah/PreInscripcionRG/querys/teachers"
	"github.com/arthurnavah/PreInscripcionRG/querys/typeOfRoads"
	"github.com/arthurnavah/PreInscripcionRG/querys/users"
	"github.com/graphql-go/graphql"
)

//RootQuery administrador de Querys
var RootQuery = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			//Querys Yo
			"me": me.QueryMe,
			//Querys Usuarios
			"users": users.QueryUsers,
			//Querys Estudiantes
			"students": students.QueryStudents,
			"student":  students.QueryStudent,
			//Querys Asignaturas
			"asignatures": asignatures.QueryAsignatures,
			"asignature":  asignatures.QueryAsignature,
			//Querys Paises
			"countries": countries.QueryCountries,
			//Querys Menciones
			"mentions": mentions.QueryMentions,
			//Querys Municipios
			"municipalities": municipalities.QueryMunicipalities,
			//Querys Parroquias
			"parishes": parishes.QueryParishes,
			//Querys Sectores
			"sectors": sectors.QuerySectors,
			//Querys Secciones
			"sections": sections.QuerySections,
			//Querys Estados
			"states": states.QueryStates,
			//Querys Profesores
			"teachers": teachers.QueryTeachers,
			//Querys Tipos de vias
			"typeOfRoads": typeOfRoads.QueryTypeOfRoads,
			//Querys Condiciones de Estudiantes
			"studentConditions": studentConditions.QueryStudentConditions,
		},
	},
)
