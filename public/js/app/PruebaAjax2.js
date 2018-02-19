var xhr = new XMLHttpRequest();
console.log(xhr);

xhr.open("GET", "/user", true);
console.log(xhr);

xhr.send();

xhr.addEventListener("load", function(e){
    console.log(e);
});