import { AccountController } from "../controllers/account";
import { BookController } from "../controllers/books";
import { ReviewController } from "../controllers/review";
import { NotFound } from "../controllers/not_found";

import { RegisterValidation, LoginValidation, UrlAccountValidation } from "../validations/account";
import { InsertBookValidation, UpdateBookValidation, UrlBookValidation } from "../validations/book";
import { InsertReviewValidation, UrlReviewValidation } from "../validations/review";

exports.Router = app => {

    app.route('/register')
        .post(RegisterValidation, AccountController.register);

    app.route('/login')
        .post(LoginValidation, AccountController.login);

    app.route('/book')
        .get(BookController.get)
        .post(InsertBookValidation, BookController.insert);
    
    app.route('/book/:bookId')
        .get(UrlBookValidation, BookController.getById)
        .put(UrlBookValidation, UpdateBookValidation, BookController.update)
        .delete(UrlBookValidation, BookController.delete);
    
    app.route('/book/:bookId/review')
        .get(ReviewController.getByBookId)
        .post(UrlBookValidation, InsertReviewValidation, ReviewController.insert);
    
    app.route('/book/:bookId/review/:reviewId')
        .get(UrlBookValidation, UrlReviewValidation, ReviewController.getById)
        .delete(UrlReviewValidation, UrlBookValidation, ReviewController.delete);
    
    app.route('/account/:accountId/review')
        .get(UrlAccountValidation, ReviewController.getByAccountId);

    app.route('*')
        .get(NotFound)
        .post(NotFound)
        .put(NotFound)
        .delete(NotFound);
}