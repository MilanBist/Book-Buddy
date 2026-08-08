import './App.css';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import SignUpForm from '../pages/Signup'
import LoginForm from '../pages/Login'
import Home from '../pages/Home'
import {ToastContainer} from "react-toastify"
function App() {
  return (
    <>
        <Routes>
          <Route path="/" element={<Home/>} />
          <Route path="/login" element={<LoginForm />} />
          <Route path="/signup" element={<SignUpForm />} />
        </Routes>

        <ToastContainer
          position="top-right"
          autoClose={3000}
        />
    </>
  );
}
export default App;