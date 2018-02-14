//Se crea una forma mas facil de seleccionar elementos semejante a JQuery.
function $(selector) {
    return document.querySelector(selector);
}

//Se declaran las variables con los elementos.
let formulario = $("#form"),
    username = $("#username"),
    password = $("#password"),
    boton = $("#btn"),
    caja = $("#caja");

//Se crea una instancia del objeto XMLHttpRequest, Esta es una instancia de nuestra peticion AJAX.
let xhr = new XMLHttpRequest();

//Con el metodo open declaramos el metodo a utilizar, la URL a donde va la peticion, Y si es asincrona.
xhr.open(formulario.method, formulario.action, true);
//Con setRequestHeader le declaramos a la cabecera unas claves basicas por convencion.
xhr.setRequestHeader("Content-Type", "application/json");
xhr.setRequestHeader("cache-control", "no-cache");

//Le agregamos el evento load a nuestra instancia AJAX, El codigo se ejecutara al cargar la peticion.
xhr.addEventListener("load", function(e) {
    //El parametro "e" nos servira como la palabra "this".

    //Parseamos el texto del atributo response a JSON.
    let informacion = JSON.parse(e.currentTarget.response);
    console.log(e.currentTarget);

    console.log(informacion.token);

    //Le agremos a la caja el contenido del atributo message del objeto informacion.
    caja.textContent = informacion.message;
});

//Le agregamos el evento click a este boton.
boton.addEventListener("click", function(e){
    //Desactivamos el comportamiento por defecto para evitar que se envie el formulario de manera tradicional.
    e.preventDefault();
    let info = {
        username: username.value,
        password: password.value
    }

    //Enviamos los datos como string.
    xhr.send(JSON.stringify(info));
})


