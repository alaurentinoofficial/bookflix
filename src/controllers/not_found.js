import { grpcCode } from '../configs/strings'

exports.NotFound = (req, res) => res.json(grpcCode.NotFound)