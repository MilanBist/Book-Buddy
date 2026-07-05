import '../styles/button.css'
import { Link } from 'react-router-dom';

export default function NavBar({fileUploadStatusChanger, status} ){
    const styles = {
        width: "80%",
        padding: "20px",
    };

    const handleUploadClick = ()=>{
        if (status == true){
            fileUploadStatusChanger(false)
            return
        } else{
            fileUploadStatusChanger(true)
        }
    }
    return(
        <div style={styles} id="navBar-display">
            <button class="button" id="navBarButton-upload" onClick={handleUploadClick}>FileUpload</button>
            <Link to='/login'><button class="button"  id="navBarButton-loginsign">Login/Signup</button></Link>
        </div>
    )
}