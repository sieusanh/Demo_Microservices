"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const common_1 = __importDefault(require("./common"));
const user_1 = __importDefault(require("./user"));
const admin_1 = __importDefault(require("./admin"));
const export_object = { common: common_1.default, user: user_1.default, admin: admin_1.default };
exports.default = export_object;
//# sourceMappingURL=index.js.map