<<<<<<< HEAD
=======
import "../styles/button.css"
>>>>>>> fc0f258 (All commit saved)
import { useState } from "react";

export default function FileInput({onFileSelect, disabled, uploadChanger, setUpUploadBar}){
    console.log("In the file input page.", disabled)
    const[fileData, setFileData] = useState(null);
    // take the data from the function 
    const handleFileChange = (e)=>{
        const file = e.target.files[0];
        // if you find the file there just send this to the funciton
        setFileData(file);
    }

    const submitFile = ()=>{
        if (fileData != null){
            setTimeout(()=>{
                onFileSelect(fileData)
                uploadChanger(true)
                setTimeout(()=>{
                    setUpUploadBar(false)
                    uploadChanger(false)
                }, 2000)
            },10)
            
        }
        else{
            alert("First select the file")
        }
    }

    const disableUploadField = ()=>{
        setUpUploadBar(false);
    }

    

    return(
        <div id="fileupload">
            <input type="file" 
                onChange={handleFileChange}
                disabled={disabled}
                className="file_input_by_user"
            />
<<<<<<< HEAD
            <button onClick={submitFile} disabled = {disabled}>
                Process
            </button>
            <button onClick={disableUploadField}>Cancel</button>
=======
            <br />
            <div className="fileuploadbuttons">
                <button onClick={submitFile} disabled = {disabled} className="button">
                    Process
                </button>
                <button onClick={disableUploadField} className="button">Cancel</button>
            </div>
>>>>>>> fc0f258 (All commit saved)
        </div>
    )
}