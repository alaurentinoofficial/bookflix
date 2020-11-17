<template>
  <div class="row justify-content-center index-carossel">
        <div class="col-12">
            <div id="carouselExampleIndicators" class="carousel slide" data-ride="carousel">
                <ol class="carousel-indicators">
                    <li data-target="#carouselExampleIndicators" data-slide-to="0" class="point-slider active"></li>
                    <li data-target="#carouselExampleIndicators" data-slide-to="1" class="point-slider"></li>
                    <li data-target="#carouselExampleIndicators" data-slide-to="2" class="point-slider"></li>
                </ol>
                <div class="carousel-inner">
                    <div class="carousel-item active">
                        <img class="banner" src="@/assets/bannerImage.jpg" alt="Banner do livro de Lancamento">
                    </div>
                    <div class="carousel-item">
                        <img class="banner" src="@/assets/bannerImage.jpg" alt="Banner do livro de Lancamento">
                    </div>
                    <div class="carousel-item">
                        <img class="banner" src="@/assets/bannerImage.jpg" alt="Banner do livro de Lancamento">
                    </div>
                </div>
                <a class="carousel-control-prev" href="#carouselExampleIndicators" role="button"
                    data-slide="prev">
                    <span class="carousel-control-prev-icon" aria-hidden="true"></span>
                    <span class="sr-only">Previous</span>
                </a>
                <a class="carousel-control-next" href="#carouselExampleIndicators" role="button"
                    data-slide="next">
                    <span class="carousel-control-next-icon" aria-hidden="true"></span>
                    <span class="sr-only">Next</span>
                </a>
            </div>
        </div>
    </div>
</template>

<script>
export default {

}
</script>

<style scoped>
.carousel-control-next {
    width: 5%;
    background-color: #000000;
}

.carousel-control-prev {
    width: 5%;
    background-color: #000000;
}


.carousel-control-next-icon {
    background-image: url("../../assets/Left.png");
    transform: rotate(180deg);
    width: 1rem;
    height:  3rem;
}

.carousel-control-prev-icon {
    background-image: url("../../assets/Left.png");
    width: 1rem;
    height:  3rem;
}

.banner {
    width: 100%;
    height: 100%;
    object-fit: cover;
}

.carousel-item {
    height: 25rem;
}

.point-slider {
    background-color: #000504 !important;
    height: 1rem !important;
    width: 1rem !important;
    border-radius: 100%;
    opacity: 1 !important;
}

ol .active {
    background-color: #db2321 !important;
    height: 1rem !important;
    width: 1rem !important;
    border-radius: 100%;
}
</style>