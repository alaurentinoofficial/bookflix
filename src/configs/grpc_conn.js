import grpc from 'grpc';

exports.AccountService = grpc.load('./protos/account.proto').AccountService(
        process.env.ACCOUNT_SERVICE || "localhost:9000"
        , grpc.credentials.createInsecure()
    )

exports.ReviewService = grpc.load('./protos/book.proto')
    .ReviewService(
        process.env.BOOK_SERVICE || "localhost:9010"
        , grpc.credentials.createInsecure()
    )

exports.BookService = grpc.load('./protos/book.proto')
    .BookService(
        process.env.BOOK_SERVICE || "localhost:9010"
        , grpc.credentials.createInsecure()
    )
