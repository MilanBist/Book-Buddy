import '../src/App.css'
import { useState } from 'react'
import axios from 'axios';
import FileInput from '../components/FileInput';
import SearchBar from '../components/SearchBar';
import NavBar from '../components/NavBar';
import ChatHistory from '../components/ChatHistory';
import ReactMarkdown from "react-markdown";

export default function Home() {

  // for the language specifically
  const [language, setLanguage] = useState("English");

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
    setMessage(prev => [
    ...prev,
    {
        role: "user",
        content: promptInput
    },
    {
        role: "assistant",
        content: ""
    },
]);
    // main task here is to  get the data from the prompt input bar and send
    // to the backend localhost/api/extractAnswer or like that
    // data to send with the prompt
    const userPrompt  = {
      query: promptInput,
      language: language,
    };

    console.log(userPrompt);

    try{
      // set the prompt to be nil
      // set is prompting to be true

      setPromptInput("");
      setIsPrompting(true);
      const response =  await fetch("http://localhost:8080/api/extractAnswer", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userPrompt),
      });

      // as the response now in the form of the stream so work according to it
      // set a reader and decoder
      const reader = response.body.getReader();
      const decoder = new TextDecoder("utf-8");

      while (true) {
        // get the streaming data
        const { done, value } = await reader.read();

        // decode the given data 
        const chunk = decoder.decode(value, {stream: true});

        // if finished reading the stream data
        if (done) break;
          setMessage(prev => {
            const updated = [...prev];
            updated[updated.length - 1].content += chunk;
            return updated;
          });

        console.log("Message is: ", message);
        }
        setIsPrompting(false);
    } catch(err){
      console.log("Reaching to this catch point.")
      console.log(err);
    }
  }
  return (
    <div id='main'>

      <div id='navBar'>
        <NavBar fileUploadStatusChanger={setUploadBar} status = {uploadBar} />
      </div>
    
      <div id='chatHistory'>
        <ChatHistory/>
      </div>

      <div id='answer-section'>

        <div id='uploadBar'>
          {uploadBar && < FileInput 
            onFileSelect={handlefileUpload} 
            disabled={uploadFile} 
            uploadChanger = {setUploadFile} 
            setUpUploadBar = {setUploadBar}/>}
        </div>
        <div id="mainContent">
          {message.map((m, index) => (
            <div key={index} className={`chat-message ${m.role}`}>
              <ReactMarkdown>{m.content}</ReactMarkdown>
            </div> 
          ))}

          {/* {isPrompting && (
            <div className="chat-message assistant typing">
              Thinking...
            </div>
          )} */}
        </div>
      </div>

      <div id='typeBar' className='grid-childs'>
        <SearchBar handlePrompt = {promptUpload} setPromptInput = {setPromptInput} promptInput={promptInput} setLanguage = {setLanguage}/>
      </div>

    </div>
  )
}