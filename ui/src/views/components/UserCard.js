import React from "react";

function UserCard(props) {
  return (
    <div className="rounded border border-secondary p-1 m-1">
      <div className="" onClick={props.onclick}>
        <strong>{props.name}</strong>
        <br />
        <small>{props.email}</small>
      </div>
    </div>
  );
}

export default UserCard;
