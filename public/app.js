new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        username: "atak", // Our username
    },

    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var msg = JSON.parse(e.data);
            // self.chatContent += '</div>'
            //         //+ '<img src="' + self.gravatarURL(msg.email) + '">' // Avatar
            //         //+ msg.username
            //     + '</div>'
            //     //+ emojione.toImage(msg.message) // Parse emojis
            //     + msg.message
            //     //+ '='
            //     + msg.username
            //     + '<br/>'; 
            self.chatContent = msg.message;

            var element = document.getElementById('chat-messages');
            //element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
    },

    methods: {
        send: function () {
            if (this.newMsg != '') {
                this.ws.send(
                    JSON.stringify({
                        //email: this.email,
                        username: this.username,
                        message: $('<p>').html(this.newMsg).text() // Strip out html
                    }
                ));
                //this.newMsg = ''; // Reset newMsg
            }
        },
        setMessage: function (x) {
            if (x != String(eval(x))) {
                this.newMsg = x + '=' + String(eval(x));}
            else { this.newMsg = x}
        },
    }
});