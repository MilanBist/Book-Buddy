import '../styles/button.css'
import '../styles/liststyle.css'

import { ArrowUpRight } from "lucide-react";
export default function ChatHistory(){
    // this will contain:
    // 1. Search Chat -> Button 
    // 2. NewChat -> Button 
    // 3. Recents -> scorll bar


    return(
        <div id="sideBar" >
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
                   
                </ul>
            </div>
        </div>
    )
}