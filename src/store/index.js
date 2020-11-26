import Vue from 'vue'
import Vuex from 'vuex'
import Repository from "@/repositories/RepositoryFactory";
import { auth } from '@/store/modules/authModule.js';

const BookRepository = Repository.get("books");
const ReviewRepository = Repository.get("review");

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    books: [],
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
      state.books = data
    },
    loadReviews: (state, res) => {
      const { data } = res;
      state.reviews = data
    },
    loadBookToReview: (state, res) => {
      const { data } = res;
      state.bookToReview = data
    },
    updateStarToReview: (state, rating) => {
      state.starToReview = rating;
    },
    addReviewToBook: (state, rating) => {
      state.reviews.body.push(rating.review);
      ReviewRepository.add(rating.review, rating.review.BookId);
    },
  },
  modules: {
    auth
  }
})
