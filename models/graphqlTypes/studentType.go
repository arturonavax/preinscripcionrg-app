package graphqlTypes

import (
	"github.com/graphql-go/graphql"
)

//StudentType Objecto GraphQL para Estudiantes.
var StudentType = grahql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"countryOfBirthID": &graphql.Field{
				Type: graphql.Int,
			},
			"stateOfBirthID": &graphql.Field{
				Type: graphql.Int,
			},
			"municOfBirthID": &graphql.Field{
				Type: graphql.Int,
			},
			"municipalityID": &graphql.Field{
				Type: graphql.Int,
			},
			"institutionOfOriginID": &graphql.Field{
				Type: graphql.Int,
			},
			"parishID": &graphql.Field{
				Type: graphql.Int,
			},
			"sectorID": &graphql.Field{
				Type: graphql.Int,
			},
			"typeOfRoadID": &graphql.Field{
				Type: graphql.Int,
			},
			"fatherID": &graphql.Field{
				Type: graphql.Int,
			},
			"motherID": &graphql.Field{
				Type: graphql.Int,
			},
			"mentionID": &graphql.Field{
				Type: graphql.Int,
			},
			"sectionID": &graphql.Field{
				Type: graphql.Int,
			},
			"studentConditionID": &graphql.Field{
				Type: graphql.Int,
			},
			"representativeID": &graphql.Field{
				Type: graphql.Int,
			},
			"teacherID": &graphql.Field{
				Type: graphql.Int,
			},
			"SIGECOD": &graphql.Field{
				Type: graphql.String,
			},
			"firstName": &graphql.Field{
				Type: graphql.String,
			},
			"lastName": &graphql.Field{
				Type: graphql.String,
			},
			"ciType": &graphql.Field{
				Type: graphql.String,
			},
			"ci": &graphql.Field{
				Type: graphql.Int,
			},
			"dateOfBirth": &graphql.Field{
				Type: graphql.DateTime,
			},
			"gender": &graphql.Field{
				Type: graphql.String,
			},
			"healthProblem": &graphql.Field{
				Type: graphql.Boolean,
			},
			"healthProblemE": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"phoneNumber": &graphql.Field{
				Type: graphql.String,
			},
			"address": &graphql.Field{
				Type: graphql.String,
			},
			"scholarship": &graphql.Field{
				Type: graphql.Boolean,
			},
			"canaima": &graphql.Field{
				Type: graphql.Boolean,
			},
			"conditionOfHousing": &graphql.Field{
				Type: graphql.String,
			},
			"year": &graphql.Field{
				Type: graphql.Int,
			},
			"age": &graphql.Field{
				Type: graphql.Int,
			},
			"size": &graphql.Field{
				Type: graphql.String,
			},
			"weight": &graphql.Field{
				Type: graphql.Int,
			},
			"repeatAsignature": &graphql.Field{
				Type: graphql.String,
			},
			"pendingAsignature": &graphql.Field{
				Type: graphql.String,
			},
			"regular": &graphql.Field{
				Type: graphql.Boolean,
			},
			"inscriptionDate": &graphql.Field{
				Type: graphql.DateTime,
			},
		},
	},
)
