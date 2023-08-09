// import { uid } from 'uid';
import EsClient from '../boot/Database/elasticsearch';
import {Request, Response} from 'express';
import { CLIENT_UTC_OFFSET, 
    AN_HOUR_MILLISECS } from '../config.json';

async function create(req: Request, res: Response)  {
    try {
        // Destructuring request data
        const create_body = req.body;
        const id = create_body?.id || '';
        const created_at = new Date().getTime();
        // Making request to database server
        // Create a new product
        // const id = uid(12);
        await EsClient.createProduct(
            id, { ...create_body }
        );

        // Format response data
        const res_time = created_at + 
            CLIENT_UTC_OFFSET * AN_HOUR_MILLISECS;
        res.status(200).json({
            message: 'Success',
            data: { id, created_at: res_time }
        })
    } catch (err) {
        res.status(500).json(err);
    }
}

async function updateByQuery(req: Request, res: Response) {
    try { 
        // Destructuring request data
        const { id = '' } = req.params;
        const fields = req.body;

        const id_list = id
        ? id.toString().split(',') : [];

        let update_string: string = "";
        for (const [k, v] of Object.entries(fields)) {
            update_string += `ctx._source[\"${k}\"]=${v};`
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
        }
        
        // Making request to database server
        await EsClient
            .updateProduct(query_body);

        // Response formatting
        res.status(200).json({
            message: 'Success.'
        })
        return;
    } catch (err)  {
        res.status(500).json(err);
    }
}

async function updateStockById(req: Request, res: Response) {
    try { 
        // Destructuring request data
        const { id = '' } = req.params;
        const {
            change = 0
        } = req.body;         
        const update_string = `ctx._source["stock"] += ${change}`;
        
        // Query body
        const query_body = {
            script: {
                source: update_string,
                lang: 'painless'
            }
        }
        
        // Making request to database server
        await EsClient
            .updateStock(id, query_body);

        // Response formatting
        res.status(200).json({
            message: 'Success.'
        })
        return;
    } catch (err)  {
        res.status(500).json(err);
    }
}

async function deleteById(req: Request, res: Response) {
    try { 
        // Destructuring request data
        const { id = '' } = req.params;

        // Making request to database server
        const result_data = await EsClient
            .removeProduct(id);
        const { deletedCount = 1 } = result_data;

        // Response formatting
        const plural = deletedCount > 1 ? true : false;
        res.status(200).json({
            message: `Deleted ${deletedCount} record${plural ? 's' : ''}.`
        })
    } catch (err)  {
        res.status(500).json(err);
    }
}

export default { create, updateByQuery, updateStockById, deleteById }
