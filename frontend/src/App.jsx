import './App.css'
import { useState } from 'react'
import axios from 'axios';
import FileInput from '../components/FileInput';
import SearchBar from '../components/SearchBar';
import NavBar from '../components/NavBar';

function App() {
  // set the headers
  const [isUploading, setIsUploading] = useState(false);
  const [result, setResult] = useState(null);
  const [uploadResponse, setUploadResponse] = useState(null);
  const [searchBar, setSearchBar] = useState(false);
  const [uploadFile, setUploadFile] = useState(false);


  // handle the fileupload 
  const handlefileUpload = async (file)=>{

    // console.log("Handle file is being called");
    // if (file != null){
    //   console.log("This is also not null.")
    // }

    setIsUploading(true);
    setResult(null);

    // create the instance of the form data
    const formData = new FormData();
    formData.append("document", file)
    // after appending the file now my task is to send the request using the axios to the browser
    try{
      const response = await axios.post("http://localhost:8080/api/handlePdf", formData,{
      // onUploadProgress: (event)=>{
      //   const percent = Math.round(
      //     (event.loaded/event.total)*100
      //   );
      //   // change the upload progress by certain percent calculated above
      // },

      
    });
    console.log(response.data)
    setResult(response.data)
    
  } catch(err){
    console.log(err)
    setResult(null)
  }

  // since successfully uploaded
    setIsUploading(false);

  }

  const handleShowUploadFile = ()=>{
    if (uploadFile){
      setUploadFile(false);
      return;
    }

    setUploadFile(true);
  }


  return (
    <div id='main'>

      <div id='navBare'>
        <NavBar/>
      </div>
      <div id='chatHistory'>ChatHistory</div>
      <div id='upload'>Upload</div>
      <div id='typeBar'>SearchBar</div>

    </div>
  )
}

export default App;
