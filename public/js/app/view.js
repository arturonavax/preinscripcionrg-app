(function(){
    self.UI = function(menu,profile){
        this.MenuView = menu;
        this.ProfileView = profile;

        this.MenuView.Menu.Username = this.ProfileView.Profile.Username;
        this.BtnBackWindowProfile = $("#btn-back");

        let self = this;

        this.MenuView.MenuBtn.addEventListener("click", function(){
            self.MenuView.MenuContainer.classList.toggle("header__nav--open");
            self.MenuView.Dark.classList.toggle("dark--move");
        });

        this.MenuView.BtnLogout.addEventListener("click", function(){
            sessionStorage.removeItem("token");

            window.location.replace("/user");
        });

        this.MenuView.BtnProfile.addEventListener("click", function(){
            self.ProfileView.ContainerProfile.classList.add("section__container--open");
        });

        this.BtnBackWindowProfile.addEventListener("click",function() {
            self.ProfileView.ContainerProfile.classList.remove("section__container--open");
        });

        this.ProfileView.BtnEditProfile.addEventListener("click", function(){
            self.ProfileView.renderEdit();
        });

    }

    self.UI.prototype = {
        render: function(){
            this.MenuView.render();
            this.ProfileView.render();
        }
    }
})();
