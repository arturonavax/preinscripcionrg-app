(function(){
    self.UI = function(menu,profile, studentRegister, students){
        this.MenuView = menu;
        this.ProfileView = profile;
        this.StudentRegister = studentRegister;
        this.StudentsView = students;

        this.MenuView.Menu.Username = this.ProfileView.Profile.Username;
        this.BtnBackWindowProfile = $("#btn-back-profile");
        this.BtnBackWindowStudentRegister = $("#btn-back-studentRegister");
        this.BtnBackWindowStudents = $("#btn-back-students");

        let self = this;

        // Eventos del menu.
        this.MenuView.MenuBtn.addEventListener("click", function(){
            self.MenuView.MenuContainer.classList.toggle("modal--open");
        });

        this.MenuView.MenuContainer.addEventListener("click", function(){
            self.MenuView.MenuContainer.classList.toggle("modal--open");
        });

        this.MenuView.BtnStudents.addEventListener("click", function(){
            self.StudentsView.ContainerStudents.classList.add("modal-article--open");
            $("#section").classList.add("tx-C");
            $("#article-studentRegister").style.height = "430px";
        });

        this.MenuView.BtnLogout.addEventListener("click", function(){
            sessionStorage.removeItem("token");

            window.location.replace("/user");
        });

        this.MenuView.BtnProfile.addEventListener("click", function(){
            self.ProfileView.ContainerProfile.classList.add("modal-article--open");
            $("#section").classList.add("tx-C");
        });

        this.StudentRegister.BtnStudentRegister.addEventListener("click", function(){
            self.StudentRegister.ContainerStudentRegister.classList.add("modal-article--open");
        });

        // Eventos de la ventana profile.
        this.BtnBackWindowProfile.addEventListener("click",function() {
            self.ProfileView.ContainerProfile.classList.remove("modal-article--open");
            $("#section").classList.remove("tx-B");
            $("#section").classList.remove("tx-C");
            $("#section").classList.remove("tx-D");
            $("#section").classList.add("tx-A");
        });

        this.ProfileView.BtnEditProfile.addEventListener("click", function(){
            self.ProfileView.renderEdit();
        });

        this.ProfileView.BtnCancelEditProfile.addEventListener("click", function(){
            self.ProfileView.cancelEdit();
        }); 

        this.ProfileView.BtnSaveEditProfile.addEventListener("click", function() {
            self.ProfileView.BtnSaveEditProfile.textContent = "Guardando..";
            self.ProfileView.saveEdit();
        });

        // Eventos de la ventana de creacion de estudiantes.
        //  Botones de avance.
        this.BtnBackWindowStudentRegister.addEventListener("click", function(){
            self.StudentRegister.ContainerStudentRegister.classList.remove("modal-article--open");
            $("#section").classList.remove("tx-B");
            $("#section").classList.remove("tx-C");
            $("#section").classList.remove("tx-D");
            $("#section").classList.add("tx-A");
        })

        this.StudentRegister.BtnNextPage1.addEventListener("click", function() {
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page2");
            $("#article-studentRegister").style.height = "720px";
            $("#section").classList.remove("tx-A");
            $("#section").classList.add("tx-C");
        });
        this.StudentRegister.BtnNextPage2.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "600px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page2");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page3");
        });
        this.StudentRegister.BtnNextPage3.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "580px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page3");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page4");
            $("#section").classList.remove("tx-C");
            $("#section").classList.add("tx-B");
        });
        this.StudentRegister.BtnNextPage4.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "750px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page4");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page5");
            $("#section").classList.remove("tx-B");
            $("#section").classList.add("tx-C");
        });
        this.StudentRegister.BtnNextPage5.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "540px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page5");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page6");
            $("#section").classList.remove("tx-C");
            $("#section").classList.add("tx-B");
        });
        this.StudentRegister.BtnNextPage6.addEventListener("click", function() {
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page6");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page7");
        });
        this.StudentRegister.BtnNextPage7.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "700px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page7");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page8");
        });
        this.StudentRegister.BtnNextPage8.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "490px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page8");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page9");
        });
        this.StudentRegister.BtnNextPage9.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "930px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page9");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page10");
            $("#section").classList.remove("tx-B");
            $("#section").classList.add("tx-D");
        });
        this.StudentRegister.BtnNextPage10.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "300px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page10");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page11");
            $("#section").classList.remove("tx-D");
            $("#section").classList.add("tx-A");
        });
        // Botones de retroceso.
        this.StudentRegister.BtnBackPage2.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "430px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page2");
            $("#section").classList.remove("tx-C");
            $("#section").classList.add("tx-A");
        });
        this.StudentRegister.BtnBackPage3.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "720px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page3");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page2");
        });
        this.StudentRegister.BtnBackPage4.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "600px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page4");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page3");
            $("#section").classList.remove("tx-B");
            $("#section").classList.add("tx-C");
        });
        this.StudentRegister.BtnBackPage5.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "580px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page5");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page4");
            $("#section").classList.remove("tx-C");
            $("#section").classList.add("tx-B");
        });
        this.StudentRegister.BtnBackPage6.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "750px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page6");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page5");
            $("#section").classList.remove("tx-B");
            $("#section").classList.add("tx-C");
        });
        this.StudentRegister.BtnBackPage7.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "540px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page7");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page6");
        });
        this.StudentRegister.BtnBackPage8.addEventListener("click", function() {
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page8");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page7");
        });
        this.StudentRegister.BtnBackPage9.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "700px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page9");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page8");
        });
        this.StudentRegister.BtnBackPage10.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "490px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page10");
            self.StudentRegister.PanelMove.classList.add("panelMove-ficha--page9");
            $("#section").classList.remove("tx-D");
            $("#section").classList.add("tx-B");
        });
        this.StudentRegister.BtnBackPage11.addEventListener("click", function() {
            $("#article-studentRegister").style.height = "300px";
            self.StudentRegister.PanelMove.classList.remove("panelMove-ficha--page11");
            $("#section").classList.remove("tx-A");
            $("#section").classList.add("tx-D");
        });

        $("#btn-prein").addEventListener("click", function(){
            $("#modal").classList.add("modal--open");
        });
        
        $("#modal-cancel").addEventListener("click", function(){
            $("#modal").classList.remove("modal--open");

        });

        // SIGE
        this.StudentRegister.SIGEYes.addEventListener("click", function(){
           self.StudentRegister.SIGETitle.classList.remove("no-visible"); 
           self.StudentRegister.SIGECOD.classList.remove("no-visible"); 
        });

        this.StudentRegister.SIGENo.addEventListener("click", function(){
           self.StudentRegister.SIGETitle.classList.add("no-visible"); 
           self.StudentRegister.SIGECOD.classList.add("no-visible"); 
        });

        this.StudentRegister.BtnPreInscribir.addEventListener("click", function(){
            self.StudentRegister.register();
        });

        // Eventos de la ventana de estudiantes.
        this.BtnBackWindowStudents.addEventListener("click", function() {
            self.StudentsView.ContainerStudents.classList.remove("modal-article--open");
            $("#section").classList.remove("tx-B");
            $("#section").classList.remove("tx-C");
            $("#section").classList.remove("tx-D");
            $("#section").classList.add("tx-A");

        });
    }

    self.UI.prototype = {
        render: function(){
            this.MenuView.render();
            this.ProfileView.render();
            this.StudentRegister.render(this.ProfileView.Profile.Token);
            this.StudentsView.render();
        }
    }
})();
