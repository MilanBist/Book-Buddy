import '../src/App.css'
import { useEffect, useState } from 'react'
import axios from 'axios';
import FileInput from '../components/FileInput';
import SearchBar from '../components/SearchBar';
import NavBar from '../components/NavBar';
import ChatHistory from '../components/ChatHistory';
import ReactMarkdown from "react-markdown";
import { useNavigate } from 'react-router-dom';
import { toast } from 'react-toastify';

export default function Home() {
  // navigate constant
  const navigate = new useNavigate();

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
  const [message, setMessage] = useState({});


  // when the processing is being done by the backend
  const [isPrompting, setIsPrompting] = useState(false);

  // set for the books that I am getting
  const [books, setBooks] = useState([]);

  // for the book id and bookName in order to receive the conversaton
  const [forConversation, setForConversation] = useState(new Map());


  // handle the fileupload 
  const handlefileUpload = async (file)=>{
    setIsUploading(true);

    // get the token string
    const token = localStorage.getItem("tokenId");
    if (!token){
      // redirect to the login page
      toast.error("Your are not authenticated. Please login/register", {
      autoClose: 3000,
       onClose: () => {
        navigate("/login");
      },
      });
    }
    // create the instance of the form data
    const formData = new FormData();
    formData.append("document", file)
    // after appending the file now my task is to send the request using the axios to the browser
    try{
      setIsUploading(true);
      const response = await axios.post("http://localhost:8080/api/handlePdf", formData, {
        headers:{
          Authorization: `Bearer ${token}`,
        }
      }
      );

      // set the uploading to be tru

      const responseData = response.data;
      // Print the error here
      const success = responseData.success;
      const successMsg = responseData.message;
      console.log("Success message: ", successMsg);


      if (success){
        alert("Your book is uploaded successfully \n Now you can ask??j");
        return;
      }
    
    } catch(err){
      console.log(err.response?.status)
      const status = err.response?.status;

      if (status === 401){
        toast.error("Your session expired please login/register.", {
          autoClose: 3000,
          onClose: () => {
          navigate("/login");
          },
          })
      }else if (status === 404) {
          console.log("Bad api request. Try agiain later.");
      } else if (status === 500) {
            alert("Server error.")
            console.log("Server error");
        }
      else {
        console.log("Network error:", error.message);
      }  
    } finally{
      setIsUploading(false);
    }
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
    // get the token


    try{
      setPromptInput("");
      setIsPrompting(true);
      const token = localStorage.getItem("tokenId");
      if (!token){
        // redirect to the login page
        toast.error("Your are not authenticated. Please login/register", {
        autoClose: 3000,
        onClose: () => {
          navigate("/login");
        },
        });
        return;
      }
      const response =  await fetch("http://localhost:8080/api/extractAnswer", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          "Authorization": `Bearer ${token}`,
        },
        body: JSON.stringify(userPrompt),
      });

      if (!response.ok){
        // check for the not authenticated and redirect to login page
        if (response.status === 401){
          toast.error("Your session expired please login/register.", {
          autoClose: 3000,
          onClose: () => {
          navigate("/login");
          },
          })
        }
        // I will get the data in the form of the string 
        const err = await response.json();
        console.log(err);
        // Print the error here
        console.log("Status Code: ", err.error.code);
        console.log("Error: ", err.error.message);

        // This is the error being obtained.
        if (err.error.message == "No table"){
          alert("Please upload the book first.")
          return;
        }
        alert(err.error.message)
        return
      }
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


  // for getting all the books titles from the server for certain user
  useEffect(()=>{
    const token = localStorage.getItem("tokenId");
    const getBooks = async ()=>{
    try{
        const response = axios.get("http://localhost:8080/api/getBooks", {
            headers: {
                Authorization: `Bearer ${token}`,
            }
        });

        const respondedData = (await response).data
        console.log("From home: ", (await response).data)
        console.log("Responed data: ", respondedData.data);


        // set the books here
        setBooks(respondedData.data);
        console.log("Book id: ", respondedData[0].bookId);
        console.log("BokName: ", respondedData[0].bookName);
    } catch(error){
        console.log(error);
    }
  }

  getBooks();
  },[])

// for getting all the conversation
  const getConversation = async (book)=>{
    // call the handler and get the data based on it
    // call the handler of the chat conversation and get the data from it
    const bookId = book.bookId;
    const bookName = book.bookName;
    const bookData = {
      "bookId": bookId,
      "bookName": bookName,
    };
    const token = localStorage.getItem("tokenId");
    try{
    axios.get("http://localhost:8080/api/getConversation", {
      params: bookData,
      headers: {
        Authorization: `Bearer ${token}`,
        }
      });
    }catch(error){
      const status = error.response?.status;
      // check for all kinds of status
      if (status === 401){
        // simply redirect to the login page
        toast.error("Your are not authenticated. Please login/register", {
        autoClose: 3000,
        onClose: () => {
        navigate("/login");
        },
        });
      }
      console.log(error);
    } finally{
      console.log("Chat coversation end here.");
    }
  }

  return (
    <div id='main'>
      <div id='navBar'>
        <NavBar fileUploadStatusChanger={setUploadBar} status = {uploadBar} />
      </div>
    
      <div id='chatHistory'>
        <ChatHistory allBooksData={books} bookConversation={getConversation}/>
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

        </div>
      </div>

      <div id='typeBar' className='grid-childs'>
        <SearchBar handlePrompt = {promptUpload} setPromptInput = {setPromptInput} promptInput={promptInput} setLanguage = {setLanguage}/>
      </div>

    </div>
  )
}