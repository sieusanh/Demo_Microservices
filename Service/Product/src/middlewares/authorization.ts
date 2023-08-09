'use strict'
import { Response, NextFunction } from 'express';
import IGetUserAuthInfoRequest 
    from '../libs/interfaces/auth-info-request';

const adminAuthorization = (req: IGetUserAuthInfoRequest, 
        res: Response, next: NextFunction) => {
            
    if (req.user.role !== 'Admin') {
        res.status(403).json('Admin authorization failed.')
        return
    } 
    
    next();
}

export default { adminAuthorization };
