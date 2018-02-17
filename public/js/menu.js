var menu = $("#menu");
var menuBtn = $("#menu__btn");
var dark = $("#dark");

menuBtn.addEventListener("click", toggleMenu);
dark.addEventListener("click", toggleMenu);

function toggleMenu() {
    menu.classList.toggle("header__nav--open");
    //dark.classList.toggle("header__nav--open");
    dark.classList.toggle("dark--move");
}