(function(){
    self.Menu = function(menuView) {
        
    }
})();

let menu = $("#menu"),
    menuBtn = $("#menu__btn"),
    dark = $("#dark");

menuBtn.addEventListener("click", toggleMenu);
dark.addEventListener("click", toggleMenu);

function toggleMenu() {
    menu.classList.toggle("header__nav--open");
    dark.classList.toggle("dark--move");
}

/* ------------------------------------------------------ */

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