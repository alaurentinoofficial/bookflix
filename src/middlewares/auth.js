import * as jwt from "jsonwebtoken"
import { grpcCode } from "../configs/strings"

exports.Authenticate = (req, res, next) => {
    const authHeader = req.headers['authorization'];
    const token = authHeader.split(' ')[1] || authHeader;

    if (token == null)
        return res.status(401).json(grpcCode.Unauthenticated);

    jwt.verify(token, process.env.TOKEN_SECRET, (err, user) => {
        if (err)
            return res.status(401).json(grpcCode.Unauthenticated);

        req.userId = user.id;
        next();
    });
}