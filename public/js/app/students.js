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
                 console.log(r);
                 $("#student-name").textContent = r.response.data.students[0].firstName + r.response.data.students[0].lastName ;
                 if (r.response.data.students[0].statusID === 1) {
                    $("#student-status").textContent = "En Proceso..";
                 }
                 $("#student-ci").textContent = r.response.data.students[0].ci;
                 $("#student-sige").textContent = r.response.data.students[0].SIGECOD;
                 $("#student-id").textContent = r.response.data.students[0].id;

             });
        }

    }
})();