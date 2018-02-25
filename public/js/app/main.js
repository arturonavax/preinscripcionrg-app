if (!sessionStorage.getItem("token")) {
    window.location.replace("/user");
}

let tokenUser = sessionStorage.getItem("token");

let profileUser = new Profile(),
    menuUser = new Menu();

profileUser.token = tokenUser;

let profileView = new ProfileView(profileUser),
    menuView = new MenuView(menuUser),
    studentRegister = new StudentRegisterView();
    students = new StudentsView();

let UserInterface = new UI(menuView,profileView,studentRegister, students);

window.addEventListener("load", function(){
    UserInterface.render();
});