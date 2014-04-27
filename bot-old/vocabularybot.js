const config = require('./config.js').settings;
const dictionary = require('./dictionary.js');
execute_bot();


/**
* Adapted from http://njsbot.simonholywell.com/
*/
function execute_bot() {
    const xmpp = require('node-xmpp');
    const util = require('util');
    const request_helper = require('request');
    const conn = new xmpp.Client(config.client);

    function log_message(stanza) {
       	//var fs = require("fs");
        console.log('msg'+stanza.getChildText('body'));	
       	//fs.appendFile('message.txt', stanza.attrs.from+";"+stanza.getChildText('body')+"\n", function (err) { });
    }

    function set_presence() {
        var presence_elem = new xmpp.Element('presence', { })
                                .c('show').t('chat').up();
        conn.send(presence_elem);
    }

    /**
     * Send a message to the supplied JID
     * @param {String} to_jid
     * @param {String} message_body
     */
    function send_message(to_jid, message_body) {
        var elem = new xmpp.Element('message', { to: to_jid, type: 'chat' })
                 .c('body').t(message_body);
        conn.send(elem);
    }



    function message_dispatcher(stanza) {
        if('error' === stanza.attrs.type) {
            util.log('[error] ' + stanza.toString());
        } else if(stanza.is('message') && stanza.getChildText('body'))  {
            log_message(stanza);
            var reply = dictionary.translate(stanza.getChildText('body'));
            send_message(stanza.attrs.from, reply);
        }
    }

    conn.addListener('stanza', message_dispatcher);

    conn.on('online', function() {
        set_presence();

        // send whitespace to keep the connection alive
        // and prevent timeouts
        setInterval(function() {
            conn.send(' ');
        }, 30000);
    });

    conn.on('error', function(stanza) {
        console.log('[error] ' + stanza.toString());
    });
}
