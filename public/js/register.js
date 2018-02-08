var form = $("#form");
var formSubmit = $("#btn-submit");

$("#username").value = sessionStorage.getItem("username");
$("#email").value = sessionStorage.getItem("email");
$("#firstName").value = sessionStorage.getItem("firstName");
$("#lastName").value = sessionStorage.getItem("lastName");
$("#address").value = sessionStorage.getItem("address");
$("#phoneNumber").value = sessionStorage.getItem("phoneNumber");

form.addEventListener("submit", function() {
    formSubmit.disabled = true;
    formSubmit.value = "Enviando...";

    var username = $("#username").value;
    var email = $("#email").value;
    var firstName = $("#firstName").value;
    var lastName = $("#lastName").value;
    var address = $("#address").value;
    var phoneNumber = $("#phoneNumber").value;
    
    sessionStorage.setItem("username", username);
    sessionStorage.setItem("email", email);
    sessionStorage.setItem("firstName", firstName);
    sessionStorage.setItem("lastName", lastName);
    sessionStorage.setItem("address", address);
    sessionStorage.setItem("phoneNumber", phoneNumber);
});