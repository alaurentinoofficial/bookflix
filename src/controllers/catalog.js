import grpc from 'grpc';
import { grpcCode, RPCtoError } from "../configs/strings";
import { validationResult } from "express-validator";

var proto = grpc.load('./protos/catalog.proto')
var CatalogService = new proto.CatalogService(process.env.CATALOG_SERVICE || "localhost:9020", grpc.credentials.createInsecure())

proto = grpc.load('./protos/account.proto')
var AccountService = new proto.AccountService(process.env.ACCOUNT_SERVICE || "localhost:9000", grpc.credentials.createInsecure())

export class CatalogController {
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
            
            if (user.profile != 'ADMIN')
                return res.status(401).json(JSON.parse(JSON.stringify(grpcCode.PermissionDenied)));

            let body_request = {
                name: req.body.name,
                items: req.body.items.map(x => {
                    return {
                        name: x.name,
                        item_id: x.item_id,
                        image_url: x.image_url,
                    }
                })
            };
            
            CatalogService.Insert(body_request, (err, result) => {
                if(err)
                    return res.json(RPCtoError(err));
                
                let response = JSON.parse(JSON.stringify(grpcCode.OK));
        
                res.json(response);
            });
        });
    }

    static update (req, res) {
        const checker = validationResult(req);
        if (!checker.isEmpty()) {
            let result = grpcCode.InvalidArgument;
            result.errors = checker.array();

            return res.status(400).json(result);
        }

        AccountService.GetById({id: req.userId}, (err, user) => {
            if (err)
                return res.json(RPCtoError(err));
            
            if (user.profile != 'ADMIN')
                return res.status(401).json(JSON.parse(JSON.stringify(grpcCode.PermissionDenied)));

            let body_request = {
                id: req.params.catalogId,
                name: req.body.name,
                items: req.body.items.map(x => {
                    return {
                        name: x.name,
                        item_id: x.item_id,
                        image_url: x.image_url,
                    }
                })
            };
            
            CatalogService.Update(body_request, (err, result) => {
                if(err)
                    return res.json(RPCtoError(err));
                
                let response = JSON.parse(JSON.stringify(grpcCode.OK));
        
                res.json(response);
            });
        });
    }

    static get (req, res) {
        CatalogService.Get({}, (err, result) => {
            if(err)
                return res.json(RPCtoError(err));
            
            let response = JSON.parse(JSON.stringify(grpcCode.OK));
            response.body = result.catalogs
    
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

        CatalogService.GetById({id: req.params.catalogId}, (err, result) => {
            if(err)
                return res.json(RPCtoError(err));
            
            let response = JSON.parse(JSON.stringify(grpcCode.OK));
            response.body = result
    
            res.json(response);
        });
    }

    static delete (req, res) {
        const checker = validationResult(req);
        if (!checker.isEmpty()) {
            let result = grpcCode.InvalidArgument;
            result.errors = checker.array();

            return res.status(400).json(result);
        }

        AccountService.GetById({id: req.userId}, (err, user) => {
            if (err)
                return res.json(RPCtoError(err));
            
            if (user.profile != 'ADMIN')
                return res.status(401).json(JSON.parse(JSON.stringify(grpcCode.PermissionDenied)));
            
            CatalogService.Delete({id: req.params.catalogId}, (err, result) => {
                if(err)
                    return res.json(RPCtoError(err));
                
                let response = JSON.parse(JSON.stringify(grpcCode.OK));
        
                res.json(response);
            });
        });
    }
}