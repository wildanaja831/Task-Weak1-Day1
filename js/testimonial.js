class Testimonial{
    #image = "";
    #paragraph = "";
    #author = "";

    constructor(image, paragraph, author){
        this.#image = image;
        this.#paragraph = paragraph;
        this.#author = author;
    }

    get image(){
        return this.#image;
    }

    get paragraph(){
        return this.#paragraph;
    }

    get author(){
        return this.#author;
    }

    get cardHTML(){
        return `<div class="testimonial-card">
                    <img src="${this.image}" alt="">
                    <p class="paragraph">" ${this.paragraph} "</p>
                    <p class="author">~ ${this.author}</p>
                </div>`;
    }
}


const data1 = new Testimonial("assets/images/5506218.jpg", "Kebersihan sebagian dari iman1!", "MWF aja1");
const data2 = new Testimonial("assets/images/5506218.jpg", "Kebersihan sebagian dari iman2!", "MWF aja2");
const data3 = new Testimonial("assets/images/5506218.jpg", "Kebersihan sebagian dari iman3!", "MWF aja3");

let datas = [data1,data2,data3];
let cardHTML = ""


for (let i = 0; i < datas.length; i++) {
    cardHTML += datas[i].cardHTML;
  }
  
document.getElementById("testimonials").innerHTML = cardHTML;