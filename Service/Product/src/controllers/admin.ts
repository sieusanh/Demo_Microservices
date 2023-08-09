import axios from 'axios';
import EsClient from '../boot/Database/elasticsearch';
import {Request, Response} from 'express';

async function aggByCategory(req: Request, res: Response)  {``
    try {
        // Destructuring request data
        const {
            from_time = 0, to_time = new Date().getTime(),
        } = req.query;
        
        const TERM_AGG_MAX = 100;
        // Query body
        const query_body = {
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
                        }] : [])
                    ]
                }
            },
            size: 0,
            aggs: {
                category_terms: {
                    terms: {
                        field: 'category.keyword',
                        size: TERM_AGG_MAX
                    },
                    aggs: {
                        brand_terms: {
                            terms: {
                                field: 'brand.keyword'
                            }
                        }
                    }
                }
            }
        };
        console.log('query string: ', JSON.stringify(query_body))

        // Making request to database server
        // Aggregate product
        const {
            body: {
                aggregations: {
                    category_terms: {
                        buckets: category_buckets = []
                    } = {}
                } = {}
            } = {}
        } = await EsClient
            .findProduct(query_body);

        // Format response data
        const data = category_buckets.map(e => {
            const {
                key: category_name = '',
                doc_count: category_count = 0,
                brand_terms = {}
            } = e;

            const {
                buckets: brand_buckets = []
            } = brand_terms;

            const brand_data = brand_buckets.map((e: any) => ({
                name: e.key,
                count: e.doc_count
            }));

            return { 
                name: category_name,
                count: category_count,
                brand_data
            };
        })
        
        res.status(200).json({
            data
        });
    } catch (err) {
        res.status(500).json(err);
    }
}

async function aggByRating(req: Request, res: Response)  {
    try {
        // Destructuring request data
        const {
            from_time = 0, to_time = new Date().getTime(),
        } = req.query;
        
        // Query body
        const query_body = {
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
                        }] : [])
                    ]
                }
            },
            size: 0,
            aggs: {
                products_over_rating: {
                    histogram: {
                        field: 'rating',
                        interval: 0.1
                    }
                }
            }
        };
        console.log('query string: ', JSON.stringify(query_body))

        // Making request to database server
        // Aggregate product
        const {
            body: {
                aggregations: {
                    products_over_rating: {
                        buckets = []
                    } = {}
                } = {}
            } = {}
        } = await EsClient
            .findProduct(query_body);

        // Format response data
        const data = buckets.map(e => ({
            number: e.key,
            count: e.doc_count
        }));
        
        res.status(200).json({
            data
        });
    } catch (err) {
        res.status(500).json(err);
    }
}

async function aggByBrand(req: Request, res: Response)  {
    try {
        // Destructuring request data
        const {
            from_time = 0, to_time = new Date().getTime(),
        } = req.query;
        
        const TERM_AGG_MAX = 100;
        // Query body
        const query_body = {
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
                        }] : [])
                    ]
                }
            },
            size: 0,
            aggs: {
                brand_terms: {
                    terms: {
                        field: 'brand.keyword',
                        size: TERM_AGG_MAX
                    },
                    aggs: {
                        product_terms: {
                            terms: {
                                field: 'title.keyword'
                            }
                        },
                        category_terms: {
                            terms: {
                                field: 'category.keyword'
                            }
                        }
                    }
                }
            }
        };
        console.log('query string: ', JSON.stringify(query_body))

        // Making request to database server
        // Aggregate product
        const {
            body: {
                aggregations: {
                    brand_terms: {
                        buckets: brand_buckets = []
                    } = {}
                } = {}
            } = {}
        } = await EsClient
            .findProduct(query_body);

        // Format response data
        const data = brand_buckets.map(e => {
            const {
                key: brand_name = '',
                doc_count: brand_count = 0,
                product_terms = {},
                category_terms = {}
            } = e;

            const {
                buckets: product_buckets = []
            } = product_terms;

            const {
                buckets: category_buckets = []
            } = category_terms;

            const product_data = product_buckets.map((e: any) => ({
                name: e.key
            }));

            const category_data = category_buckets.map((e: any) => ({
                name: e.key,
                count: e.doc_count
            }));

            return { 
                name: brand_name,
                count: brand_count,
                product_data,
                category_data
            };
        })
        
        res.status(200).json({
            data
        });
    } catch (err) {
        res.status(500).json(err);
    }
}

async function aggByTime(req: Request, res: Response)  {
    try {
        // Destructuring request data
        const {
            from_time = 0, to_time = new Date().getTime(),
            interval = '1d' // 1d = date | day | hour
        } = req.query;

        const time_mapping = {
            histogram_field: interval == 'date' 
                ? 'date_histogram' : 'histogram',
            histogram_value: interval == 'date'
                ? 'created_at' : '',
            interval: '1d',
            min: from_time,
            max: to_time
        }
        
        // Query body
        const query_body = {
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
                        }] : [])
                    ]
                }
            },
            size: 0,
            aggs: {
                products_over_time: {
                    date_histogram: {
                        field: 'created_at',
                        interval,
                        min_doc_count: 0,
                        extended_bounds: {
                            min: from_time,
                            max: to_time
                        }
                    }
                }
            }
        };
        console.log('query string: ', JSON.stringify(query_body))

        // Making request to database server
        // Aggregate product
        const {
            body: {
                aggregations: {
                    products_over_time: {
                        buckets = []
                    } = {}
                } = {}
            } = {}
        } = await EsClient
            .findProduct(query_body);

        // Format response data
        const data = buckets.map(e => ({
            timestamp: e.key,
            count: e.doc_count
        }));
        
        res.status(200).json({
            data
        });
    } catch (err) {
        res.status(500).json(err);
    }
}   

async function populatingData(req: Request, res: Response) {
    try {
        const dummy_res = await axios.get('https://dummyjson.com/products?limit=100');
        // console.log('data: ', dummy_res.data.products)
        console.log('len: ', dummy_res.data.products.length)
        const dummy_product_list = dummy_res.data
                                .products.map((e: any) => {
                                    const {
                                        discountPercentage: discount,
                                        ...others
                                    } = e;
                                    return {
                                        discount, ...others
                                    }
                                });
        const created_at = new Date().getTime();

        for (const product of dummy_product_list) {                         
            const {
                description, thumbnail, images, 
                id, discountPercentage: discount,
                ...others
            } = product; 
            
            // Making request to database server
            EsClient.createProduct(
                id, { id, discount, created_at, ...others}
            );
        }
        
        // Format response data
        res.status(200).json({
            message: 'Success'
        })
    } catch (err) {
        console.log("Loi gi: ", err)
        res.status(500).json(err);
    }
}

export default { aggByCategory, aggByRating, 
    aggByBrand, aggByTime }
