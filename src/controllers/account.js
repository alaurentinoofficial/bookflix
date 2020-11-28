import grpc from 'grpc'
import { grpcCode, RPCtoError } from "../configs/strings"
import { GenerateToken } from '../utils/generate_token'

let proto = grpc.load('./protos/account.proto')
var AccountService = new proto.AccountService(process.env.ACCOUNT_SERVICE || "localhost:9000", grpc.credentials.createInsecure())

export class AccountController {
    static login (req, res) {
        AccountService.Auth({email: req.body.email, password: req.body.password}, (err, userResult) => {
            if(err)
                return res.json(RPCtoError(err))
            var result = JSON.parse(JSON.stringify(grpcCode.OK))
            result.token = GenerateToken({id: userResult.id})
    
            res.json(result)
        })
    }

    static register (req, res) {
        AccountService.Insert(req.body, (err, token) => {
            if(err)
                return res.json(RPCtoError(err))

            var result = JSON.parse(JSON.stringify(grpcCode.OK))
            result.token = token.token
    
            res.json(result)
        })
    }

    static getById (req, res) {
        AccountService.GetById({id: req.userId}, (err, user) => {
            if (err)
                return res.json(RPCtoError(err));
            
            res.json(user);
        });
    }
}