"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
// import { uid } from 'uid';
const elasticsearch_1 = __importDefault(require("../boot/Database/elasticsearch"));
const config_json_1 = require("../config.json");
async function create(req, res) {
    try {
        // Destructuring request data
        const create_body = req.body;
        const id = create_body?.id || '';
        const created_at = new Date().getTime();
        // Making request to database server
        // Create a new product
        // const id = uid(12);
        await elasticsearch_1.default.createProduct(id, { ...create_body });
        // Format response data
        const res_time = created_at +
            config_json_1.CLIENT_UTC_OFFSET * config_json_1.AN_HOUR_MILLISECS;
        res.status(200).json({
            message: 'Success',
            data: { id, created_at: res_time }
        });
    }
    catch (err) {
        res.status(500).json(err);
    }
}
async function updateByQuery(req, res) {
    try {
        // Destructuring request data
        const { id = '' } = req.params;
        const fields = req.body;
        const id_list = id
            ? id.toString().split(',') : [];
        let update_string = "";
        for (const [k, v] of Object.entries(fields)) {
            update_string += `ctx._source[\"${k}\"]=${v};`;
        }
        // Query body
        const query_body = {
            query: {
                ...(id_list.length > 0 && {
                    terms: {
                        id: id_list
                    }
                })
            },
            script: {
                source: update_string,
                lang: 'painless'
            }
        };
        // Making request to database server
        await elasticsearch_1.default
            .updateProduct(query_body);
        // Response formatting
        res.status(200).json({
            message: 'Success.'
        });
        return;
    }
    catch (err) {
        res.status(500).json(err);
    }
}
async function updateStockById(req, res) {
    try {
        // Destructuring request data
        const { id = '' } = req.params;
        const { change = 0 } = req.body;
        const update_string = `ctx._source["stock"] += ${change}`;
        // Query body
        const query_body = {
            script: {
                source: update_string,
                lang: 'painless'
            }
        };
        // Making request to database server
        await elasticsearch_1.default
            .updateStock(id, query_body);
        // Response formatting
        res.status(200).json({
            message: 'Success.'
        });
        return;
    }
    catch (err) {
        res.status(500).json(err);
    }
}
async function deleteById(req, res) {
    try {
        // Destructuring request data
        const { id = '' } = req.params;
        // Making request to database server
        const result_data = await elasticsearch_1.default
            .removeProduct(id);
        const { deletedCount = 1 } = result_data;
        // Response formatting
        const plural = deletedCount > 1 ? true : false;
        res.status(200).json({
            message: `Deleted ${deletedCount} record${plural ? 's' : ''}.`
        });
    }
    catch (err) {
        res.status(500).json(err);
    }
}
exports.default = { create, updateByQuery, updateStockById, deleteById };
//# sourceMappingURL=user.js.map