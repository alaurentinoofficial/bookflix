import { param, body } from "express-validator";

exports.InsertReviewValidation = [
    body('rating').notEmpty().isFloat({ min: 0, max: 5 }),
    body('resume').optional().isString(),
    body('user_name').notEmpty().isString().isLength({ min: 5 }).isURL()
]

exports.UrlReviewValidation = [
    param('reviewId').notEmpty().isString().isUUID()
]