'use strict';
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const jsonwebtoken_1 = __importDefault(require("jsonwebtoken"));
const config_json_1 = require("../config.json");
const userAuthentication = async (req, res, next) => {
    try {
        // Check if account authorized
        const author_header = req.headers?.authorization || '';
        let statusCode, message;
        const splitted = author_header.split(' ');
        if (!author_header || splitted.length != 2) {
            statusCode = 403; // Forbidden
            message = 'Unauthorized.';
            return res.status(statusCode).json({
                user: null, message
            });
        }
        // Verify access token
        const accessToken = splitted[1] || '';
        jsonwebtoken_1.default.verify(accessToken, config_json_1.ACCESS_TOKEN_SECRET, (err, decodedToken) => {
            // Verify failed.
            if (err) {
                statusCode = 500; // Internal Server Error
                message = 'Access Token is not valid!';
                return res.status(500).json({
                    user: null, message
                });
            }
            // Verify passed.
            const { email = '', name = '', role = '' } = decodedToken;
            req.user = { email, name, role };
            next();
        });
    }
    catch (err) {
        return res.status(403).json({
            user: null, message: 'Invalid access token.'
        });
    }
};
exports.default = { userAuthentication };
//# sourceMappingURL=authentication.js.map