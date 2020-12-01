import { param, body } from "express-validator";

exports.InsertCategoryValidation = [
    body('name').notEmpty().isString().isLength({ min: 2 }),
]

exports.UpdateCategoryValidation = [
    body('name').notEmpty().isString().isLength({ min: 2 }),
]

exports.UrlCategoryValidation = [
    param('categoryId').notEmpty().isString().isUUID()
]