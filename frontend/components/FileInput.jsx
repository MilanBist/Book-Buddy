import { useState } from "react";

export default function FileInput({onFileSelect, disabled}){
    const[fileData, setFileData] = useState(null);
    // take the data from the function 
    const handleFileChange = (e)=>{
        const file = e.target.files[0];
        // if you find the file there just send this to the funciton
        if (!fileData) setFileData(file);
    }

    const submitFile = ()=>{
        if (fileData){
            onFileSelect(fileData)
        }
        else{
            alert("First select the file")
        }
    }

    return(
        <div>
            <input type="file" 
                onChange={handleFileChange}
                disabled={disabled}
                className="file_input_by_user"
            />
            <button onClick={submitFile} disabled={disabled}>
                Upload pdf
            </button>
        </div>
    )
}