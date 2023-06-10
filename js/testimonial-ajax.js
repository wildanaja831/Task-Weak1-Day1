const promise = new Promise((resolve, rejcet) => {
    const xhr = new XMLHttpRequest();
    xhr.open("GET", "https://api.npoint.io/e0d5bb4873d8d4c1ebf7", true);

    xhr.onload = () => {
      if (xhr.status === 200) {
        resolve(JSON.parse(xhr.response));
      } else {
        reject("Error loading data.");
      }
    };

    xhr.onerror = () => {
      reject("Network error.");
    };

    xhr.send();
});

async function showAllTestimonial(){
    const  response = await promise;
    let cardHTML = "";

    response.forEach(function(item){
        cardHTML += `<div class="d-flex cardt my-3">
                        <div class="card">
                            <img src="${item.image}" class="card-img-top" alt="..." style="width: 400px; height: 300px; object-fit: cover;">
                            <div class="card-body">
                                <p class="card-text">${item.paragraph}</p>
                                <div class="d-flex my-2 float-end">
                                    <p>${item.rate} <i class="fa-solid fa-star"></i></p>
                                    <p class="ms-2">~ ${item.author}</p>
                                </div>
                            </div>
                        </div>
                    </div>`;
    });

    document.getElementById("testimonials").innerHTML = cardHTML;
}

showAllTestimonial();

async function filterRating(rate){
    const  response = await promise;
    let cardHTML = "";

    const filtering = response.filter(function(item){
        return item.rate === rate;
    });

    if(filtering.length === 0){
        cardHTML += `<h3>Data Kosong</h3>`;
    } else{
        filtering.forEach(function(item){
            cardHTML += `<div class="d-flex cardt my-3">
                            <div class="card">
                                <img src="${item.image}" class="card-img-top" alt="..." style="width: 400px; height: 300px; object-fit: cover;">
                                <div class="card-body">
                                    <p class="card-text">${item.paragraph}</p>
                                    <div class="d-flex my-2 float-end">
                                        <p>${item.rate} <i class="fa-solid fa-star"></i></p>
                                        <p class="ms-2">~ ${item.author}</p>
                                    </div>
                                </div>
                            </div>
                        </div>`;
        });
    };

    document.getElementById("testimonials").innerHTML = cardHTML;
}