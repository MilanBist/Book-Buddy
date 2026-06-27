import './App.css'
import { useState } from 'react'
import axios from 'axios';
import FileInput from '../components/FileInput';
import SearchBar from '../components/SearchBar';
import NavBar from '../components/NavBar';
import ChatHistory from '../components/ChatHistory';
import ReactMarkdown from "react-markdown";

function App() {

 
  // for the uploading of the file through form
  const [isUploading, setIsUploading] = useState(false);

  // after getting the result from the
  const [result, setResult] = useState(null);

  // for showing the upload Bar visible or not in the page
  const [uploadBar, setUploadBar] = useState(false);

  // for the uploaded file true or not
  const [uploadFile, setUploadFile] = useState(false);

  // set for the prompt input
  const [promptInput, setPromptInput] = useState("");

  // for setting the data to the main place
  const [data, setData] = useState("");
  const [message, setMessage] = useState([]);



  // when the processing is being done by the backend
  const [isPrompting, setIsPrompting] = useState(false);


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

  // handle the promptUpload
  const promptUpload = async () =>{
    console.log("Incoming prompt: ",promptInput);
    // set the user message here
    setMessage(prev=>{
      return [
        ...prev,
        {role: "user", content: promptInput}
      ]
    });
    // main task here is to  get the data from the prompt input bar and send
    // to the backend localhost/api/extractAnswer or like that

    try{
      setPromptInput("");
      setIsPrompting(true);
      const response =  await axios.post("http://localhost:8080/api/extractDocuments", {
        Query: promptInput,
      });

      const result = response.data["Response"];
      setIsPrompting(false);
      if(result !=  null){
        setData(result);

        console.log(result);
        setMessage(prev=>{
          return [
          ...prev,
          {role: "assistant", content: result}
          ]
        });
      }
      console.log(message);
    } catch(err){
      console.log("Reaching to this catch point.")
      console.log(err);
    }
    
  }
  return (
    <div id='grid-container'>

      <div className='grid-childs' id='navBar'>
        <NavBar fileUploadStatusChanger={setUploadBar} status = {uploadBar} />
      </div>
    
      <div className='grid-childs' id='chat-history'>
        <ChatHistory/>
      </div>

      <div className='grid-childs' id='answer-section'>

      <div id='uploadBar'>
          {uploadBar && < FileInput 
            onFileSelect={handlefileUpload} 
            disabled={uploadFile} 
            uploadChanger = {setUploadFile} 
            setUpUploadBar = {setUploadBar}/>}
      </div>
        <div id='mainContent'>
          {isPrompting ? (
            (message.map((m, index) =>{
            return(
            <div key={index} className={`chat-message ${m.role}`}>
              <ReactMarkdown>{m.content}</ReactMarkdown>
            </div>
            )
          }))
          ) : (message.map((m, index) =>{
            return(
            <div key={index} className={`chat-message ${m.role}`}>
              <ReactMarkdown>{m.content}</ReactMarkdown>
            </div>
            )
          }))}
        </div>
      </div>

      <div id='typeBar' className='grid-childs'>
        <SearchBar handlePrompt = {promptUpload} setPromptInput = {setPromptInput} promptInput={promptInput}/>
      </div>

    </div>
  )
}

export default App;