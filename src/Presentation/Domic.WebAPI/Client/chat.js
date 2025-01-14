const chatBox = document.getElementById('chat-box');
const messageInput = document.getElementById('message-input');
const sendButton = document.getElementById('send-button');

const socket = new WebSocket('ws://localhost:8080/chat?connectionId=' + generateUUID());

socket.onopen = (event) => {
    console.log("connection with chat server is up now!");
};

socket.onmessage = (event) => {
    const message = document.createElement('div');

    console.log("message payload: " + event.data);

    message.className = 'message other';
    message.textContent = JSON.parse(event.data).content;
    chatBox.appendChild(message);
    chatBox.scrollTop = chatBox.scrollHeight;
};

sendButton.addEventListener('click', () => {
    const messageText = messageInput.value.trim();
    if (messageText) {
        socket.send(JSON.stringify({
            ConnectionId: '',
            Content: messageText,
            To: ''
        }));

    // const message = document.createElement('div');
    // message.className = 'message self';
    // message.textContent = messageText;
    // chatBox.appendChild(message);
    // chatBox.scrollTop = chatBox.scrollHeight;

    messageInput.value = '';
}
});

messageInput.addEventListener('keypress', (e) => {
    if (e.key === 'Enter') {
        sendButton.click();
    }
});

function generateUUID() {
    return 'xxxxxxxx-xxxx-4xxx-yxxx-xxxxxxxxxxxx'.replace(/[xy]/g, function (c) {
        const r = (Math.random() * 16) | 0;
        const v = c === 'x' ? r : (r & 0x3) | 0x8;
        return v.toString(16);
    });
}