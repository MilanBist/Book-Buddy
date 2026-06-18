export default function ChatHistory(){
    // this will contain:
    // 1. Search Chat -> Button 
    // 2. NewChat -> Button 
    // 3. Recents -> scorll bar


    return(
        <div id="sideBar" >
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
                </ul>
            </div>
        </div>
    )
}