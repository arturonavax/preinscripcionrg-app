//Profile
(function(){
    self.Profile = function(){
        this.Token = "";
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
        }
    }

})();

(function(){
    self.ProfileView = function(profile){
        this.Profile = profile;
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

    self.ProfileView.prototype = {
        render: function() {
            this.BtnEditProfile.classList.add("itemBtn");
            this.BtnEditProfile.classList.add("itemBtn--edit");

            this.BtnEditProfileContainer.appendChild(this.BtnEditProfile);
            this.UserIDDataContainer.textContent = this.Profile.UserID;
            this.UsernameDataContainer.textContent = this.Profile.Username;
            this.EmailDataContainer.textContent = this.Profile.Email;
            this.FirstNameDataContainer.textContent = this.Profile.FirstName;
            this.LastNameDataContainer.textContent = this.Profile.LastName;
            this.AddressDataContainer.textContent = this.Profile.Address;
            this.PhoneNumberDataContainer.textContent = this.Profile.PhoneNumber;
            this.KindIDDataContainer.textContent = this.Profile.KindID;
        },

        renderEdit: function() {
            this.BtnCancelEditProfile.textContent = "Cancelar";
            this.BtnCancelEditProfile.classList.add("item__submit");
            this.BtnCancelEditProfile.classList.add("item__submit--edit");
            this.BtnCancelEditProfile.addEventListener("click", function(){
                this.cancelEdit();
            });

            this.BtnSaveEditProfile.textContent = "Guardar";
            this.BtnSaveEditProfile.classList.add("item__submit");
            this.BtnSaveEditProfile.classList.add("item__submit--edit");
            this.BtnSaveEditProfile.addEventListener("click", function(){
                this.BtnSaveEditProfile.textContent = "Guardando..";
                this.saveEdit();
            });

            this.EmailDataEdit.value = this.EmailDataContainer.textContent;
            this.EmailDataEdit.id = this.EmailDataContainer.id;
            this.EmailDataEdit.classList.add("item__input");
            this.EmailDataEdit.classList.add("item__input--edit");

            this.FirstNameDataEdit.value = this.FirstNameDataContainer.textContent;
            this.FirstNameDataEdit.id = this.FirstNameDataContainer.id;
            this.FirstNameDataEdit.classList.add("item__input");
            this.FirstNameDataEdit.classList.add("item__input--edit");

            this.LastNameDataEdit.value = this.LastNameDataContainer.textContent;
            this.LastNameDataEdit.id = this.LastNameDataContainer.id;
            this.LastNameDataEdit.classList.add("item__input");
            this.LastNameDataEdit.classList.add("item__input--edit");

            this.AddressDataEdit.value = this.AddressDataContainer.textContent;
            this.AddressDataEdit.id = this.AddressDataContainer.id;
            this.AddressDataEdit.classList.add("item__input");
            this.AddressDataEdit.classList.add("item__input--edit");

            this.PhoneNumberDataEdit.value = this.PhoneNumberDataContainer.textContent;
            this.PhoneNumberDataEdit.id = this.PhoneNumberDataContainer.id;
            this.PhoneNumberDataEdit.classList.add("item__input");
            this.PhoneNumberDataEdit.classList.add("item__input--edit");

            this.EmailProfileContainer.replaceChild(this.EmailDataEdit, this.EmailDataContainer);
            this.FirstNameProfileContainer.replaceChild(this.FirstNameDataEdit, this.FirstNameDataContainer);
            this.LastNameProfileContainer.replaceChild(this.LastNameDataEdit, this.LastNameDataContainer);
            this.AddressProfileContainer.replaceChild(this.AddressDataEdit, this.AddressDataContainer);
            this.PhoneNumberProfileContainer.replaceChild(this.PhoneNumberDataEdit, this.PhoneNumberDataContainer);

            this.BtnContainerEditProfile.appendChild(this.BtnCancelEditProfile);
            this.BtnContainerEditProfile.appendChild(this.BtnSaveEditProfile);
        },

        cancelEdit: function(){
            this.EmailProfileContainer.replaceChild(this.EmailDataContainer, this.EmailDataEdit);
            this.FirstNameProfileContainer.replaceChild(this.FirstNameDataContainer, this.FirstNameDataEdit);
            this.LastNameProfileContainer.replaceChild(this.LastNameDataContainer, this.LastNameDataEdit);
            this.AddressProfileContainer.replaceChild(this.AddressDataContainer, this.AddressDataEdit);
            this.PhoneNumberProfileContainer.replaceChild(this.PhoneNumberDataContainer, this.PhoneNumberDataEdit);

            this.BtnContainerEditProfile.removeChild(this.BtnCancelEditProfile);
            this.BtnContainerEditProfile.removeChild(this.BtnSaveEditProfile);
        },

        saveEdit: function(){
            let email = this.EmailDataEdit.value,
                firstName = this.FirstNameDataEdit.value,
                lastName = this.LastNameDataEdit.value,
                address = this.AddressDataEdit.value,
                phoneNumber = this.PhoneNumberDataEdit.value;
            
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

            requestAjaxToken("POST", "/graphql", this.Profile.Token, mutation, true)
             .then( r => {
                this.BtnSaveEditProfile.textContent = "Guardar";
                if (r.response.data.meU.code === 100) {
                    this.EmailDataContainer.textContent = email;
                    this.FirstNameDataContainer.textContent = firstName;
                    this.LastNameDataContainer.textContent = lastName;
                    this.AddressDataContainer.textContent = address;
                    this.PhoneNumberDataContainer.textContent = phoneNumber;

                    this.ProfileNotification.textContent = "Usuario actualizado con exito."
                } else if (r.response.data.meU.code === 500) {
                    this.ProfileNotification.textContent = "Ocurrio un error."

                }
                console.log(r);
             });


            this.cancelEdit();
        }
    }

})();

/* -------------------------------------------------------------------------- */
