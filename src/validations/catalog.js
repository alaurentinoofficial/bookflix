import { param, body } from "express-validator";

exports.InsertCatalogValidation = [
    body('name').notEmpty().isString().isLength({ min: 2 }),
    body('items.*.name').notEmpty().isString().isLength({ min: 3 }),
    body('items.*.item_id').notEmpty().isString().isUUID(),
    body('items.*.image_url').notEmpty().isString().isURL(),
]

exports.UpdateCatalogValidation = [
    body('name').notEmpty().isString().isLength({ min: 2 }),
    body('items.*.name').notEmpty().isString().isLength({ min: 3 }),
    body('items.*.item_id').notEmpty().isString().isUUID(),
    body('items.*.image_url').notEmpty().isString().isURL(),
]

exports.UrlCatalogValidation = [
    param('catalogId').notEmpty().isString().isUUID()
]