import { param, body } from "express-validator";

exports.InsertBookValidation = [
    body('title').notEmpty().isString().isLength({ min: 2 }),
    body('author').notEmpty().isString().isLength({ min: 2 }),
    body('resume').notEmpty().isString().isLength({ min: 5 }),
    body('image_url').notEmpty().isString().isLength({ min: 5 }).isURL(),
    body('categories').notEmpty().isArray({ min: 1 }),
    body('categories.*.name').notEmpty().isString().isLength({ min: 3 }),
]

exports.UpdateBookValidation = [
    body('title').optional().isString().isLength({ min: 2 }),
    body('author').optional().isString().isLength({ min: 2 }),
    body('resume').optional().isString().isLength({ min: 5 }),
    body('image_url').optional().isString().isLength({ min: 5 }).isURL(),
    body('categories').optional().isArray({ min: 1 }),
    body('categories.*.name').optional().isString().isLength({ min: 3 }),
]

exports.UrlBookValidation = [
    param('bookId').notEmpty().isString().isUUID()
]