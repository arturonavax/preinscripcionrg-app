var menu = $("#menu");
var menuBtn = $("#menu__btn");

menuBtn.addEventListener("click", toggleMenu);

function toggleMenu() {
    menu.classList.toggle("header__nav--open");
}