"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const elasticsearch_1 = __importDefault(require("../boot/Database/elasticsearch"));
function greeting(req, res) {
    res.send('Welcome to Product Service');
}
async function find(req, res) {
    try {
        // Destructuring request data
        const { id = '' } = req.params;
        const { from_time = 0, to_time = new Date().getTime(), title = '', search_after = '', size = 0, sort_by_fields = '', get_fields = '' } = req.query;
        // const ids = JSON.parse(id_list.toString())
        // const sort_by_list = JSON.parse(sort_by.toString())
        const id_list = id
            ? id.toString().split(',') : [];
        // rating,stock,brand,category,created_at 
        const res_fields = get_fields
            ? get_fields.toString().split(',') : [];
        const sort_fields = sort_by_fields
            ? sort_by_fields.toString().split(',') : [];
        // Mapping request data
        const sort_type_mapping = {
            id: {
                id: 'asc'
            },
            name: {
                name: 'asc'
            },
            stock_count: {
                stock_count: 'asc'
            },
            time: {
                created_at: 'desc'
            }
        };
        const sort_list = sort_fields.map((type) => {
            const sort_type = sort_type_mapping[type];
            return sort_type;
        });
        // Query body
        const query_body = {
            ...(size && { size }),
            query: {
                bool: {
                    must: [
                        ...(from_time ? [{
                                range: {
                                    created_at: {
                                        gte: from_time,
                                        lte: to_time
                                    }
                                }
                            }] : []),
                        ...(id_list.length > 0 ? [{
                                terms: { id: id_list }
                            }] : []),
                        ...(title ? [{
                                match: { title }
                            }] : [])
                    ]
                }
            },
            ...(sort_list.length > 0 && { sort: sort_list }),
            ...(search_after && { search_after })
        };
        console.log('query string: ', JSON.stringify(query_body));
        // Making request to database server
        // Find product
        const result_data = await elasticsearch_1.default
            .findProduct(query_body);
        const { body: { hits: { total: { value: total_value = 0 } = {}, hits = [] } = {} } = {} } = result_data;
        // Format response data
        const res_data = hits.map((item) => {
            const { _source = {} } = item;
            // return default fields
            const { id = '', title = '', price = 0 } = _source;
            const res_item = {
                id, title, price
            };
            for (const field_name of res_fields) {
                Object.assign(res_item, {
                    [field_name]: _source[field_name]
                });
            }
            return res_item;
        });
        res.status(200).json({
            total: total_value,
            data: res_data
        });
    }
    catch (err) {
        res.status(500).json(err);
    }
}
exports.default = { greeting, find };
//# sourceMappingURL=common.js.map