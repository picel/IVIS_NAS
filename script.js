let pass = document.getElementById("pass");
let pass2 = document.getElementById("pass2");
let alert = document.getElementById("alert");
let submit = document.getElementById("submit");

function checkPass() {
    if (pass.value != pass2.value) {
        alert.style.display = "block";
        alert.style.color = "red";
        alert.innerHTML = "비밀번호가 일치하지 않습니다.";
        submit.classList.remove("btn-primary");
        submit.classList.add("btn-secondary");
        submit.disabled = true;
    } else {
        alert.style.display = "block";
        alert.style.color = "blue";
        alert.innerHTML = "비밀번호가 일치합니다.";
        submit.classList.remove("btn-secondary");
        submit.classList.add("btn-primary");
        submit.disabled = false;
    }
}

pass.addEventListener("change", checkPass);
pass2.addEventListener("change", checkPass);
