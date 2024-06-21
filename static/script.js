var x = document.getElementById("login");
var y = document.getElementById("register");
var z = document.getElementById("btn");

function register() {
    x.style.left = "-400px";
    y.style.left = "50px";
    z.style.left = "110px";
    document.querySelector('.toggle-btn:nth-child(2)').style.color = "#fff";
    document.querySelector('.toggle-btn:nth-child(3)').style.color = "#000000";
}

function login() {
    x.style.left = "50px";
    y.style.left = "450px";
    z.style.left = "0px";
    document.querySelector('.toggle-btn:nth-child(2)').style.color = "#000000";
    document.querySelector('.toggle-btn:nth-child(3)').style.color = "#fff";
}
