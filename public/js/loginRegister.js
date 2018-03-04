/*if (sessionStorage.getItem("token")) {
    window.location.replace("/app");
}*/

var btnRegister = $("#navBtn-register");
var btnLogin = $("#navBtn-login");
var loginRegister = $("#login-register");

btnRegister.addEventListener("click", function() {
    $("#section").style.minHeight = "735px";
    loginRegister.classList.add("register--open");
});

btnLogin.addEventListener("click", function(){
    $("#section").style.minHeight = "400px";
    loginRegister.classList.remove("register--open");
});

let formLogin = $("#form-login"),
    formRegister = $("#form-register"),
    usernameLogin = $("#username-login"),
    usernameRegister = $("#username-register"),

    email = $("#email"),
    firstName = $("#firstName"),
    lastName = $("#lastName"),
    address = $("#address"),
    phoneNumber = $("#phoneNumber"),
    passwLogin = $("#password-login"),
    passwRegister = $("#password-register"),
    confirmPassw = $("#confirmPassword"),
    btnSubmitLogin = $("#btn-submit-login"),
    btnSubmitRegister = $("#btn-submit-register"),
    formMessageLogin = $("#form-message-login"),
    formMessageRegister = $("#form-message-register");


usernameLogin.value = sessionStorage.getItem("username");
usernameRegister.value = sessionStorage.getItem("username");
email.value = sessionStorage.getItem("email");
firstName.value = sessionStorage.getItem("firstName");
lastName.value = sessionStorage.getItem("lastName");
address.value = sessionStorage.getItem("address");
phoneNumber.value = sessionStorage.getItem("phoneNumber");

btnSubmitLogin.addEventListener("click", function(){
    btnSubmitLogin.value = "Entrando...";
});

formLogin.addEventListener("submit", e => {
    e.preventDefault();
    btnSubmitLogin.disabled = true;
    
    let obj = {
        username: usernameLogin.value,
        password: passwLogin.value
    }
    
    requestAjax(formLogin.method, formLogin.action, JSON.stringify(obj))
    .then( respuesta => {
        
        btnSubmitLogin.disabled = false;
        btnSubmitLogin.value = "Login";
        if (respuesta.response.code === 100) {
            
            formMessageLogin.textContent = "Logueo exitoso!";
            sessionStorage.setItem("token", respuesta.response.token);
            
            usernameLogin.value = sessionStorage.getItem("username");
            console.log(respuesta.response);

            sessionStorage.removeItem("username");
            sessionStorage.removeItem("email");
            sessionStorage.removeItem("firstName");
            sessionStorage.removeItem("lastName");
            sessionStorage.removeItem("address");
            sessionStorage.removeItem("phoneNumber");

            window.location.replace("/app");

         } else if (respuesta.response.code === 203) {

             formMessageLogin.textContent = "Usuario o contraseña invalidos";
             console.log(respuesta.response);

         } else {

             console.log(respuesta.response.code);

         }
     })
     .catch( error => {
         console.log(error)
     });
    
    sessionStorage.setItem("username", usernameLogin.value);

});

btnSubmitRegister.addEventListener("click", function() {
    btnSubmitRegister.value = "Registrando...";
});

formRegister.addEventListener("submit", e => {
    e.preventDefault();
    btnSubmitRegister.disabled = true;

    let obj = {
        username: usernameRegister.value,
        email: email.value,
        firstName: firstName.value,
        lastName: lastName.value,
        address: address.value,
        phoneNumber: phoneNumber.value,
        password: passwRegister.value,
        confirmPassword: confirmPassw.value
    }

    requestAjax(formRegister.method, formRegister.action, JSON.stringify(obj))
     .then( respuesta => {
        btnSubmitRegister.disabled = false;
        btnSubmitRegister.value = "Registrar";
         if (respuesta.response.code === 201) {

             formMessageRegister.textContent = "Registro exitoso";
             console.log(respuesta.response);

         } else if (respuesta.response.code === 400) {

             formMessageRegister.textContent = "El Usuario ya existe";
             console.log(respuesta.response);

         } else if (respuesta.response.code === 406) {

             formMessageRegister.textContent = "Las contraseñas no coinciden";
             console.log(respuesta.response);

         } else {

             console.log(respuesta.response.code);

         }
     })
     .catch( error => {
         console.log(error)
     });
    
    sessionStorage.setItem("username", usernameRegister.value);
    sessionStorage.setItem("email", email.value);
    sessionStorage.setItem("firstName", firstName.value);
    sessionStorage.setItem("lastName", lastName.value);
    sessionStorage.setItem("address", address.value);
    sessionStorage.setItem("phoneNumber", phoneNumber.value);

});