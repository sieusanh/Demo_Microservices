import { Router } from 'express';
import controllers from '../controllers';
import middlewares from '../middlewares';
import api_method_mapping from '../libs/mapping/api-method';

const router = Router();

// Get controller aggregate
const { admin } = controllers;

// Get middlewares
const { userAuthentication, adminAuthorization } = middlewares;

for (const name in admin) {
    // Mapping REST API Methods 
    const method = api_method_mapping[name];

    // Routing to the specified method names
    router[method](`/${name}`, 
        userAuthentication, adminAuthorization, admin[name]);
}

export default router;