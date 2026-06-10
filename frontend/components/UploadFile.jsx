import { useState } from "react"

export default function UplaodFile(){

    // set there are no files in the selected postion
    const [file, setFile] = useState(null);

    const handleFile = async ()=>{
        // if the file is null just return
        if (!file) return;

        // if the file is not null
        const formData = new FormData();
        formData.append(
            "document",
            file,
        );

        // call my local host
        try{
            const resp = await fetch("http://localhost:8080/api/handlePdf", {
                method: "POST",
                body: formData,
            });

            if (resp.ok){
                const data = await resp.json()
                console.log(data)
            } else{
                console.log("No response.")
            }
            
    } catch(err){
        console.log("Error in fetching the data")
        return
    }

      
    }


    return (
       <div>
            <input 
                type="file" 
                accept="application/pdf"
                // sets the file to be the files[0] means like selected files
                onChange={(e)=>setFile(e.target.files[0])}
            /> <br />
            <button onClick={handleFile}>
                Upload Pdf
            </button>
       </div>
    )
}