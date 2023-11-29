import React, { useState, ChangeEvent, FormEvent } from 'react';
import { Link, useNavigate } from 'react-router-dom';
import Validation from './SignupValidation';
import axios from 'axios'

function Signup(){
  const [values, setValues] = useState({
    name: '',
    email: '',
    password: '',
  });

  const navigate = useNavigate();
  const [errors, setErrors] = useState({
    name: '',
    email: '',
    password: '',
  });

  const handleInput = (e: ChangeEvent<HTMLInputElement>) => {
    setValues((prev) => ({ ...prev, [e.target.name]: e.target.value }));
  };

  const handleSubmit = async (e: FormEvent) => {
    e.preventDefault();
    setErrors(Validation(values));
    if(errors.name===" "&&errors.email===" "&&errors.password===" "){
      axios.post('http://localhost:3000/signup', values)
      .then(res => {
        navigate('/');
    })
    .catch(err => {
        console.log(err.response); 
    });
      
    }
  };



  return (
    <div className="d-flex flex-column align-items-center justify-content-center min-vh-100 bg-dark text-white">
      <div className="login-container bg-gray-800 p-4 rounded shadow-sm" style={{ width: '400px' }}>
        <h2 className="text-center text-white mb-4">Register</h2>
        <form action='' onSubmit={handleSubmit}>
          <div className="mb-3">
            <label htmlFor="exampleInputUsername" className="form-label text-white">Username</label>
            <input
              type="username"
              className="form-control"
              id="exampleInputUsername"
              onChange={handleInput}
              name="name"
            />
            {errors.name && <span className='text-danger'> {errors.name}</span>}
          </div>
          

<div className="mb-3">
  <label htmlFor="exampleInputEmail1" className="form-label text-white">
    Email address
  </label>
  <input
    type="email"
    className="form-control"
    id="exampleInputEmail1"
    aria-describedby="emailHelp"
    onChange={handleInput}
    name="email"
  />
  {errors.email && <span className='text-danger'> {errors.email}</span>}
</div>



          <div className="mb-3">
            <label htmlFor="exampleInputPassword1" className="form-label text-white">Password</label>
            <input
              type="password"
              className="form-control"
              id="exampleInputPassword1"
              name="password"
              onChange={handleInput}
            />
            {errors.password && <span className='text-danger'> {errors.password}</span>}
          </div>
          <div className="mb-3 form-check">
            <input type="checkbox" className="form-check-input" id="exampleCheck1" />
            <label className="form-check-label text-white" htmlFor="exampleCheck1">Agree to Terms & Conditions</label>
          </div>
          <div className="d-grid gap-2">
            <button type="submit" className="btn btn-success">SignUp</button>
            <div className="d-grid gap-2">
              <label className="form-check-label text-white" htmlFor="exampleCheck1">Have an account?</label>
              <Link to="/" className="btn btn-primary">Login</Link>
            </div>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Signup;
