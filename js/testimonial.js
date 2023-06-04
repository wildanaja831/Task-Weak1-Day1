// class Testimonial{
//     #image = "";
//     #paragraph = "";
//     #author = "";

//     constructor(image, paragraph, author){
//         this.#image = image;
//         this.#paragraph = paragraph;
//         this.#author = author;
//     }

//     get image(){
//         return this.#image;
//     }

//     get paragraph(){
//         return this.#paragraph;
//     }

//     get author(){
//         return this.#author;
//     }

//     get cardHTML(){
//         return `<div class="testimonial-card">
//                     <img src="${this.image}" alt="">
//                     <p class="paragraph">" ${this.paragraph} "</p>
//                     <p class="author">~ ${this.author}</p>
//                 </div>`;
//     }
// }


// const data1 = new Testimonial("assets/images/5506218.jpg", "Kebersihan sebagian dari iman1!", "MWF aja1");
// const data2 = new Testimonial("assets/images/5506218.jpg", "Kebersihan sebagian dari iman2!", "MWF aja2");
// const data3 = new Testimonial("assets/images/5506218.jpg", "Kebersihan sebagian dari iman3!", "MWF aja3");

// let datas = [data1,data2,data3];
// let cardHTML = ""


// for (let i = 0; i < datas.length; i++) {
//     cardHTML += datas[i].cardHTML;
//   }
  
// document.getElementById("testimonials").innerHTML = cardHTML;


// let testimonialData = [
//     {
//         image : "assets/images/5506218.jpg",
//         paragraph : "Kebersihan sebagian dari iman1!",
//         author : "rate 1",
//         rate : 1
//     },
//     {
//         image : "assets/images/5506218.jpg",
//         paragraph : "Kebersihan sebagian dari iman2!",
//         author : "rate 2",
//         rate : 2
//     },
//     {
//         image : "assets/images/5506218.jpg",
//         paragraph : "Kebersihan sebagian dari iman3!",
//         author : "rate 3",
//         rate : 3
//     },
//     {
//         image : "assets/images/5506218.jpg",
//         paragraph : "Kebersihan sebagian dari iman4!",
//         author : "rate 4",
//         rate : 4
//     },
//     {
//         image : "assets/images/5506218.jpg",
//         paragraph : "Kebersihan sebagian dari iman5!",
//         author : "rate 5",
//         rate : 5
//     }
// ]

// function showAllTestimonial(){
//     let cardHTML = "";

//     testimonialData.forEach(function(item){
//         cardHTML += `<div class="testimonial-card">
//                         <img src="${item.image}" alt="">
//                         <p class="paragraph">" ${item.paragraph} "</p>
//                         <div style="display: flex; float: right;">
//                             <p style="margin-right: 10px;">${item.rate}<i class="fa-solid fa-star"></i></p>
//                             <p class="author">~ ${item.author}</p>
//                         </div>
//                     </div>`;
//     });

//     document.getElementById("testimonials").innerHTML = cardHTML;
// }

// showAllTestimonial();

// function filterRating(rate){
//     let cardHTML = "";

//     const filtering = testimonialData.filter(function(item){
//         return item.rate === rate;
//     });

//     if(filtering.length === 0){
//         cardHTML += `<h3>Data Kosong</h3>`;
//     } else{
//         filtering.forEach(function(item){
//             cardHTML += `<div class="testimonial-card">
//                             <img src="${item.image}" alt="">
//                             <p class="paragraph">" ${item.paragraph} "</p>
//                             <div style="display: flex; float: right;">
//                                 <p style="margin-right: 10px;">${item.rate}<i class="fa-solid fa-star"></i></p>
//                                 <p class="author">~ ${item.author}</p>
//                             </div>
//                          </div>`;
//         });
//     };

//     document.getElementById("testimonials").innerHTML = cardHTML;
// }