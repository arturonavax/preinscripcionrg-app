let username = document.getElementById("username"),
    password = document.getElementById("password"),
    btn = document.getElementById("btn");

btn.addEventListener("click", function(){
    
    var xhr = new XMLHttpRequest();

    xhr.open("POST", "/api/login", true);
    xhr.setRequestHeader("Content-Type", "application/json");

    let info = {
        username: username.value,
        password: password.value
    }

    xhr.send(JSON.stringify(info));


    xhr.addEventListener("load", function(e){
        let json = JSON.parse(e.currentTarget.response);

        document.getElementById("message").textContent = json.message;
        console.log(json);

        let xhr2 = new XMLHttpRequest();
        xhr2.open("POST", "/graphql", true);
        xhr2.setRequestHeader("Content-Type", "application/graphql");
        xhr2.setRequestHeader("Authorization", json.token);

        let query = `
        
        query {
            users {
                username,
                password
            }
        }
        
        `;
        xhr2.send(query);

        xhr2.addEventListener("load", function(e){
            let json2 = JSON.parse(e.currentTarget.response);
            console.log(json2);
        });
    });
});