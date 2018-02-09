let formLogin = $("#form-login"),
    username = $("#username"),
    passw = $("#password"),
    btnSubmit = $("#btn-submit"),
    formMessage = $("#form-message");

username.value = sessionStorage.getItem("username");

formLogin.addEventListener("submit", e => {
    e.preventDefault();
    btnSubmit.disabled = true;
    btnSubmit.value = "Enviando...";

    let obj = {
        username: username.value,
        password: passw.value
    }

    requestAjax(formLogin.method, formLogin.action, JSON.stringify(obj))
     .then( respuesta => {

         if (respuesta.response.code === 100) {

             formMessage.textContent = "Logueo exitoso!";
             sessionStorage.setItem("token", respuesta.response.token);
             console.log(respuesta.response);

         } else if (respuesta.response.code === 203) {

             formMessage.textContent = "Usuario o contraseña invalidos";
             console.log(respuesta.response);

         } else {

             console.log(respuesta.response.code);

         }
     })
     .catch( error => {
         console.log(error)
     });
    
    sessionStorage.setItem("username", username.value);

    btnSubmit.disabled = false;
    btnSubmit.value = "Login";
});