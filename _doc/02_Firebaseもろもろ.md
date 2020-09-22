## Firebaseについて

[JavaScript Firebase Cloud Messaging クライアント アプリを設定する](https://firebase.google.com/docs/cloud-messaging/js/client?hl=ja)
```
FCM SDK で使用している Service Worker は、HTTPS サイトでのみ利用可能です。
このため、FCM SDK は HTTPS 経由で提供されるページでのみサポートされます。
プロバイダが必要な場合は、独自ドメインで無料の HTTPS ホスティングを行うことができる Firebase Hosting をおすすめします。
```


## Firebase SDk
[Firebase CLI リファレンス](https://firebase.google.com/docs/cli?hl=ja)

```
$ nodenv local 12.18.0
$ node -v
v12.18.0
$ npm install -g firebase-tools
$ nodenv rehash
$ firebase --version
8.10.0
```

```
$ firebase login
i  Firebase optionally collects CLI usage and error reporting information to help improve our products. Data is collected in accordance with Google's privacy policy (https://policies.google.com/privacy) and is not used to identify you.

? Allow Firebase to collect CLI usage and error reporting information? Yes
i  To change your data collection preference at any time, run `firebase logout` and log in again.

Visit this URL on this device to log in:
https://accounts.google.com/o/oauth2/auth?client_id=xxxxxxxx...

Waiting for authentication...

✔  Success! Logged in as xxxxxxxxxxxxxxg@gmail.com
```