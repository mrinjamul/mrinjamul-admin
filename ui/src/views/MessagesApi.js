import React, { useState,useEffect } from "react";
import { Button, Alert } from "reactstrap";
// import Highlight from "../components/Highlight";
import { useAuth0, withAuthenticationRequired } from "@auth0/auth0-react";
import { getConfig } from "../config";
import Loading from "../components/Loading";
import Reader from "./components/Reader";
import MessageBar from "./components/MessageBar";

export const MessagesApiComponent = () => {
  const { apiOrigin = "http://localhost:3000", audience } = getConfig();

  const [state, setState] = useState({
    showResult: false,
    apiMessage: "",
    error: null,
  });

  const {
    getAccessTokenSilently,
    loginWithPopup,
    getAccessTokenWithPopup,
  } = useAuth0();

  useEffect(() => {
    const callApi = async () => {
      try {
        const token = await getAccessTokenSilently();
  
        const response = await fetch(`${apiOrigin}/api/messages`, {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        });
  
        const responseData = await response.json();
  
        setState({
          ...state,
          showResult: true,
          apiMessage: responseData,
        });
      } catch (error) {
        setState({
          ...state,
          error: error.error,
        });
      }
    };
    setInterval(() => {
      callApi();
    }, 60000);
  },[apiOrigin,getAccessTokenSilently,state])

  const handleConsent = async () => {
    try {
      await getAccessTokenWithPopup();
      setState({
        ...state,
        error: null,
      });
    } catch (error) {
      setState({
        ...state,
        error: error.error,
      });
    }

    await callApi();
  };

  const handleLoginAgain = async () => {
    try {
      await loginWithPopup();
      setState({
        ...state,
        error: null,
      });
    } catch (error) {
      setState({
        ...state,
        error: error.error,
      });
    }

    await callApi();
  };

  const callApi = async () => {
    try {
      const token = await getAccessTokenSilently();

      const response = await fetch(`${apiOrigin}/api/messages`, {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const responseData = await response.json();

      setState({
        ...state,
        showResult: true,
        apiMessage: responseData,
      });
    } catch (error) {
      setState({
        ...state,
        error: error.error,
      });
    }
  };

  const handle = (e, fn) => {
    e.preventDefault();
    fn();
  };

  return (
    <>
      <div className="mb-5">
        {state.error === "consent_required" && (
          <Alert color="warning">
            You need to{" "}
            <a
              href="#/"
              class="alert-link"
              onClick={(e) => handle(e, handleConsent)}
            >
              consent to get access to users api
            </a>
          </Alert>
        )}

        {state.error === "login_required" && (
          <Alert color="warning">
            You need to{" "}
            <a
              href="#/"
              class="alert-link"
              onClick={(e) => handle(e, handleLoginAgain)}
            >
              log in again
            </a>
          </Alert>
        )}
        {state.showResult && <div className="text-center">
        <MessageBar numberOfMessage={state.apiMessage.length} />
        </div>}

        <div className="text-right">
          <Button
            color="primary"
            className="mt-5"
            onClick={callApi}
            disabled={!audience}
          >
            Refresh
          </Button>
        </div>
      </div>

      <div className="result-block-container">
        {state.showResult && (
          <div className="result-block" data-testid="api-result">
            {/* <Highlight>
              <span>{JSON.stringify(state.apiMessage, null, 2)}</span>
            </Highlight> */}
            <Reader messages={state.apiMessage} />
          </div>
        )}
      </div>
    </>
  );
};

export default withAuthenticationRequired(MessagesApiComponent, {
  onRedirecting: () => <Loading />,
});
