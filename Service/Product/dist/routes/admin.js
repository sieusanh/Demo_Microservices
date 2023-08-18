"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = require("express");
const controllers_1 = __importDefault(require("../controllers"));
const middlewares_1 = __importDefault(require("../middlewares"));
const api_method_1 = __importDefault(require("../libs/mapping/api-method"));
const router = (0, express_1.Router)();
// Get controller aggregate
const { admin } = controllers_1.default;
// Get middlewares
const { userAuthentication, adminAuthorization } = middlewares_1.default;
for (const name in admin) {
    // Mapping REST API Methods 
    const method = api_method_1.default[name];
    // Routing to the specified method names
    router[method](`/${name}`, userAuthentication, adminAuthorization, admin[name]);
}
exports.default = router;
//# sourceMappingURL=admin.js.map