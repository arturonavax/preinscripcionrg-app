(function(){
    self.StudentsView = function() {
        this.ContainerStudents = $("#container-students");
    }

    self.StudentsView.prototype = {
        render: function(token){
            this.Token = token;
            var self = this;

            let queryStudents = `
            query {
                students {
                    id,
                    firstName,
                    lastName,
                    statusID,
                    ci,
                    SIGECOD
                }
            }
            
            `
            $("#student-listContainer").removeChild($("#student-list"));
            let studentList = document.createElement("div");
            studentList.id = "student-list";
            studentList.className = "student-list";
            $("#student-listContainer").appendChild(studentList);
            
            requestAjaxToken("POST", "/graphql", this.Token, queryStudents, false)
             .then( r => {
                 for (i=0; i < r.response.data.students.length; i++) {
                    let studentContainer = document.createElement("div");
                    studentContainer.className = "RE__container";

                    // -----------------------------------------
                    let studentContainerItem = document.createElement("div");
                    studentContainerItem.className = "RE-container__item";

                    let ItemP1 = document.createElement("p");
                    ItemP1.textContent = "Nombre: ";
                    let ItemSpan1 = document.createElement("span");
                    ItemSpan1.textContent = r.response.data.students[i].firstName + " " + r.response.data.students[i].lastName ;

                    let ItemP2 = document.createElement("p");
                    ItemP2.textContent = "C.I: ";
                    let ItemSpan2 = document.createElement("span");
                    ItemSpan2.textContent = r.response.data.students[i].ci;

                    let ItemSpan3 = document.createElement("span");
                    if (r.response.data.students[i].statusID === 1) {
                        ItemSpan3.textContent = "En Proceso..";
                    }

                    studentContainer.appendChild(studentContainerItem);
                    studentContainerItem.appendChild(ItemP1);
                    ItemP1.appendChild(ItemSpan1);
                    studentContainerItem.appendChild(ItemP2);
                    ItemP2.appendChild(ItemSpan2);
                    studentContainerItem.appendChild(ItemSpan3);
                    $("#student-list").appendChild(studentContainer);
                    // ---------------------------------------------

                    // -----------------------------------------
                    let studentContainerItem2 = document.createElement("div");
                    studentContainerItem2.className = "RE-container__item";

                    let ItemP12 = document.createElement("p");
                    ItemP12.textContent = "SIGE: ";
                    let ItemSpan12 = document.createElement("span");
                    ItemSpan12.textContent = r.response.data.students[i].SIGECOD;

                    let ItemP22 = document.createElement("p");
                    ItemP22.textContent = "ID: ";
                    let ItemSpan22 = document.createElement("span");
                    ItemSpan22.textContent = r.response.data.students[i].id;

                    let ItemButton = document.createElement("button");
                    ItemButton.textContent = "Ver";
                    ItemButton.className = "btnView";
                    ItemButton.id = "btnStudentView" + r.response.data.students[i].id;

                    studentContainer.appendChild(studentContainerItem2);
                    studentContainerItem2.appendChild(ItemP12);
                    ItemP12.appendChild(ItemSpan12);
                    studentContainerItem2.appendChild(ItemP22);
                    ItemP22.appendChild(ItemSpan22);
                    studentContainerItem2.appendChild(ItemButton);

                    ItemButton.value = r.response.data.students[i].id;

                    ItemButton.addEventListener("click", function(){
                        ClearStudentView();

                        let queryStudent = `
                            query {
                                student(id:${this.value}){
                                    statusID,
                                    id,
                                    countryOfBirthID,
                                    stateOfBirthID,
                                    municOfBirthID,
                                    institutionOfOriginID,
                                    municipalityID,
                                    parishID,
                                    sectorID,
                                    typeOfRoadID,
                                    mentionID,
                                    sectionID,
                                    conditionOfHousingID,

                                    SIGECOD,
                                    firstName,
                                    lastName,
                                    ciType,
                                    ci,
                                    dateOfBirth,
                                    gender,
                                    healthProblem,
                                    healthProblemE,
                                    email,
                                    phoneNumber,
                                    address,
                                    scholarship,
                                    canaima,
                                    year,
                                    age,
                                    size,
                                    weight,
                                    repeatAsignature,
                                    pendingAsignature,
                                    regular,
                                    repeat,
                                    asigPend,
                                    inscriptionDate,
                                    
                                    motherFirstName,
                                    motherLastName,
                                    motherPhoneNumber,
                                    motherCI,
                                
                                    fatherFirstName,
                                    fatherLastName,
                                    fatherPhoneNumber,
                                    fatherCI,
                                    
                                    representativeFirstName,
                                    representativeLastName,
                                    representativePhoneNumber,
                                    representativeCI,
                                    representativeRelationship,
                                    representativeAddress
                                }
                            }
                        `
                        requestAjaxToken("POST", "/graphql", self.Token, queryStudent, false)
                         .then(r => {
                             let student = r.response.data.student;

                             let queryCountryOfBirth = `
                                 query {
                                     country(id:${student.countryOfBirthID}){
                                         name
                                     }
                                 }
                             `
                             let queryStateOfBirth = `
                                 query {
                                     state(id:${student.stateOfBirthID}){
                                         name
                                     }
                                 }
                             `
                             let queryMunicOfBirth = `
                                 query {
                                     municipality(id:${student.municOfBirthID}){
                                         name
                                     }
                                 }
                             `
                             let queryInstitutionOfOrigin = `
                                 query {
                                     institution(id:${student.institutionOfOriginID}){
                                         name
                                     }
                                 }
                             `
                             let queryMunicipality = `
                                 query {
                                     municipality(id:${student.municipalityID}){
                                         name
                                     }
                                 }
                             `
                             let queryParish = `
                                 query {
                                     parish(id:${student.parishID}){
                                         name
                                     }
                                 }
                             `
                             let querySector = `
                                 query {
                                     sector(id:${student.sectorID}){
                                         name
                                     }
                                 }
                             `
                             let queryTypeOfRoad = `
                                 query {
                                     typeOfRoad(id:${student.typeOfRoadID}){
                                         name
                                     }
                                 }
                             `
                             let queryConditionOfHousing = `
                                 query {
                                     conditionOfHousing(id:${student.conditionOfHousingID}){
                                         name
                                     }
                                 }
                             `
                             let queryMention = `
                                 query {
                                     mention(id:${student.mentionID}){
                                         name
                                     }
                                 }
                             `
                             let querySection = `
                                 query {
                                     section(id:${student.sectionID}){
                                         name
                                     }
                                 }
                             `

                             requestAjaxToken("POST", "/graphql", self.Token, queryCountryOfBirth, true)
                             .then(r => {
                                let country = r.response.data.country;
                                $("#viewStudent-countryOfBirth").textContent = country.name;
                                $("#viewStudent-countryOfBirth2").textContent = country.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, queryStateOfBirth, true)
                             .then(r => {
                                let state = r.response.data.state;
                                $("#viewStudent-stateOfBirth").textContent = state.name;
                                $("#viewStudent-stateOfBirth2").textContent = state.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, queryMunicOfBirth, true)
                             .then(r => {
                                let municipality = r.response.data.municipality;
                                $("#viewStudent-municOfBirth").textContent = municipality.name;
                                $("#viewStudent-municOfBirth2").textContent = municipality.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, queryInstitutionOfOrigin, true)
                             .then(r => {
                                let institution = r.response.data.institution;
                                $("#viewStudent-institutionOfOrigin").textContent = institution.name;
                                $("#viewStudent-institutionOfOrigin2").textContent = institution.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, queryMunicipality, true)
                             .then(r => {
                                let municipality = r.response.data.municipality;
                                $("#viewStudent-municipality").textContent = municipality.name;
                                $("#viewStudent-municipality2").textContent = municipality.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, queryParish, true)
                             .then(r => {
                                let parish = r.response.data.parish;
                                $("#viewStudent-parish").textContent = parish.name;
                                $("#viewStudent-parish2").textContent = parish.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, querySector, true)
                             .then(r => {
                                let sector = r.response.data.sector;
                                $("#viewStudent-sector").textContent = sector.name;
                                $("#viewStudent-sector2").textContent = sector.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, queryTypeOfRoad, true)
                             .then(r => {
                                let typeOfRoad = r.response.data.typeOfRoad;
                                $("#viewStudent-typeOfRoad").textContent = typeOfRoad.name;
                                $("#viewStudent-typeOfRoad2").textContent = typeOfRoad.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, queryConditionOfHousing, true)
                             .then(r => {
                                let conditionOfHousing = r.response.data.conditionOfHousing;
                                $("#viewStudent-conditionOfHousing").textContent = conditionOfHousing.name;
                                $("#viewStudent-conditionOfHousing2").textContent = conditionOfHousing.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, queryMention, true)
                             .then(r => {
                                let mention = r.response.data.mention;
                                $("#viewStudent-mention").textContent = mention.name;
                                $("#viewStudent-mention2").textContent = mention.name;
                             });
                             requestAjaxToken("POST", "/graphql", self.Token, querySection, true)
                             .then(r => {
                                let section = r.response.data.section;
                                $("#viewStudent-section").textContent = section.name;
                                $("#viewStudent-section2").textContent = section.name;
                             });


                             $("#modal-studentView").classList.add("modal--open");

                             console.log(student);
                             if (student.statusID == 1) {
                                $("#viewStudent-status").textContent = "En proceso...";
                             }
                             $("#viewStudent-id").textContent = student.id;
                             $("#viewStudent-id2").textContent = student.id;
                             $("#viewStudent-SIGECOD").textContent = student.SIGECOD;
                             if (student.SIGECOD.length == 0) {
                                $("#viewStudent-SIGE2").textContent = "No";
                             } else {
                                $("#viewStudent-SIGE2").textContent = "Si";
                                $("#viewStudent-SIGECOD2").textContent = student.SIGECOD;
                             }
                             $("#viewStudent-firstName").textContent = student.firstName;
                             $("#viewStudent-firstName2").textContent = student.firstName;
                             $("#viewStudent-lastName").textContent = student.lastName;
                             $("#viewStudent-lastName2").textContent = student.lastName;
                             $("#viewStudent-ci").textContent = student.ciType+"-"+student.ci;
                             $("#viewStudent-ci2").textContent = student.ciType+"-"+student.ci;
                             $("#viewStudent-email").textContent = student.email;
                             $("#viewStudent-email2").textContent = student.email;
                             $("#viewStudent-phoneNumber").textContent = student.phoneNumber;
                             $("#viewStudent-phoneNumber2").textContent = student.phoneNumber;
                             $("#viewStudent-gender").textContent = student.gender;
                             $("#viewStudent-gender2").textContent = student.gender;
                             $("#viewStudent-dateOfBirth").textContent = student.dateOfBirth;
                             $("#viewStudent-dateOfBirth2").textContent = student.dateOfBirth;
                             $("#viewStudent-healthProblemE").textContent = student.healthProblemE;
                             $("#viewStudent-healthProblemE2").textContent = student.healthProblemE;
                             $("#viewStudent-address").textContent = student.address;
                             $("#viewStudent-address2").textContent = student.address;
                             $("#viewStudent-motherFirstName").textContent = student.motherFirstName;
                             $("#viewStudent-motherFirstName2").textContent = student.motherFirstName;
                             $("#viewStudent-motherLastName").textContent = student.motherLastName;
                             $("#viewStudent-motherLastName2").textContent = student.motherLastName;
                             $("#viewStudent-motherPhoneNumber").textContent = student.motherPhoneNumber;
                             $("#viewStudent-motherPhoneNumber2").textContent = student.motherPhoneNumber;
                             $("#viewStudent-motherCI").textContent = student.motherCI;
                             $("#viewStudent-motherCI2").textContent = student.motherCI;
                             $("#viewStudent-fatherFirstName").textContent = student.fatherFirstName;
                             $("#viewStudent-fatherFirstName2").textContent = student.fatherFirstName;
                             $("#viewStudent-fatherLastName").textContent = student.fatherLastName;
                             $("#viewStudent-fatherLastName2").textContent = student.fatherLastName;
                             $("#viewStudent-fatherPhoneNumber").textContent = student.fatherPhoneNumber;
                             $("#viewStudent-fatherPhoneNumber2").textContent = student.fatherPhoneNumber;
                             $("#viewStudent-fatherCI").textContent = student.fatherCI;
                             $("#viewStudent-fatherCI2").textContent = student.fatherCI;
                             $("#viewStudent-representativeFirstName").textContent = student.representativeFirstName;
                             $("#viewStudent-representativeFirstName2").textContent = student.representativeFirstName;
                             $("#viewStudent-representativeLastName").textContent = student.representativeLastName;
                             $("#viewStudent-representativeLastName2").textContent = student.representativeLastName;
                             $("#viewStudent-representativePhoneNumber").textContent = student.representativePhoneNumber;
                             $("#viewStudent-representativePhoneNumber2").textContent = student.representativePhoneNumber;
                             $("#viewStudent-representativeCI").textContent = student.representativeCI;
                             $("#viewStudent-representativeCI2").textContent = student.representativeCI;
                             $("#viewStudent-representativeRelationship").textContent = student.representativeRelationship;
                             $("#viewStudent-representativeRelationship2").textContent = student.representativeRelationship;
                             $("#viewStudent-representativeAddress").textContent = student.representativeAddress;
                             $("#viewStudent-representativeAddress2").textContent = student.representativeAddress;
                             if (student.scholarship == true) {
                                $("#viewStudent-scholarship").textContent = "Si";
                                $("#viewStudent-scholarship2").textContent = "Si";

                             } else {
                                $("#viewStudent-scholarship").textContent = "No";
                                $("#viewStudent-scholarship2").textContent = "No";
                             }

                             if (student.canaima == true) {
                                $("#viewStudent-canaima").textContent = "Si";
                                $("#viewStudent-canaima2").textContent = "Si";
                             } else {
                                $("#viewStudent-canaima").textContent = "No";
                                $("#viewStudent-canaima2").textContent = "No";
                             }

                             if (student.regular == true) {
                                $("#viewStudent-regular2").textContent = "Si";
                             } else {
                                $("#viewStudent-regular2").textContent = "No";
                             }
                             if (student.repeat == true) {
                                $("#viewStudent-repeat2").textContent = "Si";
                             } else {
                                $("#viewStudent-repeat2").textContent = "No";
                             }
                             if (student.asigPend == true) {
                                $("#viewStudent-asigPend2").textContent = "Si";
                             } else {
                                $("#viewStudent-asigPend2").textContent = "No";
                             }
                             let dat = student.inscriptionDate[0]+student.inscriptionDate[1]+student.inscriptionDate[2]+student.inscriptionDate[3];
                             $("#viewStudent-year22").textContent = (dat-1)+"-"+dat;
                             $("#viewStudent-asigRepeatE2").textContent = student.repeatAsignature;
                             $("#viewStudent-asigPendE2").textContent = student.pendingAsignature;
                             $("#viewStudent-dateOfInscription2").textContent = student.inscriptionDate;
                             $("#viewStudent-year").textContent = student.year;
                             $("#viewStudent-age").textContent = student.age;
                             $("#viewStudent-age2").textContent = student.age;
                             $("#viewStudent-size").textContent = student.size;
                             $("#viewStudent-size2").textContent = student.size;
                             $("#viewStudent-weight").textContent = student.weight;
                             $("#viewStudent-weight2").textContent = student.weight;
                         });

                    });


                    // ---------------------------------------------
                 }

             });
        }

    }
})();