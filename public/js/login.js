sessionStorage.removeItem("email");
sessionStorage.removeItem("firstName");
sessionStorage.removeItem("lastName");
sessionStorage.removeItem("address");
sessionStorage.removeItem("phoneNumber");

var form = $("#form");
var formSubmit = $("#btn-submit");

$("#username").value = sessionStorage.getItem("username");

form.addEventListener("submit", function() {
    formSubmit.disabled = true;
    formSubmit.value = "Enviando...";
    
    var username = $("#username").value;
   
    sessionStorage.setItem("username", username);
    
});