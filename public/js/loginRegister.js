var btnRegister = $("#navBtn-register");
var btnLogin = $("#navBtn-login");
var loginRegister = $("#login-register");

btnRegister.addEventListener("click", function() {
    loginRegister.classList.add("register--open");
});

btnLogin.addEventListener("click", function(){
    loginRegister.classList.remove("register--open");
});

if (sessionStorage.getItem("token")) {
    window.location.replace("/app");
}

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

formLogin.addEventListener("submit", e => {
    e.preventDefault();
    btnSubmitLogin.disabled = true;
    btnSubmitLogin.value = "Enviando...";
    
    let obj = {
        username: usernameLogin.value,
        password: passwLogin.value
    }
    
    requestAjax(formLogin.method, formLogin.action, JSON.stringify(obj))
    .then( respuesta => {
        
        if (respuesta.response.code === 100) {
            
            formMessageLogin.textContent = "Logueo exitoso!";
            sessionStorage.setItem("token", respuesta.response.token);

            if (sessionStorage.getItem("token")) {
                var data = "query {me{firstName,lastName}}";
            
                requestAjaxToken("POST", "/graphql", data)
                 .then( respuesta => {
                     dataGraph = respuesta.response.data
                     console.log(respuesta);
                     document.querySelector("#titleUser").textContent = "Bienvenido/a "+dataGraph.me.firstName+" "+dataGraph.me.lastName+"!";
                 })
                 .catch( error => {
                     console.log(error);
                 });
            }
            
            usernameLogin.value = sessionStorage.getItem("username");
            console.log(respuesta.response);

            if (sessionStorage.getItem("token")) {
                window.location.replace("/app");
            }

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

    btnSubmitLogin.disabled = false;
    btnSubmitLogin.value = "Login";
});

formRegister.addEventListener("submit", e => {
    e.preventDefault();
    btnSubmitRegister.disabled = true;
    btnSubmitRegister.value = "Enviando...";

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

    btnSubmitRegister.disabled = false;
    btnSubmitRegister.value = "Registrar";
});