import express from 'express'
import bodyParser from 'body-parser'

// import gRPC from './configs/grpc'
import { Router } from './configs/router'
import { AllowCrossDomain } from './middlewares/cors'
import { grpcCode } from "./configs/strings"
import * as dotenv from "dotenv"

dotenv.config();

//======================
//   Express CONFIG
//======================
var app = express()
app.use(AllowCrossDomain)
app.use(bodyParser.urlencoded({'extended':'true'}))
app.use(bodyParser.json())

app.use(function (err, req, res, next) {
  console.error(err.stack)
  res.status(500).json(grpcCode.Internal)
})

app.set('port', process.env.PORT || 8080)
//======================

//======================
//     APP CONFIG
//======================
Router(app)
// gRPC.init(8080)
//======================

exports.Server = app