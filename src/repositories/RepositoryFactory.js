import BookRepository from './Books/BookRepository';
import ReviewRepository from './Reviews/ReviewRepository';

const repositories = {
    'books': BookRepository,
    'review': ReviewRepository
}
export default {
    get: name => repositories[name]
};