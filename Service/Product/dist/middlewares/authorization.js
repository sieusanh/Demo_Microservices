'use strict';
Object.defineProperty(exports, "__esModule", { value: true });
const adminAuthorization = (req, res, next) => {
    if (req.user.role !== 'Admin') {
        res.status(403).json('Admin authorization failed.');
        return;
    }
    next();
};
exports.default = { adminAuthorization };
//# sourceMappingURL=authorization.js.map