<<<<<<< HEAD
=======
import '../styles/button.css'
import '../styles/liststyle.css'

import { ArrowUpRight } from "lucide-react";
>>>>>>> fc0f258 (All commit saved)
export default function ChatHistory(){
    // this will contain:
    // 1. Search Chat -> Button 
    // 2. NewChat -> Button 
    // 3. Recents -> scorll bar


    return(
        <div id="sideBar" >
<<<<<<< HEAD
            <div id="newChat">
                <button id="newChat" className="sidebarButton">New Chat</button>
            </div>
            <div id="searchChat">
                <button id="searchChat" className="sidebarButton">Search Chat</button>
            </div>
            <div id="recents">
                <h2 >Recents</h2>
                <ul>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
                    <li>Item1</li>
=======
            <h2>CHAT</h2>
            <div id="newChat">
                <button className="button">New Chat</button>
            </div>
            <div id="searchChat">
                <button className="button">Search Chat</button>
            </div>
            <div id="recents">
                <h2 >Recents</h2>
                <ul>                    
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                    <li className='list'>Item1n <ArrowUpRight className='list__icon-wrapper' size={18}/></li>
                   
>>>>>>> fc0f258 (All commit saved)
                </ul>
            </div>
        </div>
    )
}