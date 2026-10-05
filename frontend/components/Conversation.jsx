import FileUpload from "./FileUpload";
import "../styles/Conversation.css"
import axios from "axios";
import { useEffect } from "react";
import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import remarkMath from "remark-math";
import rehypeKatex from "rehype-katex";
import "katex/dist/katex.min.css";

export default function ConversationHistory({conversation, setConversation, currentBook, uploadBar, setUploadBar,  setBooksData}){

    useEffect(()=>{
        const getConversation = async (book)=>{
            // call the handler and get the data based on it
            // call the handler of the chat conversation and get the data from it
            const bookId = book["bookId"];
            const bookName = book["bookName"];
            const bookData = {
                "bookId": bookId,
                "bookName": bookName,
            };
            const token = localStorage.getItem("tokenId");

            try{
                const response = await axios.get("http://localhost:8080/api/getConversation", {
                params: bookData,
                headers: {
                    Authorization: `Bearer ${token}`,
                    }
                });

                const responseData = (await response).data;
                
                // after getting the response from the handler now the task is to show that one in the tab
                setConversation(responseData["data"]);
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
                }else if(status === 500){
                    alert("Internal server error.");
                    return;
                }
            } finally{
                console.log("Chat coversation end here.");
            }
        }
        if (currentBook) getConversation(currentBook);
    }, [currentBook]);
    
    // call the function to get the history of the conversation based on it
    return (
        <>
             <div id='conversation-section'>
                      {conversation && conversation.map((c, index) => (
                        <div key={index} className={`chat-message ${c.role}`}>
                          <ReactMarkdown 
                            remarkPlugins={[remarkGfm, remarkMath]}
                            rehypePlugins={[rehypeKatex]}>
                                {c.content}
                          </ReactMarkdown>
                        </div>
                      ))}

                      {uploadBar && <FileUpload 
                        uploadBar={uploadBar} 
                        setUploadBar={setUploadBar}
                        currentBookData={currentBook}
                        setBookData={setBooksData}
                      />}
              </div>
        </>
    )
}