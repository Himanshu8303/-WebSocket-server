import React, { useState, useEffect } from 'react';

function App() {
  const [message, setMessage] = useState('');
  const [messages, setMessages] = useState([]);
  const [ws, setWs] = useState(null);

  useEffect(() => {
    const socket = new WebSocket('ws://localhost:8080/ws');

    socket.onmessage = (event) => {
      setMessages((prev) => [...prev, `Received: ${event.data}`]);
    };

    setWs(socket);

    return () => socket.close();
  }, []);

  const sendMessage = () => {
    if (ws) {
      ws.send(message);
      setMessages((prev) => [...prev, `Sent: ${message}`]);
      setMessage('');
    }
  };

  const getHistory = () => {
    if (ws) {
      ws.send('history');
    }
  };

  return (
    <div>
      <input
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        placeholder="Enter message"
      />
      <button onClick={sendMessage}>Send</button>
      <button onClick={getHistory}>Get History</button>
      <div>
        <h3>Messages:</h3>
        {messages.map((msg, index) => (
          <div key={index}>{msg}</div>
        ))}
      </div>
    </div>
  );
}

export default App;
