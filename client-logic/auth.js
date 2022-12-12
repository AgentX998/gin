function registergin() {
    console.log(document.getElementById("name").value)
    console.log(document.getElementById("email").value)
    console.log(document.getElementById("password").value)
    var xhr = new XMLHttpRequest();
    xhr.open("POST", "http://localhost:8080/register", false);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
    Name:document.getElementById("name").value,
    Email:document.getElementById("email").value,
    Password:document.getElementById("password").value,
    }));
    /*xhr.open("POST", "http://localhost:8080/register", false);
    xhr.setRequestHeader('Content-Type', 'application/json');
    xhr.send(JSON.stringify({
    name:"lsdlfasdf",
    email:document.getElementById("email").value,
    password:document.getElementById("password").value,
    }));*/
  }