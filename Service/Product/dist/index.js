"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const express_1 = __importDefault(require("express"));
const elasticsearch_1 = __importDefault(require("./boot/Database/elasticsearch"));
const config_json_1 = require("./config.json");
const routes_1 = __importDefault(require("./routes"));
const app = (0, express_1.default)();
const { json } = express_1.default;
async function initServer() {
    try {
        // const res = await axios.post(
        //     'http://localhost:8001/services/', {
        //     name: 'product-service',
        //     path: '/product-service',
        //     url: 'http://192.168.1.110:9002'
        // });
        // console.log('SUCCESSFUL KONG CONNECTION');
        const connect_result = await elasticsearch_1.default.checkConnection();
        if (!connect_result) {
            throw new Error("DATABASE CONNECTION FAILED.");
        }
        console.log('SUCCESSFUL DATABASE CONNECTION');
        app.listen(config_json_1.PORT, () => console.log(`SERVER IS LISTENING ON PORT ${config_json_1.PORT}`));
    }
    catch (err) {
        console.log('DATABASE CONNECTION ERROR: ', err);
    }
}
initServer();
app.use(json());
for (const route in routes_1.default) {
    app.use(`/${route}`, routes_1.default[route]);
}
// for (const route in routes) {
//     app.use('/', routes[route]);
// }
//# sourceMappingURL=index.js.map