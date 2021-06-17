import React, { useState, useEffect } from "react";

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

  // Check if its a mobile
  const [width, setWidth] = useState(window.innerWidth);
  function handleWindowSizeChange() {
          setWidth(window.innerWidth);
      }
  useEffect(() => {
          window.addEventListener('resize', handleWindowSizeChange);
          return () => {
              window.removeEventListener('resize', handleWindowSizeChange);
          }
      }, []);

  let isMobile = (width <= 768);

  return (
    <div className="container">

      { !isMobile &&
      <div className="row">
      <div className="col-sm scrollbar scrollbar-primary">{UserMessage}</div>
      <div className="col-sm">
        <p>
          {messageText}
        </p>
      </div>
    </div>
      }

      { isMobile &&
      <div className="col">
      <div className="col-sm scroll-m">{UserMessage}</div>
      <div className="col-sm">
        <p>
          {messageText}
        </p>
      </div>
    </div>
      }

    </div>
  );
}

export default MessageBody;
