// push通知のイベントが発生したときの処理を登録
self.addEventListener('push', evt => {
    const data = evt.data.json();
    console.log(data);
    const title = data.title;
    const options = {
        body: data.body,
        icon: 'doggy.jpg'
    }
    evt.waitUntil(self.registration.showNotification(title, options));
});

// push通知をクリックしたときの処理を登録
self.addEventListener('notificationclick', evt => {
    evt.notification.close();
});