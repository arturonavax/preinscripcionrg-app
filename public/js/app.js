if (!sessionStorage.getItem("token")) {
    window.location.replace("/login");
}

var data = "query {me{id,username,email,firstName,lastName,address,phoneNumber,kindID}}"; 

requestAjaxToken("POST", "/graphql", data)
 .then( respuesta => {
    dataGraph = respuesta.response.data
    console.log(respuesta);

    userID.textContent = dataGraph.me.id;
    username.textContent = dataGraph.me.username;
    email.textContent = dataGraph.me.email;
    firstName.textContent = dataGraph.me.firstName;
    lastName.textContent = dataGraph.me.lastName;
    address.textContent = dataGraph.me.address;
    phoneNumber.textContent = dataGraph.me.phoneNumber;
    kindUser.textContent = dataGraph.me.kindID;

    titleUser.textContent = dataGraph.me.username;
 })
 .catch( error => {
    console.log(error);
 });

let userID = $("#userID"),
    username = $("#username"),
    email = $("#email"),
    firstName = $("#firstName"),
    lastName = $("#lastName"),
    address = $("#address"),
    phoneNumber = $("#phoneNumber"),
    kindUser = $("#kindUser"),
    titleUser = $("#titleUser");

var btnLogout = $("#btn-logout");

btnLogout.addEventListener("click", function() {
    sessionStorage.removeItem("token");

    window.location.replace("/login");
});