const express = require('express');
const bodyParser = require('body-parser');
const next = require('next');
const { createProxyMiddleware } = require('http-proxy-middleware');

const port = parseInt(process.env.PORT, 10) || 3000;
const dev = process.env.NODE_ENV !== 'production';
const app = next({ dev });
const handle = app.getRequestHandler();
const googleTrendApi = require('google-trends-api');

app.prepare().then(() => {
  const server = express();

  server.use(
    bodyParser.urlencoded({
      extended: true,
    }),
  );

  server.use(bodyParser.json());

  // Google Autocomplete APIを叩く
  server.use(
    '/complete/search',
    createProxyMiddleware({
      target: 'https://suggestqueries.google.com',
      changeOrigin: true,
    }),
  );

  // Ingressのhealthcheck用
  server.get('/health', (req, res) => {
    return res.sendStatus(200);
  });

  // Google Trend API実行
  server.post('/google/daily/trend', (req, res) => {
    const { trendDate, geo } = req.body;

    googleTrendApi.dailyTrends(
      {
        trendDate,
        geo,
      },
      function(err, results) {
        if (err) {
          return res.json({
            result: false,
            error: err,
          });
        } else {
          return res.json({
            result: true,
            data: JSON.parse(results),
          });
        }
      },
    );
  });

  server.all('*', (req, res) => {
    return handle(req, res);
  });

  server.listen(port, err => {
    if (err) throw err;
    console.log(`> Ready on http://localhost:${port}`);
  });
});
