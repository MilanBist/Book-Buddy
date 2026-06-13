import './App.css'
import { useState } from 'react'
import axios from 'axios';
import FileInput from '../components/FileInput';

function App() {
  // set the headers
  const [uploadProgress, setUploadProgress] = useState(null);
  const [isUploading, setIsUploading] = useState(false);
  const [result, setResult] = useState(null);
  const [uploadResponse, setUploadResponse] = useState(null);


  // handle the fileupload 
  const handlefileUpload = async (file)=>{

    console.log("Handle file is being called");
    if (file != null){
      console.log("This is also not null.")
    }

    setUploadProgress(0);
    setIsUploading(true);
    setResult(null);


    // create the instance of the form data
    const formData = new FormData();
    formData.append("document", file)
    // after appending the file now my task is to send the request using the axios to the browser
    try{
      const response = await axios.post("http://localhost:8080/api/handlePdf", formData,{
      // now keep the track of the uploadProgress
      onUploadProgress: (event)=>{
        const percent = Math.round(
          (event.loaded/event.total)*100
        );
        // change the upload progress by certain percent calculated above
        setUploadProgress(percent);
      },
    });
    console.log(response.data)
  } catch(err){
    console.log(err)
  }

    // now as the file data is read success fully and the setUploadProgress is dont now set uploading to be false
    setIsUploading(false);
  }


  return (
    <div id='main'>
      <div>
        <FileInput onFileSelect={handlefileUpload} disabled={isUploading}/>
      </div>


    </div>
  )
}

export default App
