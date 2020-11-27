import { AccountController } from "../controllers/account";
import { BookController } from "../controllers/books";
import { NotFound } from "../controllers/not_found";

import { RegisterValidation, LoginValidation } from "../validations/account";
import { InsertBookValidation, UpdateBookValidation, UrlBookValidation } from "../validations/book";

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
    
    app.route('*')
        .get(NotFound)
        .post(NotFound)
        .put(NotFound)
        .delete(NotFound);
}