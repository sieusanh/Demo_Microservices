import { Router } from 'express';
import controllers from '../controllers';
import middlewares from '../middlewares';
import api_method_mapping from '../libs/mapping/api-method';

const router = Router();

// Get controller data
const { user } = controllers;

// Get middlewares
const { userAuthentication } = middlewares;

for (const name in user) {
    // Mapping REST API Methods 
    const method = api_method_mapping[name];

    if (name === 'create') {
        router[method](`/${name}`, userAuthentication, user[name]);
        continue;
    }

    // Routing to the specified method names
    router[method](`/${name}/:id`, userAuthentication, user[name]);
}

export default router;
