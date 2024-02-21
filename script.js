
var music = ["My Mind & Me", "Born To Die", "You're Somebody Else"];
var artist = ['Selena Gomez', 'Flora Cash', 'Lana Del Rey'];

console.log(music[0]+" by "+artist[0]);
console.log(music[1]+" by "+artist[2]);
console.log(music[2]+" by "+artist[1]);


let artist1 ={firstName: "ERKANAT", lastName: "O", profession: "artist", we: 5};

document.querySelector("#button").onclick = function(){

    const button = document.querySelector("#button");
    button.innerText = "SOLD!!!";
}
document.querySelector("#button1").onclick = function(){

    const button = document.querySelector("#button1");
    button.innerText = "SOLD!!!";
}
document.querySelector("#button2").onclick = function(){

    const button = document.querySelector("#button2");
    button.innerText = "SOLD!!!";
}
document.querySelector("#button3").onclick = function(){

    const button = document.querySelector("#button3");
    button.innerText = "SOLD!!!";
}
document.querySelector("#button4").onclick = function(){

    const button = document.querySelector("#button4");
    button.innerText = "SOLD!!!";
}
document.querySelector("#button5").onclick = function(){

    const button = document.querySelector("#button5");
    button.innerText = "SOLD!!!";

}



function showTime(){
    var date = new Date();
    var h = date.getHours(); // 0 - 23
    var m = date.getMinutes(); // 0 - 59
    var s = date.getSeconds(); // 0 - 59
    var session = "AM";

    if(h == 0){
        h = 12;
    }

    if(h > 12){
        h = h - 12;
        session = "PM";
    }

    h = (h < 10) ? "0" + h : h;
    m = (m < 10) ? "0" + m : m;
    s = (s < 10) ? "0" + s : s;

    var time = h + ":" + m + ":" + s + " " + session;
    document.getElementById("MyClockDisplay").innerText = time;
    document.getElementById("MyClockDisplay").textContent = time;

    setTimeout(showTime, 1000);

}

showTime();

const button = document.getElementById('toggle')
const nav = document.getElementById('nav')

button.addEventListener('click', ()=> nav.classList.toggle('active'))
