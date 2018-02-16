//Profile
(function(){
    self.Profile = function(){
        this.UserID = 0;
        this.Username = "";
        this.Email = "";
        this.FirstName = "";
        this.LastName = "";
        this.Address = "";
        this.PhoneNumber = "";
        this.KindID = 0;
    }

    self.Profile.prototype = {
        getProfile : function(token){
            let query = `query {
                            me {
                                id,
                                username,
                                email,
                                firstName,
                                lastName,
                                address,
                                phoneNumber,
                                kindID
                            }
                         }`;
            
            let self = this;

            requestAjaxToken("POST", "/graphql", token, query)
             .then( r => {
                 self.UserID = r.response.data.me.id;
                 self.Username = r.response.data.me.username;
                 this.Email = r.response.data.me.email;
                 this.FirstName = r.response.data.me.firstName;
                 this.LastName = r.response.data.me.lastName;
                 this.Address = r.response.data.me.address;
                 this.PhoneNumber = r.response.data.me.phoneNumber;
                 this.KindID = r.response.data.me.kindID;
             })
             .catch( error => {
                 console.log("Hubo un error al settear la data del perfil "+error);
             })
        }
    }

})();

(function(){
    self.ProfileView = function(profile){
        this.Container;
        this.UserID = profile.UserID;
        this.Username = profile.Username;
        this.Email = profile.Email;
        this.FirstName = profile.FirstName;
        this.LastName = profile.LastName;
        this.Address = profile.Address;
        this.PhoneNumber = profile.PhoneNumber;
        this.KindID = profile.KindID;
    }
})();

var miToken = sessionStorage.getItem("token");

var miProfile = new Profile();
miProfile.getProfile(miToken);
console.log(miProfile);
console.log(miProfile.Username);
console.log("--------------");

var miView = new ProfileView(miProfile);
console.log(miView);

function GetProfile() {

    let userIDContainer = $("#userID-container"),
        usernameContainer = $("#username-container"),
        emailContainer = $("#email-container"),
        firstNameContainer = $("#firstName-container"),
        lastNameContainer = $("#lastName-container"),
        addressContainer = $("#address-container"),
        phoneNumberContainer = $("#phoneNumber-container"),
        containerBtns = $("#containerBtns");

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
}