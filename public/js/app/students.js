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
            
            requestAjaxToken("POST", "/graphql", this.Token, queryStudents)
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

                    var studentID = r.response.data.students[i].id;
                    ItemButton.addEventListener("click", function(){
                        let queryStudent = `
                            query {
                                student(id:${studentID}){
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

                             $("#modal-studentView").classList.add("modal--open");

                             $("#viewStudent-id").textContent = student.id;
                             $("#viewStudent-SIGE").textContent = student.SIGECOD;
                             $("#viewStudent-firstName").textContent = student.firstName;
                             $("#viewStudent-lastName").textContent = student.lastName;
                             $("#viewStudent-ci").textContent = student.ci;
                             $("#viewStudent-email").textContent = student.email;
                             $("#viewStudent-phoneNumber").textContent = student.phoneNumber;
                             $("#viewStudent-gender").textContent = student.gender;
                             $("#viewStudent-dateOfBirth").textContent = student.dateOfBirth;
                             $("#viewStudent-countryOfBirth").textContent = student.countryOfBirthID;
                             $("#viewStudent-stateOfBirth").textContent = student.stateOfBirthID;
                             $("#viewStudent-municOfBirth").textContent = student.municOfBirthID;
                             $("#viewStudent-institutionOfOrigin").textContent = student.institutionOfOriginID;
                             $("#viewStudent-healthProblemE").textContent = student.healthProblemE;
                             $("#viewStudent-address").textContent = student.address;
                             $("#viewStudent-municipality").textContent = student.municipalityID;
                             $("#viewStudent-parish").textContent = student.parishID;
                             $("#viewStudent-sector").textContent = student.sectorID;
                             $("#viewStudent-typeOfRoad").textContent = student.typeOfRoadID;
                             $("#viewStudent-conditionOfHousing").textContent = student.conditionOfHousingID;
                             $("#viewStudent-motherFirstName").textContent = student.motherFirstName;
                             $("#viewStudent-motherLastName").textContent = student.motherLastName;
                             $("#viewStudent-motherPhoneNumber").textContent = student.motherPhoneNumber;
                             $("#viewStudent-motherCI").textContent = student.motherCI;
                             $("#viewStudent-fatherFirstName").textContent = student.fatherFirstName;
                             $("#viewStudent-fatherLastName").textContent = student.fatherLastName;
                             $("#viewStudent-fatherPhoneNumber").textContent = student.fatherPhoneNumber;
                             $("#viewStudent-fatherCI").textContent = student.fatherCI;
                             $("#viewStudent-representativeFirstName").textContent = student.representativeFirstName;
                             $("#viewStudent-representativeLastName").textContent = student.representativeLastName;
                             $("#viewStudent-representativePhoneNumber").textContent = student.representativePhoneNumber;
                             $("#viewStudent-representativeCI").textContent = student.representativeCI;
                             $("#viewStudent-representativeRelationship").textContent = student.representativeRelationship;
                             $("#viewStudent-representativeAddress").textContent = student.representativeAddress;
                             $("#viewStudent-scholarship").textContent = student.scholarship;
                             $("#viewStudent-canaima").textContent = student.canaima;
                             $("#viewStudent-mention").textContent = student.mentionID;
                             $("#viewStudent-section").textContent = student.sectionID;
                             $("#viewStudent-year").textContent = student.year;
                             $("#viewStudent-age").textContent = student.age;
                             $("#viewStudent-size").textContent = student.size;
                             $("#viewStudent-weight").textContent = student.weight;
                         });


                    });

                    studentContainer.appendChild(studentContainerItem2);
                    studentContainerItem2.appendChild(ItemP12);
                    ItemP12.appendChild(ItemSpan12);
                    studentContainerItem2.appendChild(ItemP22);
                    ItemP22.appendChild(ItemSpan22);
                    studentContainerItem2.appendChild(ItemButton);
                    // ---------------------------------------------
                 }

             });
        }

    }
})();