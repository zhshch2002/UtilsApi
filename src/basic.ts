import axios from 'axios'
import cheerio from 'cheerio';
import sharp, {FitEnum} from 'sharp';
import express = require('express');

const app = express();
export = app

app.get("/ip", async (req, resp) => {
    resp.send(req.ip)
})

app.get("/time", async (req, resp) => {
    let t = new Date();
    resp.send({
        "unix": t.valueOf(),
        "year": t.getFullYear(),
        "month": t.getMonth() + 1,
        "day": t.getDate(),
        "week": t.getDay(),
        "hour": t.getHours(),
        "minute": t.getMinutes(),
        "second": t.getSeconds(),
        "milliseconds": t.getMilliseconds(),
        "tz": t.getTimezoneOffset(),
        "utc": t.toUTCString(),
        "utc_year": t.getUTCFullYear(),
        "utc_month": t.getUTCMonth() + 1,
        "utc_day": t.getUTCDate(),
        "utc_week": t.getUTCDay(),
        "utc_hour": t.getUTCHours(),
        "utc_minute": t.getUTCMinutes(),
        "utc_second": t.getUTCSeconds(),
        "utc_milliseconds": t.getUTCMilliseconds(),
    })
})