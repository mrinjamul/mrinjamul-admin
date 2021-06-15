import React from "react";

import logo from "../assets/logo.svg";

const Hero = () => (
  <div className="text-center hero my-5">
    <img className="mb-3 app-logo" src={logo} alt="logo" width="120" />
    <h1 className="mb-4">Welcome to Admin</h1>

    <p className="lead">
      This is Injamul's dashboard for the <a href="https://mrinjamul.github.io">mrinjamul</a>.
    </p>
  </div>
);

export default Hero;
