import { useState } from "react";
import { data, Link } from 'react-router-dom';
import apiClient from '../api/api';
import { useNavigate } from 'react-router-dom';
import "../styles/Register.css";
// check names valididty
const checkName = (name)=>{
    if (name === "" || name.length <2 || name.length >30){
        alert("Please enter the valid first name.");
        return false;
    }
    return true;
}

// check for the address
const checkAddress = (newaddress)=>{
    if (newaddress.length>30 || newaddress.length <=2){
        alert("Please enter the valid address of length between 3 and 30.");
        return false
    }
    return true;
}

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

export default function SignUpForm({setLoggedInState}){
    const navigate = useNavigate();

    // make the states for all of the given things
    const [firstName, setFirstName] = useState("");
    const [lastName, setLastName] = useState("");
    const [address, setAddress] = useState("");
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");


    // for the response

    // add the setter functions
    const UpdateFirstName = (evt)=>{
        setFirstName(evt.target.value);
    }
    const UpdateLastName = (evt)=>{
        setLastName(evt.target.value);
    }
    const UpdateAddress = (evt)=>{
        setAddress(evt.target.value);
    }
    const UpdatePassword = (evt)=>{
        setPassword(evt.target.value);
    }
    const UpdateEmail = (evt)=>{
        setEmail(evt.target.value);
    }

    const handleSignupSubmit = (evt)=>{
        evt.preventDefault();
    
        // check for each of the messages
        // check for the firstName and last Name
        if (String(firstName).length===0 || String(lastName).length===0
        || String(address).length === 0 || String(email).length === 0 || String(password).length === 0){
            alert("Fill all of credentials.");
            return;
        }
    
        let checkFirstName = checkName(firstName);
        let checkLastName = checkName(lastName);
    
        if (checkFirstName !== true || checkLastName !== true){
            // alert enter the name validly
            alert("Enter the valid name.");
            return;
        }
    
        // check for the address 
        let cadd = checkAddress(address);
        if (cadd !== true){
            alert("Enter valid address.");
            return;
        }
    
        // check for the email and the password
        
        // check for the email and the password
        let cemail = checkEmail(email);
        if (cemail === false){
            alert("Enter valid email.");
            return;
        }
    
        let cpass = checkPassword(password);
        if (cpass !== true){
            alert("Enter valid password.");
            return;
        }    

        console.log("Validation successfull.");

        const formdata = {
            userFirstName: firstName,
            userLastName: lastName,
            userAddress: address,
            userEmail: email,
            userPassword: password,
        };

        console.log(formdata)

        // send this to the frontend using the axios
        apiClient.post("/register", formdata).then((resp) =>{
            // if the response status is 202
            localStorage.setItem("tokenId", resp["data"]["data"]["token"]);
            // navigate to the homepage now
            console.log("Successfully added new user.");
            alert("Redirecting to home page.");
            setTimeout(()=>{
                console.log("Redirecting to home page.");
                navigate("/");
            },1000);
            
        }).catch((err) => {
            console.log("Reaching to this part of error section.", err.response.data.message);
            let statusCode = err.response.status;
            let message = err.response.data.message;
            console.log("Status code is: ", statusCode);
            console.log("Message is: ", message);
            switch(statusCode){
                case 500:
                    // check for the message
                    switch(message){
                        case "can't generate":
                            console.log("Can't generate tokens.");
                            alert("Server error.");
                            return;
                        case "can't insert":
                            console.log("Can't insert to database.");
                            alert("Server error.");
                            return;
                    }
                case 401:
                    let msg = err.response.data.data;
                    alert(msg);
                    return;
                case 409:
                    alert("User already exisit.\n Redirecting to login page.");
                    setTimeout(()=>{
                        navigate("/login");
                    }, 1000);
            }
        })
    }

    return(
        <div id='form-body'>
            <div className="form-container">
                <h2>Create Account</h2>
                <form onSubmit={handleSignupSubmit}>
                    <label htmlFor="firstname">FirstName</label>
                    <input 
                        type="text" 
                        onChange={UpdateFirstName} 
                        name="firstName" 
                        placeholder="Enter your first name"
                    />

                    <label htmlFor="lastname">Lastname</label>
                    <input 
                        type="text" 
                        onChange={UpdateLastName} 
                        name="lastName" 
                        placeholder="Enter your last name"
                    />

                    <label htmlFor="address">Address</label>
                    <input 
                        type="text" 
                        onChange={UpdateAddress} 
                        name="address" 
                        placeholder="Enter your address"
                    />

                    <label htmlFor="email">Email</label>
                    <input 
                        type="email" 
                        onChange={UpdateEmail}
                        name="email" 
                        id="email" 
                    />

                    <label htmlFor="password">Password</label>
                    <input 
                        type="password" 
                        name="password" 
                        id="password" 
                        onChange={UpdatePassword}
                    />
                    <button type="submit" onSubmit={handleSignupSubmit}>Sumbit Form</button>
                    <div className='link-login'>
                        <h4>Already have an account?</h4> 
                        <Link to='/login'><p>Click Here!</p></Link>
                    </div>
                </form>
            </div>
        </div>
    )
}