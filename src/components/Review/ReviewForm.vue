<template>
  <div class="row justify-content-between">
    <div class="col-12 ml-3 mt-5 book-details">
      <div class="book-details-content mt-3 ml-3">
        <div class="row ml-auto mr-auto review-stars">
          <em
            class="fa fa-star book-rate-star-review col-lg-1 col-md-12 col-sm-12 col-xs-12 offset-md-1 offset-sm-1 offset-xs-1"
          ></em>
          <em
            class="fa fa-star book-rate-star-review col-lg-1 col-md-12 col-sm-12 col-xs-12 offset-md-1 offset-sm-1 offset-xs-1"
          ></em>
          <em
            class="fa fa-star book-rate-star-review col-lg-1 col-md-12 col-sm-12 col-xs-12 offset-md-1 offset-sm-1 offset-xs-1"
          ></em>
          <em
            class="fa fa-star book-rate-star-review col-lg-1 col-md-12 col-sm-12 col-xs-12 offset-md-1 offset-sm-1 offset-xs-1"
          ></em>
          <em
            class="fa fa-star book-rate-star-review col-lg-1 col-md-12 col-sm-12 col-xs-12 offset-md-1 offset-sm-1 offset-xs-1"
          ></em>
        </div>
        <div class="form-group mt-4 ml-auto mr-auto" style="width: 80%">
          <textarea
            class="form-control review-textarea"
            id="exampleFormControlTextarea1"
            rows="4"
          ></textarea>
          <div class="col-md-12 text-center">
            <button type="submit" class="btn btn-primary mt-3 review-button">
              ENVIAR
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.book-details {
  background-color: #232323;
  width: auto;
  border-radius: 1%;
}

.book-details {
  display: flex;
  flex-direction: column;
  color: white;
}

.book-rate-star-review {
  color: #f5c30b;
  font-size: 2.5rem;
}

.review-stars {
  width: 45%;
  justify-content: space-between;
}

.review-textarea {
  color: white;
  background-color: #171717;
  border-color: #474747;
}

.review-button {
  background-color: #b7323b;
  border-color: #b7323b;
}
</style>