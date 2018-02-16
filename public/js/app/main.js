if (!sessionStorage.getItem("token")) {
    window.location.replace("/user");
}

let btnProfile = $("#btn-menuProfile"),
    profileContainer = $("#container-profile"),
    btnBack = $("#btn-back");

btnProfile.addEventListener("click", function(){
    profileContainer.classList.add("section__container--profile--open");
});

btnBack.addEventListener("click", function() {
    profileContainer.classList.remove("section__container--profile--open");
});