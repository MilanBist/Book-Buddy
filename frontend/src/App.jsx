import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import SignUpForm from '../pages/Signup';
import LoginForm from '../pages/Login';
import Home from '../pages/Home';
import {ToastContainer} from "react-toastify"
import { useState } from 'react';
function App() {

  const [loggedInState, setLoggedInState] = useState(false);
  return (
  <>
      <Routes>
        <Route path="/" element={<Home/>} />
        <Route path="/login" element={<LoginForm setLoggedInState={setLoggedInState}/>} />
        <Route path="/signup" element={<SignUpForm setLoggedInState={setLoggedInState}/>} />
      </Routes>

      <ToastContainer
        position="top-right"
        autoClose={2000}
      />
  </>
  );
  }
export default App;