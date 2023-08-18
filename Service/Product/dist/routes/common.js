"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = require("express");
const controllers_1 = __importDefault(require("../controllers"));
const api_method_1 = __importDefault(require("../libs/mapping/api-method"));
const router = (0, express_1.Router)();
// Get controller data
const { common } = controllers_1.default;
for (const name in common) {
    // Mapping REST API Methods 
    const method = api_method_1.default[name];
    if (name === 'find') {
        router[method](`/${name}/:id`, common[name]);
    }
    // Routing to the specified method names
    router[method](`/${name}`, common[name]);
}
exports.default = router;
//# sourceMappingURL=common.js.map