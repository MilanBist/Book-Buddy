import './App.css'
import { useState } from 'react'
import axios from 'axios';
import FileInput from '../components/FileInput';
import SearchBar from '../components/SearchBar';
import NavBar from '../components/NavBar';
import ChatHistory from '../components/ChatHistory';

function App() {
  // for the uploading of the file through form
  const [isUploading, setIsUploading] = useState(false);

  // after getting the result from the
  const [result, setResult] = useState(null);

  // for showing the upload Bar visible or not in the page
  const [uploadBar, setUploadBar] = useState(false);


  // for the uploaded file true or not
  const [uploadFile, setUploadFile] = useState(false);


  // handle the fileupload 
  const handlefileUpload = async (file)=>{
    setIsUploading(true);
    // setResult(null);

    // create the instance of the form data
    const formData = new FormData();
    formData.append("document", file)
    // after appending the file now my task is to send the request using the axios to the browser
    try{
      const response = await axios.post("http://localhost:8080/api/handlePdf", formData);
      console.log(response.data)
    // setResult(response.data)
    
    } catch(err){
        console.log(err)
    // setResult(null)
    }

  // since successfully uploaded
    setIsUploading(false);

  }

  return (
    <div id='main'>

      <div id='navBar'>
        <NavBar fileUploadStatusChanger={setUploadBar} status = {uploadBar} />
      </div>

    
      <div id='chatHistory'>
        <h2 id='chathead'>CHAT HISTORY</h2>
        <ChatHistory/>
      </div>

      <div id='upload'>This will contains the response from the backend.
        {uploadBar && < FileInput onFileSelect={handlefileUpload} disabled={uploadFile} uploadChanger = {setUploadFile} setUpUploadBar = {setUploadBar}/>}
      </div>

      <div id='typeBar'>
        <SearchBar/>
      </div>

    </div>
  )
}

export default App;
