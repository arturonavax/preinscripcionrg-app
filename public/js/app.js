/*if (!sessionStorage.getItem("token")) {
    window.location.replace("/user");
}*/

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

let userIDContainer = $("#userID-container"),
    usernameContainer= $("#username-container"),
    emailContainer= $("#email-container"),
    firstNameContainer= $("#firstName-container"),
    lastNameContainer= $("#lastName-container"),
    addressContainer= $("#address-container"),
    phoneNumberContainer = $("#phoneNumber-container"),
    containerBtns = $("#containerBtns");

let userID = $("#userID"),
    username = $("#username"),
    email = $("#email"),
    firstName = $("#firstName"),
    lastName = $("#lastName"),
    address = $("#address"),
    phoneNumber = $("#phoneNumber"),
    titleUser = $("#titleUser");

var btnLogout = $("#btn-logout"),
    btnEdit = $("#edit-btn");

btnLogout.addEventListener("click", function() {
    sessionStorage.removeItem("token");

    window.location.replace("/user");
});

btnEdit.addEventListener("click", function() {
    let inputUsername = document.createElement("input");
    inputUsername.id = username.id;
    inputUsername.className = "item__input";
    inputUsername.classList.add("item__input--edit");
    inputUsername.value = username.textContent;

    let inputEmail = document.createElement("input");
    inputEmail.id = email.id;
    inputEmail.className = "item__input";
    inputEmail.classList.add("item__input--edit");
    inputEmail.value = email.textContent;

    let inputFirstName = document.createElement("input");
    inputFirstName.id = firstName.id;
    inputFirstName.className = "item__input";
    inputFirstName.classList.add("item__input--edit");
    inputFirstName.value = firstName.textContent;

    let inputLastName = document.createElement("input");
    inputLastName.id = lastName.id;
    inputLastName.className = "item__input";
    inputLastName.classList.add("item__input--edit");
    inputLastName.value = lastName.textContent;

    let inputAddress = document.createElement("input");
    inputAddress.id = address.id;
    inputAddress.className = "item__input";
    inputAddress.classList.add("item__input--edit");
    inputAddress.value = address.textContent;

    let inputPhoneNumber = document.createElement("input");
    inputPhoneNumber.id = phoneNumber.id;
    inputPhoneNumber.className = "item__input";
    inputPhoneNumber.classList.add("item__input--edit");
    inputPhoneNumber.value = phoneNumber.textContent;

    let btnSave = document.createElement("button");
    btnSave.id = "btn-save";
    btnSave.textContent = "Guardar";
    btnSave.className = "item__submit";

    usernameContainer.replaceChild(inputUsername,username);
    emailContainer.replaceChild(inputEmail, email);
    firstNameContainer.replaceChild(inputFirstName, firstName);
    lastNameContainer.replaceChild(inputLastName, lastName);
    addressContainer.replaceChild(inputAddress, address);
    phoneNumberContainer.replaceChild(inputPhoneNumber, phoneNumber);

    containerBtns.appendChild(btnSave);
});