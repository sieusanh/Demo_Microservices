import { Request } from 'express';
import User from '../../models/user';

interface IGetUserAuthInfoRequest extends Request {
    user: User 
}

export default IGetUserAuthInfoRequest;