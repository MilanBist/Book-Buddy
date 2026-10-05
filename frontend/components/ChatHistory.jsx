import { ArrowUpRight } from "lucide-react";
import "../styles/ChatHistory.css"
import apiClient from "../api/api";
import { data, useNavigate } from "react-router-dom";
import { toast } from "react-toastify";
import { useEffect } from "react";

const redirection = (time, navigate)=>{
    toast.error("Your are not authenticated. Please login/register", {
            autoClose: time,
            onClose: () => {
            navigate("/login");
        },
    });
}

export default function ChatHistory({books, setBooks, currentBook, setCurrentBook}){
    const navigate =  useNavigate();

    useEffect(()=>{
        const accessToken = localStorage.getItem("tokenId");
        if (accessToken === ""){
            console.log("Please login.");
            alert("You are not authenticated. \n Please login.")
            setTimeout(()=>{
                navigate("/login");
            }, 1000);
        return;
        }
        const getBooks = async ()=>{
        try{
            const response = await apiClient.get("/getBooks", {
                headers:{
                    Authorization: `Bearer ${accessToken}`,
                },
            })
            const newData = response.data.data;
            setBooks(newData);
            setCurrentBook(newData[0]);
        }catch(err){
            const status = err.response?.status;
            switch(status){
                case 500:
                    console.log("Error is: ", err.response.data.message);
                    alert("Internal server error.");
                
                case 401:
                    redirection(2000, navigate);
                    return;

            }
        }finally{
            console.log("Finished getting books.");
        }
    }
    getBooks();
    },[])
    

    return (
        <>
        <div id="sideBar" >
            <h2>CHAT HISTORY</h2>
            <div id="newChat">
                <button className="button">New Chat</button>
            </div>
            <div id="searchChat">
                <button className="button">Search Chat</button>
            </div>

            <div id='sidebar-bottom'>
                <div id="recents">
                    <h2 >Recents</h2>
                    <ul>   
                        { books &&
                            books.map((book)=>{
                                return(
                                <li className='list' key={book.bookId} onClick={()=>setCurrentBook(book)}>{book.bookName}<ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                                );
                            })
                        }                 
                    </ul>
                 </div>
            </div>
        </div>
        </>
    )
}