"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const authentication_1 = __importDefault(require("./authentication"));
const authorization_1 = __importDefault(require("./authorization"));
const export_object = { ...authentication_1.default, ...authorization_1.default };
exports.default = export_object;
//# sourceMappingURL=index.js.map