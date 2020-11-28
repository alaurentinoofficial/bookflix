import grpc from 'grpc'
import { grpcCode, RPCtoError } from "../configs/strings"
import { validationResult } from "express-validator";

var proto = grpc.load('./protos/book.proto')
var ReviewService = new proto.ReviewService(process.env.BOOK_SERVICE || "localhost:9010", grpc.credentials.createInsecure())
var BookService = new proto.BookService(process.env.BOOK_SERVICE || "localhost:9010", grpc.credentials.createInsecure())

proto = grpc.load('./protos/account.proto')
var AccountService = new proto.AccountService(process.env.ACCOUNT_SERVICE || "localhost:9000", grpc.credentials.createInsecure())

export class ReviewController {
    static insert (req, res) {
        const checker = validationResult(req);
        if (!checker.isEmpty()) {
            let result = grpcCode.InvalidArgument;
            result.errors = checker.array();

            return res.status(400).json(result);
        }

        AccountService.GetById({id: req.userId}, (err, user) => {
            if (err)
                return res.json(RPCtoError(err));
            
            BookService.GetById({id: req.params.bookId}, (err, result) => {
                if(err)
                    return res.json(RPCtoError(err));
                
                let body_request = {
                    account_id: req.userId,
                    book_id: req.params.bookId,
                    user_name: user.name,
                    rating: req.body.rating,
                    resume: req.body.resume
                };
                
                ReviewService.Insert(body_request, (err, result) => {
                    if(err)
                        return res.json(RPCtoError(err));
                    
                    let response = JSON.parse(JSON.stringify(grpcCode.OK));
            
                    res.json(response);
                });
            });
        });
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

        ReviewService.Delete({id: req.params.reviewId, account_id: req.userId}, (err, result) => {
            if(err)
                return res.json(RPCtoError(err));
            
            let response = JSON.parse(JSON.stringify(grpcCode.OK));
    
            res.json(response);
        })
    }
}