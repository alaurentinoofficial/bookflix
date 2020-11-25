import { AccountController } from "../controllers/account";
import { NotFound } from "../controllers/not_found";

exports.Router = app => {

    app.route('/register')
        .post(AccountController.register);

    app.route('/login')
        .post(AccountController.login);
    
    app.route('*')
        .get(NotFound)
        .post(NotFound)
        .put(NotFound)
        .delete(NotFound);
}