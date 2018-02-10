if (sessionStorage.getItem("token")) {
    window.location.replace("http://localhost:8080/app");
}

let formLogin = $("#form-login"),
username = $("#username"),
passw = $("#password"),
btnSubmit = $("#btn-submit"),
formMessage = $("#form-message");


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
            
            username.value = sessionStorage.getItem("username");
            console.log(respuesta.response);

            if (sessionStorage.getItem("token")) {
                window.location.replace("/app");
            }

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