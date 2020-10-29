const functions = require('firebase-functions');

// // Create and Deploy Your First Cloud Functions
// // https://firebase.google.com/docs/functions/write-firebase-functions
//
// exports.helloWorld = functions.https.onRequest((request, response) => {
//   functions.logger.info("Hello logs!", {structuredData: true});
//   response.send("Hello from Firebase!");
// });

exports.letsJanken = functions.https.onRequest(async (req, res) => {
    // Grab the text parameter.
    const query = req.query;
    // Send back a message that we've succesfully written the message
    console.log(query.hand);

    let hands = ['gu', 'choki', 'pa'];
    let computerHand = hands[Math.floor(Math.random() * hands.length)];
    let winner = '';
    if(query.hand == 'gu') {
        if(computerHand == 'gu') {
            winner = '';
        }else if(computerHand == 'choki') {
            winner = 'あなた';
        }else if(computerHand == 'pa') {
            winner = 'コンピュータ'
        }
    }else if(query.hand == 'choki') {
        if(computerHand == 'gu') {
            winner = 'コンピュータ'
        }else if(computerHand == 'choki') {
            winner = '';
        }else if(computerHand == 'pa') {
            winner = 'あなた'
        }
    }else if(query.hand == 'pa') {
        if(computerHand == 'gu') {
            winner = 'あなた'
        }else if(computerHand == 'choki') {
            winner = 'コンピュータ';
        }else if(computerHand == 'pa') {
            winner = ''
        }
    }
    let message = 'あなたは ' + query.hand + ' コンピュータは ' + computerHand + ' 結果は...';
    if(winner == '') {
        message = message + 'あいこでした。';
    }else {
        message = message + winner + 'の勝利';
    }
    console.log(message);
    res.json({result: message});
});
