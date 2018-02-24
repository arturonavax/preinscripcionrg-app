(function(){
    self.StudentRegister = function() {
    }
})();

(function(){
    self.StudentRegisterView = function(studentRegister) {
        this.StudentRegister = studentRegister;
        this.BtnStudentRegister = $("#btn-registerStudent");
        this.ContainerStudentRegister = $("#container-studentRegister");
        this.PanelMove = $("#panelMove-ficha");
        this.BtnNextPage1 = $("#btn-next-page1");
        this.BtnNextPage2 = $("#btn-next-page2");
        this.BtnNextPage3 = $("#btn-next-page3");
        this.BtnNextPage4 = $("#btn-next-page4");
        this.BtnNextPage5 = $("#btn-next-page5");
        this.BtnNextPage6 = $("#btn-next-page6");
        this.BtnNextPage7 = $("#btn-next-page7");
        this.BtnNextPage8 = $("#btn-next-page8");
        this.BtnNextPage9 = $("#btn-next-page9");
        this.BtnNextPage10 = $("#btn-next-page10");

        this.BtnBackPage2 = $("#btn-back-page2");
        this.BtnBackPage3 = $("#btn-back-page3");
        this.BtnBackPage4 = $("#btn-back-page4");
        this.BtnBackPage5 = $("#btn-back-page5");
        this.BtnBackPage6 = $("#btn-back-page6");
        this.BtnBackPage7 = $("#btn-back-page7");
        this.BtnBackPage8 = $("#btn-back-page8");
        this.BtnBackPage9 = $("#btn-back-page9");
        this.BtnBackPage10 = $("#btn-back-page10");
        this.BtnBackPage11 = $("#btn-back-page11");

        this.SIGEYes = $("#SIGE-yes");
        this.SIGENo = $("#SIGE-no");

        this.SIGETitle = $("#SIGE-title");
        this.SIGECOD = $("#SIGE-cod");

        // CAMPOS
        this.SIGERadio = document.getElementsByName("SIGE");
        this.SIGECODInput = $("#SIGE-codInput");

        this.FirstNameInput = $("#firstName-studentRegister");
        this.LastNameInput = $("#lastName-studentRegister");
        this.CIInput = $("#ci-studentRegister");
        this.EmailInput = $("#email-studentRegister");
        this.PhoneNumberInput = $("#phoneNumber-studentRegister");
        this.GenderRadioInput = document.getElementsByName("genero");

        this.DateOfBirthInput = $("#dateOfBirth-studentRegister");
        this.CountriesOfBirth = $("#countriesOfBirth-studentRegister");
        this.StatesOfBirth = $("#statesOfBirth-studentRegister");
        this.MunicipalitiesOfBirth = $("#municipalitiesOfBirth-studentRegister");

        this.InstitutionsOfOrigin = $("#institutionsOfOrigin-studentRegister");
        this.HealthProblemRadio = document.getElementsByName("salud");
        this.HealthProblemnEInput = $("#healthProblemE-studentRegister");

        this.AddressInput = $("#address-studentRegister");
        this.Municipalities = $("#municipalities-studentRegister");
        this.Parishes = $("#parishes-studentRegister");
        this.Sectors = $("#sectors-studentRegister");
        this.TypeOfRoads = $("#typeOfRoads-studentRegister");
        this.ConditionOfHousings = $("#conditionOfHousing-studentRegister");

        this.MotherFirstNameInput = $("#motherFirstName-studentRegister");
        this.MotherLastNameInput = $("#motherLastName-studentRegister");
        this.MotherCIInput = $("#motherCI-studentRegister");
        this.MotherPhoneNumberInput = $("motherPhoneNumber-studentRegister");

        this.FatherFirstNameInput = $("#fatherFirstName-studentRegister");
        this.FatherLastNameInput = $("#fatherLastName-studentRegister");
        this.FatherCIInput = $("#fatherCI-studentRegister");
        this.FatherPhoneNumberInput = $("fatherPhoneNumber-studentRegister");

        this.RepresentativeFirstNameInput = $("#representativeFirstName-studentRegister");
        this.RepresentativeLastNameInput = $("#representativeLastName-studentRegister");
        this.RepresentativeCIInput = $("#representativeCI-studentRegister");
        this.RepresentativeRelationshipInput = $("#representativeRelationship-studentRegister");
        this.RepresentativeAddressInput = $("#representativeAddress-studentRegister");
        this.RepresentativePhoneNumberInput = $("representativePhoneNumber-studentRegister");

        this.BecaRadio = document.getElementsByName("Beca");
        this.CanaimaRadio = document.getElementsByName("Canaima");

        this.Mentions = $("#mentions-studentRegister");
        this.Sections = $("#sections-studentRegister");
        this.AgeInput = $("#age-studentRegister");
        this.SizeInput = $("#size-studentRegister");
        this.WeightInput = $("#weight-studentRegister");

        this.Regular = $("#regular-studentRegister");
        this.Repet = $("#repet-studentRegister");
        this.AsigPend = $("#asigPend-studentRegister");

        this.AsigRepetE = $("#asigRepetE-studentRegister");
        this.AsigPendE = $("#asigPendE-studentRegister");

        this.BtnPreInscribir = $("#btn-preinscribir");

    }

    self.StudentRegisterView.prototype = {
        render: function(token) {
            //Query de paises
            let queryCountries = `
            query {
                countries {
                    id,
                    name
                }
            }
            `
            //Query de estados
            let queryStates = `
            query {
                states {
                    id,
                    name
                }
            }
            `
            //Query de municipios
            let queryMunicipalities = `
            query {
                municipalities {
                    id,
                    name
                }
            }
            `
            //Query de instituciones
            let queryInstitutions = `
            query {
                institutions {
                    id,
                    name
                }
            }
            `
            //Query de parroquias
            let queryParishes = `
            query {
                parishes {
                    id,
                    name
                }
            }
            `
            //Query de sectores
            let querySectors = `
            query {
                sectors {
                    id,
                    name
                }
            }
            `
            //Query de tipos de vias
            let queryTypeOfRoads = `
            query {
                typeOfRoads {
                    id,
                    name
                }
            }
            `
            //Query de condiciones de viviendas
            let queryConditionOfHousings = `
            query {
                conditionOfHousings {
                    id,
                    name
                }
            }
            `
            //Query de menciones 
            let queryMentions = `
            query {
                mentions {
                    id,
                    name
                }
            }
            `
            //Query de secciones 
            let querySections = `
            query {
                sections {
                    id,
                    name
                }
            }
            `

            //Renderea los paises 
            requestAjaxToken("POST", "/graphql", token, queryCountries, true)
             .then( r => {
                 let countries = r.response.data.countries;
                 for (let i = 0; i < countries.length; i++){
                     let country = document.createElement("option");
                     country.textContent = countries[i].name;
                     country.value = countries[i].id;

                     this.CountriesOfBirth.appendChild(country);
                 }
             });
            //Renderea los estados 
            requestAjaxToken("POST", "/graphql", token, queryStates, true)
             .then( r => {
                 let states = r.response.data.states;
                 for (let i = 0; i < states.length; i++){
                     let state = document.createElement("option");
                     state.textContent = states[i].name;
                     state.value = states[i].id;

                     this.StatesOfBirth.appendChild(state);
                 }
             });
            //Renderea los municipios 
            requestAjaxToken("POST", "/graphql", token, queryMunicipalities, true)
             .then( r => {
                 let municipalities = r.response.data.municipalities;
                 for (let i = 0; i < municipalities.length; i++){
                     let municipalityOfBirth = document.createElement("option");
                     municipalityOfBirth.textContent = municipalities[i].name;
                     municipalityOfBirth.value = municipalities[i].id;

                     let municipality = document.createElement("option");
                     municipality.textContent = municipalities[i].name;
                     municipality.value = municipalities[i].id;

                     this.MunicipalitiesOfBirth.appendChild(municipalityOfBirth);
                     this.Municipalities.appendChild(municipality);
                 }
             });
            //Renderea las institutiones 
            requestAjaxToken("POST", "/graphql", token, queryInstitutions, true)
             .then( r => {
                 let institutions = r.response.data.institutions;
                 for (let i = 0; i < institutions.length; i++){
                     let institution = document.createElement("option");
                     institution.textContent = institutions[i].name;
                     institution.value = institutions[i].id;

                     this.InstitutionsOfOrigin.appendChild(institution);
                 }
             });
            //Renderea las parroquias 
            requestAjaxToken("POST", "/graphql", token, queryParishes, true)
             .then( r => {
                 let parishes = r.response.data.parishes;
                 for (let i = 0; i < parishes.length; i++){
                     let parish = document.createElement("option");
                     parish.textContent = parishes[i].name;
                     parish.value = parishes[i].id;

                     this.Parishes.appendChild(parish);
                 }
             });
            //Renderea los sectores 
            requestAjaxToken("POST", "/graphql", token, querySectors, true)
             .then( r => {
                 let sectors = r.response.data.sectors;
                 for (let i = 0; i < sectors.length; i++){
                     let sector = document.createElement("option");
                     sector.textContent = sectors[i].name;
                     sector.value = sectors[i].id;

                     this.Sectors.appendChild(sector);
                 }
             });
            //Renderea los tipos de vias 
            requestAjaxToken("POST", "/graphql", token, queryTypeOfRoads, true)
             .then( r => {
                 let typeOfRoads = r.response.data.typeOfRoads;
                 for (let i = 0; i < typeOfRoads.length; i++){
                     let typeOfRoad = document.createElement("option");
                     typeOfRoad.textContent = typeOfRoads[i].name;
                     typeOfRoad.value = typeOfRoads[i].id;

                     this.TypeOfRoads.appendChild(typeOfRoad);
                 }
             });
            //Renderea las condiciones de estudiantes
            requestAjaxToken("POST", "/graphql", token, queryConditionOfHousings, true)
             .then( r => {
                 let conditionOfHousings = r.response.data.conditionOfHousings;
                 for (let i = 0; i < conditionOfHousings.length; i++){
                     let conditionOfHousing = document.createElement("option");
                     conditionOfHousing.textContent = conditionOfHousings[i].name;
                     conditionOfHousing.value = conditionOfHousings[i].id;

                     this.ConditionOfHousings.appendChild(conditionOfHousing);
                 }
             });
            //Renderea las menciones
            requestAjaxToken("POST", "/graphql", token, queryMentions, true)
             .then( r => {
                 let mentions = r.response.data.mentions;
                 for (let i = 0; i < mentions.length; i++){
                     let mention = document.createElement("option");
                     mention.textContent = mentions[i].name;
                     mention.value = mentions[i].id;

                     this.Mentions.appendChild(mention);
                 }
             });
            //Renderea las secciones
            requestAjaxToken("POST", "/graphql", token, querySections, true)
             .then( r => {
                 let sections = r.response.data.sections;
                 for (let i = 0; i < sections.length; i++){
                     let section = document.createElement("option");
                     section.textContent = sections[i].name;
                     section.value = sections[i].id;

                     this.Sections.appendChild(section);
                 }
             });
        },

        register: function(){

            let mutation = `
            mutation {
                studentC (
                    countryOfBirthID : ${this.CountriesOfBirth.value},
                    stateOfBirthID : ${this.StatesOfBirth.value},
                    municOfBirthID : ${this.MunicipalitiesOfBirth.value},
                    institutionOfOriginID : ${this.InstitutionsOfOrigin.value},
                    municipalityID : ${this.Municipalities.value},
                    parishID : ${this.Parishes.value},
                    sectorID : ${this.Sectors.value},
                    typeOfRoadID : ${this.TypeOfRoads.value},
                    mentionID : ${this.Mentions.value},
                    sectionID : ${this.Sections.value},
                    conditionOfHousingID : ${this.ConditionOfHousings.value},
                    
                    SIGECOD : ${this.SIGECODInput.value},
                    firstName : ${this.FirstNameInput.value},
                    lastName : ${this.LastNameInput.value},
                    ci : ${this.CIInput.value},
                    dateOfBirth : ${this.DateOfBirthInput.value},
                    gender : "M",
                    healthProblem : false,
                    healthProblemE : ${this.HealthProblemnEInput.value},
                    email : ${this.EmailInput.value},
                    phoneNumber : ${this.PhoneNumberInput.value},
                    address : ${this.AddressInput.value},
                    scholarship : true,
                    canaima : true,
                    year : ${this.YearInput.value},
                    age : ${this.AgeInput.value},
                    size : ${this.SizeInput.value},
                    weight : ${this.WeightInput.value},
                    repeatAsignature : ${this.AsigRepetE.value},
                    pendingAsignature :${this.AsigPendE.value},
                    regular : true,
                    inscriptionDate : ${this.InscriptionDateInput.value},
                    
                    motherFirstName: ${this.MotherFirstNameInput.value},
                    motherLastName: ${this.MotherLastNameInput.value},
                    motherPhoneNumber: ${this.MotherPhoneNumberInput.value},
                    motherCI: ${this.MotherCIInput.value},
                
                    fatherFirstName: ${this.FatherFirstNameInput.value},
                    fatherLastName: ${this.FatherLastNameInput.value},
                    fatherPhoneNumber: ${this.FatherPhoneNumber.value},
                    fatherCI: ${this.FatherCIInput.value},
                    
                    representativeFirstName: ${this.RepresentativeFirstNameInput.value},
                    representativeLastName: ${this.RepresentativeLastNameInput.value},
                    representativePhoneNumber: ${this.RepresentativePhoneNumberInput.value},
                    representativeCI: ${this.RepresentativeCIInput.value},
                    representativeRelationship: ${this.RepresentativeRelationshipInput.value},
                    representativeAddress: ${this.RepresentativeAddressInput.value},
                    
                ){id,message,code}
            }
            `
            console.log(mutation);
        }
    }
})();