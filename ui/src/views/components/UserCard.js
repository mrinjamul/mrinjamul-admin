import React from "react";

function UserCard(props) {
  return (
    <div className="rounded border p-1 m-1">
      <div className="cursorpointer" onClick={props.onclick}>
        <strong className="text-dark">{props.name}</strong>
        <br />
        {/* <a href={"mailto:"+ props.email}> */}
          <small className="text-secondary">{props.email}</small>
        {/* </a> */}
      </div>
    </div>
  );
}

export default UserCard;
