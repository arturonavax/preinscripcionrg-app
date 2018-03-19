function $(select) {
    return document.querySelector(select);
}

function requestAjax(method, url, obj) {
    return new Promise(function(resolver, rechazar){
        let xhr = new XMLHttpRequest();

        xhr.open(method, url, true);
        xhr.setRequestHeader("Content-Type", "application/json");
        xhr.addEventListener("load", e => {
            let self = e.target;
            let respuesta = {
                status: self.status,
                response: JSON.parse(self.response)
            };
            resolver(respuesta)
        })
        xhr.addEventListener("error", e => {
            let self = e.target;
    
            rechazar(selft);
        });
    
        xhr.send(obj);
    });

}

function requestAjaxToken(method, url, token, obj, async) {
    return new Promise(function(resolver, rechazar){
        let xhr = new XMLHttpRequest();

        xhr.open(method, url, async);
        xhr.setRequestHeader("Content-Type", "application/graphql");
        xhr.setRequestHeader("Authorization", token);
        xhr.withCredentials = true;
        xhr.setRequestHeader("cache-control", "no-cache");
        xhr.addEventListener("load", e => {
            let self = e.target;
            let respuesta = {
                status: self.status,
                response: JSON.parse(self.response)
            };
            resolver(respuesta);
        })
        xhr.addEventListener("error", e => {
            let self = e.target;
    
            rechazar(self);
        });
    
        xhr.send(obj);
    });

}


function ClearStudentView() {
    $("#viewStudent-id").textContent = "";
    $("#viewStudent-id2").textContent = "";
    $("#viewStudent-SIGECOD").textContent = "";
    $("#viewStudent-SIGECOD2").textContent = "";
    $("#viewStudent-SIGE2").textContent = "";
    $("#viewStudent-firstName").textContent = "";
    $("#viewStudent-firstName2").textContent = "";
    $("#viewStudent-lastName").textContent = "";
    $("#viewStudent-lastName2").textContent = "";
    $("#viewStudent-ci").textContent ="";
    $("#viewStudent-ci2").textContent = "";
    $("#viewStudent-email").textContent = "";
    $("#viewStudent-email2").textContent = "";
    $("#viewStudent-phoneNumber").textContent = "";
    $("#viewStudent-phoneNumber2").textContent = "";
    $("#viewStudent-gender").textContent = "";
    $("#viewStudent-gender2").textContent = "";
    $("#viewStudent-dateOfBirth").textContent = "";
    $("#viewStudent-dateOfBirth2").textContent = "";
    $("#viewStudent-healthProblemE").textContent = "";
    $("#viewStudent-healthProblemE2").textContent = "";
    $("#viewStudent-address").textContent = "";
    $("#viewStudent-address2").textContent = "";
    $("#viewStudent-motherFirstName").textContent = "";
    $("#viewStudent-motherFirstName2").textContent = "";
    $("#viewStudent-motherLastName").textContent = "";
    $("#viewStudent-motherLastName2").textContent = "";
    $("#viewStudent-motherPhoneNumber").textContent = "";
    $("#viewStudent-motherPhoneNumber2").textContent = "";
    $("#viewStudent-motherCI").textContent = "";
    $("#viewStudent-motherCI2").textContent = "";
    $("#viewStudent-fatherFirstName").textContent = "";
    $("#viewStudent-fatherFirstName2").textContent = "";
    $("#viewStudent-fatherLastName").textContent = "";
    $("#viewStudent-fatherLastName2").textContent = "";
    $("#viewStudent-fatherPhoneNumber").textContent = "";
    $("#viewStudent-fatherPhoneNumber2").textContent = "";
    $("#viewStudent-fatherCI").textContent = "";
    $("#viewStudent-fatherCI2").textContent = "";
    $("#viewStudent-representativeFirstName").textContent = "";
    $("#viewStudent-representativeFirstName2").textContent = "";
    $("#viewStudent-representativeLastName").textContent = "";
    $("#viewStudent-representativeLastName2").textContent = "";
    $("#viewStudent-representativePhoneNumber").textContent = "";
    $("#viewStudent-representativePhoneNumber2").textContent = "";
    $("#viewStudent-representativeCI").textContent = "";
    $("#viewStudent-representativeCI2").textContent = "";
    $("#viewStudent-representativeRelationship").textContent = "";
    $("#viewStudent-representativeRelationship2").textContent = "";
    $("#viewStudent-representativeAddress").textContent = "";
    $("#viewStudent-representativeAddress2").textContent = "";
    $("#viewStudent-scholarship").textContent = "";
    $("#viewStudent-scholarship2").textContent = "";
    $("#viewStudent-canaima").textContent = "";
    $("#viewStudent-regular2").textContent = "";
    $("#viewStudent-repeat2").textContent = "";
    $("#viewStudent-asigPend2").textContent = "";
    $("#viewStudent-year22").textContent = "";
    $("#viewStudent-asigRepeatE2").textContent = "";
    $("#viewStudent-asigPendE2").textContent = "";
    $("#viewStudent-dateOfInscription2").textContent = "";
    $("#viewStudent-year").textContent = "";
    $("#viewStudent-age").textContent = "";
    $("#viewStudent-age2").textContent = "";
    $("#viewStudent-size").textContent = "";
    $("#viewStudent-size2").textContent = "";
    $("#viewStudent-weight").textContent = "";
    $("#viewStudent-weight2").textContent = "";
}