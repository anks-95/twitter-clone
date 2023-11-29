//Main entry point for your React application using React Router.
//Sets up routes for the Home and About pages.

import React from 'react'
import ReactDOM from 'react-dom/client';
import { BrowserRouter, Route, Routes } from 'react-router-dom';

import Login from './pages/Login';
import Signup from './pages/Signup'




const root = ReactDOM.createRoot(document.querySelector("#application")!);
root.render(
<>  
  <BrowserRouter>
    <Routes>
      <Route path='/' element={<Login />} />
      <Route path='/signup' element={<Signup />} />

    </Routes>
  </BrowserRouter>
    
</>
);