import express = require('express');
import proxy from "./v1/proxy"
import basic from "./v1/basic"

const app = express();
// const mc = memjs.Client.create()
// let cache = (req: express.Request, res: express.Response, next: express.NextFunction) => {
//     const key = req.url
//     mc.get(key, (err, val) => {
//         if (err == null && val != null) {
//             res.send('from cache')
//         } else {
//             res.locals.sendResponse = res.send
//             res.send = function (body: any): any {
//                 mc.set(key, body, {expires: 0}, (err, reply) => {
//                     res.setHeader("uapi-cache", "hit")
//                     res.locals.sendResponse(body)
//                 })
//             }
//             next()
//         }
//     })
// }

const v1api = express();
v1api.use("/", proxy)
v1api.use("/", basic)
app.use("/v1",  v1api)

app.all('*', function (req, res, next) {
    res.header('Access-Control-Allow-Origin', '*');
    res.header('Access-Control-Allow-Headers', 'Content-Type, Content-Length, Authorization, Accept, X-Requested-With');
    res.header('Access-Control-Allow-Methods', 'PUT, POST, GET, DELETE, OPTIONS');
    if (req.method == 'OPTIONS') {
        res.send(200);
    } else {
        next();
    }
});

app.get('/', (req, res) => {
    res.redirect("https://github.com/zhshch2002/UtilsApi");
});

app.listen(4000, () => {
    console.log('listening on port 4000!');
});