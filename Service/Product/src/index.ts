import express from 'express';
import EsClient from './boot/Database/elasticsearch';
import { PORT } from './config.json';
import routes from './routes';
const app =  express();
const { json } = express;

async function initServer() {
    try {
        const connect_result = await EsClient.checkConnection();
        if (!connect_result) {
            throw new Error("DATABASE CONNECTION FAILED.")
        }
        console.log('SUCCESSFUL DATABASE CONNECTION');
        
        app.listen(PORT, () => 
            console.log(`SERVER IS LISTENING ON PORT ${PORT}`)
        );
    } catch (err) {
        console.log('DATABASE CONNECTION ERROR: ', err);
    }
}

initServer();

app.use(json());

for (const route in routes) {
    app.use(`/${route}`, routes[route]);
}

// for (const route in routes) {
//     app.use('/', routes[route]);
// }
