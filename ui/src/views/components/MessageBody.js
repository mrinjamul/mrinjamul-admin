import React, { useState } from "react";

import UserCard from "./UserCard";

function MessageBody(props) {
  let [messageText, setMessageText] = useState("");

  const onclick = (message) => {
    setMessageText(message);
  };

  const messages = props.messages;

  let UserMessage = messages.map((message) => (
      <UserCard
        key={messages.indexOf(message)}
        id={message.id}
        name={message.name}
        email={message.email}
        onclick={() => {
          onclick(message.message);
        }}
      />
  ));
  return (
    <div className="container">
      <div className="row">
        <div className="col-sm scrollbar scrollbar-primary">{UserMessage}</div>
        <div className="col-sm">
          <p>
            {messageText}
          </p>
        </div>
      </div>
    </div>
  );
}

export default MessageBody;
