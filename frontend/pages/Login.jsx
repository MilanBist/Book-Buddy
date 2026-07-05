import axios from "axios";
import { useState } from "react";
import { Link, Navigate } from "react-router-dom";
import apiClient from "../api/api";
import { useNavigate } from "react-router-dom";




// check for the email
const checkEmail = (email)=>{
    let newEmail = String(email).trim();
        // use the regex to verify the email
    const re  = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
    const result = re.test(email);

    if (result !== true){
        return false;
    }
    return true;

}

// check for the password
const checkPassword = (password)=>{
    const re = /^[a-zA-Z0-9!@#$%^&*]{6,16}$/;
    const result = re.test(password);

    if (result !== true){
        return false;
    }
    return true;
}

export default function LoginForm(){
    const navigate = useNavigate();
    // make the states for all of the given things
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");

    // for the response
    const [response, setResponse] = useState({});
    const [error, setError] = useState();

    const UpdatePassword = (evt)=>{
        setPassword(evt.target.value);
    }
    const UpdateEmail = (evt)=>{
        setEmail(evt.target.value);
    }    

    const handleLoginSubmit = (evt)=>{
        evt.preventDefault();
       

        console.log("Button is being clicked.")
        // check for email and password
        if (String(email).length ==0){
            alert("Enter the email.");
            return
        }
        let cmail = checkEmail(email);
        if (String(password).length === 0){
            alert("Enter the password.");
            return;
        }

        let cpass = checkPassword(password);

        if (cmail === false){
            alert("Wrong email.");
            return;
        }

        if (cpass === false){
            alert("Wrong password.");
            return;
        }


    // console.log("Email is : ", email);
    // console.log("Password: ", password);
    const bodyMap = {
        userEmail: email,
        userPassword: password
    };
''
    // to be sent
    console.log("The bodymap which is going to be sent:",bodyMap);


    // if both are correct then call the handler for the login
    apiClient.post("/login",bodyMap).then((resp)=>{
            setResponse(resp);
        }).catch((err) => {
            setError(err);
        }).finally(()=>{
            console.log("Login credentials checked.");
        })

    // check for the error
    if (error){
        console.log("Error in checking the credentials.");
        return;
    }

    // // check from the response
    // // check if mail exists then 
    // if (response["mailMessage"] === false){
    //     // now the user doesn't exist
    //     alert("User don't exist. Please signup.");
    //     return;
    // }

    // if (response["mailMessage"] && response["passwordMessage"] === false){
    //     alert("Password credentials wronng.");
    //     return;
    // }

    // if all are correct give the message of successfully logged in
    // alert('Successfully logged up.');

    // navigate('/');
    
    }
       return(
        <div id="form-body">
            <div className="form-container">

                <form action="" onSubmit={handleLoginSubmit}>                
                    <label htmlFor="email">Email</label>
                    <input 
                        type="email" 
                        name="email" 
                        id="email" 
                        placeholder="johndoe@gmail.com"
                        onChange={UpdateEmail}
                    />

                    <label htmlFor="password">Password</label>
                    <input 
                        type="password" 
                        name="password" 
                        id="password" 
                        placeholder="password"
                        onChange={UpdatePassword}
                    />

                    <button type="submit">Submit Form</button>

                    <div className='link-login'>
                        <h4>Don't have an account?</h4> 
                        <Link to='/signup'><p>Click Here!</p></Link>
                    </div>
                </form>
            </div>
        </div>
    )
}