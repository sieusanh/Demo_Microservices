"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const { Client, errors } = require('@elastic/elasticsearch');
const config_json_1 = require("../../config.json");
// Class elasticsearch client 
class EsService {
    constructor() {
        this.EsClient = new Client({ node: config_json_1.DATABASE_HOST });
    }
    async checkConnection() {
        try {
            const result = await this.EsClient.ping();
            const status = result?.statusCode || 404;
            if (status != 200) {
                return false;
            }
            return true;
        }
        catch (err) {
            console.log('err: ', err);
            return false;
        }
    }
    createProduct(id, body) {
        const promise = this.EsClient.create({
            index: config_json_1.INDEX_PRODUCT, id, body
        });
        return promise;
    }
    findProduct(body) {
        const promise = this.EsClient.search({
            index: config_json_1.INDEX_PRODUCT, body
        });
        return promise;
    }
    updateProduct(body) {
        const promise = this.EsClient.update({
            index: config_json_1.INDEX_PRODUCT, id: null, body
        });
        return promise;
    }
    updateStock(id, body) {
        const promise = this.EsClient.update({
            index: config_json_1.INDEX_PRODUCT, id, body
        });
        return promise;
    }
    removeProduct(id) {
        const promise = this.EsClient.delete({
            index: config_json_1.INDEX_PRODUCT, id
        });
        return promise;
    }
    aggregateProduct(id, body) {
        const promise = this.EsClient.search({
            index: config_json_1.INDEX_PRODUCT, body
        });
        return promise;
    }
}
// Declare Elasticsearch client instance
const EsClient = new EsService();
exports.default = EsClient;
//# sourceMappingURL=elasticsearch.js.map