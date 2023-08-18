"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const mapping = {
    // Common methods
    greeting: 'get',
    find: 'get',
    // User methods
    create: 'post',
    updateByQuery: 'patch',
    updateStockById: 'patch',
    deleteById: 'delete',
    // Admin methods
    aggByCategory: 'get',
    aggByRating: 'get',
    aggByBrand: 'get',
    aggByTime: 'get',
    populatingData: 'post'
};
exports.default = mapping;
//# sourceMappingURL=api-method.js.map