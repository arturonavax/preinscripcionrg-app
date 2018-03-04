var menu = $("#modal-header"),
    menuBtn = $("#menu__btn");

menuBtn.addEventListener("click", function(){
    menu.classList.add("modal--open");
});

menu.addEventListener("click",function(){
    menu.classList.remove("modal--open");
})

