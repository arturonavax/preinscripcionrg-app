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

        this.CountriesOfBirth = $("#countriesOfBirth-studentRegister");
        this.StatesOfBirth = $("#statesOfBirth-studentRegister");
        this.MunicipalitiesOfBirth = $("#municipalitiesOfBirth-studentRegister");
        this.InstitutionsOfBirth = $("#institutionsOfBirth-studentRegister");

        this.Municipalities = $("#municipalities-studentRegister");
        this.Parishes = $("#parishes-studentRegister");
        this.Sectors = $("#sectors-studentRegister");
        this.TypeOfRoads = $("#typeOfRoads-studentRegister");
        this.ConditionOfHousings = $("#conditionOfHousing-studentRegister");

        this.Mentions = $("#mentions-studentRegister");
        this.Sections = $("#sections-studentRegister");

        // CAMPOS
        this.SIGERadio = document.getElementsByName("SIGE");
        this.SIGECODInput = $("#SIGE-codInput");

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

                     this.InstitutionsOfBirth.appendChild(institution);
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
                    countryOfBirthID : 1,
                    stateOfBirthID : 1,
                    municOfBirthID : 1,
                    institutionOfOriginID : 1,
                    municipalityID : 1,
                    parishID : 1,
                    sectorID : 1,
                    typeOfRoadID : 1,
                    mentionID : 1,
                    sectionID : 1,
                    teacherID : 1,
                    conditionOfHousingID : 1,
                    
                    SIGECOD : "A1001",
                    motherFirstName: "Francis",
                    motherLastName: "Peñaranda",
                    motherEmail: "francis@gmail.com",
                    motherPhoneNumber: "+58928273493",
                    motherCI: "V-1727392",
                
                    representativeFirstName: "Francis",
                    representativeLastName: "Peñaranda",
                    representativeEmail: "francis@gmail.com",
                    representativePhoneNumber: "+58928273493",
                    representativeCI: "V-1727392",
                    representativeRelationship: "Madre",
                    representativeAddress: "La lago",
                    
                    fatherFirstName: "Dario",
                    fatherLastName: "Peñaranda",
                    fatherEmail: "dario@gmail.com",
                    fatherPhoneNumber: "+58928273423",
                    fatherCI: "V-1727772",
                    
                    firstName : "Frank",
                    lastName : "Peñaranda",
                    ciType : "V",
                    ci : 29100800,
                    dateOfBirth : "11/11/01",
                    gender : "M",
                    healthProblem : false,
                    healthProblemE : "",
                    email : "coffemanfp@gmail.com",
                    phoneNumber : "+584167634291",
                    address : "Valle Frio",
                    scholarship : true,
                    canaima : true,
                    year : 4,
                    age : 15,
                    size : "L",
                    weight : 60,
                    repeatAsignature : "Castellano",
                    pendingAsignature : "",
                    regular : true,
                    inscriptionDate : "10/08/15"
                    
                ){id,message,code}
            }
            `
        }
    }
})();