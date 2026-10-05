import { Link } from 'react-router-dom';
import "../styles/NavBar.css"

export default function NavBar({uploadBar, setUploadBar} ){
    const handleUploadClick = ()=>{
        if (uploadBar == true){
            setUploadBar(false);
            return
        } else{
            setUploadBar(true);
        }
    }
    return(
        <div id="navBar-display">
            <button className="button" id="navBarButton-upload" onClick={handleUploadClick}>FileUpload</button>
            <Link to='/login'><button className="button"  id="navBarButton-loginsign">Login/Signup</button></Link>
        </div>
    )
}