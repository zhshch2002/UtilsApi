import axios from 'axios'
import cheerio from 'cheerio';
import sharp, {FitEnum} from 'sharp';
import express = require('express');

const app = express();
export = app

app.get("/proxy", async (req, resp) => {
    let u = req.query.url as string;
    if (u == undefined) {
        resp.send("miss params url")
        return
    }
    let r = await axios.get(u, {
        responseType: "arraybuffer"
    })
    resp.statusCode = r.status;
    resp.type(r.headers["content-type"])
    resp.send(Buffer.from(r.data))
})

/*
resize=300,300,cover 调整大小（最后一个可选）
rotate=180           旋转图像
flip                 上下翻转
flop                 左右翻转
quality=80           图像质量
 */
app.get("/img", img)

app.get("/favicon", async (req, resp) => {
    let u = req.query.url as string;
    if (u == undefined) {
        resp.send("miss params url")
        return
    }
    let r = await axios.get(u)
    let h = cheerio.load(r.data)
    let icon = new URL("/favicon.ico", u).href
    h("link[rel]").each(function (i, e) {
        if (e.attribs.href != undefined && e.attribs.rel.indexOf("icon") != -1) {
            icon = new URL(e.attribs.href, u).href
            return false
        }
        return true
    })
    resp.locals.url = icon

    await img(req, resp)
})

async function img(req: express.Request, resp: express.Response) {
    let u = resp.locals.url as string || req.query.url as string;
    console.log(u)
    if (u == undefined) {
        resp.send("miss params url")
        return
    }
    let r = await axios.get(u, {
        responseType: "arraybuffer"
    })
    resp.statusCode = r.status;

    if (r.status != 200) {
        resp.type(r.headers["content-type"])
        resp.send(Buffer.from(r.data))
        return
    }
    let img = sharp(Buffer.from(r.data))
    if (req.query.resize != undefined) {
        let rv = (req.query.resize as string).split(",")
        let f = (rv.length >= 3 ? rv[2] : "cover") as "contain" | "cover" | "fill" | "inside" | "outside"
        img = img.resize({
            width: rv[0] == "" ? undefined : parseInt(rv[0]),
            height: rv[0] == "" ? undefined : parseInt(rv[1]),
            fit: f,
            background: {r: 0, g: 0, b: 0, alpha: 0}
        })
    }
    if (req.query.rotate != undefined) {
        img = img.rotate(parseInt(req.query.rotate as string), {
            background: {r: 0, g: 0, b: 0, alpha: 0}
        })
    }
    if (req.query.flip != undefined) {
        img = img.flip()
    }
    if (req.query.flop != undefined) {
        img = img.flop()
    }

    img = img.webp({
        quality: req.query.quality != undefined ? parseInt(req.query.quality as string) : undefined,
    })
    img.toBuffer().then(b => {
        resp.type("image/webp")
        resp.send(b)
    }).catch(err => {
        console.error(err)
        resp.status(500).send("server error")
    })
}

app.get("/htmlProxy", async (req, resp) => {
    let u = req.query.url as string;
    if (u == undefined) {
        resp.send("miss params url")
        return
    }
    let r = await axios.get(u)
    resp.statusCode = r.status;
    resp.type(r.headers["content-type"])

    let h = cheerio.load(r.data)
    h("a[href]").each(function (i, e) {
        if (e.attribs.href != undefined && !e.attribs.href.startsWith("javascript:"))
            e.attribs.href = new URL(e.attribs.href, u).href
    })
    h("img[src]").each(function (i, e) {
        let a = e.attribs.src
        if (a != undefined && !a.startsWith("javascript:"))
            e.attribs.src = new URL(a, u).href
    })
    h("script[src]").each(function (i, e) {
        let a = e.attribs.src
        if (a != undefined && !a.startsWith("javascript:"))
            e.attribs.src = new URL(a, u).href
    })
    h("link[href]").each(function (i, e) {
        let a = e.attribs.href
        if (a != undefined && !a.startsWith("javascript:"))
            e.attribs.href = new URL(a, u).href
    })

    resp.send(h.html())
})