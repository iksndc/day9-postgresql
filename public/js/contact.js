function submitData() {
    let name = document.getElementById("name").value;
    let email = document.getElementById("email").value;
    let phone = document.getElementById("phone").value;
    let subject = document.getElementById("subject").value;
    let message = document.getElementById("message").value;
    console.log(name, email, phone, subject, message)

    let emailReceiver = "iksndc@gmail.com"

    if (name == "") {
        return alert('Nama harus diisi')
    } else if (email == "") {
        return alert('Email harus diisi')
    } else if (phone == "") {
        return alert('Nomor telepon harus diisi')
    } else if (subject == "") {
        return alert('Subject harus diisi')
    } else if (message == "") {
        return alert('Message harus diisi')
    }

    let link = document.createElement(`a`)
    link.href = `mailto:${email}?subject=${subject}&body=Hello my name ${name}, ${message}, Lets talk with me asap on ${phone}`
    link.click()

    let student = {
        name,
        email,
        phone,
        subject,
        message
    }

    console.log(student)

}