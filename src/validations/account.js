import { param, body } from "express-validator"

exports.LoginValidation = [
    body('email').notEmpty().isString().isEmail(),
    body('password').notEmpty().isString().isLength({ min: 6 })
]

exports.RegisterValidation = [
    body('name').notEmpty().isString().isLength({ min: 2 }).trim(),
    body('email').notEmpty().isString().isEmail(),
    body('password').notEmpty().isString().isLength({ min: 6 })
]

exports.UrlAccountValidation = [
    param('accountId').notEmpty().isString().isUUID()
]