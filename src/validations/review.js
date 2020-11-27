import { param, body } from "express-validator";

exports.InsertReviewValidation = [
    body('rating').notEmpty().isFloat({ min: 0, max: 5 }),
    body('resume').optional().isString()
]

exports.UrlReviewValidation = [
    param('reviewId').notEmpty().isString().isUUID()
]