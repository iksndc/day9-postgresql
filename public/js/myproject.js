let data = []

function addData(event) {

  event.preventDefault()

  let title = document.getElementById("input-blog-title").value;
  let startDate = document.getElementById("input-blog-start-date").value;
  let endDate = document.getElementById("input-blog-end-date").value;
  let description = document.getElementById("input-blog-description").value;
  let technologies = document.getElementsByName('input-blog-tech')
  let image = document.getElementById("input-blog-image").files;

  if (title == "") {
    return alert('title tidak boleh kosong!')
  } else if (description == "") {
    return alert('description tidak boleh kosong!')
  } else if (image.length == 0) {
    return alert('gambar tidak boleh kosong!')
  }

  let gambar = URL.createObjectURL(image[0])




  let blog = {
    title,
    startDate,
    endDate,
    description,

    gambar,
    postAt: new Date(),
    author: "iksndc"
  }

  data.push(blog)
  console.log(data)
  renderBlog()
}

function renderBlog() {
  document.getElementById("contents").innerHTML = ``
  for (let index = 0; index < data.length; index++) {
    document.getElementById("contents").innerHTML += `<div class="blog-list-item">
    <div class="blog-image">
      <img src="${data[index].gambar}" alt="" />
    </div>
    <div class="blog-content">

      <h3 style="font-weight: 700; font-size: 18pt;">
        <a href="blog-detail.html" target="_blank">${data[index].title}</a>
      </h3>
      <div class="detail-blog-content">
      <p style="color: gray;">Durasi: ${selisihWaktu(data[index].postAt)}</p>
      </div>
      <p style="max-block-size: 81px; overflow: hidden;">
      ${data[index].description}
      </p>
      <div>
        <img style="height: 40px; margin: 40px 10px 40px 10px;" src="assets/nodejs.png" alt="" />
        <img style="height: 40px; margin: 40px 10px 40px 10px;" src="assets/Socket-io.png" alt="" />
        <img style="height: 40px; margin: 40px 10px 40px 10px;" src="assets/Reactjs.png" alt="" />
        <img style="height: 40px; margin: 40px 10px 40px 10px;" src="assets/js.png" alt="" />
      </div>  
    </div>
    <div class="btn-group">
    <button class="btn-edit">edit</button>
    <button class="btn-delete">delete</button>
   </div>
  </div>`
  }
}

function konversiWaktu(time) {
  let monthName = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Des"]
  let hours = time.getHours()
  let minutes = time.getMinutes()

  if (hours < 10) {
    hours = "0" + hours
  }

  if (minutes < 10) {
    minutes = "0" + minutes
  }

  return `${time.getDate()} ${monthName[time.getMonth()]} ${time.getFullYear()} ${hours}:${minutes} WIB`
}

function selisihWaktu(time) {
  let timeNow = new Date()
  let timePost = time

  let distance = timeNow - timePost
  console.log("jarak", distance)

  let miliseconds = 1000 // 1 seconds adalah 1000 miliseconds

  let distanceDay = Math.floor(distance / (miliseconds * 60 * 60 * 24))
  let distanceHours = Math.floor(distance / (miliseconds * 60 * 60))
  let distanceMinutes = Math.floor(distance / (miliseconds * 60))
  let distanceSecond = Math.floor(distance / miliseconds) // 1 seconds ago

  if (distanceDay > 0) {
    return `${distanceDay} day ago`
  } else if (distanceHours > 0) {
    return `${distanceHours} hours ago`
  } else if (distanceMinutes > 0) {
    return `${distanceMinutes} minutes ago`
  } else {
    return `${distanceSecond} seconds ago`
  }
}

setInterval(function () {
  renderBlog()
}, 1000)


