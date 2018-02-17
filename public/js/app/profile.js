//Profile
(function(){
    self.Profile = function(profileView){
        this.Token = "";
        this.UserID = 0;
        this.Username = "";
        this.Email = "";
        this.FirstName = "";
        this.LastName = "";
        this.Address = "";
        this.PhoneNumber = "";
        this.KindID = 0;

        this.ProfileView = profileView;
    }

    self.Profile.prototype = {
        set token(token){
            this.Token = token;
            let query = `

            query {
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
            }

            `;

            let xhr = new XMLHttpRequest();

            xhr.open("POST", "/graphql", false);
            xhr.setRequestHeader("Content-Type", "application/graphql");
            xhr.setRequestHeader("Authorization", token);
            xhr.withCredentials = true;
            xhr.setRequestHeader("cache-control", "no-cache");
            xhr.addEventListener("load", e => {
                let self = e.target;
                let r = {
                    status: self.status,
                    response: JSON.parse(self.response)
                };

                this.Username = r.response.data.me.username;
                this.UserID = r.response.data.me.id;
                this.Email = r.response.data.me.email;
                this.FirstName = r.response.data.me.firstName;
                this.LastName = r.response.data.me.lastName;
                this.Address = r.response.data.me.address;
                this.PhoneNumber = r.response.data.me.phoneNumber;
                this.KindID = r.response.data.me.kindID;
            });

            xhr.addEventListener("error", e => {
                let self = e.target;
        
                console.log(self);
            });

            xhr.send(query);
        },
        render: function() {
            let pv = this.ProfileView;
            let self = this;

            pv.BtnEditProfile.classList.add("itemBtn");
            pv.BtnEditProfile.classList.add("itemBtn--edit");
            pv.BtnEditProfile.addEventListener("click", function(){
               self.renderEdit(); 
            });

            pv.BtnEditProfileContainer.appendChild(pv.BtnEditProfile);
            pv.UserIDDataContainer.textContent = this.UserID;
            pv.UsernameDataContainer.textContent = this.Username;
            pv.EmailDataContainer.textContent = this.Email;
            pv.FirstNameDataContainer.textContent = this.FirstName;
            pv.LastNameDataContainer.textContent = this.LastName;
            pv.AddressDataContainer.textContent = this.Address;
            pv.PhoneNumberDataContainer.textContent = this.PhoneNumber;
            pv.KindIDDataContainer.textContent = this.KindID;
        },

        renderEdit: function() {
            let pv = this.ProfileView;
            let self = this;
            pv.BtnCancelEditProfile.textContent = "Cancelar";
            pv.BtnCancelEditProfile.classList.add("item__submit");
            pv.BtnCancelEditProfile.classList.add("item__submit--edit");
            pv.BtnCancelEditProfile.addEventListener("click", function(){
                self.cancelEdit();
            });

            pv.BtnSaveEditProfile.textContent = "Guardar";
            pv.BtnSaveEditProfile.classList.add("item__submit");
            pv.BtnSaveEditProfile.classList.add("item__submit--edit");
            pv.BtnSaveEditProfile.addEventListener("click", function(){
                pv.BtnSaveEditProfile.textContent = "Guardando..";
                self.saveEdit();
            });

            pv.EmailDataEdit.value = pv.EmailDataContainer.textContent;
            pv.EmailDataEdit.id = pv.EmailDataContainer.id;
            pv.EmailDataEdit.classList.add("item__input");
            pv.EmailDataEdit.classList.add("item__input--edit");

            pv.FirstNameDataEdit.value = pv.FirstNameDataContainer.textContent;
            pv.FirstNameDataEdit.id = pv.FirstNameDataContainer.id;
            pv.FirstNameDataEdit.classList.add("item__input");
            pv.FirstNameDataEdit.classList.add("item__input--edit");

            pv.LastNameDataEdit.value = pv.LastNameDataContainer.textContent;
            pv.LastNameDataEdit.id = pv.LastNameDataContainer.id;
            pv.LastNameDataEdit.classList.add("item__input");
            pv.LastNameDataEdit.classList.add("item__input--edit");

            pv.AddressDataEdit.value = pv.AddressDataContainer.textContent;
            pv.AddressDataEdit.id = pv.AddressDataContainer.id;
            pv.AddressDataEdit.classList.add("item__input");
            pv.AddressDataEdit.classList.add("item__input--edit");

            pv.PhoneNumberDataEdit.value = pv.PhoneNumberDataContainer.textContent;
            pv.PhoneNumberDataEdit.id = pv.PhoneNumberDataContainer.id;
            pv.PhoneNumberDataEdit.classList.add("item__input");
            pv.PhoneNumberDataEdit.classList.add("item__input--edit");

            pv.EmailProfileContainer.replaceChild(pv.EmailDataEdit, pv.EmailDataContainer);
            pv.FirstNameProfileContainer.replaceChild(pv.FirstNameDataEdit, pv.FirstNameDataContainer);
            pv.LastNameProfileContainer.replaceChild(pv.LastNameDataEdit, pv.LastNameDataContainer);
            pv.AddressProfileContainer.replaceChild(pv.AddressDataEdit, pv.AddressDataContainer);
            pv.PhoneNumberProfileContainer.replaceChild(pv.PhoneNumberDataEdit, pv.PhoneNumberDataContainer);

            pv.BtnContainerEditProfile.appendChild(pv.BtnCancelEditProfile);
            pv.BtnContainerEditProfile.appendChild(pv.BtnSaveEditProfile);
        },

        cancelEdit: function(){
            let pv = this.ProfileView;
            pv.EmailProfileContainer.replaceChild(pv.EmailDataContainer, pv.EmailDataEdit);
            pv.FirstNameProfileContainer.replaceChild(pv.FirstNameDataContainer, pv.FirstNameDataEdit);
            pv.LastNameProfileContainer.replaceChild(pv.LastNameDataContainer, pv.LastNameDataEdit);
            pv.AddressProfileContainer.replaceChild(pv.AddressDataContainer, pv.AddressDataEdit);
            pv.PhoneNumberProfileContainer.replaceChild(pv.PhoneNumberDataContainer, pv.PhoneNumberDataEdit);

            pv.BtnContainerEditProfile.removeChild(pv.BtnCancelEditProfile);
            pv.BtnContainerEditProfile.removeChild(pv.BtnSaveEditProfile);
        },

        saveEdit: function(){
            let pv = this.ProfileView;
            let email = pv.EmailDataEdit.value,
                firstName = pv.FirstNameDataEdit.value,
                lastName = pv.LastNameDataEdit.value,
                address = pv.AddressDataEdit.value,
                phoneNumber = pv.PhoneNumberDataEdit.value;
            
            let mutation = `
            
            mutation {
                meU(
                    email: "${email}",
                    firstName: "${firstName}",
                    lastName: "${lastName}",
                    address: "${address}",
                    phoneNumber: "${phoneNumber}"
                ){id,message,code}
            }

            `

            requestAjaxToken("POST", "/graphql", this.Token, mutation, true)
             .then( r => {
                pv.BtnSaveEditProfile.textContent = "Guardar";
                if (r.response.data.meU.code === 100) {
                    pv.EmailDataContainer.textContent = email;
                    pv.FirstNameDataContainer.textContent = firstName;
                    pv.LastNameDataContainer.textContent = lastName;
                    pv.AddressDataContainer.textContent = address;
                    pv.PhoneNumberDataContainer.textContent = phoneNumber;

                    pv.ProfileNotification.textContent = "Usuario actualizado con exito."
                } else if (r.response.data.meU.code === 500) {
                    pv.ProfileNotification.textContent = "Ocurrio un error."

                }
                console.log(r);
             });


            this.cancelEdit();
        }
    }

})();

(function(){
    self.ProfileView = function(){
        this.ContainerProfile = $("#container-profile");

        this.TopProfileContainer = $("#top-profileContainer");
        this.BtnEditProfile = document.createElement("button");
        this.BtnProfileContainer = $("#btnContainer-topProfile");
        this.BtnEditProfileContainer = $("#btnContainer-editProfile");
        this.ProfileNotification = $("#profileNotification");


        this.UserIDProfileContainer = $("#userID-profileContainer");
        this.UsernameProfileContainer = $("#username-profileContainer");
        this.EmailProfileContainer = $("#email-profileContainer");
        this.FirstNameProfileContainer = $("#firstName-profileContainer");
        this.LastNameProfileContainer = $("#lastName-profileContainer");
        this.AddressProfileContainer = $("#address-profileContainer");
        this.PhoneNumberProfileContainer = $("#phoneNumber-profileContainer");
        this.KindIDProfileContainer = $("#kindID-profileContainer");

        this.UserIDDataContainer = $("#userIDData-container");
        this.UsernameDataContainer = $("#usernameData-container");
        this.EmailDataContainer = $("#emailData-container");
        this.FirstNameDataContainer = $("#firstNameData-container");
        this.LastNameDataContainer = $("#lastNameData-container");
        this.AddressDataContainer = $("#addressData-container");
        this.PhoneNumberDataContainer = $("#phoneNumberData-container");
        this.KindIDDataContainer = $("#kindIDData-container");

        this.EmailDataEdit = document.createElement("input");
        this.FirstNameDataEdit = document.createElement("input");
        this.LastNameDataEdit = document.createElement("input");
        this.AddressDataEdit = document.createElement("input");
        this.PhoneNumberDataEdit = document.createElement("input");
        this.KindIDDataEdit = document.createElement("input");

        this.BtnContainerEditProfile = $("#btnContainer-profileContainer");
        this.BtnCancelEditProfile = document.createElement("button"); 
        this.BtnSaveEditProfile = document.createElement("button"); 
    }
})();

/* -------------------------------------------------------------------------- */

let tokenUser = sessionStorage.getItem("token");

let btnEditProfile = $("#btn-editProfile");

let profileView = new ProfileView(),
    profileUser = new Profile(profileView);

profileUser.token = tokenUser;
profileUser.render();
