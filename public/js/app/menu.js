(function(){
    self.Menu = function() {
        this.Username = "";
        this.Home = "Inicio";
        this.Profile = "Mi Perfil";
        this.Logout = "Salida";
    }
})();

(function(){
    self.MenuView = function(menu) {
        this.Menu = menu;
        this.MenuBtn = $("#menu__btn");
        this.Dark = $("#dark");
        this.MenuContainer = $("#menu");
        this.MenuTitle = $("#menu-title");
        this.Username = $("#menu-titleUser");
        this.BtnHome = $("#menu-btnHome");
        this.BtnProfile = $("#menu-btnProfile");
        this.BtnLogout = $("#menu-btnLogout");
    }

    self.MenuView.prototype = {
        render: function(){
            this.Username.textContent = this.Menu.Username;
            this.BtnHome.textContent = this.Menu.Home;
            this.BtnProfile.textContent = this.Menu.Profile;
            this.BtnLogout.textContent = this.Menu.Logout;
        }
    }
})();
/*
let menu = $("#menu"),
    menuBtn = $("#menu__btn"),
    dark = $("#dark");

menuBtn.addEventListener("click", toggleMenu);
dark.addEventListener("click", toggleMenu);

function toggleMenu() {
    menu.classList.toggle("header__nav--open");
    dark.classList.toggle("dark--move");
}
*/
/* ------------------------------------------------------ */
/*
let btnLogout = $("#btn-logout"),
    btnProfile = $("#btn-menuProfile"),
    profileContainer = $("#container-profile"),
    btnBack = $("#btn-back");

btnLogout.addEventListener("click", function(){
    sessionStorage.removeItem("token");

    window.location.replace("/user");
});

btnProfile.addEventListener("click", function(){
    profileContainer.classList.add("section__container--profile--open");
});

btnBack.addEventListener("click", function() {
    profileContainer.classList.remove("section__container--profile--open");
});

$("#titleUser").textContent = profileUser.Username;
*/