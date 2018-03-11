(function(){
    self.StudentsView = function() {
        this.ContainerStudents = $("#container-students");
    }

    self.StudentsView.prototype = {
        render: function(token){
            this.Token = token;

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
                    ItemSpan1.id = "student-name";
                    ItemSpan1.textContent = r.response.data.students[i].firstName + " " + r.response.data.students[i].lastName ;

                    let ItemP2 = document.createElement("p");
                    ItemP2.textContent = "C.I: ";
                    let ItemSpan2 = document.createElement("span");
                    ItemSpan2.textContent = r.response.data.students[i].ci;

                    let ItemSpan3 = document.createElement("span");
                    ItemSpan3.id = "student-status";
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