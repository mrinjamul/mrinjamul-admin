import React from "react";

const MessageBar = (props) => {
  return (
    <div className="">
      <b>You have {props.numberOfMessage} message(s).</b>
    </div>
  );
};

export default MessageBar;
