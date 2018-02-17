if (!sessionStorage.getItem("token")) {
    window.location.replace("/user");
}

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