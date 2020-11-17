<template>
    <div class="col-lg-3 col-md-4 col-sm-6 col-xs-12 mt-3">
        <router-link to="/review">
            <img class="bookimage" src="@/assets/bookImage1.jpg" alt="Imagem do livro Antifragil">
            <div class="position-relative book-transparent-box"></div>
            <div class="book-rate d-flex flex-column">
                <div>
                    <em class="fa fa-star book-rate-star"></em>
                    <span>4.7</span>
                </div>
                <span class="book-title">Antifragil</span>
            </div>
        </router-link>
    </div>
</template>

<script>
export default {
    name: "Book"
}
</script>

<style scoped>
.bookimage {
    border-radius: 3%;
    width: 100%;
    height: 100%;
}

.booklist {
    width: 80%;
}

.book-title {
    color: white;
    margin-top: 0.2rem;
    font-size: 0.8rem;
    font-weight: normal;
}

.book-rate {
    font-weight: bold;
    color: white;
    top: -125%; 
    padding-left: 10%; 
    padding-right: 10%;
    width: 100%;
    position: relative;
}

.book-rate-star {
    color: #f5c30b;
}

.book-transparent-box {
    top: -100%; 
    left: 0; 
    width:100%; 
    height: 100%; 
    pointer-events: visible;
    background: rgb(255,255,255);
    background: linear-gradient(180deg, rgba(255,255,255,0) 0%, rgba(0,0,0,0.7035189075630253) 100%);
}

.book-transparent-box:hover {
    background: rgb(0,0,0,0.8);
    transition: 0.3s;
}
</style>