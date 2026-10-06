import "../styles/FileUpload.css"
import 'react-toastify/dist/ReactToastify.css';
import { useState } from "react";
import { toast } from "react-toastify";
import { useNavigate } from "react-router-dom";
import axios from "axios";


const redirection = (time, navigate)=>{
    toast.error("Your are not authenticated. Please login/register", {
            autoClose: time,
            onClose: () => {
            navigate("/login");
        },
    });
}
export default function FileUpload({uploadBar,setUploadBar, currentBookData, setBookData}){
    // take the data from the function 
    const[disable, setDisable] = useState(false);
    const[fileData, setFileData] = useState(null);
    const [isUploading, setIsUploading] = useState(false);
    const navigate = useNavigate();
    const uploadFile =  async ()=>{
        const token = localStorage.getItem("tokenId");
        if (!token){
            // redirect to the login page
            redirection(2000,navigate);
            return;
        }
         // create the instance of the form data
         if (!fileData){
            alert("Insert the file first.");
            return;
         }
        const formData = new FormData();
        formData.append("document", fileData)
        // after appending the file now my task is to send the request using the axios to the browser
        try{
            setIsUploading(true);
            const response = await axios.post("http://localhost:8080/api/handlePdf", formData, {
                headers:{
                Authorization: `Bearer ${token}`,
                "Content-Type": 'multipart/form-data',
                }
            });

            const responseData = response.data;
            const bookId = responseData.data.bookId;
            const bookName = responseData.data.bookName;
            const bookData = {
                "bookId": bookId,
                "bookName": bookName,
            };
            
            setBookData((previous)=>[
                bookData,
                ...previous,
            ])
            // set this book as teh current book
            currentBookData(bookData);
    
        } catch(err){
            setUploadBar(false);
            const status = err.response?.status;
            // const message =  err.response.data.message;
            switch (status){
                case 500:
                    // check for the message
                    // console.log(message);
                    alert("Internal server error. \n Please try again later.");
                    return;
                case 401:
                    // console.log(message);
                    alert("You are not validated. \n Please login.");
                    redirection(2000, navigate);
                    return;
                case 400:
                    // console.log(message);
                    alert("Wrong input.");
                    return;
            }
        
        } finally{
            console.log("Finished uploading the file.");
            setIsUploading(false);
        }
    }


    const handleFileChange = (e)=>{
        const file = e.target.files[0];
        // if you find the file there just send this to the funciton
        setFileData(file);
    }


    const submitFile = ()=>{
        if (fileData != null){
            setDisable(true);
            setTimeout(()=>{
                setDisable(false);
            }, 10000);

            // submit the file from here to the backend
            uploadFile();
            return;
        }
        else{
            alert("First select the file");
        }
    }


    const disableUploadField = ()=>{
        setUploadBar(false);
    }

    

    return (
        <div id="fileupload-overlay">
            <div id="fileupload">
                { isUploading ? (
                // Uploading state
                <div className="uploading-container">
                    <div className="upload-spinner"></div>

                    <h2>Uploading...</h2>

                    <p>
                        Please wait while your book is being processed.
                    </p>
                </div>
            )
             : (
                // Normal upload form
                <>
                    <input
                        type="file"
                        onChange={handleFileChange}
                        disabled={!uploadBar}
                        className="file_input_by_user"
                    />

                    <div className="fileuploadbuttons">
                        <button
                            onClick={submitFile}
                            className="button"
                            disabled={disable}
                        >
                            Process
                        </button>

                        <button
                            onClick={disableUploadField}
                            className="button"
                            disabled={disable}
                        >
                            Cancel
                        </button>
                    </div>
                </>
            )}
            </div>
        </div>
    )
}