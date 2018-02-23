(function(){
    self.UI = function(menu,profile){
        this.MenuView = menu;
        this.ProfileView = profile;

    }

    self.UI.prototype = {
        render: function(){
            this.MenuView.render();
            this.ProfileView.render();
        }
    }
})();
