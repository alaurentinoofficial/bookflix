exports.grpcCode = {
    OK: {code: 0, message: "Succefully"},
    Canceled: {code: 1, message: "Canceled"},
    Unknown: {code: 2, message: "Unknow"},
    InvalidArgument: {code: 3, message: "Invalid Argument"},
    DeadlineExceeded: {code: 4, message: "Deadline Exceeded"},
    NotFound: {code: 5, message: "Not Found"},
    AlreadyExists: {code: 6, message: "Already Exists"},
    PermissionDenied: {code: 7, message: "Permission Denied"},
    ResourceExhausted: {code: 8, message: "ResourceE"},
    FailedPrecondition: {code: 9, message: "Failed Precondition"},
    Aborted: {code: 10, message: "Aborted"},
    OutOfRange: {code: 11, message: "Out of Range"},
    Unimplemented: {code: 12, message: "Unimplemented"},
    Internal: {code: 13, message: "Internal error"},
    Unavailable: {code: 14, message: "Unavailavble"},
    DataLoss: {code: 15, message: "Data Loss"},
    Unauthenticated: {code: 16, message: "Unauthenticated"},
}

exports.RPCtoError = (str) => {
    return {code: str.code, message: str.details}
}