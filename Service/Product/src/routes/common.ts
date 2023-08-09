import { Router } from 'express';
import controllers from '../controllers';
import api_method_mapping from '../libs/mapping/api-method';

const router = Router();

// Get controller data
const { common } = controllers;

for (const name in common) {
    // Mapping REST API Methods 
    const method = api_method_mapping[name];

    if (name === 'find') {
        router[method](`/${name}/:id`, common[name]);
    }

    // Routing to the specified method names
    router[method](`/${name}`, common[name]);
}

export default router;