// Minimal Lambda@Edge origin-request handler to demonstrate SSR response.
// For real Next.js SSR, replace logic with a bundle that renders pages.

exports.handler = async (event, context, callback) => {
  const cf = event.Records[0].cf;
  const request = cf.request;
  const { uri, method } = request;

  const isStatic = uri.startsWith('/_next/') || uri.includes('.') || uri.startsWith('/api/');

  // Only intercept GET HTML requests for non-static paths
  if (method === 'GET' && !isStatic) {
    const body = `<!doctype html>
    <html lang="en">
    <head>
      <meta charset="utf-8" />
      <meta name="viewport" content="width=device-width, initial-scale=1" />
      <title>Edge SSR Demo</title>
      <style>body{font-family:system-ui, -apple-system, Segoe UI, Roboto, Ubuntu, Cantarell, Noto Sans, Helvetica, Arial, Apple Color Emoji, Segoe UI Emoji; padding:2rem; line-height:1.6}</style>
    </head>
    <body>
      <h1>Lambda@Edge SSR デモ</h1>
      <p>このHTMLはオリジン前で生成されました。</p>
      <p>Request URI: <code>${uri}</code></p>
      <p>本番運用ではNext.jsのSSRバンドルに置き換えてください。</p>
    </body>
    </html>`;

    const response = {
      status: '200',
      statusDescription: 'OK',
      headers: {
        'content-type': [{ key: 'Content-Type', value: 'text/html; charset=utf-8' }],
        'cache-control': [{ key: 'Cache-Control', value: 'no-store' }],
      },
      body,
    };

    return callback(null, response);
  }

  // Pass through to origin for other assets/APIs
  return callback(null, request);
};

