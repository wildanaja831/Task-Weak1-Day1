let dataProject = [];

function formAlert() {
    let projectName = document.getElementById("project-name").value;
    let startDate = document.getElementById("start-date").value;
    let finishDate = document.getElementById("end-date").value;
    let description = document.getElementById("description").value;
    let reactjs = document.getElementById("reactjs").value;
    let nodejs = document.getElementById("nodejs").value;
    let nextjs = document.getElementById("nextjs").value;
    let typeScript = document.getElementById("typescript").value;
    let image = document.getElementById("form-project-upload-image").value;
    
    if(projectName == "") {
        return alert("Isi Nama atau Judul Project Terlebih Dahulu!");
    } else if(startDate == "") {
        return alert("Isi Tanggal Dimulai Project Terlebih Dahulu!");
    } else if(finishDate == "") {
        return alert("Isi Tanggal Diakhirinya Project Terlebih Dahulu!");
    } else if(description == "") {
        return alert("Isi Deskripsi atau Isi Project Telebih Dahulu!");
    } else if( reactjs.checked == "" || nodejs.checked == "" || nextjs.checked == "" || typeScript.checked == "") {
        return alert("Isi Teknologi Yang Digunakan Telebih Dahulu!");
    } else if(image == "") {
        return alert("Pilih Gambar Yang Ingin Dimasukan Terlebih Dahulu!");
    }
};

function submitData(event) {
    event.preventDefault();

    function getDistanceTime(){
        let diff = new Date(endDate) - new Date(startDate);

        let days = Math.floor(diff / (24 * 60 * 60 * 1000));
        let months = Math.floor(days / 30);
        let years = Math.floor(months / 12);
        let remainingDays = days % 30;
        let remainingMonths = months % 12;
        let daysAffix = `$`
      
        if (years == 1 && remainingMonths > 0 && remainingDays > 0) {
            return `${years} Years ${remainingMonths} Months ${remainingDays} Days`;
        } else if (years > 0 && remainingMonths > 0 && remainingDays == 0){
            return `${years} Years ${remainingMonths} Months`;
        } else if (years > 0 && remainingMonths == 0 && remainingDays == 0){
            return `${years} Years`;
        } else if (years == 0 && remainingMonths > 0 && remainingDays > 0){
            return `${remainingMonths} Months ${remainingDays} Days`;
        } else if (years == 0 && remainingMonths == 0 && remainingDays > 0){
            return `${remainingDays} Days`;
        } 
    }
    

    // Deklarasi variable dari input
    let title = document.getElementById("project-name").value;
    let startDate = document.getElementById("start-date").value;
    let endDate = document.getElementById("end-date").value;
    let description = document.getElementById("description").value;
    let image = document.getElementById("form-project-upload-image").files;
    let duration = getDistanceTime();

    const reactjs = '<img src="assets/images/reactjs.png" alt="reactjs">';
    const nextjs = '<img src="assets/images/nextjs.png" alt="nextjs">';
    const nodejs = '<img src="assets/images/nodejs.png" alt="nodejs">';
    const typeScript = '<img src="assets/images/typescript.png" alt="typescript">';

    let reactjsImg = document.getElementById("reactjs").checked ? reactjs : "";
    let nextjsImg = document.getElementById("nextjs").checked ? nextjs : "";
    let nodejsImg = document.getElementById("nodejs").checked ? nodejs : "";
    let typeScriptImg = document.getElementById("typescript").checked ? typeScript : "";

    // Membuat url image
    let imageUrl = URL.createObjectURL(image[0]);

    // Membuat object data project
    let data = {
        title,
        duration,
        description,
        reactjsImg,
        nextjsImg,
        nodejsImg,
        typeScriptImg,
        imageUrl,
        postAt: new Date()
    };
    dataProject.push(data);
    loopProjectList();
}

function loopProjectList() {
    document.getElementById("project-list").innerHTML = "";

    for (let i = 0; i < dataProject.length; i++) {
        document.getElementById("project-list").innerHTML += `
            <div class="project-items">
                <div class="project-items-container">
                    <div class="project-list-image">
                        <img src="${dataProject[i].imageUrl}" alt="project-list">
                    </div>
                    <div class="project-list-title">
                        <p class="list-title"><a target="_blank" href="project-detail.html">${dataProject[i].title}</a></p>
                        <p class="list-duration">durasi: ${dataProject[i].duration} </p>
                    </div>
                    <div>
                        <p class="list-description">${dataProject[i].description}</p>
                    </div>
                    <div class="technology">
                        ${dataProject[i].nodejsImg}
                        ${dataProject[i].nextjsImg}
                        ${dataProject[i].reactjsImg}
                        ${dataProject[i].typeScriptImg}
                    </div>
                    <div class="project-list-button">
                        <button class="edit" type="button">edit</button>
                        <button class="delete" type="button">delete</button>
                    </div>
                </div>
            </div>`;
    }
}

// setInterval(function () {
//     loopProjectList();
//   }, 3000);