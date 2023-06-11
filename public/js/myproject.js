let dataProject = [];

function formAlert() {
    let projectName = document.getElementById("project-title").value;
    let startDate = document.getElementById("start-date").value;
    let finishDate = document.getElementById("end-date").value;
    let description = document.getElementById("description").value;
    let image = document.getElementById("form-project-upload-image").value;
    
    if(projectName == "") {
        return alert("Isi Nama atau Judul Project Terlebih Dahulu!");
    } else if(startDate == "") {
        return alert("Isi Tanggal Dimulai Project Terlebih Dahulu!");
    } else if(finishDate == "") {
        return alert("Isi Tanggal Diakhirinya Project Terlebih Dahulu!");
    } else if(description == "") {
        return alert("Isi Deskripsi atau Isi Project Telebih Dahulu!");
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
      
        if (years > 0 && remainingMonths > 0 && remainingDays > 0) {
            return `${years} Years ${remainingMonths} Months ${remainingDays} Days`;
        } else if (years > 0 && remainingMonths > 0 && remainingDays == 0){
            return `${years} Years ${remainingMonths} Months`;
        } else if (years > 0 && remainingMonths == 0 && remainingDays == 0){
            return `${years} Years`;
        } else if (years > 0 && remainingMonths == 0 && remainingDays > 0){
            return `${years} Years ${remainingDays} Days`;
        } else if (years == 0 && remainingMonths > 0 && remainingDays > 0){
            return `${remainingMonths} Months ${remainingDays} Days`;
        } else if (years == 0 && remainingMonths > 0 && remainingDays == 0){
            return `${remainingMonths} Months`;
        } else if (years == 0 && remainingMonths == 0 && remainingDays > 0){
            return `${remainingDays} Days`;
        } 
    }
    

    // Deklarasi variable dari input
    let title = document.getElementById("project-title").value;
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
        <div class="d-flex my-4 cards">
            <div class="card">
                <img src="${dataProject[i].imageUrl}" class="card-img-top" alt="">
                <div class="card-body">
                    <div class="mb-4">
                        <h5 class="">${dataProject[i].title}</h5>
                        <p class="list-duration">Durasi : ${dataProject[i].duration}</p>
                    </div>
                    <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
                    <div class="technology">
                        ${dataProject[i].nodejsImg}
                        ${dataProject[i].nextjsImg}
                        ${dataProject[i].reactjsImg}
                        ${dataProject[i].typeScriptImg}
                    </div>
                    <div class="d-flex justify-content-between mt-4">
                        <a href="#" class="btn btn-outline-primary w-50 me-1">Edit <i class="fa-solid fa-pen-to-square"></i></a>
                        <a href="#" class="btn btn-outline-danger w-50 ms-1">Delete <i class="fa-solid fa-trash"></i></a>
                    </div>
                </div>  
            </div>
        </div>`;
    }
}
