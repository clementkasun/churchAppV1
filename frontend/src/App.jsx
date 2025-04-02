import React, { useState } from "react";
import axios from "axios";
import logo from "./assets/images/church_logo.jpg";
import { MDBBtn, MDBContainer, MDBRow, MDBCol, MDBInput } from "mdb-react-ui-kit";

function App() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");

  const handleLogin = async () => {
    try {
      const response = await axios.post("http://localhost:5000/api/login", {
        email,
        password,
      });
      setMessage(response.data.message);
      console.log("Login Success:", response.data);
    } catch (error) {
      setMessage(error.response?.data?.message || "Login failed");
      console.error("Login Error:", error);
    }
  };

  return (
    <MDBContainer className="my-5 gradient-form">
      <MDBRow>
        <MDBCol col="6" className="mb-5">
          <div className="d-flex flex-column ms-5">
            <div className="text-center">
              <img src={logo} style={{ width: "185px" }} alt="logo" />
              <h4 className="mt-1 mb-5 pb-1">New Door Of Life Prayer Centre Narammala</h4>
            </div>
            <p>Please login to your account</p>

            <MDBInput
              wrapperClass="mb-4"
              label="Email address"
              id="email"
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <MDBInput
              wrapperClass="mb-4"
              label="Password"
              id="pass"
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />

            <div className="text-center pt-1 mb-5 pb-1">
              <MDBBtn className="mb-4 w-100 gradient-custom-2" style={{ maxWidth: "200px", maxHeight: "40px" }} onClick={handleLogin}>
                Sign in
              </MDBBtn>
              <p style={{ color: "red" }}>{message}</p>
              <a className="text-muted d-block" href="#!">
                Forgot password?
              </a>
            </div>

            <div className="d-flex flex-row align-items-center justify-content-center pb-4 mb-4">
              <p className="mb-0">Don't have an account?</p>
              <MDBBtn outline style={{ maxWidth: "200px", maxHeight: "40px" }} color="danger">
                Signup
              </MDBBtn>
            </div>
          </div>
        </MDBCol>

        <MDBCol col="6" className="mb-5">
          <div className="d-flex flex-column justify-content-center gradient-custom-2 h-100 mb-4">
            <div className="text-white px-3 py-4 p-md-5 mx-md-4">
              <h4 className="mb-4">Welcome to New Door Life Prayer Center - Narammala</h4>
              <p className="small mb-0">
                We are a community committed to faith, prayer, and transformation. Join us to experience spiritual growth and guidance.
              </p>
            </div>
          </div>
        </MDBCol>
      </MDBRow>
    </MDBContainer>
  );
}

export default App;
