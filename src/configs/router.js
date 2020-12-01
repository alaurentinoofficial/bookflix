import { AccountController } from "../controllers/account";
import { BookController } from "../controllers/books";
import { ReviewController } from "../controllers/review";
import { CatalogController } from "../controllers/catalog";
import { CategoryController } from "../controllers/category";
import { NotFound } from "../controllers/not_found";

import { Authenticate } from "../middlewares/auth";
import { RegisterValidation, LoginValidation } from "../validations/account";
import { InsertBookValidation, UpdateBookValidation, UrlBookValidation } from "../validations/book";
import { InsertCatalogValidation, UpdateCatalogValidation, UrlCatalogValidation } from "../validations/catalog";
import { InsertCategoryValidation, UpdateCategoryValidation, UrlCategoryValidation } from "../validations/category";
import { InsertReviewValidation, UrlReviewValidation } from "../validations/review";

exports.Router = app => {

    app.route('/register')
        .post(RegisterValidation, AccountController.register);

    app.route('/login')
        .post(LoginValidation, AccountController.login);

    app.route('/book')
        .get(BookController.get)
        .post(Authenticate, InsertBookValidation, BookController.insert);
    
    app.route('/book/:bookId')
        .get(UrlBookValidation, BookController.getById)
        .put(Authenticate, UrlBookValidation, UpdateBookValidation, BookController.update)
        .delete(Authenticate, UrlBookValidation, BookController.delete);
    
    app.route('/book/:bookId/review')
        .get(ReviewController.getByBookId)
        .post(UrlBookValidation, InsertReviewValidation, ReviewController.insert);
    
    app.route('/book/:bookId/review/:reviewId')
        .get(UrlBookValidation, UrlReviewValidation, ReviewController.getById)
        .delete(Authenticate, UrlReviewValidation, UrlBookValidation, ReviewController.delete);
    
    app.route('/account')
        .get(Authenticate, AccountController.getById);

    app.route('/account/review')
        .get(Authenticate, ReviewController.getByAccountId);
    
    app.route('/catalog')
        .get(CatalogController.get)
        .post(Authenticate, InsertCatalogValidation, CatalogController.insert);
    
    app.route('/catalog/:catalogId')
        .get(UrlCatalogValidation, CatalogController.getById)
        .put(Authenticate, UrlCatalogValidation, UpdateCatalogValidation, CatalogController.update)
        .delete(Authenticate, UrlCatalogValidation, CatalogController.delete);
    
    app.route('/category')
        .get(CategoryController.get)
        .post(Authenticate, InsertCategoryValidation, CategoryController.insert);
    
    app.route('/category/:categoryId')
        .get(UrlCategoryValidation, CategoryController.getById)
        .put(Authenticate, UrlCategoryValidation, UpdateCategoryValidation, CategoryController.update)
        .delete(Authenticate, UrlCategoryValidation, CategoryController.delete);

    app.route('*')
        .get(NotFound)
        .post(NotFound)
        .put(NotFound)
        .delete(NotFound);
}