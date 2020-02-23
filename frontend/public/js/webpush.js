let convertedVapidKey, subscription;

(async _ => {
    try {

        const registration = await navigator.serviceWorker.register('/serviceworker.js', { scope: '/' });
        const res = await fetch('http://localhost:3000/pubkey', {mode: 'cors'});
        const vapidPublicKey = await res.text();

        console.log(vapidPublicKey);
        convertedVapidKey = urlBase64ToUint8Array(vapidPublicKey);
        
        // プッシュサーバにアプリケーションサーバの公開鍵を登録する、エンドポイントとブラウザの公開鍵とソルトを取得する。
        subscription = await registration.pushManager.subscribe({
            userVisibleOnly: true,
            applicationServerKey: convertedVapidKey
        });
        
        // 通知の許可を要求
        Notification.requestPermission(permission => {
            console.log(permission); // 'default', 'granted', 'denied'
        });

        if (!subscription) return console.log('sbuscription is null');
        await fetch('http://localhost:3000/subscription', {
            method: 'POST',
            body: JSON.stringify(subscription),
            headers: {
                'Content-Type': 'application/json',
            },
            mode: 'cors',
        });

    } catch (err) {
        console.log(err);
    }
})();

btnWebPushTest.onclick = async evt => {
    // 通知をオフにしているとnullになる
    if (!subscription) return console.log('sbuscription is null');
    await fetch('/webpushtest', {
        method: 'POST',
        body: JSON.stringify(subscription),
        headers: {
            'Content-Type': 'application/json',
        },
    });
};

function urlBase64ToUint8Array(base64String) {
    const padding = '='.repeat((4 - base64String.length % 4) % 4);
    const base64 = (base64String + padding).replace(/-/g, '+').replace(/_/g, '/');
    const rawData = window.atob(base64);
    const outputArray = new Uint8Array(rawData.length);
    for (let i = 0; i < rawData.length; ++i) {
        outputArray[i] = rawData.charCodeAt(i);
    }
    return outputArray;
}