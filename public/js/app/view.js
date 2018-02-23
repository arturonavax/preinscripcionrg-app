(function(){
    self.UI = function(profile){
        this.ProfileView = profile;

    }

    self.UI.prototype = {
        render: function(){
            this.ProfileView.render();
        }
    }
})();

let tokenUser = sessionStorage.getItem("token");

let profileUser = new Profile();
profileUser.token = tokenUser;
let profileView = new ProfileView(profileUser);

let UserInterface = new UI(profileView);

window.addEventListener("load", function(){
    UserInterface.render();
});