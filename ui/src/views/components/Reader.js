import React from "react";

import MessageBody from "./MessageBody";

import "./Reader.css";

function Reader(props) {
  let messages = props.messages;
  return (
    <div className="container">
      <div className="col">
        <MessageBody messages={messages} />
      </div>
    </div>
  );
}

export default Reader;
