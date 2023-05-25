function getData(){
    let name = document.getElementById("input-name").value
    let email = document.getElementById("input-email").value
    let phone = document.getElementById("phone-number").value
    let subject = document.getElementById("input-subject").value
    let message = document.getElementById("input-message").value

    if(name == ""){
        return alert("Nama Harus Di Isi!");
    } else if(email == ""){
        return alert("Email Harus Di Isi!");
    } else if(phone == ""){
        return alert("phone Harus Di Isi!");
    } else if(subject == ""){
        return alert("Subject Harus Di Pilih!");
    } else if(message == ""){
        return alert("Message Harus Di Isi!");
    }

    // console.log(name)
    // console.log(email)
    // console.log(phone)
    // console.log(subject)
    // console.log(message)

    let emailReceiver = "wildanaja110404@gmail.com"

    let a = document.createElement('a')
    a.href = `mailto:${emailReceiver}?subject=${subject}&body=Halo, Nama saya ${name}. ${message}. Anda bisa menghubungi saya di nomor ${phone}, Terimaksih.`
    a.click()
}

let dropdown = false;

function openDropdown(){
  let dropdownContainer = document.getElementById("dropdown-container");

  if (!dropdown) {
    dropdownContainer.style.display = "block";
    dropdown = true;
  } else {
    dropdownContainer.style.display = "none";
    dropdown = false;
  }
}