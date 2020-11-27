import grpc from 'grpc'
import { grpcCode, RPCtoError } from "../configs/strings"
import { validationResult } from "express-validator";

let proto = grpc.load('./protos/book.proto')
var ReviewService = new proto.ReviewService(process.env.BOOK_SERVICE || "localhost:9010", grpc.credentials.createInsecure())

export class ReviewController {
    static insert (req, res) {
        const checker = validationResult(req);
        if (!checker.isEmpty()) {
            let result = grpcCode.InvalidArgument;
            result.errors = checker.array();

            return res.status(400).json(result);
        }

        req.body.book_id = req.params.bookId

        ReviewService.Insert(req.body, (err, result) => {
            if(err)
                return res.json(RPCtoError(err));
            
            let response = JSON.parse(JSON.stringify(grpcCode.OK));
    
            res.json(response);
        })
    }

    static getByAccountId (req, res) {
        ReviewService.GetByAccountId({id: req.userId}, (err, result) => {
            if(err)
                return res.json(RPCtoError(err));
            
            let response = JSON.parse(JSON.stringify(grpcCode.OK));
            response.body = result.reviews
    
            res.json(response);
        })
    }

    static getByBookId (req, res) {
        ReviewService.GetByBookId({id: req.params.bookId}, (err, result) => {
            if(err)
                return res.json(RPCtoError(err));
            
            let response = JSON.parse(JSON.stringify(grpcCode.OK));
            response.body = result.reviews
    
            res.json(response);
        })
    }

    static getById (req, res) {
        const checker = validationResult(req);
        if (!checker.isEmpty()) {
            let result = grpcCode.InvalidArgument;
            result.errors = checker.array();

            return res.status(400).json(result);
        }

        ReviewService.GetById({id: req.params.reviewId}, (err, result) => {
            if(err)
                return res.json(RPCtoError(err));
            
            let response = JSON.parse(JSON.stringify(grpcCode.OK));
            response.body = result
    
            res.json(response);
        })
    }

    static delete (req, res) {
        const checker = validationResult(req);
        if (!checker.isEmpty()) {
            let result = grpcCode.InvalidArgument;
            result.errors = checker.array();

            return res.status(400).json(result);
        }

        ReviewService.Delete({id: req.params.reviewId}, (err, result) => {
            if(err)
                return res.json(RPCtoError(err));
            
            let response = JSON.parse(JSON.stringify(grpcCode.OK));
    
            res.json(response);
        })
    }
}