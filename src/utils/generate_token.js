import * as jwt from "jsonwebtoken"

exports.GenerateToken = (id) => {
    return jwt.sign(id, process.env.TOKEN_SECRET, { expiresIn: '1800s' });
}