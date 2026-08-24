import { useEffect, useState } from 'react';
import '../styles/button.css'
import '../styles/liststyle.css'
import { ArrowUpRight } from "lucide-react";
import axios, { all } from 'axios';


export default function ChatHistory({allBooksData, bookConversation}){
    console.log("Chat History: ", allBooksData);
    return(
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
                        {
                            allBooksData.map((book)=>{
                                return(
                                <li className='list' key={book.bookId} onClick={()=>bookConversation(book)}>{book.bookName} Book 1<ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                                );
                            })
                        }                 
                    </ul>
                 </div>
            </div>
        </div>
    )
}