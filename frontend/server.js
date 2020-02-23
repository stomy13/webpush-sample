const Koa = require('koa');
const Router = require('koa-router');
const serve = require('koa-static');
const koaBody = require('koa-body');
const webPush = require('web-push');

const app = new Koa();
const router = new Router();
const private_key = "9LfdrbPLcGQxhEOJkQ4Ka7Oia3w2Lf2Sd6EaVkNeUQA";
const public_key = "BO-UY2C7nObUfD6MYDfw5ecSpIuf8REJsu9gISnsCCtdvC6u-FpHkC_HNjjZmjvnn1HzOiGaLJy-tzPfY6M_6ns";

// 鍵の設定、メール...?
webPush.setVapidDetails(
    "mailto:masadinti@gmail.com",
    public_key,
    private_key
);

// ルーティング
router
    .get('/key', ctx => {
        // アプリケーションサーバーの公開鍵を返す
        ctx.body = public_key;
    })
    .post('/webpushtest', koaBody(), async ctx => {
        // プッシュサーバに通知を送信する
        try {
            await webPush.sendNotification(ctx.request.body, JSON.stringify({
                    title: 'Web Push通知テスト',
                    body: 'welcome web push',
            }));
            // ブラウザに返すレスポンスいれようね
            ctx.response.body = "aaabc"
        } catch (err) {
            console.log(err);
        }
    });

app
    .use(serve(__dirname + '/public'))
    .use(router.routes())
    .use(router.allowedMethods());

app.listen(3001);
