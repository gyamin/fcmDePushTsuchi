# FirebaseからPUSH通知を行う

## はじめに
Firebase経由でPUSH通知を行うに当たって、iOS、Androidアプリの作成はハードルがあるので、
まずはウェブPUSHから試してみる。

## Firebaseについて

[JavaScript Firebase Cloud Messaging クライアント アプリを設定する](https://firebase.google.com/docs/cloud-messaging/js/client?hl=ja)
```
FCM SDK で使用している Service Worker は、HTTPS サイトでのみ利用可能です。
このため、FCM SDK は HTTPS 経由で提供されるページでのみサポートされます。
プロバイダが必要な場合は、独自ドメインで無料の HTTPS ホスティングを行うことができる Firebase Hosting をおすすめします。
```
ってことなので、Https環境でないとダメらしい。ローカルでWebサーバ立てずに、Firebase Hostingを利用する。

## Firebase SDk インストールとセットアップ
[Firebase CLI リファレンス](https://firebase.google.com/docs/cli?hl=ja)

インストール
```
$ nodenv local 12.18.0
$ node -v
v12.18.0
$ npm install -g firebase-tools
$ nodenv rehash
$ firebase --version
8.10.0
```

セットアップ      
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
`firebase login`を実行するとブラウザが立ち上がって認証処理が求められる。指示にしたがって認証すればOK。

## ホスティングの利用

```
$ mkdir web
$ cd web/
$ firebase init
...
? Which Firebase CLI features do you want to set up for this folder? Press Space to select features,
 then Enter to confirm your choices. Hosting: Configure and deploy Firebase Hosting sites
=== Project Setup

First, let's associate this project directory with a Firebase project.
You can create multiple project aliases by running firebase use --add, 
but for now we'll just set up a default project.

? Please select an option: Use an existing project
? Select a default Firebase project for this directory: fcm-de-push (fcm-de-push)
i  Using project fcm-de-push (fcm-de-push)

=== Hosting Setup

Your public directory is the folder (relative to your project directory) that
will contain Hosting assets to be uploaded with firebase deploy. If you
have a build process for your assets, use your build's output directory.

? What do you want to use as your public directory? public
? Configure as a single-page app (rewrite all urls to /index.html)? Yes
✔  Wrote public/index.html

i  Writing configuration info to firebase.json...
i  Writing project information to .firebaserc...
i  Writing gitignore file to .gitignore...

✔  Firebase initialization complete!

```
`firebase init`を実行する利用するプロジェクト、firebase sdkが提供する機能からどの機能を利用するかなど問合せがされる。

hostingのセットアップができたら、firebase deployコマンドでファイルをアップロードできる。
```
$ firebase deploy

=== Deploying to 'fcm-de-push'...

i  deploying hosting
i  hosting[fcm-de-push]: beginning deploy...
i  hosting[fcm-de-push]: found 1 files in public
✔  hosting[fcm-de-push]: file upload complete
i  hosting[fcm-de-push]: finalizing version...
✔  hosting[fcm-de-push]: version finalized
i  hosting[fcm-de-push]: releasing new version...
✔  hosting[fcm-de-push]: release complete

✔  Deploy complete!

Project Console: https://console.firebase.google.com/project/fcm-de-push/overview
Hosting URL: https://fcm-de-push.web.app
```

