import Vue from 'vue'
import Vuex from 'vuex'
import Repository from "@/repositories/RepositoryFactory";

const BookRepository = Repository.get("books");
const ReviewRepository = Repository.get("review");

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    genres: [],
    reviews: [],
    bookToReview: {},
    starToReview: {},
    logginState: false
  },
  actions: {
    async getBooks({ commit }) {
      commit('loadBooks', await BookRepository.getAll());
    },
    async getReviewsByBookId({ commit }, obj) {
      commit('loadReviews', await ReviewRepository.getAllByBookId(obj.bookId));
    },
    async getBookById({ commit }, obj) {
      commit('loadBookToReview', await BookRepository.getById(obj.id));
    },
    async updateRating({ commit }, obj) {
      commit('updateStarToReview', obj.rating);
    },
    async addReview({ commit }, obj) {
      commit('addReviewToBook', obj);
    },
  },
  mutations: {
    loadBooks: (state, res) => {
      const { data } = res;

      data.body.books.forEach(element => {
        state.genres.push(...element.categories);
        state.genres = state.genres.filter((value, index, self) => {
            return self.indexOf(self.find((genre) => genre.id == value.id)) === index;
        });
      });

      state.genres.map(function(element) {
        element.books = data.body.books.filter(e => e.categories.find(category => category.id == element.id) != null);
        return element;
      });
    },
    loadReviews: (state, res) => {
      const { data } = res;
      state.reviews = data.body;
    },
    loadBookToReview: (state, res) => {
      const { data } = res;
      state.bookToReview = data
    },
    updateStarToReview: (state, rating) => {
      state.starToReview = rating;
    },
    addReviewToBook: (state, rating) => {
      state.reviews.push(rating.review);
      ReviewRepository.add(rating.review, rating.review.BookId);
    },
  }
})
