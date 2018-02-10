if (sessionStorage.getItem("token")) {
    window.location.replace("/app");
}

let formRegister = $("#form-register"),
    username = $("#username"),
    email = $("#email"),
    firstName = $("#firstName"),
    lastName = $("#lastName"),
    address = $("#address"),
    phoneNumber = $("#phoneNumber"),
    passw = $("#password"),
    confirmPassw = $("#confirmPassword"),
    btnSubmit = $("#btn-submit"),
    formMessage = $("#form-message");

$("#username").value = sessionStorage.getItem("username");
$("#email").value = sessionStorage.getItem("email");
$("#firstName").value = sessionStorage.getItem("firstName");
$("#lastName").value = sessionStorage.getItem("lastName");
$("#address").value = sessionStorage.getItem("address");
$("#phoneNumber").value = sessionStorage.getItem("phoneNumber");

formRegister.addEventListener("submit", e => {
    e.preventDefault();
    btnSubmit.disabled = true;
    btnSubmit.value = "Enviando...";

    let obj = {
        username: username.value,
        email: email.value,
        firstName: firstName.value,
        lastName: lastName.value,
        address: address.value,
        phoneNumber: phoneNumber.value,
        password: passw.value,
        confirmPassword: confirmPassw.value
    }

    requestAjax(formRegister.method, formRegister.action, JSON.stringify(obj))
     .then( respuesta => {

         if (respuesta.response.code === 201) {

             formMessage.textContent = "Registro exitoso";
             console.log(respuesta.response);

         } else if (respuesta.response.code === 400) {

             formMessage.textContent = "El Usuario ya existe";
             console.log(respuesta.response);

         } else if (respuesta.response.code === 406) {

             formMessage.textContent = "Las contraseñas no coinciden";
             console.log(respuesta.response);

         } else {

             console.log(respuesta.response.code);

         }
     })
     .catch( error => {
         console.log(error)
     });
    
    sessionStorage.setItem("username", username.value);
    sessionStorage.setItem("email", email.value);
    sessionStorage.setItem("firstName", firstName.value);
    sessionStorage.setItem("lastName", lastName.value);
    sessionStorage.setItem("address", address.value);
    sessionStorage.setItem("phoneNumber", phoneNumber.value);

    btnSubmit.disabled = false;
    btnSubmit.value = "Registrar";
});