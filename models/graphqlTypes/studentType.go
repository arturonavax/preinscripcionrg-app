package graphqlTypes

import (
	"github.com/graphql-go/graphql"
)

//StudentType Objeto GraphQL para Estudiantes.
var StudentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Student",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.Int,
			},
			"userID": &graphql.Field{
				Type: graphql.Int,
			},
			"statusID": &graphql.Field{
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
			"motherFirstName": &graphql.Field{
				Type: graphql.String,
			},
			"motherLastName": &graphql.Field{
				Type: graphql.String,
			},
			"motherEmail": &graphql.Field{
				Type: graphql.String,
			},
			"motherPhoneNumber": &graphql.Field{
				Type: graphql.String,
			},
			"motherCIType": &graphql.Field{
				Type: graphql.String,
			},
			"motherCI": &graphql.Field{
				Type: graphql.Int,
			},
			"fatherFirstName": &graphql.Field{
				Type: graphql.String,
			},
			"fatherLastName": &graphql.Field{
				Type: graphql.String,
			},
			"fatherEmail": &graphql.Field{
				Type: graphql.String,
			},
			"fatherPhoneNumber": &graphql.Field{
				Type: graphql.String,
			},
			"fatherCIType": &graphql.Field{
				Type: graphql.String,
			},
			"fatherCI": &graphql.Field{
				Type: graphql.Int,
			},
			"representativeFirstName": &graphql.Field{
				Type: graphql.String,
			},
			"representativeLastName": &graphql.Field{
				Type: graphql.String,
			},
			"representativeEmail": &graphql.Field{
				Type: graphql.String,
			},
			"representativePhoneNumber": &graphql.Field{
				Type: graphql.String,
			},
			"representativeCIType": &graphql.Field{
				Type: graphql.String,
			},
			"representativeCI": &graphql.Field{
				Type: graphql.Int,
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
				Type: graphql.String,
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
				Type: graphql.String,
			},
			"message": &graphql.Field{
				Type: graphql.String,
			},
			"code": &graphql.Field{
				Type: graphql.Int,
			},
		},
	},
)
