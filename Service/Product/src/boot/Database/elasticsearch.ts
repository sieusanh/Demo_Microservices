const { Client, errors } = require('@elastic/elasticsearch')
import { DATABASE_HOST, INDEX_PRODUCT } from '../../config.json';

// Define elasticsearch client 
class EsService {
    private EsClient: any
    constructor() {
        this.EsClient = new Client({ node: DATABASE_HOST });
    }

    async checkConnection(): Promise<boolean> {
        try {
            const result = await this.EsClient.ping();
            const status = result?.statusCode || 404;
            if (status != 200) {
                return false;
            }
            return true;
        } catch (err) {
            console.log('err: ', err)
            return false;
        }
    
    }

    createProduct(id: string, body: {}) {
        const promise = this.EsClient.create({
            index: INDEX_PRODUCT, id, body
        });
        
        return promise;
    }
    
    findProduct(body: {}) {
        const promise = this.EsClient.search({
            index: INDEX_PRODUCT, body
        });
        
        return promise;
    }

    updateProduct(body: {}) {
        const promise = this.EsClient.update({
            index: INDEX_PRODUCT, id: null, body
        });
        
        return promise;
    }

    updateStock(id: string, body: {}) {
        const promise = this.EsClient.update({
            index: INDEX_PRODUCT, id, body
        });
        
        return promise;
    }
    
    removeProduct(id: string) {
        const promise = this.EsClient.delete({
            index: INDEX_PRODUCT, id
        });
        
        return promise;
    }
    
    aggregateProduct(id: string, body: {}) {
        const promise = this.EsClient.search({
            index: INDEX_PRODUCT, body
        });
        
        return promise;
    }
    
}
    
const EsClient = new EsService();
export default EsClient;