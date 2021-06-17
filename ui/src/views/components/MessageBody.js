import React, { useState, useEffect } from "react";

import UserCard from "./UserCard";

function MessageBody(props) {
  let [messegeMeta, setMessegeMeta] = useState("");

  const onclick = (messege) => {
    setMessegeMeta({
      ...messegeMeta,
      name: messege.name,
      email: messege.email,
      subject: messege.subject = messege.subject ? messege.subject : "No Subject",
      message: messege.message,
    });
  };

  const messages = props.messages;

  let UserMessage = messages.map((message) => (
      <UserCard
        key={messages.indexOf(message)}
        id={message.id}
        name={message.name}
        email={message.email}
        onclick={() => {
          onclick(message);
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
      <div className="col-sm cursorarrow">
        <div className="">
          <h6 className="text-info"><a href={"mailto:"+messegeMeta.email}>{messegeMeta.name}</a></h6>
          <h6 className="text-dark">{messegeMeta.subject}</h6>
        </div>
        <br/>
        <div className="">
          <p>
            {messegeMeta.message}
          </p>
        </div>
      </div>
    </div>
      }

      { isMobile &&
      <div className="col">
      <div className="col-sm scroll-m">{UserMessage}</div>
      <div className="col-sm cursorarrow">
        <div className="">
          <h6 className="text-info">{messegeMeta.name}</h6>
          <h6 className="text-dark">{messegeMeta.subject}</h6>
        </div>
        <br/>
        <div className="">
          <p>
            {messegeMeta.message}
          </p>
        </div>
      </div>
    </div>
      }

    </div>
  );
}

export default MessageBody;
