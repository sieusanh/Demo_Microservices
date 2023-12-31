'use strict'
import jwt from 'jsonwebtoken';
import { Response, NextFunction } from 'express';
import { ACCESS_TOKEN_SECRET } from '../config.json';

import IGetUserAuthInfoRequest 
    from '../libs/interfaces/auth-info-request';

const userAuthentication = async (req: IGetUserAuthInfoRequest, 
        res: Response, next: NextFunction) => {
    try {
        // Check if account authorized
        const author_header = req.headers?.authorization || '';
        let statusCode: number, message: string;
        const splitted = author_header.split(' ');
        if (!author_header || splitted.length != 2) {
            statusCode = 403; // Forbidden
            message = 'Unauthorized.';
            return res.status(statusCode).json({ 
                user: null, message
            })
        }

        // Verify access token
        const accessToken = splitted[1] || '';
        jwt.verify(accessToken, ACCESS_TOKEN_SECRET, (err, 
                decodedToken: jwt.JwtPayload) => {

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
            next()
        })
        
    } catch (err) {
        return res.status(403).json({ 
            user: null, message: 'Invalid access token.' 
        });   
    }
}

export default { userAuthentication };
