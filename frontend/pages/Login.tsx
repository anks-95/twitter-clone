// login.tsx

import React, { useState, ChangeEvent, FormEvent } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Validation from './LoginValidation';
import axios from 'axios';




function Login(){
  const [values, setValues] = useState({
    email: '',
    password: '',
  });

  const [errors, setErrors] = useState({
    email: '',
    password: '',
  });

  const handleInput = (e: ChangeEvent<HTMLInputElement>) => {
    setValues((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };
  const navigate = useNavigate();




  const handleSubmit = async (e: FormEvent) => {
  e.preventDefault();
  setErrors(Validation(values));

  // Check if there are no validation errors and both email and password are not empty
  if (errors.email===" " && errors.password===" ") {
    axios.post('http://localhost:3000/login', values)
      .then(res => {
        if (res.data.message === "success") {
          // User found, navigate to home
          navigate('/home');
        } else  {
          // User not found, show invalid credentials message
          alert("Invalid credentials");
          // You can also set an error state to display a message on the UI
        } 
      })
      .catch(err => {
        console.log(err.response);
      });
  }
};


 

  return (
    <>
    <div className="d-flex flex-column align-items-center justify-content-center min-vh-100 bg-dark text-white">
      <div className="login-container bg-gray-800 p-4 rounded shadow-sm" style={{ width: '400px' }}>
        <h2 className="text-center text-white mb-4">Login to Your Account</h2>
        <form action="" onSubmit={handleSubmit}>
          <div className="mb-3">
            <label htmlFor="exampleInputEmailLogin" className="form-label text-white">
              Email address
            </label>
            <input
              type="email"
              className="form-control"
              id="exampleInputEmailLogin"
              placeholder="email"
              onChange={handleInput}
              name="email"
              required 
            />
            {errors.email && <span className="text-danger"> {errors.email}</span>}
          </div>
          <div className="mb-3">
            <label htmlFor="exampleInputPasswordLogin" className="form-label text-white">
              Password
            </label>
            <input
              type="password"
              className="form-control"
              id="exampleInputPasswordLogin"
              placeholder="password"
              onChange={handleInput}
              name="password"
              required 
            />
            {errors.password && <span className="text-danger"> {errors.password}</span>}
          </div>
          <div className="mb-3 form-check">
            <input type="checkbox" className="form-check-input" id="exampleCheck1" />
            <label className="form-check-label text-white" htmlFor="exampleCheck1">
              Save Credentials!
            </label>
          </div>
          <div className="d-grid gap-2">
            <button type="submit" className="btn btn-primary">
              Log In
            </button>
            <Link to="/signup" className="btn btn-success text-decoration-none">
              Create Account
            </Link>
          </div>
        </form>
      </div>
    </div>
    
    </>
  );
};

export default Login;
